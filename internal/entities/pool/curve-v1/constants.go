package curvev1

import (
	"errors"

	"github.com/holiman/uint256"
)

const (
	DexType       = "curve-v1"
	MaxLoopLimit  = 256
	MaxTokenCount = 8

	PoolTypePlain    = "plain"
	PoolTypeMeta     = "meta"
	PoolTypeAave     = "aave"
	PoolTypeOracle   = "oracle"
	PoolTypeCompound = "compound"
)

var (
	U_1e10 = uint256.MustFromDecimal("10000000000")
	U_1e18 = uint256.MustFromDecimal("1000000000000000000")

	Precision      = U_1e18
	FeeDenominator = U_1e10
)

var (
	ErrInvalidReserve               = errors.New("invalid reserve")
	ErrInvalidStoredRates           = errors.New("invalid stored rates")
	ErrInvalidNumToken              = errors.New("invalid number of token")
	ErrInvalidAValue                = errors.New("invalid A value")
	ErrZero                         = errors.New("zero")
	ErrBalancesMustMatchMultipliers = errors.New("balances must match multipliers")
	ErrDDoesNotConverge             = errors.New("d does not converge")
	ErrTokenFromEqualsTokenTo       = errors.New("can't compare token to itself")
	ErrTokenIndexesOutOfRange       = errors.New("token index out of range")
	ErrAmountOutNotConverge         = errors.New("approximation did not converge")
	ErrTokenNotFound                = errors.New("token not found")
	ErrWithdrawMoreThanAvailable    = errors.New("cannot withdraw more than available")
	ErrD1LowerThanD0                = errors.New("d1 <= d0")
	ErrDenominatorZero              = errors.New("denominator should not be 0")
	ErrReserveTooSmall              = errors.New("reserve too small")
	ErrInvalidFee                   = errors.New("invalid fee")
	ErrNewReserveInvalid            = errors.New("invalid new reserve")
)
