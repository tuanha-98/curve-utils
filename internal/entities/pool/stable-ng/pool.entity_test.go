package stableng

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	stableNGContract "github.com/tuanha-98/curve-utils/contract/stableswap-ng"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
)

func NewContract(client *ethclient.Client, poolAddress common.Address) (*stableNGContract.ContractCaller, error) {
	pool, err := stableNGContract.NewContractCaller(poolAddress, client)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func TestGetDYNgPool(t *testing.T) {
	// Pool addresses
	stableNgPoolAddr := "0x4f493B7dE8aAC7d55F71853688b1F7C8F0243C85"

	// Connect to Ethereum node
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Create a new contract instance
	contract, _ := NewContract(client, common.HexToAddress(stableNgPoolAddr))

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

	pool := NewPool(
		stableNgPoolAddr,
		"StableNG",
		xp,
		tokens,
		*uint256.MustFromBig(new(big.Int).Div(APrecise, A)),
		*uint256.MustFromBig(offPegFeeMultiplier),
		*uint256.MustFromBig(InitialA),
		*uint256.MustFromBig(FutureA),
		*uint256.MustFromBig(fee),
		*uint256.MustFromBig(adminFee),
		_rates,
		InitialATime.Int64(),
		FutureATime.Int64(),
	)

	var amountOut, amount uint256.Int
	amount.SetFromDecimal("1000000000000000000")

	if err := pool.GetDy(0, 1, &amount, &amountOut); err != nil {
		t.Fatalf("Failed to get DY: %v", err)
	}
	fmt.Println("DY calculated:", amountOut.String())

	contractDy, err := contract.GetDy(nil, big.NewInt(int64(0)), big.NewInt(int64(1)), amount.ToBig())
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
