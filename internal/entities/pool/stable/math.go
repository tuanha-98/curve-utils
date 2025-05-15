package stable

import (
	"math"
	"time"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

func XpMem(
	rates []uint256.Int,
	balances []uint256.Int,
) []uint256.Int {
	// try to put `result` in caller's stack (this func will be inlined)
	var result [MaxTokenCount]uint256.Int
	count := xpMem_inplace(rates, balances, result[:])
	return result[:count]
}

func xpMem_inplace(
	rates []uint256.Int,
	balances []uint256.Int,
	xp []uint256.Int,
) int {
	numTokens := len(rates)
	for i := 0; i < numTokens; i++ {
		xp[i].Div(number.SafeMul(&rates[i], &balances[i]), Precision)
	}
	return numTokens
}

func (p *Pool) _A() *uint256.Int {
	var t1 = p.Extra.FutureATime
	var a1 = p.Extra.FutureA
	var now = time.Now().Unix()
	if t1 > now {
		var t0 = p.Extra.InitialATime
		var a0 = p.Extra.InitialA
		if a1.Cmp(a0) > 0 {
			return number.Add(
				a0,
				number.Div(
					number.Mul(
						number.Sub(a1, a0),
						number.SetUint64(uint64(now-t0)),
					),
					number.SetUint64(uint64(t1-t0)),
				),
			)
		} else {
			return number.Sub(
				a0,
				number.Div(
					number.Mul(
						number.Sub(a0, a1),
						number.SetUint64(uint64(now-t0)),
					),
					number.SetUint64(uint64(t1-t0)),
				),
			)
		}
	}
	return a1
}

// func (p *Pool) getD(xp []*m.Uint256, amp *m.Uint256) (*m.Uint256, error) {
// 	n := new(m.Uint256).SetUint64(uint64(len(xp)))
// 	ann := new(m.Uint256).Mul(amp, n)
// 	S := new(m.Uint256).SetUint64(0)
// 	for _, x := range xp {
// 		S.Add(S, x)
// 	}

// 	if S.IsZero() {
// 		return new(m.Uint256).SetUint64(0), nil
// 	}

// 	D := new(m.Uint256).Set(S)
// 	for i := 0; i < 255; i++ {
// 		D_P := new(m.Uint256).Set(D)
// 		for _, x := range xp {
// 			D_P = new(m.Uint256).Div(new(m.Uint256).Mul(D_P, D), new(m.Uint256).Mul(x, n))
// 		}
// 		D_prev := new(m.Uint256).Set(D)
// 		numerator := new(m.Uint256).Add(new(m.Uint256).Div(new(m.Uint256).Mul(ann, S), p.A_precise), new(m.Uint256).Mul(D_P, n))
// 		denominator := new(m.Uint256).Add(
// 			new(m.Uint256).Div(new(m.Uint256).Mul(new(m.Uint256).Sub(ann, p.A_precise), D), p.A_precise),
// 			new(m.Uint256).Mul(new(m.Uint256).Add(n, new(m.Uint256).SetUint64(1)), D_P))
// 		D = new(m.Uint256).Div(new(m.Uint256).Mul(numerator, D), denominator)

// 		if D.Cmp(D_prev) > 0 {
// 			if new(m.Uint256).Sub(D, D_prev).Cmp(new(m.Uint256).SetUint64(1)) <= 0 {
// 				return D, nil
// 			}
// 		} else {
// 			if new(m.Uint256).Sub(D_prev, D).Cmp(new(m.Uint256).SetUint64(1)) <= 0 {
// 				return D, nil
// 			}
// 		}
// 	}
// 	return nil, fmt.Errorf("getD did not converge")
// }

func (p *Pool) getD(
	xp []uint256.Int, a *uint256.Int,
	// output
	D *uint256.Int,
) error {
	var S uint256.Int
	S.Clear()
	for i := range xp {
		if xp[i].IsZero() {
			return ErrZero
		}
		S.Add(&S, &xp[i])
	}
	if S.IsZero() {
		D.Clear()
		return nil
	}

	var D_P, Ann, Ann_mul_S_div_APrec, Ann_sub_APrec, Dprev uint256.Int
	D.Set(&S)
	Ann.Mul(a, &p.NumTokensU256)
	Ann_mul_S_div_APrec.Div(number.SafeMul(&Ann, &S), p.Extra.APrecision)
	Ann_sub_APrec.Sub(&Ann, p.Extra.APrecision)

	numTokensPlus1 := uint256.NewInt(uint64(p.NumTokens + 1))
	numTokensPow := uint256.NewInt(uint64(math.Pow(float64(p.NumTokens), float64(p.NumTokens))))

	for i := 0; i < 255; i += 1 {
		D_P.Set(D)

		for j := range xp {
			D_P.Div(
				number.SafeMul(&D_P, D),
				&xp[j],
			)
		}

		D_P.Div(&D_P, numTokensPow)

		Dprev.Set(D)

		D.Div(
			number.SafeMul(
				number.SafeAdd(&Ann_mul_S_div_APrec, number.SafeMul(&D_P, &p.NumTokensU256)),
				D,
			),
			number.SafeAdd(
				number.Div(number.SafeMul(&Ann_sub_APrec, D), p.Extra.APrecision),
				number.SafeMul(&D_P, numTokensPlus1),
			),
		)

		if number.WithinDelta(D, &Dprev, 1) {
			return nil
		}
	}
	return ErrDDoesNotConverge
}

func (p *Pool) getY(
	i, j int, x *uint256.Int, xp []uint256.Int, dCached *uint256.Int,
	// output
	y *uint256.Int,
) error {
	if i == j {
		return ErrTokenFromEqualsTokenTo
	}
	if i >= p.NumTokens && j >= p.NumTokens {
		return ErrTokenIndexOutOfRange
	}

	var a = p._A()
	if a == nil {
		return ErrInvalidAValue
	}

	var d uint256.Int
	if dCached != nil {
		d.Set(dCached)
	} else {
		err := p.getD(xp, a, &d)
		if err != nil {
			return err
		}
	}

	var c = number.Set(&d)
	var Ann = number.Mul(a, &p.NumTokensU256)
	var _x, s uint256.Int
	s.Clear()

	for _i := 0; _i < p.NumTokens; _i += 1 {
		if _i == i {
			_x.Set(x)
		} else if _i != j {
			_x.Set(&xp[i])
		} else {
			continue
		}
		if _x.IsZero() {
			return ErrZero
		}
		s.Add(&s, &_x)
		c.Div(
			number.SafeMul(c, &d),
			number.SafeMul(&_x, &p.NumTokensU256),
		)
	}

	if Ann.IsZero() {
		return ErrZero
	}

	c.Div(
		number.SafeMul(number.SafeMul(c, &d), p.Extra.APrecision),
		number.SafeMul(Ann, &p.NumTokensU256),
	)

	var b = number.SafeAdd(
		&s,
		number.Div(number.SafeMul(&d, p.Extra.APrecision), Ann),
	)

	var yPrev uint256.Int
	y.Set(&d)
	for i := 0; i < 255; i += 1 {
		yPrev.Set(y)

		y.Div(
			number.SafeAdd(number.SafeMul(y, y), c),
			number.SafeSub(
				number.SafeAdd(
					number.SafeAdd(y, y),
					b,
				),
				&d,
			),
		)

		if number.WithinDelta(y, &yPrev, 1) {
			return nil
		}
	}

	return ErrAmountOutNotConverge
}
