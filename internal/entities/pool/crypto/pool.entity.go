package crypto

import (
	"github.com/holiman/uint256"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type (
	Extra struct {
		InitialAGamma     *uint256.Int
		InitialAGammaTime int64
		FutureAGamma      *uint256.Int
		FutureAGammaTime  int64

		D *uint256.Int

		PriceScale         []uint256.Int
		PriceOracle        []uint256.Int
		LastPrices         []uint256.Int
		LastPriceTimestamp int64

		FeeGamma *uint256.Int
		MidFee   *uint256.Int
		OutFee   *uint256.Int

		LpSupply           *uint256.Int
		XcpProfit          *uint256.Int
		VirtualPrice       *uint256.Int
		AllowedExtraProfit *uint256.Int
		AdjustmentStep     *uint256.Int
	}

	Pool struct {
		precisionMultipliers []uint256.Int
		Address, Exchange    string
		Reserves             []uint256.Int
		Tokens               []token.Token
		Extra                Extra
	}
)

func NewPool(address, exchange string, reserves []uint256.Int, tokens []token.Token, initial_a_gamma, future_a_gamma, d, fee_gamma, mid_fee, out_fee uint256.Int, price_scale []uint256.Int, initial_a_gamma_time, future_A_gamma_time int64) *Pool {
	pool := &Pool{
		Address:  address,
		Exchange: exchange,
		Reserves: reserves,
		Tokens:   tokens,
		Extra: Extra{
			InitialAGamma:     &initial_a_gamma,
			InitialAGammaTime: initial_a_gamma_time,
			FutureAGamma:      &future_a_gamma,
			FutureAGammaTime:  future_A_gamma_time,
			D:                 &d,
			PriceScale:        price_scale,
			FeeGamma:          &fee_gamma,
			MidFee:            &mid_fee,
			OutFee:            &out_fee,
		},
	}

	pool.precisionMultipliers = make([]uint256.Int, NumTokens)
	for i := 0; i < NumTokens; i++ {
		pool.precisionMultipliers[i].Set(number.TenPow(18 - tokens[i].Decimals))
	}
	return pool
}

func (p *Pool) FeeCalculate(xp []uint256.Int, fee *uint256.Int) error {
	var f uint256.Int
	var err = reductionCoefficient(xp, p.Extra.FeeGamma, &f)
	if err != nil {
		return err
	}
	fee.Div(
		number.SafeAdd(
			number.SafeMul(p.Extra.MidFee, &f),
			number.SafeMul(p.Extra.OutFee, number.SafeSub(U_1e18, &f)),
		),
		U_1e18,
	)
	return nil
}

func (p *Pool) GetDy(
	i, j int, dx *uint256.Int,

	dy, fee *uint256.Int, xp []uint256.Int,
) error {
	if i == j {
		return ErrSameCoin
	}
	if i >= NumTokens && j >= NumTokens {
		return ErrCoinIndexOutOfRange
	}

	for k := 0; k < NumTokens; k += 1 {
		if k == i {
			number.SafeAddZ(&p.Reserves[k], dx, &xp[k])
			continue
		}
		xp[k].Set(&p.Reserves[k])
	}

	number.SafeMulZ(&xp[0], &p.precisionMultipliers[0], &xp[0])
	for k := 0; k < NumTokens-1; k += 1 {
		xp[k+1].Div(
			number.SafeMul(number.SafeMul(&xp[k+1], &p.Extra.PriceScale[k]), &p.precisionMultipliers[k+1]),
			Precision,
		)
	}

	A, gamma := p._A_gamma()
	var y uint256.Int

	var err = newton_y(A, gamma, xp[:], p.Extra.D, j, &y)
	if err != nil {
		return err
	}
	number.SafeSubZ(number.SafeSub(&xp[j], &y), number.Number_1, dy)
	xp[j] = y
	if j > 0 {
		dy.Div(number.SafeMul(dy, Precision), &p.Extra.PriceScale[j-1])
	} else {
		dy.Div(dy, &p.precisionMultipliers[0])
	}

	err = p.FeeCalculate(xp[:], fee)
	if err != nil {
		return err
	}

	fee.Div(number.SafeMul(fee, dy), U_1e10)
	dy.Sub(dy, fee)

	return nil
}
