package tricryptong

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	triCrytoContract "github.com/tuanha-98/curve-utils/contract/tricryptoswap"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/entities/tracer"
)

type PoolResult struct {
	poolAddr string
	err      error
}

func TestGetTriCryptoPool(t *testing.T) {
	rpcs := []string{
		"https://eth.drpc.org",
		"https://eth.blockrazor.xyz",
		"https://rpc.therpc.io/ethereum",
		"https://eth-pokt.nodies.app",
		"https://mainnet.gateway.tenderly.co",
		"https://ethereum-rpc.publicnode.com",
	}

	rpcManager := tracer.NewRPCManager(rpcs)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Load pool addresses from JSON file
	var pools []string
	file, err := os.Open("tricrypto_pools.json")
	if err != nil {
		t.Fatalf("Failed to open pools file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		t.Fatalf("Failed to decode pools JSON: %v", err)
	}

	// Worker pool setup
	const maxWorkers = 5 // Reduced to avoid RPC contention
	results := make(chan PoolResult, len(pools))
	var wg sync.WaitGroup

	worker := func(poolAddr string) {
		defer wg.Done()

		poolCtx, poolCancel := context.WithTimeout(ctx, 20*time.Second)
		defer poolCancel()
		defer fmt.Printf("Finished processing pool %s\n", poolAddr)

		fmt.Printf("Processing pool %s\n", poolAddr)
		poolAddress := common.HexToAddress(poolAddr)

		client, err := rpcManager.Dial(poolCtx)
		if err != nil {
			results <- PoolResult{poolAddr: poolAddr, err: fmt.Errorf("failed to dial RPC: %w", err)}
			return
		}
		defer client.Close()

		factory := func(address common.Address, client *ethclient.Client) (tracer.ContractCallerInterface, error) {
			return triCrytoContract.NewContractCaller(address, client)
		}

		contractCaller, err := tracer.NewGenericContractCaller(poolCtx, client, poolAddress, rpcManager, factory)
		if err != nil {
			results <- PoolResult{poolAddr: poolAddr, err: fmt.Errorf("failed to create contract caller: %w", err)}
			return
		}

		numTokens := 3

		params := []struct {
			name   string
			call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
			result interface{}
		}{
			{"D", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.D(nil)
			}, nil},
			{"InitialAGamma", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.InitialAGamma(nil)
			}, nil},
			{"InitialAGammaTime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.InitialAGammaTime(nil)
			}, nil},
			{"FutureAGamma", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.FutureAGamma(nil)
			}, nil},
			{"FutureAGammaTime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.FutureAGammaTime(nil)
			}, nil},
			{"MidFee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.MidFee(nil)
			}, nil},
			{"OutFee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.OutFee(nil)
			}, nil},
			{"FeeGamma", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*triCrytoContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.FeeGamma(nil)
			}, nil},
		}

		// Dynamically add Balance and Coin calls based on N_COINS
		for i := 0; i < numTokens; i++ {
			iBig := big.NewInt(int64(i))
			params = append(params,
				struct {
					name   string
					call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
					result interface{}
				}{
					name: fmt.Sprintf("Balance%d", i),
					call: func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
						stableContract, ok := c.(*triCrytoContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return stableContract.Balances(nil, iBig)
					},
					result: nil,
				},
				struct {
					name   string
					call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
					result interface{}
				}{
					name: fmt.Sprintf("Coin%d", i),
					call: func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
						stableContract, ok := c.(*triCrytoContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return stableContract.Coins(nil, iBig)
					},
					result: nil,
				},
			)
		}

		// Add PriceScale calls
		for i := 0; i < numTokens-1; i++ {
			iBig := big.NewInt(int64(i))
			params = append(params,
				struct {
					name   string
					call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
					result interface{}
				}{
					name: fmt.Sprintf("PriceScale%d", i),
					call: func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
						stableContract, ok := c.(*triCrytoContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return stableContract.PriceScale(nil, iBig)
					},
					result: nil,
				},
			)
		}

		for i, param := range params {
			result, err := contractCaller.CallWithRetry(poolCtx, param.call)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch %s: %v", param.name, err)}
				return
			}
			params[i].result = result
		}

		D := params[0].result.(*big.Int)
		initialAGamma := params[1].result.(*big.Int)
		initialAGammaTime := params[2].result.(*big.Int)
		futureAGamma := params[3].result.(*big.Int)
		futureAGammaTime := params[4].result.(*big.Int)
		midFee := params[5].result.(*big.Int)
		outFee := params[6].result.(*big.Int)
		feeGamma := params[7].result.(*big.Int)

		priceScale := make([]uint256.Int, numTokens-1)
		balances := make([]uint256.Int, numTokens)
		tokens := make([]token.Token, numTokens)

		for i := 0; i < numTokens; i++ {
			balance := params[8+i*2].result.(*big.Int)
			coin := params[9+i*2].result.(common.Address)
			balances[i] = *uint256.MustFromBig(balance)
			token, err := token.NewToken(poolCtx, client, rpcManager, coin)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to create token%d: %v", i, err)}
				return
			}
			tokens[i] = *token
		}

		for i := 0; i < numTokens-1; i++ {
			pricescale := params[14+i].result.(*big.Int)
			priceScale[i] = *uint256.MustFromBig(pricescale)
		}

		pool := NewPool(
			poolAddr,
			"TriCryptoSwap",
			balances,
			tokens,
			*uint256.MustFromBig(initialAGamma),
			*uint256.MustFromBig(futureAGamma),
			*uint256.MustFromBig(D),
			*uint256.MustFromBig(feeGamma),
			*uint256.MustFromBig(midFee),
			*uint256.MustFromBig(outFee),
			priceScale,
			initialAGammaTime.Int64(),
			futureAGammaTime.Int64(),
		)

		var amountOut, fee, amount, K0 uint256.Int
		amount.SetFromDecimal("1000000000")

		if err := pool.GetDy(0, 1, &amount, &amountOut, &fee, &K0, balances); err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to get DY: %v", err)}
			return
		}
		fmt.Printf("DY calculated for pool %s: %s\n", poolAddr, amountOut.String())

		contractDyResult, err := contractCaller.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			contract, ok := c.(*triCrytoContract.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return contract.GetDy(nil, big.NewInt(int64(0)), big.NewInt(int64(1)), amount.ToBig())
		})
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to get DY from contract: %v", err)}
			return
		}
		contractDy := contractDyResult.(*big.Int)

		contractDyInt := uint256.MustFromBig(contractDy)
		diff := new(big.Int).Sub(amountOut.ToBig(), contractDyInt.ToBig())
		if new(big.Int).Abs(diff).Cmp(big.NewInt(2)) > 0 {
			results <- PoolResult{poolAddr, fmt.Errorf("DY values differ by more than 2 wei: got %s, want %s", amountOut.String(), contractDyInt.String())}
		} else {
			fmt.Printf("DY from contract for pool %s: %s\n", poolAddr, contractDyInt.String())
			results <- PoolResult{poolAddr, nil}
		}
	}
	// Start worker pool
	semaphore := make(chan struct{}, maxWorkers)
	for _, poolAddr := range pools {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire semaphore
		go func(addr string) {
			defer func() { <-semaphore }() // Release semaphore
			worker(addr)
		}(poolAddr)
	}

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and report results
	for res := range results {
		if res.err != nil {
			t.Errorf("Error processing pool %s: %v", res.poolAddr, res.err)
		}
	}
}
