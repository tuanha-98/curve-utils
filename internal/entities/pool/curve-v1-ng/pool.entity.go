package curvev1ng

import (
	"strconv"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/entities"
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
	Static            Static
	Extra             Extra
}

func (p *PoolSimulator) GetBasePoolType() string {
	return "curvev1ng"
}

func (p *PoolSimulator) GetNumTokens() int {
	return p.NumTokens
}

func (p *PoolSimulator) XpMem(rate_multipliers []uint256.Int, reserves []uint256.Int) []uint256.Int {
	return XpMem(rate_multipliers, reserves)
}

func (p *PoolSimulator) GetFeeInfo() entities.FeeInfo {
	return entities.FeeInfo{
		SwapFee:   *p.Extra.SwapFee,
		AdminFee:  *p.Extra.AdminFee,
		OffPegFee: *p.Extra.OffPegFee,
	}
}

func NewPool(
	entityPool entities.Pool,
) (*PoolSimulator, error) {
	var aPrecision, initialA, futureA, swapFee, adminFee, offPegFee, lpSupply uint256.Int

	lpSupply.SetFromDecimal(entityPool.TotalSupply)
	swapFee.SetFromDecimal(entityPool.SwapFee)
	adminFee.SetFromDecimal(entityPool.AdminFee)
	aPrecision.SetFromDecimal(entityPool.APrecision)
	offPegFee.SetFromDecimal(entityPool.OffPegFee)

	initialA.SetFromDecimal(entityPool.InitialA)
	futureA.SetFromDecimal(entityPool.FutureA)
	initialATime, _ := strconv.ParseInt(entityPool.InitialATime, 10, 64)
	futureATime, _ := strconv.ParseInt(entityPool.FutureATime, 10, 64)

	rates := make([]uint256.Int, entityPool.NTokens)
	for i, rmStr := range entityPool.Rates {
		rates[i].SetFromDecimal(rmStr)
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
		Exchange:      "CurveV1NG",
		Reserves:      reserves,
		LpSupply:      lpSupply,
		NumTokens:     entityPool.NTokens,
		NumTokensU256: *number.SetUint64(uint64(entityPool.NTokens)),
		Tokens:        tokens,
		Static: Static{
			PoolType:   entityPool.Kind,
			APrecision: &aPrecision,
		},

		Extra: Extra{
			InitialA:     &initialA,
			FutureA:      &futureA,
			InitialATime: initialATime,
			FutureATime:  futureATime,
			SwapFee:      &swapFee,
			AdminFee:     &adminFee,
			OffPegFee:    &offPegFee,
			Rates:        rates,
			Precisions:   precisions,
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
	_off_peg_fee_multiplier := p.Extra.OffPegFee
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
	// adminFee *uint256.Int,
) error {
	var xp = XpMem(p.Extra.Rates, p.Reserves)
	var x = number.SafeAdd(&xp[i], number.Div(number.SafeMul(dx, &p.Extra.Rates[i]), Precision))
	return p.GetDyByX(i, j, x, xp, dy)
}
