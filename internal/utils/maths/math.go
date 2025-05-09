package maths

import (
	"github.com/holiman/uint256"
)

type Uint256 = uint256.Int

// func Sqrt(x *big.Int) *big.Int {
// 	if x.Cmp(big.NewInt(0)) < 0 {
// 		return nil
// 	}
// 	if x.Cmp(big.NewInt(2)) < 0 {
// 		return x
// 	}

// 	left := big.NewInt(1)
// 	right := new(big.Int).Set(x)
// 	result := big.NewInt(1)

// 	for left.Cmp(right) <= 0 {
// 		mid := new(big.Int).Add(left, right)
// 		mid.Div(mid, big.NewInt(2))

// 		square := new(big.Int).Mul(mid, mid)

// 		switch square.Cmp(x) {
// 		case 0:
// 			return mid
// 		case -1:
// 			result.Set(mid)
// 			left.Add(mid, big.NewInt(1))
// 		case 1:
// 			right.Sub(mid, big.NewInt(1))
// 		}
// 	}

// 	return result
// }

// func ToUint256(x *big.Int) *big.Int {
// 	max_uint256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
// 	return new(big.Int).And(x, max_uint256)
// }

// func ToInt256(x *big.Int) *big.Int {
// 	max_uint256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
// 	max_int256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 255), big.NewInt(1))
// 	uint_value := new(big.Int).And(x, max_uint256)
// 	if uint_value.Cmp(max_int256) > 0 {
// 		return new(big.Int).Sub(uint_value, new(big.Int).Lsh(big.NewInt(1), 256))
// 	} else {
// 		return uint_value
// 	}
// }

// func SnekmateLog2(x *big.Int) *big.Int {
// 	value := new(big.Int).Set(x)
// 	result := big.NewInt(0)

// 	if new(big.Int).Rsh(value, 128).Cmp(big.NewInt(0)) != 0 {
// 		value.Rsh(value, 128)
// 		result.SetInt64(128)
// 	}
// 	if new(big.Int).Rsh(value, 64).Cmp(big.NewInt(0)) != 0 {
// 		value.Rsh(value, 64)
// 		result.Add(result, big.NewInt(64))
// 	}
// 	if new(big.Int).Rsh(value, 32).Cmp(big.NewInt(0)) != 0 {
// 		value.Rsh(value, 32)
// 		result.Add(result, big.NewInt(32))
// 	}
// 	if new(big.Int).Rsh(value, 16).Cmp(big.NewInt(0)) != 0 {
// 		value.Rsh(value, 16)
// 		result.Add(result, big.NewInt(16))
// 	}
// 	if new(big.Int).Rsh(value, 8).Cmp(big.NewInt(0)) != 0 {
// 		value.Rsh(value, 8)
// 		result.Add(result, big.NewInt(8))
// 	}
// 	if new(big.Int).Rsh(value, 4).Cmp(big.NewInt(0)) != 0 {
// 		value.Rsh(value, 4)
// 		result.Add(result, big.NewInt(4))
// 	}
// 	if new(big.Int).Rsh(value, 2).Cmp(big.NewInt(0)) != 0 {
// 		value.Rsh(value, 2)
// 		result.Add(result, big.NewInt(2))
// 	}
// 	if new(big.Int).Rsh(value, 1).Cmp(big.NewInt(0)) != 0 {
// 		result.Add(result, big.NewInt(1))
// 	}

// 	return result
// }

// func Cbrt(x *big.Int) *big.Int {
// 	xx := new(big.Int)
// 	thresh1 := new(big.Int).Mul(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1e18))
// 	thresh2 := new(big.Int).Lsh(big.NewInt(1), 256)

// 	if x.Cmp(thresh1) >= 0 {
// 		xx.Set(x)
// 	} else if x.Cmp(thresh2) >= 0 {
// 		xx.Mul(x, big.NewInt(Pow(big.NewInt(10), big.NewInt(18))))
// 	} else {
// 		xx.Mul(x, big.NewInt(Pow(big.NewInt(10), big.NewInt(36))))
// 	}

// 	log2x := SnekmateLog2(xx)
// 	remainder := new(big.Int).Mod(log2x, big.NewInt(3))

// 	divisor := new(big.Int).Exp(big.NewInt(1000), remainder, nil)
// 	multiplier := new(big.Int).Exp(big.NewInt(1260), remainder, nil)
// 	base := new(big.Int).Exp(big.NewInt(2), new(big.Int).Div(log2x, big.NewInt(3)), nil)

// 	a := new(big.Int).Mul(base, multiplier)
// 	a.Div(a, divisor)

// 	for i := 0; i < 7; i++ {
// 		squared := new(big.Int).Mul(a, a)
// 		quotient := new(big.Int).Div(xx, squared)
// 		sum := new(big.Int).Mul(big.NewInt(2), a)
// 		sum.Add(sum, quotient)
// 		a.Div(sum, big.NewInt(3))
// 	}

// 	if x.Cmp(thresh1) >= 0 {
// 		a.Mul(a, big.NewInt(1e12))
// 	} else if x.Cmp(thresh2) >= 0 {
// 		a.Mul(a, big.NewInt(1e6))
// 	}

// 	return a
// }
