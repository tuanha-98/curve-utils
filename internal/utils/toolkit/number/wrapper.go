package number

import (
	"math/big"

	"github.com/holiman/uint256"
)

/*
these are wrappers to make it a little bit easier to work with uint256
for example the code below checks if `x*y + (x-y)` is zero or not
it works fine, but will allocate many uint256 instances so maybe not suitable in hot path
	{
		var x, y *uint256.Int // calculated somewhere

		res := new(uint256.Int).Add(new(uint256.Int).Mul(x, y), new(uint256.Int).Sub(x, y))
		return res.IsZero()
	}

we can use temp var to remove those allocations:
	{
		var x, y *uint256.Int // calculated somewhere

		var res, tmp1, tmp2 uint256.Int
		tmp1.Mul(x, y)
		tmp2.Sub(x, y)
		res.Add(&tmp1, &tmp2)
		return res.IsZero()
	}
but the code will be hard to read, and it get worse as the expression has more nested layers

these wrappers try to keep the original style while still being efficient:
	{
		var x, y *uint256.Int // calculated somewhere

		res := number.Add(number.Mul(x, y), number.Sub(x, y))
		return res.IsZero()
	}
normally `number.Mul(x, y)` will allocate new uint256, but here the function `number.Mul`, `number.Sub`... are small/simple,
so the compiler will inline them, and the uint256 var will be placed in caller's stack instead of heap

note that this is only true if the results are used within the caller only (not escaped to heap)
the rule for that is a bit complex, so please run benchmark or Go escape analysis to be sure

https://words.filippo.io/efficient-go-apis-with-the-inliner/
*/

func AddUint64(x *uint256.Int, y uint64) *uint256.Int {
	var res uint256.Int
	res.AddUint64(x, y)
	return &res
}

func SubUint64(x *uint256.Int, y uint64) *uint256.Int {
	var res uint256.Int
	res.SubUint64(x, y)
	return &res
}

func Mul(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	res.Mul(x, y)
	return &res
}
func SafeMul(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	if _, overflow := res.MulOverflow(x, y); overflow {
		panic(ErrOverflow)
	}
	return &res
}
func SafeMulZ(x, y, z *uint256.Int) *uint256.Int {
	if _, overflow := z.MulOverflow(x, y); overflow {
		panic(ErrOverflow)
	}
	return z
}

func Div(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	res.Div(x, y)
	return &res
}
func SafeDiv(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	return SafeDivZ(x, y, &res)
}
func SafeDivZ(x, y, z *uint256.Int) *uint256.Int {
	if y.IsZero() {
		panic(ErrDivByZero)
	}
	return z.Div(x, y)
}

func SetUint64(x uint64) *uint256.Int {
	var res uint256.Int
	res.SetUint64(x)
	return &res
}

func Set(x *uint256.Int) *uint256.Int {
	var res uint256.Int
	res.Set(x)
	return &res
}

func SetFromBig(x *big.Int) *uint256.Int {
	if x == nil {
		return nil
	}
	var res uint256.Int
	res.SetFromBig(x)
	return &res
}

// Add and Sub are a bit different: uint256.Add and uint256.Sub are too simple, so they will be inlined into number.Add/number.Sub
// then number.Add/number.Sub will become complex and not inline-able
// so we need to force by go:noinline directive

//go:noinline
func _add(x, y, z *uint256.Int) bool {
	_, overflow := z.AddOverflow(x, y)
	return overflow
}
func Add(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	_add(x, y, &res)
	return &res
}
func SafeAdd(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	if _add(x, y, &res) {
		panic(ErrOverflow)
	}
	return &res
}
func SafeAddZ(x, y, z *uint256.Int) *uint256.Int {
	if _, overflow := z.AddOverflow(x, y); overflow {
		panic(ErrOverflow)
	}
	return z
}

//go:noinline
func _sub(x, y, z *uint256.Int) bool {
	_, underflow := z.SubOverflow(x, y)
	return underflow
}
func Sub(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	_sub(x, y, &res)
	return &res
}
func SafeSub(x, y *uint256.Int) *uint256.Int {
	var res uint256.Int
	if _sub(x, y, &res) {
		panic(ErrUnderflow)
	}
	return &res
}
func SafeSubZ(x, y, z *uint256.Int) *uint256.Int {
	if _, underflow := z.SubOverflow(x, y); underflow {
		panic(ErrUnderflow)
	}
	return z
}

func WithinDelta(x, y *uint256.Int, delta uint64) bool {
	var diff uint256.Int
	if x.Cmp(y) > 0 {
		diff.Sub(x, y)
	} else {
		diff.Sub(y, x)
	}
	if diff.CmpUint64(delta) <= 0 {
		return true
	}
	return false
}
