package i256

import "github.com/tuanha-98/curve-utils/internal/utils/maths/int256"

var (
	Number_1  = int256.NewInt(1)
	Number_2  = int256.NewInt(2)
	Number_3  = int256.NewInt(3)
	Number_4  = int256.NewInt(4)
	Number_10 = int256.NewInt(10)
)

func Abs(x *int256.Int) *int256.Int {
	var res int256.Int
	return AbsZ(x, &res)
}

func AbsZ(x *int256.Int, result *int256.Int) *int256.Int {
	if x.Sign() >= 0 {
		result.Set(x)
	} else {
		result.Neg(x)
	}
	return result
}

func Mul(x, y *int256.Int) *int256.Int {
	var res int256.Int
	res.Mul(x, y)
	return &res
}

func Div(x, y *int256.Int) *int256.Int {
	var res int256.Int
	res.Quo(x, y)
	return &res
}

func Lsh(x *int256.Int, n uint) *int256.Int {
	var res int256.Int
	res.Lsh(x, n)
	return &res
}

func Rsh(x *int256.Int, n uint) *int256.Int {
	var res int256.Int
	res.Rsh(x, n)
	return &res
}

func Set(x *int256.Int) *int256.Int {
	var res int256.Int
	res.Set(x)
	return &res
}

func SetInt64(x int64) *int256.Int {
	var res int256.Int
	res.SetInt64(x)
	return &res
}

func MustFromDecimal(s string) *int256.Int {
	var res int256.Int
	if err := res.SetFromDec(s); err != nil {
		panic(err)
	}
	return &res
}

func Neg(x *int256.Int) *int256.Int {
	var res int256.Int
	res.Neg(x)
	return &res
}

//go:noinline
func _add(x, y, z *int256.Int) bool {
	_, overflow := z.AddOverflow(x, y)
	return overflow
}
func Add(x, y *int256.Int) *int256.Int {
	var res int256.Int
	_add(x, y, &res)
	return &res
}

//go:noinline
func _sub(x, y, z *int256.Int) bool {
	_, underflow := z.SubOverflow(x, y)
	return underflow
}
func Sub(x, y *int256.Int) *int256.Int {
	var res int256.Int
	_sub(x, y, &res)
	return &res
}
