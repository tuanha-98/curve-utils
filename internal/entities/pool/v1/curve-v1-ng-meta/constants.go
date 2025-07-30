package curvev1ngmeta

import (
	"errors"

	"github.com/holiman/uint256"
)

const (
	DexType       = "curve-v1-ng-meta"
	MaxLoopLimit  = 256
	MaxTokenCount = 8

	PoolTypePlain  = "plain"
	PoolTypeMeta   = "meta"
	PoolTypeOracle = "oracle"
)

var (
	Precision      = uint256.MustFromDecimal("1000000000000000000")
	FeeDenominator = uint256.MustFromDecimal("10000000000")
)

var (
	ErrInvalidBasePool              = errors.New("invalid base pool")
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

	ErrTokenToUnderlyingNotSupported = errors.New("not support exchange from base pool token to its underlying")
	ErrAllBasePoolTokens             = errors.New("base pool swap should be done at base pool")
	ErrAllMetaPoolTokens             = errors.New("meta pool swap should be done using GetDy")
)
