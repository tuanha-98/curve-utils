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
		InitialA            *uint256.Int
		FutureA             *uint256.Int
		InitialATime        int64
		FutureATime         int64
		SwapFee             *uint256.Int
		AdminFee            *uint256.Int
		RateMultipliers     []uint256.Int
	}

	Pool struct {
		Address, Exchange string
		Reserves          []uint256.Int
		LpSupply          uint256.Int
		NumTokens         int
		NumTokensU256     uint256.Int
		Tokens            []token.Token
		Extra             Extra
	}
)

func (p *Pool) GetTokens() []token.Token {
	return p.Tokens
}

func (p *Pool) XpMem(rate_multipliers []uint256.Int, reserves []uint256.Int) []uint256.Int {
	return XpMem(rate_multipliers, reserves)
}

func NewPool(address, exchange string, reserves []uint256.Int, tokens []token.Token, a_precision, off_peg_fee_multipler, initial_a, future_a, swap_fee, admin_fee, total_supply uint256.Int, rate_multipliers []uint256.Int, initial_a_time, future_a_time int64) *Pool {
	numtokens := len(tokens)

	return &Pool{
		Address:       address,
		Exchange:      exchange,
		Reserves:      reserves,
		LpSupply:      total_supply,
		NumTokens:     numtokens,
		NumTokensU256: *number.SetUint64(uint64(numtokens)),
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

func (p *Pool) DynamicFee(xpi, xpj, swapFee, feeOutput *uint256.Int) {
	_off_peg_fee_multiplier := p.Extra.OffPegFeeMultiplier
	if _off_peg_fee_multiplier.Cmp(FeeDenominator) <= 0 {
		feeOutput.Set(swapFee)
		return
	}

	sum := number.SafeAdd(xpi, xpj)
	prod := number.SafeMul(xpi, xpj)
	xps2 := number.SafeMul(sum, sum)
	feeOutput.Div(
		number.Mul(_off_peg_fee_multiplier, swapFee),
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
	adminFee *uint256.Int,
) error {
	var xp = XpMem(p.Extra.RateMultipliers, p.Reserves)
	var x = number.SafeAdd(&xp[i], number.Div(number.SafeMul(dx, &p.Extra.RateMultipliers[i]), Precision))
	return p.GetDyByX(i, j, x, xp, dy, adminFee)
}
