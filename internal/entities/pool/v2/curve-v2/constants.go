package curvev2

import (
	"errors"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

const (
	DexType      = "curve-v2"
	MaxLoopLimit = 256
)

var (
	PriceMask = uint256.MustFromHex("0xffffffffffffffffffffffffffffffff")

	Precision   = number.Number_1e18
	AMultiplier = number.Number_10000

	MinGamma    = number.TenPow(10)
	MaxGammaTwo = number.Mul(number.Number_2, number.TenPow(16))
	MaxGammaTri = number.Mul(number.Number_5, number.TenPow(16))

	MinATwo = number.Div(number.Mul(number.Number_4, AMultiplier), number.Number_10) // 4 = NCoins ** NCoins, NCoins = 2
	MaxATwo = number.Mul(number.Mul(number.Number_4, AMultiplier), number.Number_100000)
	MinATri = number.Div(number.Mul(number.Number_27, AMultiplier), number.Number_100) // 27 = NCoins ** NCoins, NCoins = 3
	MaxATri = number.Mul(number.Mul(number.Number_27, AMultiplier), number.Number_1000)
)

var (
	ErrInvalidReserve      = errors.New("invalid reserve")
	ErrInvalidNumToken     = errors.New("invalid number of token")
	ErrZero                = errors.New("zero")
	ErrLoss                = errors.New("loss")
	ErrDidNotConverge      = errors.New("did not converge")
	ErrDDoesNotConverge    = errors.New("d does not converge")
	ErrYDoesNotConverge    = errors.New("y does not converge")
	ErrWadExpOverflow      = errors.New("wad_exp overflow")
	ErrUnsafeY             = errors.New("unsafe value for y")
	ErrUnsafeA             = errors.New("unsafe values A")
	ErrUnsafeGamma         = errors.New("unsafe values gamma")
	ErrUnsafeD             = errors.New("unsafe values D")
	ErrUnsafeX0            = errors.New("unsafe values x[0]")
	ErrUnsafeXi            = errors.New("unsafe values x[i]")
	ErrCoinIndexOutOfRange = errors.New("coin index out of range")
	ErrSameCoin            = errors.New("coins is same")
	ErrExchange0Coins      = errors.New("do not exchange 0 coins")
	ErrTweakPrice          = errors.New("tweak price")
	ErrDenominatorZero     = errors.New("denominator is zero")
)
