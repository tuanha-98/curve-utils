package stable

import (
	"github.com/holiman/uint256"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type (
	Extra struct {
		APrecision *uint256.Int

		InitialA     *uint256.Int
		FutureA      *uint256.Int
		InitialATime int64
		FutureATime  int64
		SwapFee      *uint256.Int
		AdminFee     *uint256.Int

		RateMultipliers []uint256.Int
	}

	Pool struct {
		precisionMultipliers []uint256.Int
		Address, Exchange    string
		Reserves             []uint256.Int
		NumTokens            int
		NumTokensU256        uint256.Int
		Tokens               []token.Token
		Extra                Extra
	}
)

func NewPool(address, exchange string, rate_multipliers, reserves []uint256.Int, tokens []token.Token, a_precision, initial_a, future_a, swap_fee, admin_fee uint256.Int, initial_a_time, future_a_time int64) *Pool {
	numtokens := len(tokens)

	pool := &Pool{
		Address:       address,
		Exchange:      exchange,
		Reserves:      reserves,
		NumTokens:     numtokens,
		NumTokensU256: *number.SetUint64(uint64(numtokens)),
		Tokens:        tokens,
		Extra: Extra{
			APrecision:      &a_precision,
			InitialA:        &initial_a,
			FutureA:         &future_a,
			InitialATime:    initial_a_time,
			FutureATime:     future_a_time,
			SwapFee:         &swap_fee,
			AdminFee:        &admin_fee,
			RateMultipliers: rate_multipliers,
		},
	}
	useStandardRate := false
	if len(rate_multipliers) == 0 {
		pool.Extra.RateMultipliers = make([]uint256.Int, numtokens)
		useStandardRate = true
	}
	pool.precisionMultipliers = make([]uint256.Int, numtokens)
	for i := 0; i < numtokens; i++ {
		if useStandardRate {
			pool.Extra.RateMultipliers[i].Set(number.TenPow(36 - tokens[i].Decimals))
		}

		pool.precisionMultipliers[i].Set(number.TenPow(18 - tokens[i].Decimals))
	}

	return pool
}

func (p *Pool) FeeCalculate(dy, fee *uint256.Int) {
	fee.Div(
		number.SafeMul(dy, p.Extra.SwapFee),
		FeeDenominator,
	)
}

func (p *Pool) GetDy(
	i, j int, dx *uint256.Int,
	// output
	dy *uint256.Int,
) error {
	var xp = XpMem(p.Extra.RateMultipliers, p.Reserves)
	var x = number.SafeAdd(&xp[i], number.Div(number.Mul(dx, &p.Extra.RateMultipliers[i]), Precision))
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
	var fee uint256.Int
	p.FeeCalculate(
		dy,
		&fee,
	)

	if dy.Cmp(&fee) < 0 {
		return ErrInvalidReserve
	}

	dy.Div(number.Mul(dy.Sub(dy, &fee), Precision), &p.Extra.RateMultipliers[j])

	return nil
}
