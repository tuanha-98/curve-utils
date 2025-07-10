package curvev1meta

import (
	"time"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

var NowFunc = time.Now

func (p *PoolSimulator) _A() *uint256.Int {
	var t1 = p.Extra.FutureATime
	var a1 = p.Extra.FutureA
	var now = NowFunc().Unix()
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

func (p *PoolSimulator) _getD(
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
	Ann_mul_S_div_APrec.Div(number.SafeMul(&Ann, &S), p.Static.APrecision)
	Ann_sub_APrec.Sub(&Ann, p.Static.APrecision)

	numTokensPlus1 := uint256.NewInt(uint64(p.NumTokens + 1))

	for i := 0; i < MaxLoopLimit; i++ {
		D_P.Set(D)

		for j := range xp {
			D_P.Div(
				number.SafeMul(&D_P, D),
				number.Mul(&xp[j], &p.NumTokensU256),
			)
		}

		Dprev.Set(D)

		D.Div(
			number.SafeMul(
				number.SafeAdd(&Ann_mul_S_div_APrec, number.SafeMul(&D_P, &p.NumTokensU256)),
				D,
			),
			number.SafeAdd(
				number.Div(number.SafeMul(&Ann_sub_APrec, D), p.Static.APrecision),
				number.SafeMul(&D_P, numTokensPlus1),
			),
		)

		if number.WithinDelta(D, &Dprev, 1) {
			return nil
		}
	}
	return ErrDDoesNotConverge
}

// func (p *PoolSimulator) _get_D_mem(rates []uint256.Int, balances []uint256.Int, amp *uint256.Int, D *uint256.Int) error {
// 	var xp = p.BasePool.XpMem(rates, balances)
// 	return p._getD(xp, amp, D)
// }

func (p *PoolSimulator) _getY(
	i, j int, x *uint256.Int, xp []uint256.Int,
	// output
	y *uint256.Int,
) error {
	if i == j {
		return ErrTokenFromEqualsTokenTo
	}
	if i >= p.NumTokens && j >= p.NumTokens {
		return ErrTokenIndexesOutOfRange
	}

	var a = p._A()

	if a == nil {
		return ErrInvalidAValue
	}

	var d uint256.Int

	err := p._getD(xp, a, &d)
	if err != nil {
		return err
	}

	var c = number.Set(&d)
	var Ann = number.Mul(a, &p.NumTokensU256)
	var _x, s uint256.Int
	s.Clear()

	for _i := 0; _i < p.NumTokens; _i += 1 {
		if _i == i {
			_x.Set(x)
		} else if _i != j {
			_x.Set(&xp[_i])
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
		number.SafeMul(number.SafeMul(c, &d), p.Static.APrecision),
		number.SafeMul(Ann, &p.NumTokensU256),
	)

	var b = number.SafeAdd(
		&s,
		number.Div(number.SafeMul(&d, p.Static.APrecision), Ann),
	)

	var yPrev uint256.Int
	y.Set(&d)
	for i := 0; i < MaxLoopLimit; i += 1 {
		yPrev.Set(y)

		y.Div(
			number.SafeAdd(number.SafeMul(y, y), c), // this overflows if y is too big
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
