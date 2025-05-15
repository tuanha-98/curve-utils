package number

import "github.com/holiman/uint256"

func NewUint256(s string) (res *uint256.Int) {
	res = new(uint256.Int)
	_ = res.SetFromDecimal(s)
	return
}
