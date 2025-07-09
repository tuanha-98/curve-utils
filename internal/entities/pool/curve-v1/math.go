package curvev1

import (
	"math"
	"time"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

var NowFunc = time.Now

func xp(
	precisions []uint256.Int,
	balances []uint256.Int,
) []uint256.Int {
	xp := make([]uint256.Int, 0)
	var numTokens = len(balances)
	if numTokens != len(precisions) {
		return nil
	}
	for i := 0; i < numTokens; i++ {
		result := number.SafeMul(&balances[i], &precisions[i])
		xp = append(xp, *result)
	}
	return xp
}

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

func (p *PoolSimulator) getD(
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
	numTokensPow := uint256.NewInt(uint64(math.Pow(float64(p.NumTokens), float64(p.NumTokens))))

	for i := 0; i < 255; i += 1 {
		D_P.Set(D)

		switch p.Static.PoolType {
		case PoolTypeAave:
			fallthrough
		case PoolTypeCompound:
			for j := range xp {
				D_P.Div(
					number.SafeMul(&D_P, D),
					number.SafeAdd(number.Mul(&xp[j], &p.NumTokensU256), number.Number_1),
				)
			}
		default:
			for j := range xp {
				D_P.Div(
					number.SafeMul(&D_P, D),
					&xp[j],
				)
			}
			D_P.Div(&D_P, numTokensPow)
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

func (p *PoolSimulator) get_D_mem(rates []uint256.Int, balances []uint256.Int, amp *uint256.Int, D *uint256.Int) error {
	var xp = XpMem(rates, balances)
	return p.getD(xp, amp, D)
}

func (p *PoolSimulator) getY(
	i, j int, x *uint256.Int, xp []uint256.Int, dCached *uint256.Int,
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

func (p *PoolSimulator) getYD(
	a *uint256.Int,
	tokenIndex int,
	xp []uint256.Int,
	d *uint256.Int,

	//output
	y *uint256.Int,
) error {
	var numTokens = len(xp)
	if tokenIndex >= numTokens {
		return ErrTokenNotFound
	}
	var c, s uint256.Int
	c.Set(d)
	s.Clear()
	var nA = number.Mul(a, &p.NumTokensU256)
	for i := 0; i < numTokens; i++ {
		if i != tokenIndex {
			s.Add(&s, &xp[i])
			c.Div(
				number.Mul(&c, d),
				number.Mul(&xp[i], &p.NumTokensU256),
			)
		}
	}
	if nA.IsZero() {
		return ErrZero
	}
	c.Div(
		number.Mul(number.Mul(&c, d), p.Static.APrecision),
		number.Mul(nA, &p.NumTokensU256),
	)
	var b = number.Add(
		&s,
		number.Div(number.Mul(d, p.Static.APrecision), nA),
	)
	var yPrev uint256.Int
	y.Set(d)
	for i := 0; i < MaxLoopLimit; i++ {
		yPrev.Set(y)
		y.Div(
			number.Add(
				number.Mul(y, y),
				&c,
			),
			number.Sub(
				number.Add(
					number.Add(y, y),
					b,
				),
				d,
			),
		)
		if number.WithinDelta(y, &yPrev, 1) {
			return nil
		}
	}
	return ErrAmountOutNotConverge
}

func (p *PoolSimulator) CalculateWithdrawOneCoin(
	tokenAmount *uint256.Int,
	i int,

	// output
	dy *uint256.Int, dyFee *uint256.Int,
) error {
	var amp = p._A()
	var xp = XpMem(p.Extra.RateMultipliers, p.Reserves)
	var D0, newY, newYD uint256.Int
	err := p.getD(xp, amp, &D0)
	if err != nil {
		return err
	}
	var totalSupply = &p.LpSupply
	var D1 = number.Sub(&D0, number.Div(number.Mul(tokenAmount, &D0), totalSupply))
	err = p.getYD(amp, i, xp, D1, &newY)
	if err != nil {
		return err
	}
	var nCoins = len(p.Reserves)
	var xpReduced [MaxTokenCount]uint256.Int
	var nCoinBI = number.SetUint64(uint64(nCoins))
	var fee = number.Div(number.Mul(p.Extra.SwapFee, nCoinBI), number.Mul(uint256.NewInt(4), number.SubUint64(nCoinBI, 1)))
	for j := 0; j < nCoins; j += 1 {
		var dxExpected uint256.Int
		if j == i {
			dxExpected.Sub(number.Div(number.Mul(&xp[j], D1), &D0), &newY)
		} else {
			dxExpected.Sub(&xp[j], number.Div(number.Mul(&xp[j], D1), &D0))
		}
		xpReduced[j].Sub(&xp[j], number.Div(number.Mul(fee, &dxExpected), FeeDenominator))
	}
	err = p.getYD(amp, i, xpReduced[:nCoins], D1, &newYD)
	if err != nil {
		return err
	}
	dy.Sub(&xpReduced[i], &newYD)
	if dy.Sign() <= 0 {
		return ErrZero
	}
	dy.Div(number.SubUint64(dy, 1), &p.Extra.PrecisionMultipliers[i])
	var dy0 = number.Div(number.Sub(&xp[i], &newY), &p.Extra.PrecisionMultipliers[i])
	dyFee.Sub(dy0, dy)
	return nil
}

func (p *PoolSimulator) CalculateTokenAmount(
	amounts []uint256.Int,
	deposit bool,

	// output
	mintAmount *uint256.Int,
	feeAmounts []uint256.Int,
) error {
	var numTokens = len(p.Tokens)
	var a = p._A()
	var d0, d1, d2 uint256.Int
	err := p.get_D_mem(p.Extra.RateMultipliers, p.Reserves, a, &d0)
	if err != nil {
		return err
	}
	var balances1 [MaxTokenCount]uint256.Int
	for i := 0; i < numTokens; i++ {
		if deposit {
			balances1[i].Add(&p.Reserves[i], &amounts[i])
		} else {
			if p.Reserves[i].Cmp(&amounts[i]) < 0 {
				return ErrWithdrawMoreThanAvailable
			}
			balances1[i].Sub(&p.Reserves[i], &amounts[i])
		}
	}

	err = p.get_D_mem(p.Extra.RateMultipliers, balances1[:], a, &d1)
	if err != nil {
		return err
	}
	// in SC, this method won't take fee into account, so the result is different than the actual add_liquidity method
	// we'll copy that code here

	// We need to recalculate the invariant accounting for fees
	// to calculate fair user's share
	var totalSupply = p.LpSupply
	var difference uint256.Int
	if !totalSupply.IsZero() {
		var _fee = number.Div(number.Mul(p.Extra.SwapFee, &p.NumTokensU256),
			number.Mul(number.Number_4, uint256.NewInt(uint64(p.NumTokens-1))))
		var _admin_fee = p.Extra.AdminFee
		for i := 0; i < p.NumTokens; i += 1 {
			var ideal_balance = number.Div(number.Mul(&d1, &p.Reserves[i]), &d0)
			if ideal_balance.Cmp(&balances1[i]) > 0 {
				difference.Sub(ideal_balance, &balances1[i])
			} else {
				difference.Sub(&balances1[i], ideal_balance)
			}
			var fee = number.Div(number.Mul(_fee, &difference), FeeDenominator)
			feeAmounts[i].Set(number.Div(number.Mul(fee, _admin_fee), FeeDenominator))
			balances1[i].Sub(&balances1[i], fee)
		}
		_ = p.get_D_mem(p.Extra.RateMultipliers, balances1[:p.NumTokens], a, &d2)
		mintAmount.Div(number.Mul(&totalSupply, number.Sub(&d2, &d0)), &d0)
	} else {
		mintAmount.Set(&d1)
	}

	return nil
}
