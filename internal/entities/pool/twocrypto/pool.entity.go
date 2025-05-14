package twocrypto

import (
	"fmt"

	c "github.com/tuanha-98/curve-utils/internal/constants"
	m "github.com/tuanha-98/curve-utils/internal/utils/maths"
)

var (
	N_COINS      = new(m.Uint256).SetUint64(2)
	A_MULTIPLIER = new(m.Uint256).SetUint64(10000)
	MIN_GAMMA    = c.Wei10
	MAX_GAMMA    = new(m.Uint256).Mul(new(m.Uint256).SetUint64(2), c.Wei16)

	MIN_A = new(m.Uint256).Div(new(m.Uint256).Mul(new(m.Uint256).Exp(N_COINS, N_COINS), A_MULTIPLIER), new(m.Uint256).SetUint64(10))
	MAX_A = new(m.Uint256).Mul(new(m.Uint256).Mul(new(m.Uint256).Exp(N_COINS, N_COINS), A_MULTIPLIER), new(m.Uint256).SetUint64(100000))
)

type Pool struct {
	Address, Exchange string
	A                 *m.Uint256 // amplification coefficient
	Gamma             *m.Uint256 // gamma value of the pool
	D                 *m.Uint256 // D value of the pool
	PriceScale        *m.Uint256 // price scale
	FutureAGammaTime  *m.Uint256 // future A gamma time
	MidFee            *m.Uint256
	OutFee            *m.Uint256
	FeeGamma          *m.Uint256
	Coins             []string     // address of coins in the pool in correct order
	XP                []*m.Uint256 // balance of each coin in the pool
	Precisions        []*m.Uint256 // precision of each coin in the pool
}

func NewCryptoSwapPool(address, exchange string, a, gamma, d, priceScale, futureAGammaTime, midFee, outFee, feeGamma *m.Uint256, coins []string, xp, precisions []*m.Uint256) *Pool {
	return &Pool{
		Address:          address,
		Exchange:         exchange,
		A:                a,
		Gamma:            gamma,
		D:                d,
		PriceScale:       priceScale,
		FutureAGammaTime: futureAGammaTime,
		MidFee:           midFee,
		OutFee:           outFee,
		FeeGamma:         feeGamma,
		Coins:            coins,
		XP:               xp,
		Precisions:       precisions,
	}
}

func (p *Pool) GetDy(i, j int, dx *m.Uint256) (*m.Uint256, error) {
	xp := make([]*m.Uint256, len(p.XP))
	copy(xp, p.XP)
	D := new(m.Uint256).Set(p.D)
	if p.FutureAGammaTime.GtUint64(0) {
		var err error
		D, err = p.newtonD(p.A, p.FeeGamma, xp)
		if err != nil {
			return nil, err
		}
	}
	xp[i] = new(m.Uint256).Add(xp[i], dx)
	xp[0] = new(m.Uint256).Mul(xp[0], p.Precisions[0])
	xp[1] = new(m.Uint256).Div(
		new(m.Uint256).Mul(
			new(m.Uint256).Mul(xp[1], p.PriceScale),
			p.Precisions[1],
		),
		c.Precision,
	)

	y, err := p.newtonY(p.A, p.Gamma, D, xp, j)
	if err != nil {
		return nil, err
	}
	dy := new(m.Uint256).Sub(new(m.Uint256).Sub(xp[j], y), new(m.Uint256).SetUint64(1))
	xp[j] = new(m.Uint256).Set(y)

	if j > 0 {
		dy = new(m.Uint256).Div(
			new(m.Uint256).Mul(
				dy,
				c.Precision,
			),
			p.PriceScale,
		)
	} else {
		dy = new(m.Uint256).Div(
			dy,
			p.Precisions[0],
		)
	}

	fee := p.feeCalculate(xp)
	dy = new(m.Uint256).Sub(
		dy,
		new(m.Uint256).Div(
			new(m.Uint256).Mul(
				fee,
				dy,
			),
			c.FeeDenominator,
		),
	)

	return dy, nil
}

