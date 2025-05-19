package stablemetang

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	stableContract "github.com/tuanha-98/curve-utils/contract/stableswap"
	stableMetaNGContract "github.com/tuanha-98/curve-utils/contract/stableswap-meta-ng"

	stableng "github.com/tuanha-98/curve-utils/internal/entities/pool/stable-ng"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

func NewContract(client *ethclient.Client, poolAddress common.Address) (*stableMetaNGContract.ContractCaller, error) {
	pool, err := stableMetaNGContract.NewContractCaller(poolAddress, client)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func NewContractStableNG(client *ethclient.Client, poolAddress common.Address) (*stableContract.ContractCaller, error) {
	pool, err := stableContract.NewContractCaller(poolAddress, client)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func TestGetDYNgPool(t *testing.T) {
	// Pool addresses
	stableNgPoolAddr := "0xC83b79C07ECE44b8b99fFa0E235C00aDd9124f9E"

	// Connect to Ethereum node
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Create a new contract instance
	contract, _ := NewContract(client, common.HexToAddress(stableNgPoolAddr))

	baseAddr, err := contract.BASEPOOL(nil)
	if err != nil {
		t.Fatalf("Failed to get base address: %v", err)
	}

	baseContract, _ := NewContractStableNG(client, baseAddr)

	// Fetch the A value
	APrecise, err := contract.APrecise(nil)
	if err != nil {
		t.Fatalf("Failed to get APrecise: %v", err)
	}

	A, err := contract.A(nil)
	if err != nil {
		t.Fatalf("Failed to get A: %v", err)
	}

	InitialA, err := contract.InitialA(nil)
	if err != nil {
		t.Fatalf("Failed to get InitialA: %v", err)
	}
	FutureA, err := contract.FutureA(nil)
	if err != nil {
		t.Fatalf("Failed to get FutureA: %v", err)
	}
	InitialATime, err := contract.InitialATime(nil)
	if err != nil {
		t.Fatalf("Failed to get InitialATime: %v", err)
	}
	FutureATime, err := contract.FutureATime(nil)
	if err != nil {
		t.Fatalf("Failed to get FutureATime: %v", err)
	}

	xi, err := contract.Balances(nil, big.NewInt(int64(0)))
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}
	xj, err := contract.Balances(nil, big.NewInt(int64(1)))
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}
	xp := []uint256.Int{*uint256.MustFromBig(xi), *uint256.MustFromBig(xj)}

	x, err := contract.Coins(nil, big.NewInt(int64(0)))
	if err != nil {
		t.Fatalf("Failed to get coin: %v", err)
	}
	y, err := contract.Coins(nil, big.NewInt(int64(1)))
	if err != nil {
		t.Fatalf("Failed to get coin: %v", err)
	}
	token1, err := token.NewToken(x)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}
	token2, err := token.NewToken(y)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}
	tokens := []token.Token{*token1, *token2}

	fee, err := contract.Fee(nil)
	if err != nil {
		t.Fatalf("Failed to get fee: %v", err)
	}

	adminFee, err := contract.AdminFee(nil)
	if err != nil {
		t.Fatalf("Failed to get adminFee: %v", err)
	}

	rates, err := contract.StoredRates(nil)
	if err != nil {
		t.Fatalf("Failed to get rates: %v", err)
	}

	_rates := []uint256.Int{*uint256.MustFromBig(rates[0]), *uint256.MustFromBig(rates[1])}

	offPegFeeMultiplier, err := contract.OffpegFeeMultiplier(nil)
	if err != nil {
		t.Fatalf("Failed to get offPegFeeMultiplier: %v", err)
	}

	totalSupply, err := contract.TotalSupply(nil)
	if err != nil {
		t.Fatalf("Failed to get total supply: %v", err)
	}

	baseInitialA, err := baseContract.InitialA(nil)
	if err != nil {
		t.Fatalf("Failed to get base InitialA: %v", err)
	}
	baseFutureA, err := baseContract.FutureA(nil)
	if err != nil {
		t.Fatalf("Failed to get base FutureA: %v", err)
	}
	baseInitialATime, err := baseContract.InitialATime(nil)
	if err != nil {
		t.Fatalf("Failed to get base InitialATime: %v", err)
	}
	baseFutureATime, err := baseContract.FutureATime(nil)
	if err != nil {
		t.Fatalf("Failed to get base FutureATime: %v", err)
	}
	baseXi, err := baseContract.Balances(nil, big.NewInt(int64(0)))
	if err != nil {
		t.Fatalf("Failed to get base balance: %v", err)
	}
	baseXj, err := baseContract.Balances(nil, big.NewInt(int64(1)))
	if err != nil {
		t.Fatalf("Failed to get base balance: %v", err)
	}
	baseXk, err := baseContract.Balances(nil, big.NewInt(int64(2)))
	if err != nil {
		t.Fatalf("Failed to get base balance: %v", err)
	}
	baseXp := []uint256.Int{*uint256.MustFromBig(baseXi), *uint256.MustFromBig(baseXj), *uint256.MustFromBig(baseXk)}
	baseX, err := baseContract.Coins(nil, big.NewInt(int64(0)))
	if err != nil {
		t.Fatalf("Failed to get base coin: %v", err)
	}
	baseY, err := baseContract.Coins(nil, big.NewInt(int64(1)))
	if err != nil {
		t.Fatalf("Failed to get base coin: %v", err)
	}
	baseZ, err := baseContract.Coins(nil, big.NewInt(int64(2)))
	if err != nil {
		t.Fatalf("Failed to get base coin: %v", err)
	}
	baseToken1, err := token.NewToken(baseX)
	if err != nil {
		t.Fatalf("Failed to create base token: %v", err)
	}
	baseToken2, err := token.NewToken(baseY)
	if err != nil {
		t.Fatalf("Failed to create base token: %v", err)
	}
	baseToken3, err := token.NewToken(baseZ)
	if err != nil {
		t.Fatalf("Failed to create base token: %v", err)
	}
	baseTokens := []token.Token{*baseToken1, *baseToken2, *baseToken3}

	baseFee, err := baseContract.Fee(nil)
	if err != nil {
		t.Fatalf("Failed to get base fee: %v", err)
	}
	baseAdminFee, err := baseContract.AdminFee(nil)
	if err != nil {
		t.Fatalf("Failed to get base adminFee: %v", err)
	}

	pool := NewPool(
		stableNgPoolAddr, baseAddr.String(),
		DexType,
		stableng.DexType,
		xp,
		baseXp,
		tokens,
		baseTokens,
		*uint256.MustFromBig(new(big.Int).Div(APrecise, A)),
		*uint256.MustFromBig(offPegFeeMultiplier),
		*uint256.MustFromBig(InitialA),
		*uint256.MustFromBig(FutureA),
		*uint256.MustFromBig(fee),
		*uint256.MustFromBig(adminFee),
		*uint256.MustFromBig(totalSupply),
		*uint256.MustFromBig(big.NewInt(1)),
		*number.Zero,
		*uint256.MustFromBig(baseInitialA),
		*uint256.MustFromBig(baseFutureA),
		*uint256.MustFromBig(baseFee),
		*uint256.MustFromBig(baseAdminFee),
		token2.TotalSupply,
		_rates,
		nil,
		InitialATime.Int64(),
		FutureATime.Int64(),
		baseInitialATime.Int64(),
		baseFutureATime.Int64(),
	)

	var amountOut, amount uint256.Int
	amount.SetFromDecimal("1000000000000000000")

	if err := pool.GetDyUnderlying(0, 2, &amount, &amountOut); err != nil {
		t.Fatalf("Failed to get DY: %v", err)
	}
	fmt.Println("DY calculated:", amountOut.String())

	contractDy, err := contract.GetDyUnderlying(nil, big.NewInt(int64(0)), big.NewInt(int64(2)), amount.ToBig())
	if err != nil {
		t.Fatalf("Failed to get DY: %v", err)
	}
	fmt.Println("DY from contract:", contractDy.String())

	contractDyInt := new(uint256.Int)
	contractDyInt.SetFromBig(contractDy)
	if amountOut.Cmp(contractDyInt) != 0 {
		t.Errorf("Values are not equal: got %s, want %s", amountOut.String(), contractDyInt.String())
	}
}
