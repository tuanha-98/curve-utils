package token

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	erc20 "github.com/tuanha-98/curve-utils/contract/erc20"

	tracer "github.com/tuanha-98/curve-utils/internal/entities/tracer"
)

// NewToken creates a new Token instance for the given address
func NewToken(
	ctx context.Context,
	client *ethclient.Client,
	rpcManager *tracer.RPCManager,
	address common.Address,
) (*Token, error) {
	// Special case for Ether
	if address == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
		return &Token{
			Address:  address.String(),
			Name:     "Ether",
			Symbol:   "ETH",
			Decimals: 18,
		}, nil
	}

	// Define factory for ERC20 contract
	factory := func(addr common.Address, client *ethclient.Client) (tracer.ContractCallerInterface, error) {
		return erc20.NewContractCaller(addr, client)
	}

	// Create contract caller with retries
	contractCaller, err := tracer.NewGenericContractCaller(ctx, client, address, rpcManager, factory)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract caller for %s: %v", address.Hex(), err)
	}

	// Fetch token details with retries
	type callResult struct {
		name   string
		result interface{}
	}
	calls := []callResult{
		{"Name", nil},
		{"Symbol", nil},
		{"Decimals", nil},
		{"TotalSupply", nil},
	}

	for i, call := range []func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error){
		func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			erc20Contract, ok := c.(*erc20.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return erc20Contract.Name(nil)
		},
		func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			erc20Contract, ok := c.(*erc20.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return erc20Contract.Symbol(nil)
		},
		func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			erc20Contract, ok := c.(*erc20.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return erc20Contract.Decimals(nil)
		},
		func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			erc20Contract, ok := c.(*erc20.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return erc20Contract.TotalSupply(nil)
		},
	} {
		result, err := contractCaller.CallWithRetry(ctx, call)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch %s for %s: %v", calls[i].name, address.Hex(), err)
		}
		calls[i].result = result
	}

	// Extract results
	name := calls[0].result.(string)
	symbol := calls[1].result.(string)
	decimals := calls[2].result.(uint8)
	totalSupply := calls[3].result.(*big.Int)

	return &Token{
		Address:     address.String(),
		Name:        name,
		Symbol:      symbol,
		Decimals:    decimals,
		TotalSupply: *uint256.MustFromBig(totalSupply),
	}, nil
}