func (p *Pool) newtonD(ann, gamma *m.Uint256, x_unsorted []*m.Uint256) (*m.Uint256, error) {
	if (ann.Cmp(new(m.Uint256).Sub(MIN_A, new(m.Uint256).SetUint64(1))) <= 0) || (ann.Cmp(new(m.Uint256).Add(MAX_A, new(m.Uint256).SetUint64(1))) >= 0) {
		return nil, fmt.Errorf("dev: unsafe values ann")
	}
	if (gamma.Cmp(new(m.Uint256).Sub(MIN_GAMMA, new(m.Uint256).SetUint64(1))) <= 0) || (ann.Cmp(new(m.Uint256).Add(MAX_GAMMA, new(m.Uint256).SetUint64(1))) >= 0) {
		return nil, fmt.Errorf("dev: unsafe values gamma")
	}

	x := make([]*m.Uint256, len(x_unsorted))
	copy(x, x_unsorted)

	if x[0].Cmp(x[1]) < 0 {
		x[0], x[1] = x[1], x[0]
	}

	if x[0].Cmp(new(m.Uint256).Sub(c.Wei9, new(m.Uint256).SetUint64(1))) <= 0 || x[0].Cmp(new(m.Uint256).Add(new(m.Uint256).Mul(c.Wei15, c.Wei18), new(m.Uint256).SetUint64(1))) >= 0 {
		return nil, fmt.Errorf("dev: unsafe values x[0]")
	}

	if new(m.Uint256).Div(new(m.Uint256).Mul(x[1], c.Wei18), x[0]).Cmp(new(m.Uint256).Sub(c.Wei14, new(m.Uint256).SetUint64(1))) > 0 {
		return nil, fmt.Errorf("dev: unsafe values x[1] (input)")
	}
	gm2c, err := m.GeometricMean2Coins(x, false)
	if err != nil {
		return nil, err
	}
	D := new(m.Uint256).Mul(N_COINS, gm2c)
	S := new(m.Uint256).Add(x[0], x[1])

	for i := 0; i < 255; i++ {
		D_prev := new(m.Uint256).Set(D)
		k0 := new(m.Uint256).Div(new(m.Uint256).Mul(new(m.Uint256).Div(new(m.Uint256).Mul(new(m.Uint256).Mul(c.Wei18, new(m.Uint256).Exp(N_COINS, new(m.Uint256).SetUint64(2))), x[0]), D), x[1]), D)
		g1k0 := new(m.Uint256).Add(gamma, c.Wei18)

		if g1k0.Cmp(k0) > 0 {
			g1k0 = new(m.Uint256).Sub(new(m.Uint256).Sub(g1k0, k0), new(m.Uint256).SetUint64(1))
		} else {
			g1k0 = new(m.Uint256).Sub(new(m.Uint256).Sub(k0, g1k0), new(m.Uint256).SetUint64(1))
		}

		mul1 := new(m.Uint256).Div(
			new(m.Uint256).Mul(
				new(m.Uint256).Mul(
					new(m.Uint256).Div(
						new(m.Uint256).Mul(
							new(m.Uint256).Div(
								new(m.Uint256).Mul(
									c.Wei18,
									D,
								),
								gamma,
							),
							g1k0,
						),
						gamma,
					),
					g1k0,
				),
				A_MULTIPLIER,
			),
			ann,
		)

		mul2 := new(m.Uint256).Div(
			new(m.Uint256).Mul(
				new(m.Uint256).Mul(
					new(m.Uint256).Mul(
						new(m.Uint256).SetUint64(2),
						c.Wei18,
					),
					N_COINS,
				),
				k0,
			),
			g1k0,
		)

		negFPrime := new(m.Uint256).Sub(
			new(m.Uint256).Add(
				new(m.Uint256).Add(
					S,
					new(m.Uint256).Div(
						new(m.Uint256).Mul(
							S,
							mul2,
						),
						c.Wei18,
					),
				),
				new(m.Uint256).Div(
					new(m.Uint256).Mul(
						mul1,
						N_COINS,
					),
					k0,
				),
			),
			new(m.Uint256).Div(
				new(m.Uint256).Mul(
					mul2,
					D,
				),
				c.Wei18,
			),
		)

		DPlus := new(m.Uint256).Mul(
			D,
			new(m.Uint256).Div(
				new(m.Uint256).Add(
					negFPrime,
					S,
				),
				negFPrime,
			),
		)
		DMinus := new(m.Uint256).Mul(
			D,
			new(m.Uint256).Div(
				D,
				negFPrime,
			),
		)

		if c.Wei18.Cmp(k0) > 0 {
			DMinus = new(m.Uint256).Add(
				DMinus,
				new(m.Uint256).Mul(
					new(m.Uint256).Div(
						new(m.Uint256).Mul(
							D,
							new(m.Uint256).Div(
								mul1,
								negFPrime,
							),
						),
						c.Wei18,
					),
					new(m.Uint256).Div(
						new(m.Uint256).Sub(
							c.Wei18,
							k0,
						),
						k0,
					),
				),
			)
		} else {
			DMinus = new(m.Uint256).Sub(
				DMinus,
				new(m.Uint256).Mul(
					new(m.Uint256).Div(
						new(m.Uint256).Mul(
							D,
							new(m.Uint256).Div(
								mul1,
								negFPrime,
							),
						),
						c.Wei18,
					),
					new(m.Uint256).Div(
						new(m.Uint256).Sub(
							k0,
							c.Wei18,
						),
						k0,
					),
				),
			)
		}

		if DPlus.Cmp(DMinus) > 0 {
			D = new(m.Uint256).Sub(DPlus, DMinus)
		} else {
			D = new(m.Uint256).Div(new(m.Uint256).Sub(DMinus, DPlus), new(m.Uint256).SetUint64(2))
		}

		diff := new(m.Uint256).SetUint64(0)
		if D.Cmp(D_prev) > 0 {
			diff = new(m.Uint256).Sub(D, D_prev)
		} else {
			diff = new(m.Uint256).Sub(D_prev, D)
		}

		if new(m.Uint256).Mul(diff, c.Wei14).Cmp(m.MaxUint256(c.Wei16, D)) < 0 {
			for _, _x := range x {
				frac := new(m.Uint256).Div(new(m.Uint256).Mul(_x, c.Wei18), D)
				if frac.Cmp(new(m.Uint256).Sub(c.Wei16, new(m.Uint256).SetUint64(1))) <= 0 || frac.Cmp(new(m.Uint256).Add(c.Wei20, new(m.Uint256).SetUint64(1))) >= 0 {
					return nil, fmt.Errorf("dev: unsafe values x[i]")
				}
				return D, nil
			}
		}
	}

	return nil, fmt.Errorf("dev: newtonD did not converge")
}

