package twocrypto

import (
	"time"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

func sort(x []uint256.Int) {
	if x[0].Cmp(&x[1]) < 0 {
		tmp := number.Set(&x[0])
		x[0].Set(&x[1])
		x[1].Set(tmp)
	}
}

func geometric_mean(unsorted_x []uint256.Int) (*uint256.Int, error) {
	var x [NumTokens]uint256.Int
	for i := range unsorted_x {
		x[i].Set(&unsorted_x[i])
	}
	sort(x[:])
	var D = &x[0]
	var diff uint256.Int
	for i := 0; i < 255; i += 1 {
		var D_prev = D
		var tmp = U_1e18
		for _, _x := range x {
			tmp.Set(number.Div(number.Mul(tmp, &_x), D))
		}
		D.Set(
			number.Div(
				number.Mul(
					D,
					number.Add(
						number.Mul(
							number.Sub(
								NumTokensU256,
								number.Number_1,
							),
							U_1e18,
						),
						tmp,
					),
				),
				number.Mul(NumTokensU256, U_1e18),
			),
		)
		if D.Cmp(D_prev) > 0 {
			diff.Set(number.Sub(D, D_prev))
		} else {
			diff.Set(number.Sub(D_prev, D))
		}
		if diff.Cmp(number.Number_1) <= 0 || number.Mul(&diff, U_1e18).Cmp(D) < 0 {
			return D, nil
		}
	}
	return nil, ErrDidNotConverge
}

func newton_D(ANN *uint256.Int, gamma *uint256.Int, x_unsorted []uint256.Int) (*uint256.Int, error) {
	if ANN.Cmp(MinA) < 0 || ANN.Cmp(MaxA) > 0 {
		return nil, ErrUnsafeA
	}
	if gamma.Cmp(MinGamma) < 0 || gamma.Cmp(MaxGamma) > 0 {
		return nil, ErrUnsafeGamma
	}
	var x [NumTokens]uint256.Int
	for i := range x_unsorted {
		x[i].Set(&x_unsorted[i])
	}
	sort(x[:])
	if x[0].Cmp(MinX0) < 0 || x[0].Cmp(MaxX1) > 0 {
		return nil, ErrUnsafeX0
	}
	for i := 1; i < NumTokens; i += 1 {
		var frac = number.Div(
			number.Mul(
				&x[i],
				U_1e18,
			),
			&x[0],
		)
		if frac.Cmp(number.TenPow(11)) < 0 {
			return nil, ErrUnsafeXi
		}
	}

	var mean, err = geometric_mean(x[:])
	if err != nil {
		return nil, err
	}
	var D = number.Mul(NumTokensU256, mean)
	var S *uint256.Int

	for _, xI := range x {
		S.Set(number.Add(S, &xI))
	}

	for i := 0; i < 255; i += 1 {
		var D_prev = D
		var K0 = U_1e18
		for _, _x := range x {
			K0.Set(number.Div(number.Mul(number.Mul(K0, &_x), NumTokensU256), D))
		}
		var _g1k0 = number.Add(gamma, U_1e18)
		if _g1k0.Cmp(K0) > 0 {
			_g1k0.Set(number.Add(number.Sub(_g1k0, K0), number.Number_1))
		} else {
			_g1k0.Set(number.Add(number.Sub(K0, _g1k0), number.Number_1))
		}
		var mul1 = number.Div(
			number.Mul(
				number.Mul(
					number.Div(
						number.Mul(
							number.Div(
								number.Mul(U_1e18, D),
								gamma,
							),
							_g1k0,
						),
						gamma,
					),
					_g1k0,
				),
				AMultiplier,
			),
			ANN,
		)
		var mul2 = number.Div(
			number.Mul(
				number.Mul(
					number.Mul(
						number.Number_2,
						U_1e18,
					),
					NumTokensU256,
				),
				K0,
			),
			_g1k0,
		)
		var neg_fprime = number.Sub(
			number.Add(
				number.Add(S, number.Div(number.Mul(S, mul2), U_1e18)),
				number.Div(number.Mul(mul1, NumTokensU256), K0),
			),
			number.Div(number.Mul(mul2, D), U_1e18),
		)
		var D_plus = number.Div(number.Mul(D, number.Add(neg_fprime, S)), neg_fprime)
		var D_minus = number.Div(number.Mul(D, D), neg_fprime)
		if U_1e18.Cmp(K0) > 0 {
			D_minus.Set(
				number.Add(
					D_minus,
					number.Div(
						number.Mul(
							number.Div(
								number.Mul(D, number.Div(mul1, neg_fprime)),
								U_1e18,
							),
							number.Sub(U_1e18, K0),
						),
						K0,
					),
				),
			)
		} else {
			D_minus.Set(
				number.Sub(
					D_minus,
					number.Div(
						number.Mul(
							number.Div(
								number.Mul(D, number.Div(mul1, neg_fprime)),
								U_1e18,
							),
							number.Sub(K0, U_1e18),
						),
						K0,
					),
				),
			)
		}
		if D_plus.Cmp(D_minus) > 0 {
			D.Set(number.Sub(D_plus, D_minus))
		} else {
			D.Set(number.Div(number.Sub(D_minus, D_plus), number.Number_2))
		}
		var diff *uint256.Int
		if D.Cmp(D_prev) > 0 {
			diff.Set(number.Sub(D, D_prev))
		} else {
			diff.Set(number.Sub(D_prev, D))
		}
		temp := U_1e16
		if D.Cmp(temp) > 0 {
			temp.Set(D)
		}
		if number.Mul(diff, U_1e14).Cmp(temp) < 0 {
			for _, _x := range x {
				var frac = number.Div(number.Mul(&_x, U_1e18), D)
				if frac.Cmp(U_1e16) < 0 || frac.Cmp(U_1e20) > 0 {
					return nil, ErrUnsafeXi
				}
			}
			return D, nil
		}
	}
	return nil, ErrDidNotConverge
}

func newton_y(ann, gamma *uint256.Int, x []uint256.Int, D *uint256.Int, i int,
	// output
	y *uint256.Int,
) error {
	if ann.Cmp(MinA) < 0 || ann.Cmp(MaxA) > 0 {
		return ErrUnsafeA
	}
	if gamma.Cmp(MinGamma) < 0 || gamma.Cmp(MaxGamma) > 0 {
		return ErrUnsafeGamma
	}
	if D.Cmp(MinD) < 0 || D.Cmp(MaxD) > 0 {
		return ErrUnsafeD
	}
	x_j := &x[1-i]
	y.Div(number.Mul(D, D), number.Mul(x_j, number.Mul(NumTokensU256, NumTokensU256)))
	K0i := number.Div(number.Mul(number.Mul(U_1e18, NumTokensU256), x_j), D)
	if K0i.Cmp(number.Mul(U_1e16, NumTokensU256)) < 0 || K0i.Cmp(number.Mul(U_1e20, NumTokensU256)) > 0 {
		return ErrUnsafeXi
	}
	var convergenceLimit = number.Div(x_j, U_1e14)

	var temp = number.Div(D, U_1e14)

	if temp.Cmp(convergenceLimit) > 0 {
		convergenceLimit = temp
	}
	if convergenceLimit.CmpUint64(100) < 0 {
		convergenceLimit.SetUint64(100)
	}
	var yPrev, K0, S, _g1k0, mul1, yfprime uint256.Int
	De18 := number.SafeMul(D, U_1e18)

	for j := 0; j < 255; j += 1 {
		yPrev.Set(y)
		K0.Div(number.Mul(number.Mul(K0i, y), NumTokensU256), D)
		S.Add(x_j, y)

		_g1k0.Add(gamma, U_1e18)
		if _g1k0.Cmp(&K0) > 0 {
			number.SafeAddZ(number.SafeSub(&_g1k0, &K0), number.Number_1, &_g1k0)
		} else {
			number.SafeAddZ(number.SafeSub(&K0, &_g1k0), number.Number_1, &_g1k0)
		}

		// mul1 = 10**18 * D / gamma * _g1k0 / gamma * _g1k0 * A_MULTIPLIER / ANN
		mul1.Div(
			number.SafeMul(
				number.Div(
					number.SafeMul(
						number.Div(De18, gamma),
						&_g1k0,
					), gamma,
				),
				number.SafeMul(&_g1k0, AMultiplier),
			), ann)

		// mul2 = 10**18 + (2 * 10**18) * K0 / _g1k0
		var mul2 = number.SafeAdd(
			U_1e18,
			number.Div(number.SafeMul(U_2e18, &K0), &_g1k0),
		)

		// yfprime = 10**18 * y + S * mul2 + mul1
		number.SafeAddZ(
			number.SafeAdd(number.SafeMul(U_1e18, y), number.SafeMul(&S, mul2)),
			&mul1, &yfprime)
		var _dyfprime = number.SafeMul(D, mul2)

		if yfprime.Cmp(_dyfprime) < 0 {
			y.Set(number.Div(&yPrev, number.Number_2))
			continue
		} else {
			number.SafeSubZ(&yfprime, _dyfprime, &yfprime)
		}

		var fprime = number.Div(&yfprime, y)

		var yMinus = number.Div(&mul1, fprime)
		var yPlus = number.SafeAdd(number.Div(
			number.SafeAdd(&yfprime, De18),
			fprime),
			number.Div(number.SafeMul(yMinus, U_1e18), &K0))
		number.SafeAddZ(yMinus, number.Div(number.SafeMul(U_1e18, &S), fprime), yMinus)
		if yPlus.Cmp(yMinus) < 0 {
			y.Div(&yPrev, number.Number_2)
		} else {
			number.SafeSubZ(yPlus, yMinus, y)
		}
		var diff uint256.Int
		if y.Cmp(&yPrev) > 0 {
			diff.Sub(y, &yPrev)
		} else {
			diff.Sub(&yPrev, y)
		}
		var t = number.Div(y, U_1e14)
		if convergenceLimit.Cmp(t) > 0 {
			t = convergenceLimit
		}
		if diff.Cmp(t) < 0 {
			var frac = number.Div(number.Mul(y, U_1e18), D)
			if frac.Cmp(U_1e16) < 0 || frac.Cmp(U_1e20) > 0 {
				return ErrUnsafeY
			}
			return nil
		}
	}
	return ErrDidNotConverge
}

func reductionCoefficient(x []uint256.Int, feeGamma *uint256.Int, K *uint256.Int) error {
	var S uint256.Int
	number.SafeAddZ(&x[0], &x[1], &S)
	if S.IsZero() {
		return ErrZero
	}

	K.Mul(U_1e18, number.Mul(NumTokensU256, NumTokensU256))
	K.Div(number.SafeMul(K, &x[0]), &S)
	K.Div(number.SafeMul(K, &x[1]), &S)

	K.Div(
		number.SafeMul(feeGamma, U_1e18),
		number.SafeSub(number.SafeAdd(feeGamma, U_1e18), K))
	return nil
}

func halfpow(power *uint256.Int) (*uint256.Int, error) {
	intpow := number.Div(power, U_1e18)
	otherpow := number.Sub(power, number.Mul(intpow, U_1e18))
	if intpow.CmpUint64(59) > 0 {
		return number.Zero, nil
	}
	result := number.Div(U_1e18, new(uint256.Int).Exp(number.Number_2, intpow))
	if otherpow.IsZero() {
		return result, nil
	}
	term := U_1e18
	x := number.Mul(number.Number_5, U_1e17)
	S := U_1e18
	var neg = false

	for i := 1; i < 256; i += 1 {
		K := number.Mul(uint256.NewInt(uint64(i)), U_1e18)
		c := number.Sub(K, U_1e18)
		if otherpow.Cmp(c) > 0 {
			c.Set(number.Sub(otherpow, c))
			neg = !neg
		} else {
			c.Set(number.Sub(c, otherpow))
		}
		term = number.Div(number.Mul(term, number.Div(number.Mul(c, x), U_1e18)), K)
		if neg {
			S.Set(number.Sub(S, term))
		} else {
			S.Set(number.Add(S, term))
		}
		if term.Cmp(U_1e10) < 0 {
			return number.Div(number.Mul(result, S), U_1e18), nil
		}
	}
	return nil, ErrDidNotConverge
}

func (p *Pool) _A_gamma() (*uint256.Int, *uint256.Int) {
	var A, gamma uint256.Int
	p._A_gamma_inplace(&A, &gamma)
	return &A, &gamma
}

func (p *Pool) _A_gamma_inplace(A, gamma *uint256.Int) {
	var t1 = p.Extra.FutureAGammaTime
	var AGamma1 = p.Extra.FutureAGamma
	gamma.Set(new(uint256.Int).And(AGamma1, PriceMask))
	A.Set(new(uint256.Int).Rsh(AGamma1, 128))
	var now = time.Now().Unix()
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
