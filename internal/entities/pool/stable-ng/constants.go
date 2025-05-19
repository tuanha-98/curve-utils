package stableng

import (
	"errors"

	"github.com/holiman/uint256"
)

const (
	DexType       = "curve-stable-ng"
	MaxLoopLimit  = 255
	MaxTokenCount = 8
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
	ErrTokenIndexOutOfRange         = errors.New("token index out of range")
	ErrAmountOutNotConverge         = errors.New("approximation did not converge")
	ErrExecutionReverted            = errors.New("execution reverted")
)
