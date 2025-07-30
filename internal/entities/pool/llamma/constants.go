package llamma

import (
	"errors"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

const (
	DexType = "curve-llamma"

	maxTicksUnit int64 = 50
	maxTicks     int64 = 50
	maxSkipTicks int64 = 1024
)

var (
	Number_1e36    = number.TenPow(36)
	tenPow18Minus1 = new(uint256.Int).Sub(number.Number_1e18, number.Number_1)
	tenPow18Div4   = new(uint256.Int).Div(number.Number_1e18, number.Number_4)
)

var (
	ErrMulDivOverflow      = errors.New("mul div overflow")
	ErrWrongIndex          = errors.New("wrong index")
	ErrZeroSwapAmount      = errors.New("zero swap amount")
	ErrWadExpOverflow      = errors.New("wad_exp overflow")
	ErrInsufficientBalance = errors.New("insufficient balance")
)
