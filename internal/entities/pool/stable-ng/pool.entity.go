package stableng

import (
	"github.com/holiman/uint256"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type (
	Extra struct {
		APrecision          *uint256.Int
		OffPegFeeMultiplier *uint256.Int

		InitialA     *uint256.Int
		FutureA      *uint256.Int
		InitialATime int64
		FutureATime  int64
		SwapFee      *uint256.Int
		AdminFee     *uint256.Int

		RateMultipliers []uint256.Int
	}

	Pool struct {
		Address, Exchange string
		Reserves          []uint256.Int
		NumTokens         int
		NumTokensU256     uint256.Int
		Tokens            []token.Token
		Extra             Extra
	}
)

func NewPool(address, exchange string, reserves []uint256.Int, tokens []token.Token, a_precision, off_peg_fee_multipler, initial_a, future_a, swap_fee, admin_fee uint256.Int, rate_multipliers []uint256.Int, initial_a_time, future_a_time int64) *Pool {
	return &Pool{
		Address:       address,
		Exchange:      exchange,
		Reserves:      reserves,
		NumTokens:     len(tokens),
		NumTokensU256: *number.SetUint64(uint64(len(tokens))),
		Extra: Extra{
			APrecision:          &a_precision,
			OffPegFeeMultiplier: &off_peg_fee_multipler,
			InitialA:            &initial_a,
			FutureA:             &future_a,
			InitialATime:        initial_a_time,
			FutureATime:         future_a_time,
			SwapFee:             &swap_fee,
			AdminFee:            &admin_fee,
			RateMultipliers:     rate_multipliers,
		},
	}
}

func (p *Pool) DynamicFee(xpi, xpj, fee *uint256.Int) {
	_swap_fee := p.Extra.SwapFee
	_off_peg_fee_multiplier := p.Extra.OffPegFeeMultiplier
	if _off_peg_fee_multiplier.Cmp(FeeDenominator) <= 0 {
		fee.Set(_swap_fee)
		return
	}

	sum := number.SafeAdd(xpi, xpj)
	prod := number.SafeMul(xpi, xpj)
	xps2 := number.SafeMul(sum, sum)
	fee.Div(
		number.Mul(_off_peg_fee_multiplier, _swap_fee),
		number.Add(
			number.Div(
				number.SafeMul(
					number.SafeMul(
						number.Sub(_off_peg_fee_multiplier, FeeDenominator),
						number.Number_4,
					),
					prod,
				),
				xps2,
			),
			FeeDenominator,
		),
	)
}

func (p *Pool) GetDy(
	i, j int, dx *uint256.Int,
	// output
	dy *uint256.Int,
) error {
	var xp = XpMem(p.Extra.RateMultipliers, p.Reserves)
	var x = number.SafeAdd(&xp[i], number.Div(number.SafeMul(dx, &p.Extra.RateMultipliers[i]), Precision))
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
	var fee, dyFee uint256.Int
	p.DynamicFee(
		number.Div(number.SafeAdd(&xp[i], x), number.Number_2),
		number.Div(number.SafeAdd(&xp[j], &y), number.Number_2),
		&fee,
	)
	dyFee.Div(
		number.SafeMul(dy, &fee),
		FeeDenominator,
	)

	dy.Div(number.SafeMul(number.SafeSub(dy, &dyFee), Precision), &p.Extra.RateMultipliers[j])

	return nil
}
