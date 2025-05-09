package pool

import (
	"fmt"

	c "github.com/tuanha-98/curve-utils/internal/constants"
	m "github.com/tuanha-98/curve-utils/internal/utils/maths"
)

type Pool struct {
	Address, Exchange   string
	Fee                 *m.Uint256
	A                   *m.Uint256   // amplification coefficient
	A_precision         *m.Uint256   // precision of A
	Coins               []string     // address of coins in the pool in correct order
	XP                  []*m.Uint256 // balance of each coin in the pool
	Rates               []*m.Uint256 // rate of each coin in the NG pool
	OffPegFeeMultiplier *m.Uint256   // multiplier for off-peg in fee NG pool
}

func NewStableSwapPool(address, exchange string, fee *m.Uint256, a, a_precision *m.Uint256, coins []string, xp []*m.Uint256, rates []*m.Uint256, offPegFeeMultiplier *m.Uint256) *Pool {
	return &Pool{
		Address:             address,
		Exchange:            exchange,
		Fee:                 fee,
		A:                   a,
		A_precision:         a_precision,
		Coins:               coins,
		XP:                  xp,
		Rates:               rates,
		OffPegFeeMultiplier: offPegFeeMultiplier,
	}
}

func (p *Pool) getD(xp []*m.Uint256, amp *m.Uint256) (*m.Uint256, error) {
	n := new(m.Uint256).SetUint64(uint64(len(xp)))
	ann := new(m.Uint256).Mul(amp, n)
	S := new(m.Uint256).SetUint64(0)
	for _, x := range xp {
		S.Add(S, x)
	}

	if S.IsZero() {
		return new(m.Uint256).SetUint64(0), nil
	}

	D := new(m.Uint256).Set(S)
	for i := 0; i < 255; i++ {
		D_P := new(m.Uint256).Set(D)
		for _, x := range xp {
			D_P = new(m.Uint256).Div(new(m.Uint256).Mul(D_P, D), new(m.Uint256).Mul(x, n))
		}
		D_prev := new(m.Uint256).Set(D)
		numerator := new(m.Uint256).Add(new(m.Uint256).Div(new(m.Uint256).Mul(ann, S), p.A_precision), new(m.Uint256).Mul(D_P, n))
		denominator := new(m.Uint256).Add(
			new(m.Uint256).Div(new(m.Uint256).Mul(new(m.Uint256).Sub(ann, p.A_precision), D), p.A_precision),
			new(m.Uint256).Mul(new(m.Uint256).Add(n, new(m.Uint256).SetUint64(1)), D_P))
		D = new(m.Uint256).Div(new(m.Uint256).Mul(numerator, D), denominator)

		if D.Cmp(D_prev) > 0 {
			if new(m.Uint256).Sub(D, D_prev).Cmp(new(m.Uint256).SetUint64(1)) <= 0 {
				return D, nil
			}
		} else {
			if new(m.Uint256).Sub(D_prev, D).Cmp(new(m.Uint256).SetUint64(1)) <= 0 {
				return D, nil
			}
		}
	}
	return nil, fmt.Errorf("getD did not converge")
}

func (p *Pool) getY(i, j int, x *m.Uint256, xp []*m.Uint256, amp *m.Uint256) (*m.Uint256, error) {
	n := len(xp)
	if i == j {
		return nil, fmt.Errorf("i and j must be different")
	}
	if i < 0 || i >= n || j < 0 || j >= n {
		return nil, fmt.Errorf("index out of bounds")
	}

	N := new(m.Uint256).SetUint64(uint64(n))
	D, err := p.getD(xp, amp)

	if err != nil {
		return nil, err
	}

	Ann := new(m.Uint256).Mul(amp, N)
	c := new(m.Uint256).Set(D)
	S_ := new(m.Uint256).SetUint64(0)

	for k := 0; k < n; k++ {
		var _x = new(m.Uint256).SetUint64(0)
		if k == i {
			_x = new(m.Uint256).Set(x)
		} else if k == j {
			_x = new(m.Uint256).SetUint64(0)
		} else {
			_x = new(m.Uint256).Set(xp[k])
		}

		if k != j {
			S_ = new(m.Uint256).Add(S_, _x)
			c = new(m.Uint256).Mul(c, D)
			c = new(m.Uint256).Div(c, new(m.Uint256).Mul(_x, N))
		}
	}

	c = new(m.Uint256).Mul(c, D)
	c = new(m.Uint256).Mul(c, p.A_precision)
	c = new(m.Uint256).Div(c, new(m.Uint256).Mul(Ann, N))
	b := new(m.Uint256).Add(S_, new(m.Uint256).Div(new(m.Uint256).Mul(D, p.A_precision), Ann))
	y := new(m.Uint256).Set(D)

	for i := 0; i < 255; i++ {
		yPrev := new(m.Uint256).Set(y)
		numerator := new(m.Uint256).Add(new(m.Uint256).Mul(y, y), c)
		denominator := new(m.Uint256).Sub(new(m.Uint256).Add(new(m.Uint256).Mul(new(m.Uint256).SetUint64(2), y), b), D)
		y = new(m.Uint256).Div(numerator, denominator)

		if y.Cmp(yPrev) == 0 {
			return y, nil
		}

		if y.Cmp(yPrev) > 0 {
			if new(m.Uint256).Sub(y, yPrev).Cmp(new(m.Uint256).SetUint64(1)) <= 0 {
				return y, nil
			}
		} else {
			if new(m.Uint256).Sub(yPrev, y).Cmp(new(m.Uint256).SetUint64(1)) <= 0 {
				return y, nil
			}
		}
	}
	return nil, fmt.Errorf("getY did not converge")
}

