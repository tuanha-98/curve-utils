package curvev1ng

import (
	"math"
	"time"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

var NowFunc = time.Now

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
	Ann_mul_S_div_APrec.Div(number.Mul(&Ann, &S), p.Static.APrecision)
	Ann_sub_APrec.Sub(&Ann, p.Static.APrecision)

	numTokensPlus1 := uint256.NewInt(uint64(p.NumTokens + 1))
	numTokensPow := uint256.NewInt(uint64(math.Pow(float64(p.NumTokens), float64(p.NumTokens))))
	for i := 0; i < 255; i++ {
		D_P.Set(D)

		for _, x := range xp {
			D_P.Div(
				number.SafeMul(&D_P, D),
				&x,
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

func (p *PoolSimulator) getY(
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

	for _i := 0; _i < p.NumTokens; _i++ {
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
	yPrev.Clear()
	y.Set(&d)
	for i := 0; i < 255; i++ {
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

func (p *PoolSimulator) getYD(
	A *uint256.Int,
	i int,
	xp []uint256.Int,
	D *uint256.Int,
	// output
	y *uint256.Int,
) error {
	if i >= p.NumTokens {
		return ErrTokenIndexOutOfRange
	}
	var c, S uint256.Int
	c.Set(D)
	S.Clear()
	var Ann = number.Mul(A, &p.NumTokensU256)
	for _i := 0; _i < p.NumTokens; _i += 1 {
		if _i != i {
			// S.Add(&S, &xp[_i])
			number.SafeAddZ(&S, &xp[_i], &S)
			c.Div(
				number.Mul(&c, D),
				number.Mul(&xp[_i], &p.NumTokensU256),
			)
		}
	}
	if Ann.IsZero() {
		return ErrZero
	}
	c.Div(
		number.Mul(number.Mul(&c, D), p.Static.APrecision),
		number.Mul(Ann, &p.NumTokensU256),
	)
	var b = number.Add(
		&S,
		number.Div(number.Mul(D, p.Static.APrecision), Ann),
	)
	var yPrev uint256.Int
	y.Set(D)
	for _i := 0; _i < MaxLoopLimit; _i += 1 {
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
				D,
			),
		)
		if number.WithinDelta(y, &yPrev, 1) {
			return nil
		}
	}
	return ErrAmountOutNotConverge
}

func (p *PoolSimulator) GetDyByX(
	i, j int,
	x *uint256.Int,
	xp []uint256.Int,
	// output
	dy *uint256.Int,
	// adminFee *uint256.Int,
) error {
	var y uint256.Int

	var err = p.getY(i, j, x, xp, nil, &y)
	if err != nil {
		return err
	}
	number.SafeSubZ(&xp[j], &y, dy)
	if dy.Sign() <= 0 {
		return ErrZero
	}
	dy.SubUint64(dy, 1)
	var dynamicFee, dyFee uint256.Int
	p.DynamicFee(
		number.Div(number.SafeAdd(&xp[i], x), number.Number_2),
		number.Div(number.SafeAdd(&xp[j], &y), number.Number_2),
		p.Extra.SwapFee,
		&dynamicFee,
	)
	dyFee.Div(
		number.SafeMul(dy, &dynamicFee),
		FeeDenominator,
	)

	dy.Div(number.SafeMul(number.SafeSub(dy, &dyFee), Precision), &p.Extra.Rates[j])

	// adminFee.Div(
	// 	number.SafeMul(
	// 		number.Div(
	// 			number.SafeMul(&dyFee, p.Extra.AdminFee),
	// 			FeeDenominator,
	// 		),
	// 		Precision,
	// 	),
	// 	&p.Extra.Rates[j],
	// )

	return nil
}

func (p *PoolSimulator) CalculateWithdrawOneCoin(
	burnAmount *uint256.Int,
	i int,
	// output
	dy *uint256.Int, dyFee *uint256.Int,
) error {
	var amp = p._A()
	var xp = XpMem(p.Extra.Rates, p.Reserves)

	var D0, D1, newY, newYD uint256.Int
	var err = p.getD(xp, amp, &D0)
	if err != nil {
		return err
	}
	var totalSupply = &p.LpSupply
	number.SafeSubZ(&D0, number.Div(number.Mul(burnAmount, &D0), totalSupply), &D1)
	err = p.getYD(amp, i, xp, &D1, &newY)
	if err != nil {
		return err
	}
	var baseFee = number.Div(
		number.Mul(p.Extra.SwapFee, &p.NumTokensU256),
		number.Mul(number.Number_4, uint256.NewInt(uint64(p.NumTokens-1))),
	)
	var xpReduced [MaxTokenCount]uint256.Int
	var ys = number.Div(number.SafeAdd(&D0, &D1), uint256.NewInt(uint64(p.NumTokens*2)))

	var dxExpected, xavg, dynamicFee uint256.Int
	for j := 0; j < p.NumTokens; j++ {
		if j == i {
			number.SafeSubZ(number.Div(number.SafeMul(&xp[j], &D1), &D0), &newY, &dxExpected)
			xavg.Div(number.SafeAdd(&xp[j], &newY), number.Number_2)
		} else {
			number.SafeSubZ(&xp[j], number.Div(number.SafeMul(&xp[j], &D1), &D0), &dxExpected)
			xavg.Set(&xp[j])
		}
		p.DynamicFee(&xavg, ys, baseFee, &dynamicFee)
		number.SafeSubZ(&xp[j], number.Div(number.SafeMul(&dynamicFee, &dxExpected), FeeDenominator), &xpReduced[j])
	}

	err = p.getYD(amp, i, xpReduced[:p.NumTokens], &D1, &newYD)
	if err != nil {
		return err
	}
	number.SafeSubZ(&xpReduced[i], &newYD, dy)

	var dy0 = number.Div(number.SafeMul(number.SafeSub(&xp[i], &newY), Precision), &p.Extra.Rates[i])
	if dy.Sign() <= 0 {
		return ErrZero
	}

	dy.Div(number.SafeMul(number.SafeSub(dy, number.Number_1), Precision), &p.Extra.Rates[i])
	number.SafeSubZ(dy0, dy, dyFee)

	return nil
}

func (p *PoolSimulator) CalculateTokenAmount(
	amounts []uint256.Int,
	deposit bool,

	// output
	mintAmount *uint256.Int,
	// feeAmounts []uint256.Int,
) error {
	var a = p._A()
	var d0, d1, d2 uint256.Int
	var xp = XpMem(p.Extra.Rates, p.Reserves)

	// Initial invariant
	err := p.getD(xp, a, &d0)
	if err != nil {
		return err
	}

	var newBalances [MaxTokenCount]uint256.Int
	for i := 0; i < p.NumTokens; i++ {
		if deposit {
			number.SafeAddZ(&p.Reserves[i], &amounts[i], &newBalances[i])
		} else {
			number.SafeSubZ(&p.Reserves[i], &amounts[i], &newBalances[i])
		}
	}

	// Invariant after change
	xp = XpMem(p.Extra.Rates, newBalances[:p.NumTokens])
	err = p.getD(xp, a, &d1)
	if err != nil {
		return err
	}

	// We need to recalculate the invariant accounting for fees
	// to calculate fair user's share
	var totalSupply = &p.LpSupply
	if !totalSupply.IsZero() {
		// Only account for fees if we are not the first to deposit
		var baseFee = number.Div(
			number.Mul(p.Extra.SwapFee, &p.NumTokensU256),
			uint256.NewInt(4*uint64(p.NumTokens-1)),
		)
		var _dynamic_fee_i, difference, xs, ys uint256.Int
		// ys: uint256 = (D0 + D1) / N_COINS
		ys.Div(number.SafeAdd(&d0, &d1), &p.NumTokensU256)
		for i := 0; i < p.NumTokens; i++ {
			// ideal_balance: uint256 = D1 * old_balances[i] / D0
			ideal_balance := number.Div(number.SafeMul(&d1, &p.Reserves[i]), &d0)
			if ideal_balance.Cmp(&newBalances[i]) > 0 {
				difference.Sub(ideal_balance, &newBalances[i])
			} else {
				difference.Sub(&newBalances[i], ideal_balance)
			}

			// xs = old_balances[i] + new_balance
			number.SafeAddZ(&p.Reserves[i], &newBalances[i], &xs)

			// this line is from `add_liquidity` method, the `calc_token_amount` method doesn't have it (might be a bug)
			// xs = unsafe_div(rates[i] * (old_balances[i] + new_balance), PRECISION)
			xs.Div(number.SafeMul(&p.Extra.Rates[i], &xs), Precision)

			// _dynamic_fee_i = self._dynamic_fee(xs, ys, base_fee, fee_multiplier)
			p.DynamicFee(&xs, &ys, baseFee, &_dynamic_fee_i)

			// new_balances[i] -= _dynamic_fee_i * difference / FEE_DENOMINATOR
			fee := number.Div(number.SafeMul(&_dynamic_fee_i, &difference), FeeDenominator)
			number.SafeSubZ(&newBalances[i], fee, &newBalances[i])

			// record fee so we can update balance later
			// self.admin_balances[i] += unsafe_div(fees[i] * admin_fee, FEE_DENOMINATOR)
			// feeAmounts[i].Div(number.SafeMul(fee, p.Extra.AdminFee), FeeDenominator)
		}

		for i := 0; i < p.NumTokens; i++ {
			// xp[idx] = rates[idx] * new_balances[idx] / PRECISION
			xp[i].Div(number.SafeMul(&p.Extra.Rates[i], &newBalances[i]), Precision)
		}
		// D2 = self.get_D(xp, amp, N_COINS)
		err = p.getD(xp, a, &d2)
		if err != nil {
			return err
		}
	} else {
		// Take the dust if there was any
		mintAmount.Set(&d1)
		return nil
	}

	var diff uint256.Int
	if deposit {
		number.SafeSubZ(&d2, &d0, &diff)
	} else {
		number.SafeSubZ(&d0, &d2, &diff)
	}
	// return diff * total_supply / D0
	mintAmount.Div(number.Mul(&diff, totalSupply), &d0)
	return nil
}
