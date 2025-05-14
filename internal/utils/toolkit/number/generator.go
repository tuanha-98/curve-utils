package number

import (
	"github.com/holiman/uint256"
)

func TenPow(decimals uint8) *uint256.Int {
	return new(uint256.Int).Exp(
		Number_10,
		uint256.NewInt(uint64(decimals)),
	)
}
