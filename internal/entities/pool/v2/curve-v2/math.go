package curvev2

import (
	"sort"
	"time"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

var NowFunc = time.Now

func sortArray(x []uint256.Int) []uint256.Int {
	sort.Slice(x, func(i, j int) bool {
		return x[i].Cmp(&x[j]) > 0
	})
	return x
}

func newton_y(ann, gamma *uint256.Int, x []uint256.Int, D *uint256.Int, i int,
	// output
	y *uint256.Int,
) error {
	var NumTokens = len(x)
	var NumTokensU256 = uint256.NewInt(uint64(NumTokens))

	if NumTokens > 2 {
		if ann.Cmp(number.SubUint64(MinATri, 1)) <= 0 || ann.Cmp(number.AddUint64(MaxATri, 1)) >= 0 {
			return ErrUnsafeA
		}
		if gamma.Cmp(number.SubUint64(MinGamma, 1)) <= 0 || gamma.Cmp(number.AddUint64(MaxGammaTri, 1)) >= 0 {
			return ErrUnsafeGamma
		}
		for k := 0; k < NumTokens; k += 1 {
			if k == i {
				continue
			}
			var frac = number.Div(
				number.Mul(&x[k], number.Number_1e18),
				D,
			)
			if frac.Cmp(number.SubUint64(number.TenPow(16), 1)) <= 0 || frac.Cmp(number.AddUint64(number.TenPow(20), 1)) >= 0 {
				return ErrUnsafeXi
			}
		}
	} else {
		if ann.Cmp(number.SubUint64(MinATwo, 1)) <= 0 || ann.Cmp(number.AddUint64(MaxATwo, 1)) >= 0 {
			return ErrUnsafeA
		}
		if gamma.Cmp(number.SubUint64(MinGamma, 1)) <= 0 || gamma.Cmp(number.AddUint64(MaxGammaTwo, 1)) >= 0 {
			return ErrUnsafeGamma
		}
	}
	if D.Cmp(number.SubUint64(number.TenPow(17), 1)) <= 0 || D.Cmp(number.AddUint64(number.Mul(number.TenPow(15), number.Number_1e18), 1)) >= 0 {
		return ErrUnsafeD
	}

	y.Div(D, NumTokensU256)
	var K0i = number.TenPow(18)
	var Si uint256.Int

	var xSorted = make([]uint256.Int, NumTokens)
	for j := 0; j < NumTokens; j++ {
		xSorted[j].Set(&x[j])
	}
	xSorted[i].Clear()
	xSorted = sortArray(xSorted)

	var convergenceLimit = number.Div(&xSorted[0], number.TenPow(14))
	var temp = number.Div(D, number.TenPow(14))

	if temp.Cmp(convergenceLimit) > 0 {
		convergenceLimit.Set(temp)
	}

	if convergenceLimit.CmpUint64(100) < 0 {
		convergenceLimit.SetUint64(100)
	}

	for j := 2; j < NumTokens+1; j++ {
		var _x = xSorted[NumTokens-j]
		y.Set(
			number.Div(
				number.Mul(y, D),
				number.Mul(&_x, NumTokensU256),
			),
		)
		Si.Set(number.Add(&Si, &_x))
	}

	for j := 0; j < NumTokens-1; j++ {
		K0i.Set(
			number.Div(
				number.Mul(
					number.Mul(
						K0i,
						&xSorted[j],
					),
					NumTokensU256,
				),
				D,
			),
		)
	}

	// if y use the above formula will lead 1 wei wrong -> lead wrong newton_y result
	if NumTokens < 3 {
		y.Div(
			new(uint256.Int).Exp(
				D,
				number.Number_2,
			),
			number.SafeMul(
				&Si,
				new(uint256.Int).Exp(
					NumTokensU256,
					number.Number_2,
				),
			),
		)
	}

	var yPrev, K0, S, _g1k0, mul1, yfprime uint256.Int
	for j := 0; j < MaxLoopLimit; j++ {
		yPrev.Set(y)
		K0.Div(
			number.Mul(
				number.Mul(K0i, y),
				NumTokensU256,
			),
			D,
		)
		S.Add(&Si, y)

		_g1k0.Add(gamma, number.Number_1e18)
		if _g1k0.Cmp(&K0) > 0 {
			number.SafeAddZ(number.SafeSub(&_g1k0, &K0), number.Number_1, &_g1k0)
		} else {
			number.SafeAddZ(number.SafeSub(&K0, &_g1k0), number.Number_1, &_g1k0)
		}

		mul1.Div(
			number.SafeMul(
				number.Div(
					number.SafeMul(
						number.Div(
							number.Mul(
								number.Number_1e18,
								D,
							),
							gamma,
						),
						&_g1k0,
					),
					gamma,
				),
				number.SafeMul(&_g1k0, AMultiplier),
			),
			ann,
		)

		var mul2 = number.SafeAdd(
			number.Div(
				number.SafeMul(
					number.Mul(
						number.Number_2,
						number.Number_1e18,
					),
					&K0,
				),
				&_g1k0,
			),
			number.Number_1e18,
		)

		number.SafeAddZ(
			number.SafeAdd(
				number.SafeMul(
					number.Number_1e18,
					y,
				),
				number.SafeMul(
					&S,
					mul2,
				),
			),
			&mul1,
			&yfprime,
		)

		var _dyfprime = number.SafeMul(D, mul2)

		if yfprime.Cmp(_dyfprime) < 0 {
			y.Div(&yPrev, number.Number_2)
			continue
		} else {
			number.SafeSubZ(&yfprime, _dyfprime, &yfprime)
		}

		if y.IsZero() {
			return ErrDenominatorZero
		}

		var fprime = number.Div(&yfprime, y)

		if fprime.IsZero() {
			return ErrDenominatorZero
		}

		var yMinus = number.Div(&mul1, fprime)
		var yPlus = number.SafeAdd(
			number.Div(
				number.SafeAdd(
					&yfprime,
					number.Mul(number.Number_1e18, D),
				),
				fprime,
			),
			number.Div(
				number.SafeMul(yMinus, number.Number_1e18),
				&K0,
			),
		)
		number.SafeAddZ(yMinus, number.Div(number.SafeMul(number.Number_1e18, &S), fprime), yMinus)
		if yPlus.Cmp(yMinus) < 0 {
			y.Set(number.Div(&yPrev, number.Number_2))
		} else {
			number.SafeSubZ(yPlus, yMinus, y)
		}
		var diff uint256.Int
		if y.Cmp(&yPrev) > 0 {
			diff.Sub(y, &yPrev)
		} else {
			diff.Sub(&yPrev, y)
		}
		var t = number.Div(y, number.TenPow(14))
		if convergenceLimit.Cmp(t) > 0 {
			t.Set(convergenceLimit)
		}
		if diff.Cmp(t) < 0 {
			var frac = number.Div(number.Mul(y, number.Number_1e18), D)
			if frac.Cmp(number.TenPow(16)) < 0 || frac.Cmp(number.TenPow(20)) > 0 {
				return ErrUnsafeY
			}
			return nil
		}
	}
	return ErrDidNotConverge
}

func reductionCoefficient(x []uint256.Int, feeGamma *uint256.Int, K *uint256.Int) error {
	var NumTokens = len(x)
	var NumTokensU256 = uint256.NewInt(uint64(NumTokens))

	var S uint256.Int

	K.Set(number.TenPow(18))

	for _, xi := range x {
		S.Add(&S, &xi)
	}

	for _, xi := range x {
		K.Div(
			number.SafeMul(
				number.SafeMul(
					K,
					NumTokensU256,
				),
				&xi,
			),
			&S,
		)
	}

	K.Div(
		number.SafeMul(
			feeGamma,
			number.Number_1e18,
		),
		number.SafeSub(
			number.SafeAdd(feeGamma, number.Number_1e18),
			K,
		),
	)
	return nil
}

func (p *PoolSimulator) _A_gamma() (*uint256.Int, *uint256.Int) {
	var A, gamma uint256.Int
	p._A_gamma_inplace(&A, &gamma)
	return &A, &gamma
}

func (p *PoolSimulator) _A_gamma_inplace(A, gamma *uint256.Int) {
	var t1 = p.Extra.FutureAGammaTime
	var AGamma1 = p.Extra.FutureAGamma
	gamma.Set(new(uint256.Int).And(AGamma1, PriceMask))
	A.Set(new(uint256.Int).Rsh(AGamma1, 128))
	var now = NowFunc().Unix()
	if now < t1 {
		var AGamma0 = p.Extra.InitialAGamma
		var gamma0 = new(uint256.Int).And(AGamma0, PriceMask)
		var A0 = new(uint256.Int).Rsh(AGamma0, 128)
		var t0 = p.Extra.InitialAGammaTime
		t1 -= t0
		t0 = now - t0
		var t2 = t1 - t0
		A.Div(
			number.Add(
				number.Mul(A0, uint256.NewInt(uint64(t2))),
				number.Mul(A, uint256.NewInt(uint64(t0))),
			),
			uint256.NewInt(uint64(t1)),
		)

		gamma.Div(
			number.Add(
				number.Mul(gamma0, uint256.NewInt(uint64(t2))),
				number.Mul(gamma, uint256.NewInt(uint64(t0))),
			),
			uint256.NewInt(uint64(t1)),
		)
	}
}
