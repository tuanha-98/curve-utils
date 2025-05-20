package tracer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractCallerInterface defines the interface for contract operations
type ContractCallerInterface interface {
	// Add methods as needed for specific contract interactions
	// For now, we use a generic CallWithRetry, so no specific methods are required
}

// ContractFactory is a function that creates a new contract instance
type ContractFactory func(address common.Address, client *ethclient.Client) (ContractCallerInterface, error)

// GenericContractCaller handles contract calls with retry logic for any contract type
type GenericContractCaller struct {
	client     *ethclient.Client
	contract   ContractCallerInterface
	rpcManager *RPCManager
	address    common.Address
	factory    ContractFactory
}

func NewGenericContractCaller(
	ctx context.Context,
	client *ethclient.Client,
	poolAddress common.Address,
	rpcManager *RPCManager,
	factory ContractFactory,
) (*GenericContractCaller, error) {
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		contract, err := factory(poolAddress, client)
		if err == nil {
			return &GenericContractCaller{
				client:     client,
				contract:   contract,
				rpcManager: rpcManager,
				address:    poolAddress,
				factory:    factory,
			}, nil
		}
		log.Printf("Contract creation attempt %d failed for %s: %v", i+1, poolAddress.Hex(), err)
		client, err = rpcManager.Dial(ctx)
		if err != nil {
			log.Printf("RPC reconnection attempt %d failed: %v", i+1, err)
			continue
		}
	}
	return nil, fmt.Errorf("failed to create contract instance for %s after retries", poolAddress.Hex())
}

// CallWithRetry executes a contract call with retries
func (cc *GenericContractCaller) CallWithRetry(
	ctx context.Context,
	call func(contract ContractCallerInterface, client *ethclient.Client) (interface{}, error),
) (interface{}, error) {
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		result, err := call(cc.contract, cc.client)
		if err == nil {
			return result, nil
		}
		log.Printf("Contract call attempt %d failed for %s: %v", i+1, cc.address.Hex(), err)
		client, err := cc.rpcManager.Dial(ctx)
		if err != nil {
			log.Printf("RPC reconnection attempt %d failed: %v", i+1, err)
			continue
		}
		cc.client = client
		contract, err := cc.factory(cc.address, client)
		if err != nil {
			log.Printf("Contract recreation attempt %d failed for %s: %v", i+1, cc.address.Hex(), err)
			continue
		}
		cc.contract = contract
	}
	return nil, fmt.Errorf("failed to execute contract call for %s after retries", cc.address.Hex())
}
