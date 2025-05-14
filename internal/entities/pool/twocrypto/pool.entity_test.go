package twocrypto

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	crytoContract "github.com/tuanha-98/curve-utils/contract/cryptoswap"

	c "github.com/tuanha-98/curve-utils/internal/constants"
	m "github.com/tuanha-98/curve-utils/internal/utils/maths"
)

func NewContract(client *ethclient.Client, poolAddress common.Address) (*crytoContract.ContractCaller, error) {
	pool, err := crytoContract.NewContractCaller(poolAddress, client)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func TestGetDYCryptoPool(t *testing.T) {
	// Pool addresses
	cryptoPoolAddr := "0xB576491F1E6e5E62f1d8F26062Ee822B40B0E0d4"

	// Connect to Ethereum node
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Create a new contract instance
	contract, _ := NewContract(client, common.HexToAddress(cryptoPoolAddr))

	// Fetch the A value
	A, err := contract.A(nil)
	if err != nil {
		t.Fatalf("Failed to get A: %v", err)
	}

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
	xp := []*m.Uint256{new(m.Uint256).SetBytes(xi.Bytes()), new(m.Uint256).SetBytes(xj.Bytes())}

	x, err := contract.Coins(nil, big.NewInt(int64(0)))
	if err != nil {
		t.Fatalf("Failed to get coin: %v", err)
	}
	y, err := contract.Coins(nil, big.NewInt(int64(1)))
	if err != nil {
		t.Fatalf("Failed to get coin: %v", err)
	}
	coins := []string{x.String(), y.String()}

	precisions := []*m.Uint256{new(m.Uint256).SetUint64(1), new(m.Uint256).SetUint64(1)}

	price_scale, err := contract.PriceScale(nil)
	if err != nil {
		t.Fatalf("Failed to get price_scale: %v", err)
	}

	gamma, err := contract.Gamma(nil)
	if err != nil {
		t.Fatalf("Failed to get gamma: %v", err)
	}

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

	futureAGammaTime, err := contract.FutureAGammaTime(nil)
	if err != nil {
		t.Fatalf("Failed to get futureAGammaTime: %v", err)
	}

	pool := NewCryptoSwapPool(
		cryptoPoolAddr,
		"CryptoSwap",
		new(m.Uint256).SetBytes(A.Bytes()),
		new(m.Uint256).SetBytes(gamma.Bytes()),
		new(m.Uint256).SetBytes(D.Bytes()),
		new(m.Uint256).SetBytes(price_scale.Bytes()),
		new(m.Uint256).SetBytes(futureAGammaTime.Bytes()),
		new(m.Uint256).SetBytes(midFee.Bytes()),
		new(m.Uint256).SetBytes(outFee.Bytes()),
		new(m.Uint256).SetBytes(feeGamma.Bytes()),
		coins,
		xp,
		precisions,
	)

	value1, err := pool.GetDy(0, 1, c.Wei18)
	if err != nil {
		t.Fatalf("Failed to get Y: %v", err)
	}
	fmt.Println("DY calculated:", value1.String())

	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10)
	contractDy, err := contract.GetDy(nil, big.NewInt(int64(0)), big.NewInt(int64(1)), amount)
	if err != nil {
		t.Fatalf("Failed to get DY: %v", err)
	}
	value2 := new(m.Uint256).SetBytes(contractDy.Bytes())
	fmt.Println("DY from contract:", value2.String())

	if value1.Cmp(value2) != 0 {
		t.Errorf("Values are not equal: got %s, want %s", value1.String(), value2.String())
	}
}
