package stableng

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
	stableNGContract "github.com/tuanha-98/curve-utils/contract/stableswap-ng"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	tracer "github.com/tuanha-98/curve-utils/internal/entities/tracer"
)

type PoolResult struct {
	poolAddr string
	err      error
}

func TestGetDYNgPool(t *testing.T) {
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
	file, err := os.Open("stable_ng_pools.json")
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
		// Per-pool timeout
		poolCtx, poolCancel := context.WithTimeout(ctx, 20*time.Second)
		defer poolCancel()
		defer fmt.Printf("Finished processing pool %s\n", poolAddr)

		fmt.Printf("Processing pool: %s\n", poolAddr)
		poolAddress := common.HexToAddress(poolAddr)

		// Initialize client and contract
		client, err := rpcManager.Dial(poolCtx)
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to connect to RPC: %v", err)}
			return
		}
		defer client.Close()

		factory := func(addr common.Address, client *ethclient.Client) (tracer.ContractCallerInterface, error) {
			return stableNGContract.NewContractCaller(addr, client)
		}

		contractCaller, err := tracer.NewGenericContractCaller(poolCtx, client, poolAddress, rpcManager, factory)
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to create contract instance: %v", err)}
			return
		}

		numTokensResult, err := contractCaller.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			stableContract, ok := c.(*stableNGContract.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return stableContract.NCOINS(nil)
		})
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch NCoins: %v", err)}
			return
		}
		numTokens := numTokensResult.(*big.Int).Int64()

		// Fetch pool parameters
		parameters := []struct {
			name   string
			call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
			result interface{}
		}{
			{"APrecise", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.APrecise(nil)
			}, nil},
			{"A", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.A(nil)
			}, nil},
			{"InitialA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.InitialA(nil)
			}, nil},
			{"FutureA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.FutureA(nil)
			}, nil},
			{"InitialATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.InitialATime(nil)
			}, nil},
			{"FutureATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.FutureATime(nil)
			}, nil},
			{"TotalSupply", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.TotalSupply(nil)
			}, nil},
			{"Fee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.Fee(nil)
			}, nil},
			{"AdminFee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.AdminFee(nil)
			}, nil},
			{"StoredRates", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.StoredRates(nil)
			}, nil},
			{"OffpegFeeMultiplier", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.OffpegFeeMultiplier(nil)
			}, nil},
		}

		// Dynamically add Balance and Coin calls based on N_COINS
		for i := int64(0); i < numTokens; i++ {
			iBig := big.NewInt(i)
			parameters = append(parameters,
				struct {
					name   string
					call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
					result interface{}
				}{
					name: fmt.Sprintf("Balance%d", i),
					call: func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
						stableContract, ok := c.(*stableNGContract.ContractCaller)
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
						stableContract, ok := c.(*stableNGContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return stableContract.Coins(nil, iBig)
					},
					result: nil,
				},
			)
		}

		for i, param := range parameters {
			result, err := contractCaller.CallWithRetry(poolCtx, param.call)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch %s: %v", param.name, err)}
				return
			}
			parameters[i].result = result
		}

		// Extract results
		APrecise := parameters[0].result.(*big.Int)
		A := parameters[1].result.(*big.Int)
		InitialA := parameters[2].result.(*big.Int)
		FutureA := parameters[3].result.(*big.Int)
		InitialATime := parameters[4].result.(*big.Int)
		FutureATime := parameters[5].result.(*big.Int)
		totalSupply := parameters[6].result.(*big.Int)
		fee := parameters[7].result.(*big.Int)
		adminFee := parameters[8].result.(*big.Int)
		rates := parameters[9].result.([]*big.Int)
		offPegFeeMultiplier := parameters[10].result.(*big.Int)

		balances := make([]uint256.Int, numTokens)
		tokens := make([]token.Token, numTokens)

		for i := int64(0); i < numTokens; i++ {
			balance := parameters[11+2*i].result.(*big.Int)
			coin := parameters[12+2*i].result.(common.Address)
			balances[i] = *uint256.MustFromBig(balance)
			token, err := token.NewToken(coin)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to create token%d: %v", i, err)}
				return
			}
			tokens[i] = *token
		}

		// Convert rates to uint256
		_rates := make([]uint256.Int, len(rates))
		for i, rate := range rates {
			_rates[i] = *uint256.MustFromBig(rate)
		}

		// Create pool
		stablePool := NewPool(
			poolAddr,
			DexType,
			balances,
			tokens,
			*uint256.MustFromBig(new(big.Int).Div(APrecise, A)),
			*uint256.MustFromBig(offPegFeeMultiplier),
			*uint256.MustFromBig(InitialA),
			*uint256.MustFromBig(FutureA),
			*uint256.MustFromBig(fee),
			*uint256.MustFromBig(adminFee),
			*uint256.MustFromBig(totalSupply),
			_rates,
			InitialATime.Int64(),
			FutureATime.Int64(),
		)

		// Calculate DY
		var amountOut, amount, feeAdmin uint256.Int
		amount.SetFromDecimal("1000")

		if err := stablePool.GetDy(0, 1, &amount, &amountOut, &feeAdmin); err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to get DY: %v", err)}
			return
		}
		fmt.Printf("DY calculated for pool %s: %s\n", poolAddr, amountOut.String())

		// Fetch DY from contract with retry
		contractDyResult, err := contractCaller.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			stableContract, ok := c.(*stableNGContract.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return stableContract.GetDy(nil, big.NewInt(0), big.NewInt(1), amount.ToBig())
		})
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch GetDy: %v", err)}
			return
		}
		contractDy := contractDyResult.(*big.Int)

		// Compare results
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