func (p *Pool) isNG() bool {
	if p.Rates == nil && p.OffPegFeeMultiplier == nil {
		return false
	}
	return true
}

func (p *Pool) stableFee(dy *m.Uint256) *m.Uint256 {
	fee := new(m.Uint256).Div(new(m.Uint256).Mul(dy, p.Fee), c.FeeDenominator)
	return fee
}

func (p *Pool) dynamicFee(xpi *m.Uint256, xpj *m.Uint256) *m.Uint256 {
	if p.OffPegFeeMultiplier.Cmp(c.FeeDenominator) <= 0 {
		return p.Fee
	}

	xps2 := new(m.Uint256).Exp(new(m.Uint256).Add(xpi, xpj), new(m.Uint256).SetUint64(2))

	numerator := new(m.Uint256).Mul(p.OffPegFeeMultiplier, p.Fee)

	temp := new(m.Uint256).Sub(p.OffPegFeeMultiplier, c.FeeDenominator)
	temp = new(m.Uint256).Mul(temp, new(m.Uint256).SetUint64(4))

	temp = new(m.Uint256).Mul(temp, xpi)

	temp = new(m.Uint256).Mul(temp, xpj)

	temp = new(m.Uint256).Div(temp, xps2)

	temp = new(m.Uint256).Add(temp, c.FeeDenominator)

	return new(m.Uint256).Div(numerator, temp)
}

func (p *Pool) getXP() []*m.Uint256 {
	xp := make([]*m.Uint256, len(p.XP))
	for i, x := range p.XP {
		if p.isNG() {
			xp[i] = new(m.Uint256).Div(new(m.Uint256).Mul(x, p.Rates[i]), c.Precision)
		} else {
			xp[i] = x
		}
	}
	return xp
}

func (p *Pool) getX(xp []*m.Uint256, i int, dx *m.Uint256) *m.Uint256 {
	if p.isNG() {
		return new(m.Uint256).Add(xp[i], new(m.Uint256).Div(new(m.Uint256).Mul(dx, p.Rates[i]), c.Precision))
	} else {
		return new(m.Uint256).Add(p.XP[i], dx)
	}
}

func (p *Pool) GetDy(i, j int, dx *m.Uint256) (*m.Uint256, error) {
	xp := p.getXP()
	x := p.getX(xp, i, dx)
	y, err := p.getY(i, j, x, xp, new(m.Uint256).Mul(p.A, p.A_precision))
	if err != nil {
		return nil, err
	}
	dy := new(m.Uint256).Sub(new(m.Uint256).Sub(xp[j], y), new(m.Uint256).SetUint64(1))
	if p.isNG() {
		fee := new(m.Uint256).Div(
			new(m.Uint256).Mul(
				dy,
				p.dynamicFee(
					new(m.Uint256).Div(new(m.Uint256).Add(xp[i], x), new(m.Uint256).SetUint64(2)),
					new(m.Uint256).Div(new(m.Uint256).Add(xp[j], y), new(m.Uint256).SetUint64(2)),
				),
			),
			c.FeeDenominator,
		)
		return new(m.Uint256).Div(new(m.Uint256).Mul(new(m.Uint256).Sub(dy, fee), c.Precision), p.Rates[j]), nil
	} else {
		return new(m.Uint256).Sub(dy, p.stableFee(dy)), nil
	}
}
