package stable

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
	stableContract "github.com/tuanha-98/curve-utils/contract/stableswap"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/entities/tracer"
)

type PoolResult struct {
	poolAddr string
	err      error
}

func TestGetDYStablePool(t *testing.T) {
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
	file, err := os.Open("stable_pools.json")
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
			return stableContract.NewContractCaller(addr, client)
		}

		contractCaller, err := tracer.NewGenericContractCaller(poolCtx, client, poolAddress, rpcManager, factory)
		if err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to create contract instance: %v", err)}
			return
		}

		var numTokens int64
		for i := int64(0); ; i++ {
			_, err := contractCaller.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.Balances(nil, big.NewInt(i))
			})
			if err != nil {
				break
			}
			numTokens++
		}

		// Fetch pool parameters
		parameters := []struct {
			name   string
			call   func(tracer.ContractCallerInterface, *ethclient.Client) (interface{}, error)
			result interface{}
		}{
			{"APrecise", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.APrecise(nil)
			}, nil},
			{"A", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.A(nil)
			}, nil},
			{"InitialA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.InitialA(nil)
			}, nil},
			{"FutureA", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.FutureA(nil)
			}, nil},
			{"InitialATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.InitialATime(nil)
			}, nil},
			{"FutureATime", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.FutureATime(nil)
			}, nil},
			{"LpToken", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.LpToken(nil)
			}, nil},
			{"Fee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.Fee(nil)
			}, nil},
			{"AdminFee", func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
				stableContract, ok := c.(*stableContract.ContractCaller)
				if !ok {
					return nil, fmt.Errorf("invalid contract type")
				}
				return stableContract.AdminFee(nil)
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
						stableContract, ok := c.(*stableContract.ContractCaller)
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
						stableContract, ok := c.(*stableContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return stableContract.Coins(nil, iBig)
					},
					result: nil,
				},
			)
		}

		var totalSupply *big.Int
		for i, param := range parameters {
			result, err := contractCaller.CallWithRetry(poolCtx, param.call)
			if err != nil {
				if param.name == "LpToken" {
					result, err := contractCaller.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
						stableContract, ok := c.(*stableContract.ContractCaller)
						if !ok {
							return nil, fmt.Errorf("invalid contract type")
						}
						return stableContract.TotalSupply(nil)
					})
					if err != nil {
						results <- PoolResult{poolAddr, fmt.Errorf("failed to fetch TotalSupply: %v", err)}
						return
					}
					parameters[i].result = common.Address{}
					totalSupply = result.(*big.Int)
					continue
				}

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
		LpTokenAddr := parameters[6].result.(common.Address)
		Fee := parameters[7].result.(*big.Int)
		AdminFee := parameters[8].result.(*big.Int)

		balances := make([]uint256.Int, numTokens)
		tokens := make([]token.Token, numTokens)

		for i := int64(0); i < numTokens; i++ {
			balance := parameters[9+2*i].result.(*big.Int)
			coin := parameters[10+2*i].result.(common.Address)
			balances[i] = *uint256.MustFromBig(balance)
			token, err := token.NewToken(poolCtx, client, rpcManager, coin)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to create token%d: %v", i, err)}
				return
			}
			tokens[i] = *token
		}

		var lpToken *token.Token
		if (LpTokenAddr != common.Address{}) {
			var err error
			lpToken, err = token.NewToken(poolCtx, client, rpcManager, LpTokenAddr)
			if err != nil {
				results <- PoolResult{poolAddr, fmt.Errorf("failed to create LP token: %v", err)}
				return
			}
			totalSupply = lpToken.TotalSupply.ToBig()
		}

		// Create pool
		stablePool := NewPool(
			poolAddr,
			DexType,
			nil,
			balances,
			tokens,
			*uint256.MustFromBig(new(big.Int).Div(APrecise, A)),
			*uint256.MustFromBig(InitialA),
			*uint256.MustFromBig(FutureA),
			*uint256.MustFromBig(Fee),
			*uint256.MustFromBig(AdminFee),
			*uint256.MustFromBig(totalSupply),
			InitialATime.Int64(),
			FutureATime.Int64(),
		)

		// Calculate DY
		var amountOut, amount uint256.Int
		amount.SetFromDecimal("1000")

		if err := stablePool.GetDy(0, 1, &amount, &amountOut); err != nil {
			results <- PoolResult{poolAddr, fmt.Errorf("failed to get DY: %v", err)}
			return
		}
		fmt.Printf("DY calculated for pool %s: %s\n", poolAddr, amountOut.String())

		// Fetch DY from contract with retry
		contractDyResult, err := contractCaller.CallWithRetry(poolCtx, func(c tracer.ContractCallerInterface, _ *ethclient.Client) (interface{}, error) {
			stableContract, ok := c.(*stableContract.ContractCaller)
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

	// // Pool addresses
	// stablePoolAddr := "0xDC24316b9AE028F1497c275EB9192a3Ea0f67022"

	// // Connect to Ethereum node
	// client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	// if err != nil {
	// 	t.Fatalf("Failed to connect to Ethereum client: %v", err)
	// }

	// // Create a new contract instance
	// contract, _ := NewContract(client, common.HexToAddress(stablePoolAddr))

	// // Fetch the A value
	// APrecise, err := contract.APrecise(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get APrecise: %v", err)
	// }

	// LpToken, err := contract.LpToken(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get LPToken: %v", err)
	// }

	// lpToken, err := token.NewToken(LpToken)
	// if err != nil {
	// 	t.Fatalf("Failed to create token: %v", err)
	// }

	// A, err := contract.A(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get A: %v", err)
	// }

	// InitialA, err := contract.InitialA(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get InitialA: %v", err)
	// }
	// FutureA, err := contract.FutureA(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get FutureA: %v", err)
	// }
	// InitialATime, err := contract.InitialATime(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get InitialATime: %v", err)
	// }
	// FutureATime, err := contract.FutureATime(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get FutureATime: %v", err)
	// }

	// xi, err := contract.Balances(nil, big.NewInt(int64(0)))
	// if err != nil {
	// 	t.Fatalf("Failed to get balance: %v", err)
	// }
	// xj, err := contract.Balances(nil, big.NewInt(int64(1)))
	// if err != nil {
	// 	t.Fatalf("Failed to get balance: %v", err)
	// }
	// xp := []uint256.Int{*uint256.MustFromBig(xi), *uint256.MustFromBig(xj)}

	// x, err := contract.Coins(nil, big.NewInt(int64(0)))
	// if err != nil {
	// 	t.Fatalf("Failed to get coin: %v", err)
	// }
	// y, err := contract.Coins(nil, big.NewInt(int64(1)))
	// if err != nil {
	// 	t.Fatalf("Failed to get coin: %v", err)
	// }
	// token1, err := token.NewToken(x)
	// if err != nil {
	// 	t.Fatalf("Failed to create token: %v", err)
	// }
	// token2, err := token.NewToken(y)
	// if err != nil {
	// 	t.Fatalf("Failed to create token: %v", err)
	// }
	// tokens := []token.Token{*token1, *token2}

	// fee, err := contract.Fee(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get fee: %v", err)
	// }

	// adminFee, err := contract.AdminFee(nil)
	// if err != nil {
	// 	t.Fatalf("Failed to get adminFee: %v", err)
	// }

	// pool := NewPool(
	// 	stablePoolAddr,
	// 	DexType,
	// 	nil,
	// 	xp,
	// 	tokens,
	// 	*uint256.MustFromBig(new(big.Int).Div(APrecise, A)),
	// 	*uint256.MustFromBig(InitialA),
	// 	*uint256.MustFromBig(FutureA),
	// 	*uint256.MustFromBig(fee),
	// 	*uint256.MustFromBig(adminFee),
	// 	lpToken.TotalSupply,
	// 	InitialATime.Int64(),
	// 	FutureATime.Int64(),
	// )

	// var amountOut, amount uint256.Int
	// amount.SetFromDecimal("1000000000000000000")

	// if err := pool.GetDy(0, 1, &amount, &amountOut); err != nil {
	// 	t.Fatalf("Failed to get DY: %v", err)
	// }
	// fmt.Println("DY calculated:", amountOut.String())

	// contractDy, err := contract.GetDy(nil, big.NewInt(int64(0)), big.NewInt(int64(1)), amount.ToBig())
	// if err != nil {
	// 	t.Fatalf("Failed to get DY: %v", err)
	// }
	// fmt.Println("DY from contract:", contractDy.String())

	// contractDyInt := new(uint256.Int)
	// contractDyInt.SetFromBig(contractDy)
	// if amountOut.Cmp(contractDyInt) != 0 {
	// 	t.Errorf("Values are not equal: got %s, want %s", amountOut.String(), contractDyInt.String())
	// }
}
