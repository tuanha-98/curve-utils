package maths

import (
	"fmt"
	"math/big"

	"github.com/holiman/uint256"
)

type Uint256 = uint256.Int

var (
	MAX_UINT256 = new(Uint256).Sub(new(Uint256).Lsh(new(Uint256).SetUint64(1), 256), new(Uint256).SetUint64(1))
	MAX_INT256  = new(Uint256).Sub(new(Uint256).Lsh(new(Uint256).SetUint64(1), 255), new(Uint256).SetUint64(1))
)

func ToInt256(value *Uint256) *big.Int {
	val := new(Uint256).Set(value)
	uval := new(Uint256).And(val, MAX_UINT256)
	if uval.Cmp(MAX_INT256) <= 0 {
		return uval.ToBig()
	}
	return uval.Sub(uval, new(Uint256).Lsh(new(Uint256).SetUint64(1), 256)).ToBig()
}

func ToUint256(value *big.Int) *Uint256 {
	return new(Uint256).And(new(Uint256).SetBytes(value.Bytes()), MAX_UINT256)
}

func Sqrt(x *Uint256) *Uint256 {
	return new(Uint256).Sqrt(x)
}

func Cbrt(x *Uint256) *Uint256 {
	xx := new(Uint256).SetUint64(0)

	n, _ := new(big.Int).SetString("115792089237316195423570985008687907853269", 10)
	if x.Cmp(new(Uint256).SetBytes(new(big.Int).Mul(n, new(big.Int).Exp(new(big.Int).SetUint64(10), new(big.Int).SetUint64(18), nil)).Bytes())) >= 0 {
		xx.Set(x)
	} else if x.Cmp(new(Uint256).SetBytes(n.Bytes())) >= 0 {
		xx.Set(new(Uint256).Mul(x, new(Uint256).Exp(new(Uint256).SetUint64(10), new(Uint256).SetUint64(18))))
	} else {
		xx.Set(new(Uint256).Mul(x, new(Uint256).Exp(new(Uint256).SetUint64(10), new(Uint256).SetUint64(36))))
	}

	log2x := ToInt256(snekmateLog2(xx, false))
	remainder := new(Uint256).Mod(ToUint256(log2x), new(Uint256).SetUint64(3))

	numerator := new(Uint256).Mul(new(Uint256).Mod(new(Uint256).Exp(new(Uint256).SetUint64(2), new(Uint256).Div(ToUint256(log2x), new(Uint256).SetUint64(3))), new(Uint256).Exp(new(Uint256).SetUint64(2), new(Uint256).SetUint64(256))), new(Uint256).Mod(new(Uint256).Exp(new(Uint256).SetUint64(1260), remainder), new(Uint256).Exp(new(Uint256).SetUint64(2), new(Uint256).SetUint64(256))))
	denominator := new(Uint256).Mod(new(Uint256).Exp(new(Uint256).SetUint64(1000), remainder), new(Uint256).Exp(new(Uint256).SetUint64(2), new(Uint256).SetUint64(256)))
	a := new(Uint256).Div(numerator, denominator)

	for i := 0; i < 7; i++ {
		aSquared := new(Uint256).Mul(a, a)
		xDividedByASquared := new(Uint256).Div(xx, aSquared)
		twoTimesA := new(Uint256).Mul(new(Uint256).SetUint64(2), a)
		twoTimesA.Add(twoTimesA, xDividedByASquared)
		a.Div(twoTimesA, new(Uint256).SetUint64(3))
	}

	if x.Cmp(new(Uint256).SetBytes(new(big.Int).Mul(n, new(big.Int).Exp(new(big.Int).SetUint64(10), new(big.Int).SetUint64(18), nil)).Bytes())) >= 0 {
		a = new(Uint256).Mul(a, new(Uint256).Exp(new(Uint256).SetUint64(10), new(Uint256).SetUint64(12)))
	} else if x.Cmp(new(Uint256).SetBytes(n.Bytes())) >= 0 {
		a = new(Uint256).Mul(a, new(Uint256).Exp(new(Uint256).SetUint64(10), new(Uint256).SetUint64(6)))
	}

	return a
}

