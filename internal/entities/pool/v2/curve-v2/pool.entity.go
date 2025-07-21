package curvev2

import (
	"strconv"

	"github.com/holiman/uint256"
	entities "github.com/tuanha-98/curve-utils/internal/entities/pool/v2"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type PoolSimulator struct {
	Address, Exchange string
	Reserves          []uint256.Int
	LpSupply          uint256.Int
	NumTokens         int
	NumTokensU256     uint256.Int
	Tokens            []token.Token
	Extra             Extra
}

func NewPool(
	entityPool entities.Pool,
) (*PoolSimulator, error) {
	var initialAGamma, futureAGamma, midFee, outFee, feeGamma, adminFee, D, lpSupply uint256.Int

	lpSupply.SetFromDecimal(entityPool.TotalSupply)
	midFee.SetFromDecimal(entityPool.MidFee)
	outFee.SetFromDecimal(entityPool.OutFee)
	feeGamma.SetFromDecimal(entityPool.FeeGamma)
	adminFee.SetFromDecimal(entityPool.AdminFee)

	initialAGamma.SetFromDecimal(entityPool.InitialAGamma)
	futureAGamma.SetFromDecimal(entityPool.FutureAGamma)
	initialAGammaTime, _ := strconv.ParseInt(entityPool.InitialAGammaTime, 10, 64)
	futureAGammaTime, _ := strconv.ParseInt(entityPool.FutureAGammaTime, 10, 64)
	D.SetFromDecimal(entityPool.D)

	priceScales := make([]uint256.Int, entityPool.NTokens-1)
	for i, rmStr := range entityPool.PriceScales {
		priceScales[i].SetFromDecimal(rmStr)
	}

	precisions := make([]uint256.Int, entityPool.NTokens)
	for i, pStr := range entityPool.Precisions {
		precisions[i].SetFromDecimal(pStr)
	}

	reserves := make([]uint256.Int, entityPool.NTokens)
	for i, rStr := range entityPool.Reserves {
		reserves[i].SetFromDecimal(rStr)
	}

	tokens := make([]token.Token, entityPool.NTokens)
	for i, t := range entityPool.Tokens {
		tokens[i] = token.Token{
			Address:  t.ID,
			Name:     t.Symbol,
			Symbol:   t.Symbol,
			Decimals: uint8(t.Decimals),
		}
	}

	pool := &PoolSimulator{
		Address:       entityPool.Address,
		Exchange:      "CurveV2",
		Reserves:      reserves,
		LpSupply:      lpSupply,
		NumTokens:     entityPool.NTokens,
		NumTokensU256: *number.SetUint64(uint64(entityPool.NTokens)),
		Tokens:        tokens,

		Extra: Extra{
			InitialAGamma:     &initialAGamma,
			FutureAGamma:      &futureAGamma,
			InitialAGammaTime: initialAGammaTime,
			FutureAGammaTime:  futureAGammaTime,
			D:                 &D,
			MidFee:            &midFee,
			OutFee:            &outFee,
			FeeGamma:          &feeGamma,
			AdminFee:          &adminFee,
			PriceScales:       priceScales,
			Precisions:        precisions,
		},
	}

	return pool, nil
}

func (p *PoolSimulator) FeeCalculate(xp []uint256.Int, fee *uint256.Int) error {
	var f uint256.Int
	var err = reductionCoefficient(xp, p.Extra.FeeGamma, &f)
	if err != nil {
		return err
	}
	fee.Div(
		number.SafeAdd(
			number.SafeMul(p.Extra.MidFee, &f),
			number.SafeMul(p.Extra.OutFee, number.SafeSub(number.TenPow(18), &f)),
		),
		number.TenPow(18),
	)
	return nil
}

func (p *PoolSimulator) GetDy(
	i, j int, dx *uint256.Int,

	dy *uint256.Int,
) error {
	if i == j {
		return ErrSameCoin
	}
	if i >= p.NumTokens && j >= p.NumTokens {
		return ErrCoinIndexOutOfRange
	}

	var xp = make([]uint256.Int, p.NumTokens)

	for k := 0; k < p.NumTokens; k += 1 {
		if k == i {
			number.SafeAddZ(&p.Reserves[k], dx, &xp[k])
			continue
		}
		xp[k].Set(&p.Reserves[k])
	}

	number.SafeMulZ(&xp[0], &p.Extra.Precisions[0], &xp[0])
	for k := 0; k < p.NumTokens-1; k += 1 {
		xp[k+1].Div(
			number.SafeMul(
				number.SafeMul(
					&xp[k+1],
					&p.Extra.PriceScales[k],
				),
				&p.Extra.Precisions[k+1],
			),
			Precision,
		)
	}

	A, gamma := p._A_gamma()

	var y uint256.Int
	var err = newton_y(A, gamma, xp, p.Extra.D, j, &y)

	if err != nil {
		return err
	}

	number.SafeSubZ(number.SafeSub(&xp[j], &y), number.SetUint64(1), dy)
	xp[j].Set(&y)
	if j > 0 {
		dy.Div(number.SafeMul(dy, Precision), &p.Extra.PriceScales[j-1])
	}

	dy.Div(dy, &p.Extra.Precisions[j])

	var fee uint256.Int
	err = p.FeeCalculate(xp, &fee)
	if err != nil {
		return err
	}

	fee.Div(number.SafeMul(&fee, dy), number.TenPow(10))

	dy.Sub(dy, &fee)

	return nil
}
