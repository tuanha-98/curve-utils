package curvev2ng

import (
	"fmt"
	"time"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/maths/int256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/i256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

var NowFunc = time.Now

// export for testing
func SortArray(x []uint256.Int) []uint256.Int {
	return sortArray(x)
}

func Newton_D(ann, gamma *uint256.Int, x []uint256.Int, K0_prev *uint256.Int,
	// output
	D *uint256.Int,
) error {
	return newton_D(ann, gamma, x, K0_prev, D)
}

func GeometricMean(
	unsorted_x []uint256.Int,
) *uint256.Int {
	return geometric_mean(unsorted_x)
}

func sortArray(x []uint256.Int) []uint256.Int {
	var n = len(x)
	var ret = make([]uint256.Int, n)
	for i := 0; i < n; i++ {
		ret[i].Set(&x[i])
	}

	for i := 1; i < n; i++ {
		var key uint256.Int
		key.Set(&ret[i])
		var j = i - 1
		for j >= 0 && ret[j].Cmp(&key) < 0 {
			ret[j+1].Set(&ret[j])
			j--
		}
		ret[j+1].Set(&key)
	}
	return ret
}

func cbrt(x *uint256.Int) *uint256.Int {
	var res uint256.Int
	_cbrt(x, &res)
	return &res
}

func _cbrt(x *uint256.Int, a *uint256.Int) {
	var xx *uint256.Int
	if x.Cmp(CbrtConst1) >= 0 {
		xx = x
	} else if x.Cmp(CbrtConst2) >= 0 {
		xx = number.Mul(x, number.TenPow(18))
	} else {
		xx = number.Mul(x, number.TenPow(36))
	}

	// log2x: int256 = convert(self._snekmate_log_2(xx, False), int256)
	_log2x := i256.SafeToInt256(snekmate_log_2(xx, false))
	log2x := i256.SafeConvertToUInt256(_log2x)

	// # When we divide log2x by 3, the remainder is (log2x % 3).
	// # So if we just multiply 2**(log2x/3) and discard the remainder to calculate our
	// # guess, the newton method will need more iterations to converge to a solution,
	// # since it is missing that precision. It's a few more calculations now to do less
	// # calculations later:
	// # pow = log2(x) // 3
	// # remainder = log2(x) % 3
	// # initial_guess = 2 ** pow * cbrt(2) ** remainder
	// # substituting -> 2 = 1.26 ≈ 1260 / 1000, we get:
	// #
	// # initial_guess = 2 ** pow * 1260 ** remainder // 1000 ** remainder

	remainder := new(uint256.Int).Mod(log2x, number.Number_3)
	a.Div(
		number.Mul(
			pow_mod256(number.Number_2, number.Div(log2x, number.Number_3)), //# <- pow
			pow_mod256(uint256.NewInt(1260), remainder),
		),
		pow_mod256(uint256.NewInt(1000), remainder),
	)

	// # Because we chose good initial values for cube roots, 7 newton raphson iterations
	// # are just about sufficient. 6 iterations would result in non-convergences, and 8
	// # would be one too many iterations. Without initial values, the iteration count
	// # can go up to 20 or greater. The iterations are unrolled. This reduces gas costs
	// # but takes up more bytecode:
	a.Div(number.Add(number.Mul(number.Number_2, a), number.Div(xx, number.Mul(a, a))), number.Number_3)
	a.Div(number.Add(number.Mul(number.Number_2, a), number.Div(xx, number.Mul(a, a))), number.Number_3)
	a.Div(number.Add(number.Mul(number.Number_2, a), number.Div(xx, number.Mul(a, a))), number.Number_3)
	a.Div(number.Add(number.Mul(number.Number_2, a), number.Div(xx, number.Mul(a, a))), number.Number_3)
	a.Div(number.Add(number.Mul(number.Number_2, a), number.Div(xx, number.Mul(a, a))), number.Number_3)
	a.Div(number.Add(number.Mul(number.Number_2, a), number.Div(xx, number.Mul(a, a))), number.Number_3)
	a.Div(number.Add(number.Mul(number.Number_2, a), number.Div(xx, number.Mul(a, a))), number.Number_3)

	if x.Cmp(CbrtConst1) >= 0 {
		a.Mul(a, number.TenPow(12))
	} else if x.Cmp(CbrtConst2) >= 0 {
		a.Mul(a, number.TenPow(6))
	}
}

func geometric_mean(_x []uint256.Int) *uint256.Int {
	var result uint256.Int
	_geometric_mean(_x, &result)
	return &result
}

// calculates a geometric mean for three numbers.
func _geometric_mean(_x []uint256.Int, result *uint256.Int) {
	var NumTokens = len(_x)

	if NumTokens > 2 {
		prod := number.Div(
			number.SafeMul(number.Div(number.SafeMul(&_x[0], &_x[1]), number.Number_1e18), &_x[2]),
			number.Number_1e18,
		)

		if prod.IsZero() {
			result.Clear()
			return
		}

		_cbrt(prod, result)
	} else {
		result.Sqrt(result.Mul(&_x[0], &_x[1]))
	}
}

func get_y(
	_ann, _gamma *uint256.Int, x []uint256.Int, _D *uint256.Int, i int,
	// output
	y *uint256.Int,
) error {
	var NumTokens = len(x)
	var NumTokensU256 = uint256.NewInt(uint64(NumTokens))

	switch NumTokens {
	case 2:
		if _ann.Cmp(MinATwo) < 0 || _ann.Cmp(MaxATwo) > 0 {
			return ErrUnsafeA
		}

		if _gamma.Cmp(MinGamma) < 0 || _gamma.Cmp(MaxGammaTwo) > 0 {
			return ErrUnsafeGamma
		}

		if _D.Cmp(MinD) < 0 || _D.Cmp(MaxD) > 0 {
			return ErrUnsafeD
		}

		lim_mul := number.TenPow(20)

		if _gamma.Cmp(MaxGammaTwoSmall) > 0 {
			lim_mul = number.Div(
				number.Mul(
					lim_mul,
					MaxGammaTwoSmall,
				),
				_gamma,
			)
		}
		lim_mul_signed := i256.SafeToInt256(lim_mul)

		ann := i256.SafeToInt256(_ann)
		gamma := i256.SafeToInt256(_gamma)
		D := i256.SafeToInt256(_D)
		x_j := i256.SafeToInt256(&x[1-i])
		gamma2 := i256.Mul(gamma, gamma)

		K0_i := i256.SafeToInt256(number.Div(number.Mul(number.Mul(number.Number_1e18, NumTokensU256), &x[1-i]), _D))
		if K0_i.Cmp(i256.Div(I_1e36, lim_mul_signed)) < 0 || K0_i.Cmp(lim_mul_signed) > 0 {
			return ErrUnsafeXi
		}

		ann_gamma2 := i256.Mul(ann, gamma2)

		a := i256.Set(I_1e32)
		b := i256.Sub(
			i256.Sub(
				i256.Div(
					i256.Div(
						i256.Mul(D, ann_gamma2),
						I_4e8,
					),
					x_j,
				),
				I_3e32,
			),
			i256.Mul(gamma, I_2e14),
		)
		c := i256.Sub(
			i256.Add(
				i256.Add(
					i256.Add(
						I_3e32,
						i256.Mul(gamma, I_4e14),
					),
					i256.Div(gamma2, I_1e4),
				),
				i256.Div(i256.Mul(i256.Div(i256.Mul(i256.Number_4, ann_gamma2), I_4e8), x_j), D),
			),
			i256.Div(i256.Mul(i256.Number_4, ann_gamma2), I_4e8),
		)
		tmp := i256.Add(I_1e18, gamma)
		d := i256.Neg(i256.Div(i256.Mul(tmp, tmp), I_1e4))

		delta0 := i256.Sub(
			i256.Div(
				i256.Mul(
					i256.Mul(i256.Number_3, a),
					c,
				),
				b,
			),
			b,
		)
		delta1 := i256.Sub(
			i256.Sub(
				i256.Div(
					i256.Mul(
						i256.Mul(I_9, a),
						c,
					),
					b,
				),
				i256.Mul(i256.Number_2, b),
			),
			i256.Div(i256.Mul(i256.Div(i256.Mul(I_27, i256.Mul(a, a)), b), d), b),
		)

		var divider int256.Int
		divider.SetUint64(1)
		threshold := i256.Abs(delta0)
		if threshold.Cmp(i256.Abs(delta1)) > 0 {
			threshold = i256.Abs(delta1)
		}
		if threshold.Cmp(a) > 0 {
			threshold = a
		}
		if threshold.Cmp(I_1e48) > 0 {
			divider.Set(I_1e30)
		} else if threshold.Cmp(I_1e46) > 0 {
			divider.Set(I_1e28)
		} else if threshold.Cmp(I_1e44) > 0 {
			divider.Set(I_1e26)
		} else if threshold.Cmp(I_1e42) > 0 {
			divider.Set(I_1e24)
		} else if threshold.Cmp(I_1e40) > 0 {
			divider.Set(I_1e22)
		} else if threshold.Cmp(I_1e38) > 0 {
			divider.Set(I_1e20)
		} else if threshold.Cmp(I_1e36) > 0 {
			divider.Set(I_1e18)
		} else if threshold.Cmp(I_1e34) > 0 {
			divider.Set(I_1e16)
		} else if threshold.Cmp(I_1e32) > 0 {
			divider.Set(I_1e14)
		} else if threshold.Cmp(I_1e30) > 0 {
			divider.Set(I_1e12)
		} else if threshold.Cmp(I_1e28) > 0 {
			divider.Set(I_1e10)
		} else if threshold.Cmp(I_1e26) > 0 {
			divider.Set(I_1e8)
		} else if threshold.Cmp(I_1e24) > 0 {
			divider.Set(I_1e6)
		} else if threshold.Cmp(I_1e20) > 0 {
			divider.Set(I_1e2)
		}
		a = i256.Div(a, &divider)
		b = i256.Div(b, &divider)
		c = i256.Div(c, &divider)
		d = i256.Div(d, &divider)
		// # delta0 = 3*a*c/b - b
		_3ac := i256.Mul(i256.Mul(i256.Number_3, a), c)
		delta0 = i256.Sub(i256.Div(_3ac, b), b)

		// # delta1 = 9*a*c/b - 2*b - 27*a**2/b*d/b
		delta1 = i256.Sub(
			i256.Sub(
				i256.Div(i256.Mul(i256.Number_3, _3ac), b),
				i256.Mul(i256.Number_2, b),
			),
			i256.Div(i256.Mul(
				i256.Div(
					i256.Mul(I_27, i256.Mul(a, a)),
					b),
				d),
				b),
		)

		// # delta1**2 + 4*delta0**2/b*delta0
		sqrt_arg := i256.Add(
			i256.Mul(delta1, delta1),
			i256.Mul(
				i256.Div(
					i256.Mul(i256.Number_4, i256.Mul(delta0, delta0)),
					b),
				delta0),
		)

		sqrt_val := new(int256.Int)
		if sqrt_arg.Sign() > 0 {
			sqrt_val.Sqrt(sqrt_arg)
		} else {
			return newton_y(_ann, _gamma, x, _D, i, lim_mul, y)
		}

		var b_cbrt *int256.Int
		if b.Sign() >= 0 {
			b_cbrt = i256.SafeToInt256(cbrt(i256.SafeConvertToUInt256(b)))
		} else {
			b_cbrt = i256.Neg(i256.SafeToInt256(cbrt(i256.SafeConvertToUInt256(i256.Neg(b)))))
		}

		var second_cbrt *int256.Int
		if delta1.Sign() > 0 {
			// # convert(self._cbrt(convert((delta1 + sqrt_val), uint256)/2), int256)
			second_cbrt = i256.SafeToInt256(
				cbrt(number.Div(
					i256.SafeConvertToUInt256(i256.Add(delta1, sqrt_val)),
					number.Number_2)))
		} else {
			second_cbrt = i256.Neg(i256.SafeToInt256(
				cbrt(number.Div(
					i256.SafeConvertToUInt256(i256.Sub(sqrt_val, delta1)),
					number.Number_2))),
			)
		}

		// # C1: int256 = b_cbrt**2/10**18*second_cbrt/10**18
		C1 := i256.Div(
			i256.Mul(i256.Div(i256.Mul(b_cbrt, b_cbrt), I_1e18), second_cbrt),
			I_1e18,
		)

		// # root: int256 = (10**18*C1 - 10**18*b - 10**18*b*delta0/C1)/(3*a), keep 2 safe ops here.
		root := i256.Div(
			i256.Sub(i256.Sub(
				i256.Mul(I_1e18, C1),
				i256.Mul(I_1e18, b)),
				i256.Mul(i256.Div(i256.Mul(I_1e18, b), C1), delta0)),
			i256.Mul(i256.Number_3, a),
		)

		// # y_out: uint256[2] =  [
		// #     convert(D**2/x_j*root/4/10**18, uint256),   # <--- y
		// #     convert(root, uint256)  # <----------------------- K0Prev
		// # ]
		y.Set(i256.SafeConvertToUInt256(i256.Div(i256.Mul(i256.Div(i256.Mul(D, D), x_j), root), I_4e18)))

		frac := number.Div(number.Mul(y, number.TenPow(18)), _D)
		// assert (frac >= unsafe_div(10**36 / N_COINS, lim_mul)) and (frac <= unsafe_div(lim_mul, N_COINS))  # dev: unsafe value for y
		if frac.Cmp(number.Div(number.Div(number.TenPow(36), NumTokensU256), lim_mul)) < 0 || frac.Cmp(number.Div(lim_mul,
			NumTokensU256)) > 0 {
			return ErrUnsafeY
		}

		return nil
	case 3:
		if _ann.Cmp(MinATri) < 0 || _ann.Cmp(MaxATri) > 0 {
			return ErrUnsafeA
		}

		if _gamma.Cmp(MinGamma) < 0 || _gamma.Cmp(MaxGammaTri) > 0 {
			return ErrUnsafeGamma
		}

		if _D.Cmp(MinD) < 0 || _D.Cmp(MaxD) > 0 {
			return ErrUnsafeGamma
		}

		for k := 0; k < NumTokens; k++ {
			if k == i {
				continue
			}
			frac := number.Div(number.Mul(&x[k], number.Number_1e18), _D)
			if frac.Cmp(MinFrac) < 0 || frac.Cmp(MaxFrac) > 0 {
				return fmt.Errorf("unsafe values x[%d] %s", i, frac.Dec())
			}
		}

		j := 0
		k := 0
		if i == 0 {
			j = 1
			k = 2
		} else if i == 1 {
			j = 0
			k = 2
		} else if i == 2 {
			j = 0
			k = 1
		}

		ann := i256.SafeToInt256(_ann)
		gamma := i256.SafeToInt256(_gamma)
		D := i256.SafeToInt256(_D)
		x_j := i256.SafeToInt256(&x[j])
		x_k := i256.SafeToInt256(&x[k])
		gamma2 := i256.Mul(gamma, gamma)
		AMultiplier_ := i256.SafeToInt256(AMultiplier)

		a := i256.Set(TenPow36Div27)
		b := i256.Sub(
			i256.Add(TenPow36Div9, i256.Div(i256.Mul(I_2e18, gamma), I_27)),
			i256.Div(
				i256.Div(
					i256.Div(
						i256.Mul(
							i256.Mul(
								i256.Div(i256.Mul(D, D), x_j),
								gamma2,
							),
							ann,
						),
						I_27x27,
					),
					AMultiplier_,
				),
				x_k,
			),
		)
		c := i256.Add(
			i256.Add(
				TenPow36Div9,
				i256.Div(i256.Mul(gamma, i256.Add(gamma, I_4e18)), I_27),
			),
			i256.Div(
				i256.Div(
					i256.Mul(
						i256.Div(
							i256.Mul(gamma2,
								i256.Sub(i256.Add(x_j, x_k), D)),
							D,
						),
						ann,
					),
					I_27,
				),
				AMultiplier_,
			),
		)
		tmp := i256.Add(I_1e18, gamma)
		d := i256.Div(i256.Mul(tmp, tmp), I_27)
		d0 := i256.Abs(
			i256.Sub(
				i256.Div(
					i256.Mul(
						i256.Mul(i256.Number_3, a),
						c,
					),
					b,
				),
				b,
			),
		)
		var divider int256.Int
		divider.SetUint64(1)
		if d0.Cmp(I_1e48) > 0 {
			divider.Set(I_1e30)
		} else if d0.Cmp(I_1e44) > 0 {
			divider.Set(I_1e26)
		} else if d0.Cmp(I_1e40) > 0 {
			divider.Set(I_1e22)
		} else if d0.Cmp(I_1e36) > 0 {
			divider.Set(I_1e18)
		} else if d0.Cmp(I_1e32) > 0 {
			divider.Set(I_1e14)
		} else if d0.Cmp(I_1e28) > 0 {
			divider.Set(I_1e10)
		} else if d0.Cmp(I_1e24) > 0 {
			divider.Set(I_1e6)
		} else if d0.Cmp(I_1e20) > 0 {
			divider.Set(I_1e2)
		}

		var additional_prec *int256.Int
		if i256.Abs(a).Cmp(i256.Abs(b)) > 0 {
			additional_prec = i256.Abs(i256.Div(a, b))
			a = i256.Div(i256.Mul(a, additional_prec), &divider)
			b = i256.Div(i256.Mul(b, additional_prec), &divider)
			c = i256.Div(i256.Mul(c, additional_prec), &divider)
			d = i256.Div(i256.Mul(d, additional_prec), &divider)
		} else {
			additional_prec = i256.Abs(i256.Div(b, a))
			a = i256.Div(i256.Div(a, additional_prec), &divider)
			b = i256.Div(i256.Div(b, additional_prec), &divider)
			c = i256.Div(i256.Div(c, additional_prec), &divider)
			d = i256.Div(i256.Div(d, additional_prec), &divider)
		}

		// # 3*a*c/b - b
		_3ac := i256.Mul(i256.Mul(i256.Number_3, a), c)
		delta0 := i256.Sub(i256.Div(_3ac, b), b)

		// # 9*a*c/b - 2*b - 27*a**2/b*d/b
		delta1 := i256.Sub(
			i256.Sub(
				i256.Div(i256.Mul(i256.Number_3, _3ac), b),
				i256.Mul(i256.Number_2, b),
			),
			i256.Div(i256.Mul(
				i256.Div(
					i256.Mul(I_27, i256.Mul(a, a)),
					b),
				d),
				b),
		)

		// # delta1**2 + 4*delta0**2/b*delta0
		sqrt_arg := i256.Add(
			i256.Mul(delta1, delta1),
			i256.Mul(
				i256.Div(
					i256.Mul(i256.Number_4, i256.Mul(delta0, delta0)),
					b),
				delta0),
		)

		sqrt_val := int256.NewInt(0)
		if sqrt_arg.Sign() > 0 {
			sqrt_val.Sqrt(sqrt_arg)
		} else {
			return newton_y(_ann, _gamma, x, _D, i, y, nil)
		}

		var b_cbrt *int256.Int
		if b.Sign() >= 0 {
			b_cbrt = i256.SafeToInt256(cbrt(i256.SafeConvertToUInt256(b)))
		} else {
			b_cbrt = i256.Neg(i256.SafeToInt256(cbrt(i256.SafeConvertToUInt256(i256.Neg(b)))))
		}

		var second_cbrt *int256.Int
		if delta1.Sign() > 0 {
			// # convert(self._cbrt(convert((delta1 + sqrt_val), uint256)/2), int256)
			second_cbrt = i256.SafeToInt256(
				cbrt(number.Div(
					i256.SafeConvertToUInt256(i256.Add(delta1, sqrt_val)),
					number.Number_2)))
		} else {
			second_cbrt = i256.Neg(i256.SafeToInt256(
				cbrt(number.Div(
					i256.SafeConvertToUInt256(new(int256.Int).Neg(i256.Sub(delta1, sqrt_val))),
					number.Number_2))),
			)
		}

		// # b_cbrt*b_cbrt/10**18*second_cbrt/10**18
		C1 := i256.Div(
			i256.Mul(i256.Div(i256.Mul(b_cbrt, b_cbrt), I_1e18), second_cbrt),
			I_1e18,
		)

		// # (b + b*delta0/C1 - C1)/3
		root_K0 := i256.Div(
			i256.Sub(
				i256.Add(b,
					i256.Div(i256.Mul(b, delta0), C1)),
				C1),
			i256.Number_3)

		// # D*D/27/x_k*D/x_j*root_K0/a
		root := i256.Div(
			i256.Mul(
				i256.Div(
					i256.Mul(
						i256.Div(
							i256.Div(
								i256.Mul(D, D),
								I_27),
							x_k),
						D),
					x_j),
				root_K0),
			a,
		)

		y.Set(i256.SafeConvertToUInt256(root))

		frac := number.Div(number.Mul(y, number.Number_1e18), _D)
		if frac.Cmp(MinFrac) < 0 || frac.Cmp(MaxFrac) > 0 {
			return ErrUnsafeY
		}

		return nil
	default:
		return fmt.Errorf("unsupported number of tokens: %d", NumTokens)
	}
}

func newton_D(ann, gamma *uint256.Int, x_unsorted []uint256.Int, K0_prev *uint256.Int,
	// output
	D *uint256.Int,
) error {
	var NumTokens = len(x_unsorted)
	var NumTokensU256 = uint256.NewInt(uint64(NumTokens))

	x := make([]uint256.Int, len(x_unsorted))
	for i := range x_unsorted {
		x[i].Set(&x_unsorted[i])
	}
	x = sortArray(x)

	if NumTokens > 2 {
		if x[0].IsZero() || x[0].Cmp(MaxX) >= 0 {
			return ErrUnsafeX0
		}
	} else {
		if ann.Cmp(number.SubUint64(MinATwo, 1)) <= 0 || ann.Cmp(number.AddUint64(MaxATwo, 1)) >= 0 {
			return ErrUnsafeA
		}
		if gamma.Cmp(number.SubUint64(MinGamma, 1)) <= 0 || gamma.Cmp(number.AddUint64(MaxGammaTwo, 1)) >= 0 {
			return ErrUnsafeGamma
		}
		if x[0].Cmp(MinX0) < 0 || x[0].Cmp(MaxX1) > 0 {
			return ErrUnsafeX0
		}
		if number.Div(number.Mul(&x[1], number.TenPow(18)), &x[0]).Cmp(number.TenPow(14)) < 0 {
			return ErrUnsafeXi
		}
	}

	var S uint256.Int
	D.Clear()
	for _, x_i := range x {
		S.Add(&S, &x_i)
	}

	if K0_prev.IsZero() {
		D.Mul(NumTokensU256, geometric_mean(x))
	} else {
		if NumTokens > 2 {
			if S.Cmp(number.TenPow(36)) > 0 {
				_cbrt(
					number.Mul(
						number.Div(
							number.Mul(number.Div(number.Mul(&x[0], &x[1]), number.TenPow(36)), &x[2]),
							K0_prev,
						),
						number.Mul(number.Number_27, number.TenPow(12)),
					),
					D,
				)
			} else if S.Cmp(number.TenPow(24)) > 0 {
				_cbrt(
					number.Mul(
						number.Div(
							number.Mul(number.Div(number.Mul(&x[0], &x[1]), number.TenPow(24)), &x[2]),
							K0_prev,
						),
						number.Mul(number.Number_27, number.TenPow(6)),
					),
					D,
				)
			} else {
				_cbrt(
					number.Mul(
						number.Div(
							number.Mul(number.Div(number.Mul(&x[0], &x[1]), number.TenPow(18)), &x[2]),
							K0_prev,
						),
						number.Number_27,
					),
					D,
				)
			}
		} else { // twocrypto
			D.Sqrt(
				number.Mul(
					number.Div(
						number.Mul(number.Mul(number.Number_4, &x[0]), &x[1]),
						K0_prev,
					),
					number.Number_1e18,
				),
			)

			if S.Cmp(D) < 0 {
				D.Set(&S)
			}
		}
	}

	var __g1k0 = number.Add(gamma, number.Number_1e18)
	var D_prev, K0, diff uint256.Int

	for i := 0; i < 255; i++ {
		var _g1k0, mul1, mul2, neg_fprime, D_plus, D_minus uint256.Int

		if D.Sign() <= 0 {
			return ErrUnsafeD
		}
		D_prev.Set(D)

		if NumTokens > 2 {
			K0.Div(
				number.Mul(
					number.Mul(
						number.Div(
							number.Mul(
								number.Mul(
									number.Div(
										number.Mul(
											number.Mul(number.Number_1e18, &x[0]),
											NumTokensU256,
										),
										D,
									),
									&x[1],
								),
								NumTokensU256,
							),
							D,
						),
						&x[2],
					),
					NumTokensU256,
				),
				D,
			)
		} else {
			K0.Div(
				number.Mul(
					number.Div(
						number.Mul(
							number.Mul(
								number.Number_1e18,
								number.Mul(NumTokensU256, NumTokensU256),
							),
							&x[0],
						),
						D,
					),
					&x[1],
				),
				D,
			)
		}

		_g1k0.Set(__g1k0)
		if _g1k0.Cmp(&K0) > 0 {
			_g1k0.AddUint64(number.Sub(&_g1k0, &K0), 1)
		} else {
			_g1k0.AddUint64(number.Sub(&K0, &_g1k0), 1)
		}

		mul1.Div(
			number.Mul(
				number.Mul(
					number.Div(
						number.Mul(
							number.Div(number.Mul(number.Number_1e18, D), gamma), &_g1k0,
						),
						gamma,
					),
					&_g1k0,
				),
				AMultiplier,
			),
			ann,
		)

		mul2.Div(
			number.Mul(
				number.Mul(
					number.SafeMul(
						number.Number_2,
						number.TenPow(18),
					),
					NumTokensU256,
				),
				&K0,
			),
			&_g1k0,
		)

		neg_fprime.Sub(
			number.Add(
				number.Add(&S, number.Div(number.Mul(&S, &mul2), number.Number_1e18)),
				number.SafeDiv(number.Mul(&mul1, NumTokensU256), &K0),
			),
			number.Div(number.Mul(&mul2, D), number.Number_1e18),
		)

		D_plus.Div(
			number.SafeMul(D, number.Add(&neg_fprime, &S)),
			&neg_fprime,
		)

		D_minus.Div(
			number.SafeMul(D, D),
			&neg_fprime,
		)

		if number.Number_1e18.Cmp(&K0) > 0 {
			D_minus.Set(
				number.SafeAdd(
					&D_minus,
					number.Div(
						number.Mul(
							number.Div(number.SafeMul(D, number.Div(&mul1, &neg_fprime)), number.Number_1e18),
							number.Sub(number.Number_1e18, &K0),
						),
						&K0,
					),
				),
			)
		} else {
			D_minus.Set(
				number.SafeSub(
					&D_minus,
					number.Div(
						number.Mul(
							number.Div(number.SafeMul(D, number.Div(&mul1, &neg_fprime)), number.Number_1e18),
							number.Sub(&K0, number.Number_1e18),
						),
						&K0,
					),
				),
			)
		}

		if D_plus.Cmp(&D_minus) > 0 {
			D.Sub(&D_plus, &D_minus)
		} else {
			D.Div(
				number.Sub(&D_minus, &D_plus),
				number.Number_2,
			)
		}

		if D.Cmp(&D_prev) > 0 {
			diff.Sub(D, &D_prev)
		} else {
			diff.Sub(&D_prev, D)
		}

		temp := number.TenPow(16)
		if D.Cmp(number.TenPow(16)) > 0 {
			temp.Set(D)
		}

		if number.Mul(&diff, number.TenPow(14)).Cmp(temp) < 0 {
			for i := range x {
				var frac = number.Div(number.Mul(&x[i], number.TenPow(18)), D)
				if frac.Cmp(number.Div(MinFrac, NumTokensU256)) < 0 || frac.Cmp(number.Div(MaxFrac, NumTokensU256)) > 0 {
					return ErrUnsafeXi
				}
			}
			return nil
		}
	}
	return ErrDDoesNotConverge
}

func pow_mod256(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Exp(x, y)
}

func snekmate_log_2(x *uint256.Int, roundup bool) *uint256.Int {
	var result uint256.Int
	_snekmate_log_2(x, roundup, &result)
	return &result
}

func _snekmate_log_2(x *uint256.Int, roundup bool, result *uint256.Int) {
	/*
		@notice An `internal` helper function that returns the log in base 2
				of `x`, following the selected rounding direction.
		@dev This implementation is derived from Snekmate, which is authored
				by pcaversaccio (Snekmate), distributed under the AGPL-3.0 license.
				https://github.com/pcaversaccio/snekmate
		@dev Note that it returns 0 if given 0. The implementation is
				inspired by OpenZeppelin's implementation here:
				https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/math/Math.sol.
		@param x The 32-byte variable.
		@param roundup The Boolean variable that specifies whether
				to round up or not. The default `False` is round down.
		@return uint256 The 32-byte calculation result.
	*/
	value := number.Set(x)
	result.Clear()

	// # The following lines cannot overflow because we have the well-known
	// # decay behaviour of `log_2(max_value(uint256)) < max_value(uint256)`.
	if !new(uint256.Int).Rsh(x, 128).IsZero() {
		value.Rsh(x, 128)
		result.SetUint64(128)
	}
	if !new(uint256.Int).Rsh(value, 64).IsZero() {
		value.Rsh(value, 64)
		result.Add(result, uint256.NewInt(64))
	}
	if !new(uint256.Int).Rsh(value, 32).IsZero() {
		value.Rsh(value, 32)
		result.Add(result, uint256.NewInt(32))
	}
	if !new(uint256.Int).Rsh(value, 16).IsZero() {
		value.Rsh(value, 16)
		result.Add(result, uint256.NewInt(16))
	}
	if !new(uint256.Int).Rsh(value, 8).IsZero() {
		value.Rsh(value, 8)
		result.Add(result, uint256.NewInt(8))
	}
	if !new(uint256.Int).Rsh(value, 4).IsZero() {
		value.Rsh(value, 4)
		result.Add(result, uint256.NewInt(4))
	}
	if !new(uint256.Int).Rsh(value, 2).IsZero() {
		value.Rsh(value, 2)
		result.Add(result, uint256.NewInt(2))
	}
	if !new(uint256.Int).Rsh(value, 1).IsZero() {
		result.Add(result, uint256.NewInt(1))
	}
	const1 := new(uint256.Int).Lsh(number.Number_1, uint(result.Uint64()))
	if roundup && const1.Cmp(x) < 0 {
		result.Add(result, uint256.NewInt(1))
	}
}

func newton_y(ann, gamma *uint256.Int, x []uint256.Int, D *uint256.Int, i int, lim_mul *uint256.Int,
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

		if K0i.Cmp(number.Div(number.TenPow(36), lim_mul)) < 0 || K0i.Cmp(lim_mul) > 0 {
			return ErrUnsafeXi
		}
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
