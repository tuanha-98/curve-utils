package twocrypto

import (
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
	for i := 0; i < MaxLoopLimit; i += 1 {
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

	for i := 0; i < MaxLoopLimit; i += 1 {
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