func (p *Pool) newtonY(ann, gamma, D *m.Uint256, x []*m.Uint256, i int) (*m.Uint256, error) {
	if (ann.Cmp(new(m.Uint256).Sub(MIN_A, new(m.Uint256).SetUint64(1))) <= 0) || (ann.Cmp(new(m.Uint256).Add(MAX_A, new(m.Uint256).SetUint64(1))) >= 0) {
		return nil, fmt.Errorf("dev: unsafe values ann")
	}
	if (gamma.Cmp(new(m.Uint256).Sub(MIN_GAMMA, new(m.Uint256).SetUint64(1))) <= 0) || (ann.Cmp(new(m.Uint256).Add(MAX_GAMMA, new(m.Uint256).SetUint64(1))) >= 0) {
		return nil, fmt.Errorf("dev: unsafe values gamma")
	}
	if (D.Cmp(new(m.Uint256).Sub(c.Wei17, new(m.Uint256).SetUint64(1))) <= 0) || (D.Cmp(new(m.Uint256).Add(new(m.Uint256).Mul(c.Wei15, c.Wei18), new(m.Uint256).SetUint64(1))) >= 0) {
		return nil, fmt.Errorf("dev: unsafe values D")
	}

	xj := new(m.Uint256).Set(x[1-i])
	y := new(m.Uint256).Div(
		new(m.Uint256).Exp(
			D,
			new(m.Uint256).SetUint64(2),
		),
		new(m.Uint256).Mul(
			xj,
			new(m.Uint256).Exp(
				N_COINS,
				new(m.Uint256).SetUint64(2),
			),
		),
	)
	k0i := new(m.Uint256).Div(
		new(m.Uint256).Mul(
			new(m.Uint256).Mul(
				c.Wei18,
				N_COINS,
			),
			xj,
		),
		D,
	)

	if k0i.Cmp(new(m.Uint256).Sub(new(m.Uint256).Mul(c.Wei16, N_COINS), new(m.Uint256).SetUint64(1))) <= 0 || k0i.Cmp(new(m.Uint256).Add(new(m.Uint256).Mul(c.Wei20, N_COINS), new(m.Uint256).SetUint64(1))) >= 0 {
		return nil, fmt.Errorf("dev: unsafe values x[i]")
	}

	convergenceLimit := m.MaxUint256(m.MaxUint256(new(m.Uint256).Div(xj, c.Wei14), new(m.Uint256).Div(D, c.Wei14)), new(m.Uint256).SetUint64(100))

	for j := 0; j < 255; j++ {
		yPrev := new(m.Uint256).Set(y)
		k0 := new(m.Uint256).Div(
			new(m.Uint256).Mul(
				new(m.Uint256).Mul(
					k0i,
					y,
				),
				N_COINS,
			),
			D,
		)
		S := new(m.Uint256).Add(xj, y)
		g1k0 := new(m.Uint256).Add(gamma, c.Wei18)
		if g1k0.Cmp(k0) > 0 {
			g1k0 = new(m.Uint256).Add(new(m.Uint256).Sub(g1k0, k0), new(m.Uint256).SetUint64(1))
		} else {
			g1k0 = new(m.Uint256).Add(new(m.Uint256).Sub(k0, g1k0), new(m.Uint256).SetUint64(1))
		}
		mul1 := new(m.Uint256).Div(
			new(m.Uint256).Mul(
				new(m.Uint256).Mul(
					new(m.Uint256).Div(
						new(m.Uint256).Mul(
							new(m.Uint256).Div(
								new(m.Uint256).Mul(
									c.Wei18,
									D,
								),
								gamma,
							),
							g1k0,
						),
						gamma,
					),
					g1k0,
				),
				A_MULTIPLIER,
			),
			ann,
		)

		mul2 := new(m.Uint256).Add(
			c.Wei18,
			new(m.Uint256).Div(
				new(m.Uint256).Mul(
					new(m.Uint256).Mul(
						new(m.Uint256).SetUint64(2),
						c.Wei18,
					),
					k0,
				),
				g1k0,
			),
		)

		yfprime := new(m.Uint256).Add(
			new(m.Uint256).Add(
				new(m.Uint256).Mul(
					c.Wei18,
					y,
				),
				new(m.Uint256).Mul(
					S,
					mul2,
				),
			),
			mul1,
		)

		dyfprime := new(m.Uint256).Mul(D, mul2)

		if yfprime.Cmp(dyfprime) < 0 {
			y = new(m.Uint256).Div(
				yPrev,
				new(m.Uint256).SetUint64(2),
			)
			continue
		} else {
			yfprime = new(m.Uint256).Sub(yfprime, dyfprime)
		}
		fprime := new(m.Uint256).Div(yfprime, y)
		yMinus := new(m.Uint256).Div(mul1, fprime)
		yPlus := new(m.Uint256).Add(
			new(m.Uint256).Div(
				new(m.Uint256).Add(
					yfprime,
					new(m.Uint256).Mul(
						c.Wei18,
						D,
					),
				),
				fprime,
			),
			new(m.Uint256).Div(
				new(m.Uint256).Mul(
					yMinus,
					c.Wei18,
				),
				k0,
			),
		)
		yMinus = new(m.Uint256).Add(
			yMinus,
			new(m.Uint256).Div(
				new(m.Uint256).Mul(
					c.Wei18,
					S,
				),
				fprime,
			),
		)

		if yPlus.Cmp(yMinus) < 0 {
			y = new(m.Uint256).Div(yPrev, new(m.Uint256).SetUint64(2))
		} else {
			y = new(m.Uint256).Sub(yPlus, yMinus)
		}
		diff := new(m.Uint256).SetUint64(0)

		if y.Cmp(yPrev) > 0 {
			diff = new(m.Uint256).Sub(y, yPrev)
		} else {
			diff = new(m.Uint256).Sub(yPrev, y)
		}

		if diff.Cmp(m.MaxUint256(convergenceLimit, new(m.Uint256).Div(y, c.Wei14))) < 0 {
			frac := new(m.Uint256).Div(new(m.Uint256).Mul(y, c.Wei18), D)
			if frac.Cmp(new(m.Uint256).Sub(c.Wei16, new(m.Uint256).SetUint64(1))) <= 0 || frac.Cmp(new(m.Uint256).Add(c.Wei20, new(m.Uint256).SetUint64(1))) >= 0 {
				return nil, fmt.Errorf("dev: unsafe values y")
			}
			return y, nil
		}
	}
	return nil, fmt.Errorf("dev: newtonY did not converge")
}

func (p *Pool) feeCalculate(xp []*m.Uint256) *m.Uint256 {
	f := new(m.Uint256).Add(xp[0], xp[1])
	numerator := new(m.Uint256).Mul(p.FeeGamma, c.Wei18)
	denominator := new(m.Uint256).Sub(
		new(m.Uint256).Add(
			p.FeeGamma,
			c.Wei18,
		),
		new(m.Uint256).Div(
			new(m.Uint256).Mul(
				new(m.Uint256).Div(
					new(m.Uint256).Mul(
						new(m.Uint256).Mul(
							c.Wei18,
							new(m.Uint256).Exp(
								N_COINS,
								N_COINS,
							),
						),
						xp[0],
					),
					f,
				),
				xp[1],
			),
			f,
		),
	)
	f = new(m.Uint256).Div(numerator, denominator)

	return new(m.Uint256).Div(
		new(m.Uint256).Add(
			new(m.Uint256).Mul(
				p.MidFee,
				f,
			),
			new(m.Uint256).Mul(
				p.OutFee,
				new(m.Uint256).Sub(
					c.Wei18,
					f,
				),
			),
		),
		c.Wei18,
	)
}
