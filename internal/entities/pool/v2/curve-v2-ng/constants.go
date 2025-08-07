package curvev2ng

import (
	"errors"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/maths/int256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/i256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

const (
	DexType      = "curve-v2-ng"
	MaxLoopLimit = 256
)

var (
	PriceMask = uint256.MustFromHex("0xffffffffffffffffffffffffffffffff")

	I_1e48, _ = int256.FromDec("1000000000000000000000000000000000000000000000000")
	I_1e46, _ = int256.FromDec("10000000000000000000000000000000000000000000000")
	I_1e44, _ = int256.FromDec("100000000000000000000000000000000000000000000")
	I_1e42, _ = int256.FromDec("1000000000000000000000000000000000000000000")
	I_1e40, _ = int256.FromDec("10000000000000000000000000000000000000000")
	I_1e38, _ = int256.FromDec("100000000000000000000000000000000000000")
	I_1e36, _ = int256.FromDec("1000000000000000000000000000000000000")
	I_1e34, _ = int256.FromDec("10000000000000000000000000000000000")
	I_3e32, _ = int256.FromDec("300000000000000000000000000000000")
	I_1e32, _ = int256.FromDec("100000000000000000000000000000000")
	I_1e30, _ = int256.FromDec("1000000000000000000000000000000")
	I_1e28, _ = int256.FromDec("10000000000000000000000000000")
	I_1e26, _ = int256.FromDec("100000000000000000000000000")
	I_1e24, _ = int256.FromDec("1000000000000000000000000")
	I_1e22, _ = int256.FromDec("10000000000000000000000")
	I_1e20, _ = int256.FromDec("100000000000000000000")
	I_4e18, _ = int256.FromDec("4000000000000000000")
	I_2e18, _ = int256.FromDec("2000000000000000000")
	I_1e18, _ = int256.FromDec("1000000000000000000")
	I_1e16, _ = int256.FromDec("10000000000000000")
	I_4e14, _ = int256.FromDec("400000000000000")
	I_2e14, _ = int256.FromDec("200000000000000")
	I_1e14, _ = int256.FromDec("100000000000000")
	I_1e12, _ = int256.FromDec("1000000000000")
	I_1e10, _ = int256.FromDec("10000000000")
	I_4e8, _  = int256.FromDec("400000000")
	I_1e8, _  = int256.FromDec("100000000")
	I_1e6, _  = int256.FromDec("1000000")
	I_1e4, _  = int256.FromDec("10000")
	I_1e2, _  = int256.FromDec("100")
	I_27, _   = int256.FromDec("27")
	I_9, _    = int256.FromDec("9")

	TenPow36Div27 = i256.Div(I_1e36, I_27)
	TenPow36Div9  = i256.Div(I_1e36, I_9)
	I_27x27       = int256.NewInt(27 * 27)

	Precision   = number.Number_1e18
	AMultiplier = number.Number_10000

	MinGamma = number.TenPow(10)

	MaxGammaTwoSmall = number.Mul(number.Number_2, number.TenPow(16))
	MaxGammaTwo      = number.Mul(uint256.MustFromDecimal("199"), number.TenPow(15))
	MaxGammaTri      = number.Mul(number.Number_5, number.TenPow(16))

	MinATwo = number.Div(number.Mul(number.Number_4, AMultiplier), number.Number_10) // 4 = NCoins ** NCoins, NCoins = 2
	MaxATwo = number.Mul(number.Mul(number.Number_4, AMultiplier), number.Number_100000)
	MinATri = number.Div(number.Mul(number.Number_27, AMultiplier), number.Number_100) // 27 = NCoins ** NCoins, NCoins = 3
	MaxATri = number.Mul(number.Mul(number.Number_27, AMultiplier), number.Number_1000)

	MinD = number.TenPow(17)
	MaxD = number.Mul(number.TenPow(15), number.Number_1e18)

	MinFrac = number.TenPow(16)
	MaxFrac = number.TenPow(20)

	MinX0 = number.TenPow(9)
	MaxX1 = number.TenPow(33)
	MaxX  = number.Mul(
		number.Div(uint256.MustFromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"), number.TenPow(18)),
		number.Number_27,
	)

	CbrtConst1 = uint256.MustFromDecimal("115792089237316195423570985008687907853269000000000000000000")
	CbrtConst2 = uint256.MustFromDecimal("115792089237316195423570985008687907853269")
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