func GeometricMean2Coins(x_unsorted []*Uint256, sort bool) (*Uint256, error) {
	x := make([]*Uint256, len(x_unsorted))
	copy(x, x_unsorted)
	if sort && x[0].Cmp(x[1]) < 0 {
		x[0], x[1] = x[1], x[0]
	}
	D := new(Uint256).Set(x[0])
	diff := new(Uint256).SetUint64(0)
	for i := 0; i < 255; i++ {
		D_prev := new(Uint256).Set(D)
		D = new(Uint256).Div(new(Uint256).Add(D, new(Uint256).Div(new(Uint256).Mul(x[0], x[1]), D)), new(Uint256).SetUint64(2))
		if D.Cmp(D_prev) > 0 {
			diff = new(Uint256).Sub(D, D_prev)
		} else {
			diff = new(Uint256).Sub(D_prev, D)
		}
		if diff.Cmp(new(Uint256).SetUint64(1)) <= 0 || new(Uint256).Mul(diff, new(Uint256).Exp(new(Uint256).SetUint64(8), new(Uint256).SetUint64(18))).Cmp(D) < 0 {
			return D, nil
		}
	}
	return nil, fmt.Errorf("dev: did not converge")
}

func GeometricMean3Coins(x []*Uint256) (*Uint256, error) {
	prod := new(Uint256).Div(new(Uint256).Mul(new(Uint256).Div(new(Uint256).Mul(x[0], x[1]), new(Uint256).Exp(new(Uint256).SetUint64(10), new(Uint256).SetUint64(18))), x[2]), new(Uint256).Exp(new(Uint256).SetUint64(10), new(Uint256).SetUint64(18)))
	if prod.IsZero() {
		return new(Uint256).SetUint64(0), nil
	}
	return Cbrt(prod), nil
}

func Max(x, y *Uint256) *Uint256 {
	if x.Cmp(y) > 0 {
		return x
	}
	return y
}

func Min(x, y *Uint256) *Uint256 {
	if x.Cmp(y) < 0 {
		return x
	}
	return y
}

func snekmateLog2(x *Uint256, roundup bool) *Uint256 {
	value := new(Uint256).Set(x)
	result := new(Uint256).SetUint64(0)

	if new(Uint256).Rsh(x, 128).Cmp(new(Uint256).SetUint64(0)) != 0 {
		value.Rsh(x, 128)
		result.SetUint64(128)
	}
	if new(Uint256).Rsh(value, 64).Cmp(new(Uint256).SetUint64(0)) != 0 {
		value.Rsh(value, 64)
		result.Add(result, new(Uint256).SetUint64(64))
	}
	if new(Uint256).Rsh(value, 32).Cmp(new(Uint256).SetUint64(0)) != 0 {
		value.Rsh(value, 32)
		result.Add(result, new(Uint256).SetUint64(32))
	}
	if new(Uint256).Rsh(value, 16).Cmp(new(Uint256).SetUint64(0)) != 0 {
		value.Rsh(value, 16)
		result.Add(result, new(Uint256).SetUint64(16))
	}
	if new(Uint256).Rsh(value, 8).Cmp(new(Uint256).SetUint64(0)) != 0 {
		value.Rsh(value, 8)
		result.Add(result, new(Uint256).SetUint64(8))
	}
	if new(Uint256).Rsh(value, 4).Cmp(new(Uint256).SetUint64(0)) != 0 {
		value.Rsh(value, 4)
		result.Add(result, new(Uint256).SetUint64(4))
	}
	if new(Uint256).Rsh(value, 2).Cmp(new(Uint256).SetUint64(0)) != 0 {
		value.Rsh(value, 2)
		result.Add(result, new(Uint256).SetUint64(2))
	}
	if new(Uint256).Rsh(value, 1).Cmp(new(Uint256).SetUint64(0)) != 0 {
		result.Add(result, new(Uint256).SetUint64(1))
	}

	if roundup && new(Uint256).Lsh(new(Uint256).SetUint64(1), uint(result.ToBig().Int64())).Cmp(x) < 0 {
		result.Add(result, new(Uint256).SetUint64(1))
	}

	return result
}
