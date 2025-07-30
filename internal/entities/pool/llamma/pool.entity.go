package llamma

import (
	"github.com/holiman/uint256"
	"github.com/samber/lo"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/maths/int256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type PoolSimulator struct {
	Address, Exchange string
	Reserves          []uint256.Int
	NumTokens         int
	NumTokensU256     uint256.Int
	Tokens            []token.Token
	Extra             Extra
}

func NewPool(
	entityPool Pool,
) (*PoolSimulator, error) {
	var A, swapFee, adminFee, borrowedPrecision, collateralPrecision, basePrice, priceOracle uint256.Int

	A.SetFromDecimal(entityPool.A)
	basePrice.SetFromDecimal(entityPool.BasePrice)
	priceOracle.SetFromDecimal(entityPool.PriceOracle)
	swapFee.SetFromDecimal(entityPool.SwapFee)
	adminFee.SetFromDecimal(entityPool.AdminFee)
	borrowedPrecision.SetFromDecimal(entityPool.BorrowedPrecision)
	collateralPrecision.SetFromDecimal(entityPool.CollateralPrecision)

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

	Aminus1 := number.Sub(&A, number.Number_1)
	ARatio := number.SafeDiv(
		number.Mul(&A, number.Number_1e18),
		Aminus1,
	)

	maxOracleDnPow := number.TenPow(18)
	for range maxTicks {
		maxOracleDnPow = number.SafeDiv(
			number.Mul(maxOracleDnPow, number.Number_1e18),
			Aminus1,
		)
	}

	logARatio := lnInt(ARatio)

	pool := &PoolSimulator{
		Address:       entityPool.Address,
		Exchange:      "CurveLlamma",
		Reserves:      reserves,
		NumTokens:     entityPool.NTokens,
		NumTokensU256: *number.SetUint64(uint64(entityPool.NTokens)),
		Tokens:        tokens,

		Extra: Extra{
			A:                   &A,
			Aminus1:             Aminus1,
			LogARatio:           logARatio,
			MaxOracleDnPow:      maxOracleDnPow,
			BasePrice:           &basePrice,
			PriceOracle:         &priceOracle,
			BorrowedPrecision:   &borrowedPrecision,
			CollateralPrecision: &collateralPrecision,
			SwapFee:             &swapFee,
			AdminFee:            &adminFee,
			ActiveBand:          entityPool.ActiveBand,
			MinBand:             entityPool.MinBand,
			MaxBand:             entityPool.MaxBand,
			BandsX: lo.SliceToMap(entityPool.Bands, func(b Band) (int64, *uint256.Int) {
				return b.Index, uint256.MustFromDecimal(b.BandX)
			}),
			BandsY: lo.SliceToMap(entityPool.Bands, func(b Band) (int64, *uint256.Int) {
				return b.Index, uint256.MustFromDecimal(b.BandY)
			}),
		},
	}

	return pool, nil
}

func (p *PoolSimulator) GetDy(
	i, j int, dx *uint256.Int,
	// output
	dy *uint256.Int,
) error {
	out, err := p.getDxDy(i, j, dx)
	if err != nil {
		return err
	}
	dy.Set(&out.OutAmount)
	return nil
}

func (p *PoolSimulator) getDxDy(
	i, j int, amount *uint256.Int,
) (*DetailedTrade, error) {
	if i^j != 1 {
		return nil, ErrWrongIndex
	}

	if amount.Sign() == 0 {
		return nil, ErrZeroSwapAmount
	}

	inPrecision, outPrecision := p.Extra.CollateralPrecision, p.Extra.BorrowedPrecision
	if i == 0 {
		inPrecision, outPrecision = outPrecision, inPrecision
	}

	out := &DetailedTrade{}
	var err error
	out, err = p.calcSwapOut(i, new(uint256.Int).Mul(amount, inPrecision), p.Extra.PriceOracle, inPrecision, outPrecision)

	if err != nil {
		return nil, err
	}

	out.InAmount.Div(&out.InAmount, inPrecision)
	out.OutAmount.Div(&out.OutAmount, outPrecision)

	// ? Should this ignore checked if amount out is zero
	// if out.InAmount.Sign() == 0 || out.OutAmount.Sign() == 0 {
	// 	return nil, ErrZeroSwapAmount
	// }

	return out, nil
}

func (p *PoolSimulator) calcSwapOut(
	inIdx int, inAmount, po, inPrecision, outPrecision *uint256.Int,
) (*DetailedTrade, error) {
	pump := inIdx == 0
	minBand := p.Extra.MinBand
	maxBand := p.Extra.MaxBand

	out := &DetailedTrade{N2: p.Extra.ActiveBand}
	poUp, err := p.pOracleUp(out.N2)
	if err != nil {
		return nil, err
	}

	var x, y, inAmountLeft, antifee, temp uint256.Int
	x.Set(p.getBandX(out.N2))
	y.Set(p.getBandY(out.N2))
	inAmountLeft.Set(inAmount)

	antifee.Div(
		Number_1e36,
		temp.Sub(number.Number_1e18, minUint256(p.Extra.SwapFee, tenPow18Minus1)),
	)

	j := maxTicksUnit
	for i := range maxTicks + maxSkipTicks {
		var y0, f, g, inv, dynamicFee uint256.Int
		dynamicFee.Set(p.Extra.SwapFee)

		if x.Sign() > 0 || y.Sign() > 0 {
			if j == maxTicksUnit {
				out.N1 = out.N2
				j = 0
			}
			y0.Set(p.getY0(&x, &y, po, poUp))
			f.Mul(p.Extra.A, &y0).Mul(&f, po).Div(&f, poUp).Mul(&f, po).Div(&f, number.Number_1e18)
			g.Mul(p.Extra.Aminus1, &y0).Mul(&g, poUp).Div(&g, po)
			inv.Add(&f, &x).Mul(&inv, temp.Add(&g, &y))
			dynamicFee.Set(maxUint256(p.getDynamicFee(po, poUp), p.Extra.SwapFee))
		}

		antifee.Div(
			Number_1e36,
			temp.Sub(number.Number_1e18, minUint256(&dynamicFee, tenPow18Minus1)),
		)

		if j != maxTicksUnit {
			var tick uint256.Int
			tick.Set(&y)
			if pump {
				tick.Set(&x)
			}
			out.TicksIn = append(out.TicksIn, tick)
		}

		// Need this to break if price is too far
		var pRatio uint256.Int
		pRatio.Mul(poUp, number.Number_1e18).Div(&pRatio, po)

		if pump {
			if y.Sign() != 0 {
				if g.Sign() != 0 {
					var xDest, dx uint256.Int
					xDest.Div(&inv, &g).Sub(&xDest, &f).Sub(&xDest, &x)
					dx.Mul(&xDest, &antifee).Div(&dx, number.Number_1e18)

					if dx.Cmp(&inAmountLeft) >= 0 {
						// This is the last band
						xDest.Mul(&inAmountLeft, number.Number_1e18).Div(&xDest, &antifee)

						out.LastTickJ.Div(&inv, temp.Add(&x, &xDest).Add(&temp, &f)).
							Sub(&out.LastTickJ, &g).Add(&out.LastTickJ, number.Number_1)
						if out.LastTickJ.Cmp(&y) > 0 {
							out.LastTickJ.Set(&y)
						}

						xDest.Sub(&inAmountLeft, &xDest).Mul(&xDest, p.Extra.AdminFee).Div(&xDest, number.Number_1e18)
						x.Add(&x, &inAmountLeft)

						// Round down the output
						out.OutAmount.Add(&out.OutAmount, &y).Sub(&out.OutAmount, &out.LastTickJ)
						out.TicksIn[j].Sub(&x, &xDest)
						out.InAmount.Set(inAmount)
						out.AdminFee.Add(&out.AdminFee, &xDest)
						break
					} else { // We go into the next band
						// Prevents from leaving dust in the band
						dx.Set(maxUint256(&dx, number.Number_1))

						xDest.Sub(&dx, &xDest).Mul(&xDest, p.Extra.AdminFee).Div(&xDest, number.Number_1e18)
						inAmountLeft.Sub(&inAmountLeft, &dx)

						out.TicksIn[j].Add(&x, &dx).Sub(&out.TicksIn[j], &xDest)
						out.InAmount.Add(&out.InAmount, &dx)
						out.OutAmount.Add(&out.OutAmount, &y)
						out.AdminFee.Add(&out.AdminFee, &xDest)
					}
				}
			}

			if i != maxTicks+maxSkipTicks-1 {
				if out.N2 == maxBand {
					break
				}
				if j == maxTicksUnit-1 {
					break
				}
				if pRatio.Lt(temp.Div(Number_1e36, p.Extra.MaxOracleDnPow)) {
					break
				}
				out.N2 += 1
				poUp.Mul(poUp, p.Extra.Aminus1).Div(poUp, p.Extra.A)
				x.Set(number.Zero)
				y.Set(p.getBandY(out.N2))
			}
		} else { // dump
			if x.Sign() != 0 {
				if f.Sign() != 0 {
					var yDest, dy uint256.Int
					yDest.Div(&inv, &f).Sub(&yDest, &g).Sub(&yDest, &y)
					dy.Mul(&yDest, &antifee).Div(&dy, number.Number_1e18)

					if dy.Cmp(&inAmountLeft) >= 0 {
						// This is the last band
						yDest.Mul(&inAmountLeft, number.Number_1e18).Div(&yDest, &antifee)

						out.LastTickJ.Div(&inv, temp.Add(&y, &yDest).Add(&temp, &g)).
							Sub(&out.LastTickJ, &f).Add(&out.LastTickJ, number.Number_1)
						if out.LastTickJ.Cmp(&x) > 0 {
							out.LastTickJ.Set(&x)
						}

						yDest.Sub(&inAmountLeft, &yDest).Mul(&yDest, p.Extra.AdminFee).Div(&yDest, number.Number_1e18)
						y.Add(&y, &inAmountLeft)

						out.OutAmount.Add(&out.OutAmount, &x).Sub(&out.OutAmount, &out.LastTickJ)
						out.TicksIn[j].Sub(&y, &yDest)
						out.InAmount.Set(inAmount)
						out.AdminFee.Add(&out.AdminFee, &yDest)
						break
					} else { // We go into the next band
						// Prevents from leaving dust in the band
						dy.Set(maxUint256(&dy, number.Number_1))

						yDest.Sub(&dy, &yDest).Mul(&yDest, p.Extra.AdminFee).Div(&yDest, number.Number_1e18)
						inAmountLeft.Sub(&inAmountLeft, &dy)

						out.TicksIn[j].Add(&y, &dy).Sub(&out.TicksIn[j], &yDest)
						out.InAmount.Add(&out.InAmount, &dy)
						out.OutAmount.Add(&out.OutAmount, &x)
						out.AdminFee.Add(&out.AdminFee, &yDest)
					}
				}
			}
			if i != maxTicks+maxSkipTicks-1 {
				if out.N2 == minBand {
					break
				}
				if j == maxTicksUnit-1 {
					break
				}
				if pRatio.Gt(p.Extra.MaxOracleDnPow) {
					// Don't allow to be away by more than ~50 ticks
					break
				}
				out.N2 -= 1
				poUp.Mul(poUp, p.Extra.A).Div(poUp, p.Extra.Aminus1)
				x.Set(p.getBandX(out.N2))
				y.Set(number.Zero)
			}
		}

		if j != maxTicksUnit {
			j += 1
		}
	}

	inPrecisionMinus1 := new(uint256.Int).Sub(inPrecision, number.Number_1)
	out.InAmount.Add(&out.InAmount, inPrecisionMinus1).Div(&out.InAmount, inPrecision).Mul(&out.InAmount, inPrecision)
	out.OutAmount.Div(&out.OutAmount, outPrecision).Mul(&out.OutAmount, outPrecision)

	return out, nil
}

func (p *PoolSimulator) getY0(x, y, po, poUp *uint256.Int) *uint256.Int {
	var b, temp uint256.Int
	if x.Sign() != 0 {
		b.Mul(poUp, p.Extra.Aminus1).Mul(&b, x).Div(&b, po)
	}
	if y.Sign() != 0 {
		temp.Mul(p.Extra.A, po).Mul(&temp, po).Div(&temp, poUp).Mul(&temp, y).Div(&temp, number.Number_1e18)
		b.Add(&b, &temp)
	}
	var numerator, denominator uint256.Int
	if x.Sign() > 0 && y.Sign() > 0 {
		var D uint256.Int
		D.Mul(number.Number_4, p.Extra.A).Mul(&D, po).Mul(&D, y).Div(&D, number.Number_1e18)
		D.Mul(&D, x)

		D.Add(&D, temp.Mul(&b, &b))

		numerator.Add(&b, temp.Sqrt(&D)).Mul(&numerator, number.Number_1e18)
		denominator.Mul(p.Extra.A, number.Number_2).Mul(&denominator, po)
	} else {
		numerator.Mul(&b, number.Number_1e18)
		denominator.Mul(p.Extra.A, po)
	}

	return numerator.Div(&numerator, &denominator)
}

func (p *PoolSimulator) pOracleUp(n int64) (*uint256.Int, error) {
	var power int256.Int
	power.SetInt64(-n).Mul(&power, p.Extra.LogARatio)

	expPower, err := wadExp(&power)
	if err != nil {
		return nil, err
	}

	var ret uint256.Int
	return ret.Mul(p.Extra.BasePrice, expPower).Div(&ret, number.Number_1e18), nil
}

func (p *PoolSimulator) getDynamicFee(po, poUp *uint256.Int) *uint256.Int {
	var ret, pcd, pcu uint256.Int
	pcd.Mul(po, po).Div(&pcd, poUp).Mul(&pcd, po).Div(&pcd, poUp)
	pcu.Mul(&pcd, p.Extra.A).Div(&pcu, p.Extra.Aminus1).Mul(&pcu, p.Extra.A).Div(&pcu, p.Extra.Aminus1)

	if po.Lt(&pcd) {
		ret.Sub(&pcd, po).Mul(&ret, tenPow18Div4).Div(&ret, &pcd)
	} else if po.Gt(&pcu) {
		ret.Sub(po, &pcu).Mul(&ret, tenPow18Div4).Div(&ret, po)
	}

	return &ret
}

func (p *PoolSimulator) getBandX(index int64) *uint256.Int {
	if _, ok := p.Extra.BandsX[index]; ok {
		return p.Extra.BandsX[index]
	}
	return uint256.NewInt(0)
}

func (p *PoolSimulator) getBandY(index int64) *uint256.Int {
	if _, ok := p.Extra.BandsY[index]; ok {
		return p.Extra.BandsY[index]
	}
	return uint256.NewInt(0)
}
