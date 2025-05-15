package i256

import (
	"errors"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/maths/int256"
)

// cast the bytes directly without checking for negative (-1 -> 2^256-1)
func UnsafeToUInt256(value *int256.Int) *uint256.Int {
	var res uint256.Int
	UnsafeCastToUInt256(value, &res)
	return &res
}

// cast the bytes directly without checking for negative (-1 -> 2^256-1)
func UnsafeCastToUInt256(value *int256.Int, result *uint256.Int) {
	var ba [32]byte
	value.WriteToArray32(&ba)
	result.SetBytes32(ba[:])
}

// safely convert to u256 (negative numbers will panic)
func SafeConvertToUInt256(value *int256.Int) *uint256.Int {
	var res uint256.Int
	return safeConvertToUInt256(value, &res)
}
func safeConvertToUInt256(value *int256.Int, result *uint256.Int) *uint256.Int {
	if value.Sign() < 0 {
		panic(ErrNegative)
	}
	var ba [32]byte
	value.WriteToArray32(&ba)
	result.SetBytes32(ba[:])
	return result
}

// safely convert uint256 to int256
func SafeToInt256(value *uint256.Int) *int256.Int {
	var res int256.Int
	if err := SafeConvertToInt256(value, &res); err != nil {
		panic(err)
	}
	return &res
}

// safely convert uint256 to int256
func SafeConvertToInt256(value *uint256.Int, result *int256.Int) error {
	// if value (interpreted as a two's complement signed number) is negative -> it must be larger than max int256
	if value.Sign() < 0 {
		return errors.New("overflow u256 to i256")
	}
	var ba [32]byte
	value.WriteToArray32(&ba)
	result.SetBytes32(ba[:])
	return nil
}
