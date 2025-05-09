package pool

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	stableContract "github.com/tuanha-98/curve-utils/contract/stableswap"
	stableNGContract "github.com/tuanha-98/curve-utils/contract/stableswap-ng"

	c "github.com/tuanha-98/curve-utils/internal/constants"
	m "github.com/tuanha-98/curve-utils/internal/utils/maths"
)

func NewContract(client *ethclient.Client, poolAddress common.Address) (*stableContract.ContractCaller, error) {
	pool, err := stableContract.NewContractCaller(poolAddress, client)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func NewNGContract(client *ethclient.Client, poolAddress common.Address) (*stableNGContract.ContractCaller, error) {
	pool, err := stableNGContract.NewContractCaller(poolAddress, client)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func TestGetDYNgPool(t *testing.T) {
	// Pool addresses
	stableNgPoolAddr := "0x4f493B7dE8aAC7d55F71853688b1F7C8F0243C85"
	// stablePoolAddr := "0xDC24316b9AE028F1497c275EB9192a3Ea0f67022"

	// Connect to Ethereum node
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Create a new contract instance
	contract, _ := NewNGContract(client, common.HexToAddress(stableNgPoolAddr))

	// Fetch the A value
	A, err := contract.A(nil)
	if err != nil {
		t.Fatalf("Failed to get A: %v", err)
	}

	A_precision, err := contract.APrecise(nil)
	if err != nil {
		t.Fatalf("Failed to get A_PRECISION: %v", err)
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

	fee, err := contract.Fee(nil)
	if err != nil {
		t.Fatalf("Failed to get fee: %v", err)
	}

	rates, err := contract.StoredRates(nil)
	if err != nil {
		t.Fatalf("Failed to get rates: %v", err)
	}

	rates_ := []*m.Uint256{new(m.Uint256).SetBytes(rates[0].Bytes()), new(m.Uint256).SetBytes(rates[1].Bytes())}

	offPegFeeMultiplier, err := contract.OffpegFeeMultiplier(nil)
	if err != nil {
		t.Fatalf("Failed to get offPegFeeMultiplier: %v", err)
	}

	pool := NewStableSwapPool(
		stableNgPoolAddr,
		"StableNG",
		new(m.Uint256).SetBytes(fee.Bytes()),
		new(m.Uint256).SetBytes(A.Bytes()),
		new(m.Uint256).SetBytes(A_precision.Bytes()),
		coins,
		xp,
		rates_,
		new(m.Uint256).SetBytes(offPegFeeMultiplier.Bytes()))

	value1, err := pool.GetDy(0, 1, c.Precision)
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

func TestGetDYPool(t *testing.T) {
	// Pool addresses
	stablePoolAddr := "0xDC24316b9AE028F1497c275EB9192a3Ea0f67022"

	// Connect to Ethereum node
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		t.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Create a new contract instance
	contract, _ := NewContract(client, common.HexToAddress(stablePoolAddr))

	// Fetch the A value
	A, err := contract.A(nil)
	if err != nil {
		t.Fatalf("Failed to get A: %v", err)
	}

	A_precision, err := contract.APrecise(nil)
	if err != nil {
		t.Fatalf("Failed to get A_PRECISION: %v", err)
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

	fee, err := contract.Fee(nil)
	if err != nil {
		t.Fatalf("Failed to get fee: %v", err)
	}

	pool := NewStableSwapPool(
		stablePoolAddr,
		"Stable",
		new(m.Uint256).SetBytes(fee.Bytes()),
		new(m.Uint256).SetBytes(A.Bytes()),
		new(m.Uint256).SetBytes(A_precision.Bytes()),
		coins,
		xp,
		nil,
		nil)

	value1, err := pool.GetDy(0, 1, c.Precision)
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
