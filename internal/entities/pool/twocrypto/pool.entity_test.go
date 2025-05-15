package twocrypto

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	twoCrytoContract "github.com/tuanha-98/curve-utils/contract/twocryptoswap"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

func NewContract(client *ethclient.Client, poolAddress common.Address) (*twoCrytoContract.ContractCaller, error) {
	pool, err := twoCrytoContract.NewContractCaller(poolAddress, client)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func TestGetDYTwoCryptoPool(t *testing.T) {
	// Pool addresses
	cryptoPoolAddr := "0xB576491F1E6e5E62f1d8F26062Ee822B40B0E0d4"

	// Connect to Ethereum node
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Create a new contract instance
	contract, _ := NewContract(client, common.HexToAddress(cryptoPoolAddr))

	D, err := contract.D(nil)
	if err != nil {
		t.Fatalf("Failed to get D: %v", err)
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

	precisions := []uint256.Int{*new(uint256.Int).Exp(number.Number_10, uint256.NewInt(uint64(18-token1.Decimals))), *new(uint256.Int).Exp(number.Number_10, uint256.NewInt(uint64(18-token2.Decimals)))}

	_price_scale, err := contract.PriceScale(nil)
	if err != nil {
		t.Fatalf("Failed to get price_scale: %v", err)
	}
	price_scale := []uint256.Int{*uint256.MustFromBig(_price_scale)}

	midFee, err := contract.MidFee(nil)
	if err != nil {
		t.Fatalf("Failed to get midFee: %v", err)
	}

	outFee, err := contract.OutFee(nil)
	if err != nil {
		t.Fatalf("Failed to get outFee: %v", err)
	}

	feeGamma, err := contract.FeeGamma(nil)
	if err != nil {
		t.Fatalf("Failed to get feeGamma: %v", err)
	}

	initialAGamma, err := contract.InitialAGamma(nil)
	if err != nil {
		t.Fatalf("Failed to get initialA: %v", err)
	}

	initialAGammaTime, err := contract.InitialAGammaTime(nil)
	if err != nil {
		t.Fatalf("Failed to get initialGamma: %v", err)
	}

	futureAGamma, err := contract.FutureAGamma(nil)
	if err != nil {
		t.Fatalf("Failed to get futureA: %v", err)
	}

	futureAGammaTime, err := contract.FutureAGammaTime(nil)
	if err != nil {
		t.Fatalf("Failed to get futureAGammaTime: %v", err)
	}

	pool := NewPool(
		cryptoPoolAddr,
		"TwoCryptoSwap",
		xp,
		tokens,
		*uint256.MustFromBig(initialAGamma),
		*uint256.MustFromBig(futureAGamma),
		*uint256.MustFromBig(D),
		*uint256.MustFromBig(feeGamma),
		*uint256.MustFromBig(midFee),
		*uint256.MustFromBig(outFee),
		price_scale,
		precisions,
		initialAGammaTime.Int64(),
		futureAGammaTime.Int64(),
	)

	var amountOut, fee, amount uint256.Int
	amount.SetFromDecimal("1000000000000000000")

	if err := pool.GetDy(0, 1, &amount, &amountOut, &fee, xp); err != nil {
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
