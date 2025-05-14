package number

import (
	"errors"

	"github.com/holiman/uint256"
)

var (
	Zero      = uint256.NewInt(0)
	Number_1  = uint256.NewInt(1)
	Number_2  = uint256.NewInt(2)
	Number_3  = uint256.NewInt(3)
	Number_4  = uint256.NewInt(4)
	Number_5  = uint256.NewInt(5)
	Number_6  = uint256.NewInt(6)
	Number_7  = uint256.NewInt(7)
	Number_8  = uint256.NewInt(8)
	Number_9  = uint256.NewInt(9)
	Number_10 = uint256.NewInt(10)

	Number_18 = uint256.NewInt(18)

	Number_10000 = uint256.NewInt(10000)

	Number_1e18 = new(uint256.Int).Exp(Number_10, Number_18)

	MaxU256 = uint256.MustFromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
)

var (
	ErrOverflow  = errors.New("uint256 overflow")
	ErrUnderflow = errors.New("uint256 underflow")
	ErrDivByZero = errors.New("divide by zero")
)
