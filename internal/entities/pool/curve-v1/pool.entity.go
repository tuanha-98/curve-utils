package curvev1

import (
	"strconv"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/entities"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type FeeInfo struct {
	SwapFee, AdminFee, OffPegFeeMultiplier uint256.Int
}

type PoolSimulator struct {
	Address, Exchange string
	Reserves          []uint256.Int
	LpSupply          uint256.Int
	NumTokens         int
	NumTokensU256     uint256.Int
	Tokens            []token.Token
	Static            Static
	Extra             Extra
}

func (p *PoolSimulator) GetTokens() []token.Token {
	return p.Tokens
}

func (p *PoolSimulator) XpMem(rate_multipliers []uint256.Int, reserves []uint256.Int) []uint256.Int {
	return XpMem(rate_multipliers, reserves)
}

func (p *PoolSimulator) GetFeeInfo() FeeInfo {
	return FeeInfo{
		SwapFee:             *p.Extra.SwapFee,
		AdminFee:            *p.Extra.AdminFee,
		OffPegFeeMultiplier: *p.Extra.OffPegFeeMultiplier,
	}
}

func NewPool(
	entityPool entities.Pool,
) (*PoolSimulator, error) {
	var aPrecision, initialA, futureA, swapFee, adminFee, offPegFeeMultiplier, lpSupply uint256.Int

	lpSupply.SetFromDecimal(entityPool.TotalSupply)
	swapFee.SetFromDecimal(entityPool.Fee.SwapFee)
	adminFee.SetFromDecimal(entityPool.Fee.AdminFee)
	aPrecision.SetFromDecimal(entityPool.Multipliers.APrecision)
	offPegFeeMultiplier.SetFromDecimal(entityPool.Fee.OffPegFeeMultiplier)

	initialA.SetFromDecimal(entityPool.Amplification.Initial)
	futureA.SetFromDecimal(entityPool.Amplification.Future)
	initialATime, _ := strconv.ParseInt(entityPool.Amplification.InitialTime, 10, 64)
	futureATime, _ := strconv.ParseInt(entityPool.Amplification.FutureTime, 10, 64)

	rateMultipliers := make([]uint256.Int, entityPool.Ncoins)
	for i, rmStr := range entityPool.Multipliers.RateMultipliers {
		rateMultipliers[i].SetFromDecimal(rmStr)
	}

	precisionMultipliers := make([]uint256.Int, entityPool.Ncoins)
	for i, pStr := range entityPool.Multipliers.PrecisionMultipliers {
		precisionMultipliers[i].SetFromDecimal(pStr)
	}

	reserves := make([]uint256.Int, entityPool.Ncoins)
	for i, rStr := range entityPool.Reserves {
		reserves[i].SetFromDecimal(rStr)
	}

	tokens := make([]token.Token, entityPool.Ncoins)
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
		Exchange:      "CurveV1",
		Reserves:      reserves,
		LpSupply:      lpSupply,
		NumTokens:     entityPool.Ncoins,
		NumTokensU256: *number.SetUint64(uint64(entityPool.Ncoins)),
		Tokens:        tokens,
		Static: Static{
			PoolType:   entityPool.Type,
			APrecision: &aPrecision,
		},

		Extra: Extra{
			InitialA:             &initialA,
			FutureA:              &futureA,
			InitialATime:         initialATime,
			FutureATime:          futureATime,
			SwapFee:              &swapFee,
			AdminFee:             &adminFee,
			OffPegFeeMultiplier:  &offPegFeeMultiplier,
			RateMultipliers:      rateMultipliers,
			PrecisionMultipliers: precisionMultipliers,
		},
	}

	return pool, nil
}

func (p *PoolSimulator) FeeCalculate(dy, fee *uint256.Int) {
	fee.Div(
		number.SafeMul(dy, p.Extra.SwapFee),
		FeeDenominator,
	)
}

func (p *PoolSimulator) DynamicFee(xpi, xpj, swapFee, feeOutput *uint256.Int) {
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

func (p *PoolSimulator) GetDy(
	i, j int, dx *uint256.Int,
	// output
	dy *uint256.Int,
) error {
	switch p.Static.PoolType {
	case PoolTypeAave:
		var xp = xp(p.Extra.PrecisionMultipliers, p.Reserves)
		var x = number.SafeAdd(&xp[i], number.Mul(dx, &p.Extra.PrecisionMultipliers[i]))
		var y uint256.Int
		var err = p.getY(i, j, x, xp, nil, &y)
		if err != nil {
			return err
		}
		dy.Div(number.Sub(&xp[j], &y), &p.Extra.PrecisionMultipliers[j])
		var fee uint256.Int
		p.DynamicFee(number.Div(number.Add(&xp[i], x), number.Number_2), number.Div(number.Add(&xp[j], &y), number.Number_2), p.Extra.SwapFee, &fee)
		fee.Div(number.Mul(&fee, dy), FeeDenominator)
		dy.Sub(dy, &fee)
	default:
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
		p.FeeCalculate(dy, &fee)
		if dy.Cmp(&fee) < 0 {
			return ErrInvalidReserve
		}
		dy.Div(number.Mul(dy.Sub(dy, &fee), Precision), &p.Extra.RateMultipliers[j])
	}
	return nil
}
