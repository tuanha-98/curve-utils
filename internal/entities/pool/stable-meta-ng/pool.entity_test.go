package stablemetang

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	stableMetaNGContract "github.com/tuanha-98/curve-utils/contract/stableswap-meta-ng"
	stableNGContract "github.com/tuanha-98/curve-utils/contract/stableswap-ng"
	"github.com/tuanha-98/curve-utils/internal/entities/pool/stable"
	stableng "github.com/tuanha-98/curve-utils/internal/entities/pool/stable-ng"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/entities/tracer"
)

type PoolResult struct {
	poolAddr string
	err      error
}

func TestGetDYNgPool(t *testing.T) {
	rpcs := []string{
		"https://eth.drpc.org",
		"https://eth.blockrazor.xyz",
		"https://eth.meowrpc.com",
		"https://api.zan.top/eth-mainnet",
		"https://ethereum.blockpi.network/v1/rpc/public",
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
	file, err := os.Open("stable_ng_meta_pools.json")
	if err != nil {
		t.Fatalf("Failed to open pools file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&pools); err != nil {
		t.Fatalf("Failed to decode pools JSON: %v", err)
	}

	// Worker pool setup
	const maxWorkers = 5
	results := make(chan PoolResult, len(pools))
	var wg sync.WaitGroup

	worker := func(poolAddr string) {
		defer wg.Done()
		poolCtx, poolCancel := context.WithTimeout(ctx, 20*time.Second)
		defer poolCancel()
		defer log.Printf("Finished processing pool %s", poolAddr)

		log.Printf("Processing pool: %s", poolAddr)
		poolAddress := common.HexToAddress(poolAddr)

		// Initialize client
		client, err := rpcManager.Dial(poolCtx)
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to connect to RPC: %v", err)}
			return
		}
		defer client.Close()

		// Meta pool contract caller
		metaFactory := func(addr common.Address, client *ethclient.Client) (tracer.ContractCallerInterface, error) {
			return stableMetaNGContract.NewContractCaller(addr, client)
		}
		metaContract, err := tracer.NewGenericContractCaller(poolCtx, client, poolAddress, rpcManager, metaFactory)
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to create meta contract caller: %v", err)}
			return
		}

		// Fetch meta pool token count
		numTokensResult, err := metaContract.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			contract, ok := c.(*stableMetaNGContract.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return contract.NCOINS(nil)
		})
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch NCOINS: %v", err)}
			return
		}
		numTokens := numTokensResult.(*big.Int).Int64()
		if numTokens != 2 {
			results <- PoolResult{poolAddr, fmt.Errorf("expected 2 tokens for meta pool, got %d", numTokens)}
			return
		}

		// Fetch base pool address
		baseAddrResult, err := metaContract.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			contract, ok := c.(*stableMetaNGContract.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return contract.BASEPOOL(nil)
		})
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch BASEPOOL: %v", err)}
			return
		}
		baseAddr := baseAddrResult.(common.Address)

		// Fetch base pool token count
		baseNumTokensResult, err := metaContract.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			contract, ok := c.(*stableMetaNGContract.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return contract.BASENCOINS(nil)
		})
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch base NCOINS: %v", err)}
			return
		}
		baseNumTokens := baseNumTokensResult.(*big.Int).Int64()

		// Base pool contract caller
		baseFactory := func(addr common.Address, client *ethclient.Client) (tracer.ContractCallerInterface, error) {
			return stableNGContract.NewContractCaller(addr, client)
		}
		baseContract, err := tracer.NewGenericContractCaller(poolCtx, client, baseAddr, rpcManager, baseFactory)
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to create base contract caller: %v", err)}
			return
		}

		// Fetch meta pool parameters
		metaParams := []struct {
			name   string
			call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
			result interface{}
		}{
			{"APrecise", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.APrecise(nil)
			}, nil},
			{"A", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.A(nil)
			}, nil},
			{"InitialA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.InitialA(nil)
			}, nil},
			{"FutureA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.FutureA(nil)
			}, nil},
			{"InitialATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.InitialATime(nil)
			}, nil},
			{"FutureATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.FutureATime(nil)
			}, nil},
			{"Fee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.Fee(nil)
			}, nil},
			{"AdminFee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.AdminFee(nil)
			}, nil},
			{"StoredRates", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.StoredRates(nil)
			}, nil},
			{"OffpegFeeMultiplier", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.OffpegFeeMultiplier(nil)
			}, nil},
			{"TotalSupply", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableMetaNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.TotalSupply(nil)
			}, nil},
		}

		// Add meta pool Balance and Coin calls
		for i := int64(0); i < numTokens; i++ {
			iBig := big.NewInt(i)
			metaParams = append(metaParams,
				struct {
					name   string
					call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
					result interface{}
				}{
					name: fmt.Sprintf("Balance%d", i),
					call: func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
						contract, ok := c.(*stableMetaNGContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return contract.Balances(nil, iBig)
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
						contract, ok := c.(*stableMetaNGContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return contract.Coins(nil, iBig)
					},
					result: nil,
				},
			)
		}

		// Fetch meta pool parameters
		for i, param := range metaParams {
			result, err := metaContract.CallWithRetry(poolCtx, param.call)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch %s: %v", param.name, err)}
				return
			}
			metaParams[i].result = result
		}

		// Extract meta pool results
		APrecise := metaParams[0].result.(*big.Int)
		A := metaParams[1].result.(*big.Int)
		InitialA := metaParams[2].result.(*big.Int)
		FutureA := metaParams[3].result.(*big.Int)
		InitialATime := metaParams[4].result.(*big.Int)
		FutureATime := metaParams[5].result.(*big.Int)
		fee := metaParams[6].result.(*big.Int)
		adminFee := metaParams[7].result.(*big.Int)
		rates := metaParams[8].result.([]*big.Int)
		offPegFeeMultiplier := metaParams[9].result.(*big.Int)
		totalSupply := metaParams[10].result.(*big.Int)

		balances := make([]uint256.Int, numTokens)
		tokens := make([]token.Token, numTokens)
		for i := int64(0); i < numTokens; i++ {
			balance := metaParams[11+2*i].result.(*big.Int)
			coin := metaParams[12+2*i].result.(common.Address)
			balances[i] = *uint256.MustFromBig(balance)
			token, err := token.NewToken(poolCtx, client, rpcManager, coin)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to create base token %d: %v", i, err)}
				return
			}
			tokens[i] = *token
		}

		_rates := make([]uint256.Int, len(rates))
		for i, rate := range rates {
			_rates[i] = *uint256.MustFromBig(rate)
		}

		// Fetch base pool parameters
		baseParams := []struct {
			name   string
			call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
			result interface{}
		}{
			{"APrecise", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.APrecise(nil)
			}, nil},
			{"A", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.A(nil)
			}, nil},
			{"InitialA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.InitialA(nil)
			}, nil},
			{"FutureA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.FutureA(nil)
			}, nil},
			{"InitialATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.InitialATime(nil)
			}, nil},
			{"FutureATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.FutureATime(nil)
			}, nil},
			{"Fee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.Fee(nil)
			}, nil},
			{"AdminFee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.AdminFee(nil)
			}, nil},
			{"StoredRates", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.StoredRates(nil)
			}, nil},
			{"OffpegFeeMultiplier", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				contract, ok := c.(*stableNGContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return contract.OffpegFeeMultiplier(nil)
			}, nil},
		}

		// Add base pool Balance and Coin calls
		for i := int64(0); i < baseNumTokens; i++ {
			iBig := big.NewInt(i)
			baseParams = append(baseParams,
				struct {
					name   string
					call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
					result interface{}
				}{
					name: fmt.Sprintf("Balance%d", i),
					call: func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
						contract, ok := c.(*stableNGContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return contract.Balances(nil, iBig)
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
						contract, ok := c.(*stableNGContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return contract.Coins(nil, iBig)
					},
					result: nil,
				},
			)
		}

		// Fetch base pool parameters
		for i, param := range baseParams {
			result, err := baseContract.CallWithRetry(poolCtx, param.call)
			if err != nil {
				if param.name == "APrecise" {
					baseParams[i].result = big.NewInt(0)
					continue
				}
				if param.name == "StoredRates" {
					baseParams[i].result = nil
					continue
				}
				if param.name == "OffpegFeeMultiplier" {
					baseParams[i].result = big.NewInt(0)
					continue
				}
				results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch base %s: %v", param.name, err)}
				return
			}
			baseParams[i].result = result
		}

		// Extract base pool results
		baseAPrecise := baseParams[0].result.(*big.Int)
		baseA := baseParams[1].result.(*big.Int)
		baseInitialA := baseParams[2].result.(*big.Int)
		baseFutureA := baseParams[3].result.(*big.Int)
		baseInitialATime := baseParams[4].result.(*big.Int)
		baseFutureATime := baseParams[5].result.(*big.Int)
		baseFee := baseParams[6].result.(*big.Int)
		baseAdminFee := baseParams[7].result.(*big.Int)
		baseRates := baseParams[8].result
		baseOffPegFeeMultiplier := baseParams[9].result.(*big.Int)

		// Fetch base pool balances and tokens
		baseBalances := make([]uint256.Int, baseNumTokens)
		baseTokens := make([]token.Token, baseNumTokens)
		for i := int64(0); i < baseNumTokens; i++ {
			balance := baseParams[10+2*i].result.(*big.Int)
			coin := baseParams[11+2*i].result.(common.Address)
			baseBalances[i] = *uint256.MustFromBig(balance)
			token, err := token.NewToken(poolCtx, client, rpcManager, coin)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to create base token %d: %v", i, err)}
				return
			}
			baseTokens[i] = *token
		}

		// Convert base rates (handle nil case)
		var _baseRates []uint256.Int
		if baseRates != nil {
			ratesSlice, ok := baseRates.([]*big.Int)
			if !ok {
				results <- PoolResult{poolAddr, fmt.Errorf("invalid base rates type")}
				return
			}
			_baseRates = make([]uint256.Int, len(ratesSlice))
			for i, rate := range ratesSlice {
				_baseRates[i] = *uint256.MustFromBig(rate)
			}
		}

		// Create pool
		pool := NewPool(
			poolAddr,
			baseAddr.String(),
			DexType,
			func() string {
				if baseOffPegFeeMultiplier.Cmp(big.NewInt(0)) == 0 {
					return stable.DexType
				} else {
					return stableng.DexType
				}
			}(),
			balances,
			baseBalances,
			tokens,
			baseTokens,
			*uint256.MustFromBig(new(big.Int).Div(APrecise, A)),
			*uint256.MustFromBig(offPegFeeMultiplier),
			*uint256.MustFromBig(InitialA),
			*uint256.MustFromBig(FutureA),
			*uint256.MustFromBig(fee),
			*uint256.MustFromBig(adminFee),
			*uint256.MustFromBig(totalSupply),
			*uint256.MustFromBig(func() *big.Int {
				if baseAPrecise.Cmp(big.NewInt(0)) == 0 {
					return big.NewInt(1)
				}
				return new(big.Int).Div(baseAPrecise, baseA)
			}()),
			*uint256.MustFromBig(baseOffPegFeeMultiplier),
			*uint256.MustFromBig(baseInitialA),
			*uint256.MustFromBig(baseFutureA),
			*uint256.MustFromBig(baseFee),
			*uint256.MustFromBig(baseAdminFee),
			tokens[1].TotalSupply,
			_rates,
			_baseRates,
			InitialATime.Int64(),
			FutureATime.Int64(),
			baseInitialATime.Int64(),
			baseFutureATime.Int64(),
		)

		// Calculate DY
		var amountOut, amount uint256.Int
		amount.SetFromDecimal("1000000000000000000") // 1 token (10^18)

		if err := pool.GetDyUnderlying(0, 2, &amount, &amountOut); err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to get DY: %v", err)}
			return
		}
		fmt.Printf("DY calculated for pool %s: %s", poolAddr, amountOut.String())

		// Fetch DY from contract
		contractDyResult, err := metaContract.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			contract, ok := c.(*stableMetaNGContract.ContractCaller)
			if !ok {
				return nil, fmt.Errorf("invalid contract type")
			}
			return contract.GetDyUnderlying(nil, big.NewInt(0), big.NewInt(2), amount.ToBig())
		})
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch GetDyUnderlying: %v", err)}
			return
		}
		contractDy := contractDyResult.(*big.Int)

		// Compare results
		contractDyInt := uint256.MustFromBig(contractDy)
		if amountOut.Cmp(contractDyInt) != 0 {
			results <- PoolResult{poolAddr, fmt.Errorf("DY values differ: got %s, want %s", amountOut.String(), contractDyInt.String())}
		} else {
			fmt.Printf("DY from contract for pool %s: %s", poolAddr, contractDyInt.String())
			results <- PoolResult{poolAddr, nil}
		}
	}

	// Start worker pool
	semaphore := make(chan struct{}, maxWorkers)
	for _, poolAddr := range pools {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(addr string) {
			defer func() { <-semaphore }()
			worker(addr)
		}(poolAddr)
	}

	// Close results channel
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
