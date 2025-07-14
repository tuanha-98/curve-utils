package curvev1meta

import (
	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/entities"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"

	curvev1 "github.com/tuanha-98/curve-utils/internal/entities/pool/curve-v1"
	token "github.com/tuanha-98/curve-utils/internal/entities/token"
)

type Static = curvev1.Static
type Extra = curvev1.Extra

type BasePool interface {
	GetNumTokens() int
	GetFeeInfo() entities.FeeInfo
	GetDy(i, j int, dx, dy *uint256.Int) error
	XpMem(rates []uint256.Int, reserves []uint256.Int) []uint256.Int
	CalculateTokenAmount(amounts []uint256.Int, deposit bool, mintAmount *uint256.Int) error
	CalculateWithdrawOneCoin(tokenAmount *uint256.Int, index int, dy *uint256.Int, fee *uint256.Int) error
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
	BasePool          BasePool
}

func NewPool(
	entityPool entities.Pool,
	basePool entities.Pool,
) (*PoolSimulator, error) {
	poolSim, err := curvev1.NewPool(entityPool)
	if err != nil {
		return nil, err
	}

	basePoolSim, err := curvev1.NewPool(basePool)
	if err != nil {
		return nil, err
	}

	pool := &PoolSimulator{
		Address:       poolSim.Address,
		Exchange:      poolSim.Exchange,
		Reserves:      poolSim.Reserves,
		LpSupply:      poolSim.LpSupply,
		NumTokens:     poolSim.NumTokens,
		NumTokensU256: poolSim.NumTokensU256,
		Tokens:        poolSim.Tokens,
		Static:        poolSim.Static,
		Extra:         poolSim.Extra,
		BasePool:      basePoolSim,
	}
	return pool, nil
}

func (p *PoolSimulator) FeeCalculate(dy, fee *uint256.Int) {
	fee.Div(
		number.SafeMul(dy, p.Extra.SwapFee),
		FeeDenominator,
	)
}

func (p *PoolSimulator) GetDyUnderlying(
	i, j int, _dx *uint256.Int,
	//out put
	dy *uint256.Int,
) error {
	var maxCoins = p.NumTokens - 1
	var baseNCoins = p.BasePool.GetNumTokens()
	xp := p.BasePool.XpMem(p.Extra.Rates, p.Reserves)

	var base_i = i - maxCoins
	var base_j = j - maxCoins
	var meta_i = maxCoins
	var meta_j = maxCoins
	if base_i < 0 {
		meta_i = i
	}
	if base_j < 0 {
		meta_j = j
	}

	var x *uint256.Int
	var amountOut, withdrawDy, withdrawDyFee uint256.Int
	if base_i < 0 {
		x = number.SafeAdd(&xp[i], number.SafeMul(_dx, number.Div(&p.Extra.Rates[i], Precision)))
	} else {
		if base_j < 0 {
			var base_inputs = make([]uint256.Int, baseNCoins)
			for k := 0; k < baseNCoins; k++ {
				base_inputs[k].Clear()
			}
			base_inputs[base_i].Set(_dx)
			var err = p.BasePool.CalculateTokenAmount(base_inputs, true, &amountOut)
			if err != nil {
				return err
			}
			x = number.Div(number.SafeMul(&amountOut, &p.Extra.Rates[maxCoins]), Precision)
			feeInfo := p.BasePool.GetFeeInfo()
			x = number.Sub(x, number.Div(number.Mul(x, &feeInfo.SwapFee), number.Mul(number.Number_2, FeeDenominator)))
			x = number.SafeAdd(&xp[maxCoins], x)
		} else {
			if err := p.BasePool.GetDy(base_i, base_j, _dx, dy); err != nil {
				return err
			}
			return nil
		}
	}

	if err := p._getY(meta_i, meta_j, x, xp, &amountOut); err != nil {
		return err
	}

	number.SafeSubZ(&xp[meta_j], &amountOut, dy)

	if dy.Sign() <= 0 {
		return ErrZero
	}
	dy.SubUint64(dy, 1)

	var fee uint256.Int
	p.FeeCalculate(dy, &fee)
	if dy.Cmp(&fee) < 0 {
		return ErrInvalidReserve
	}

	dy.Sub(dy, &fee)

	if base_j < 0 {
		dy.Set(number.Div(number.Mul(dy, Precision), &p.Extra.Rates[j]))
	} else {
		if err := p.BasePool.CalculateWithdrawOneCoin(number.Div(number.Mul(dy, Precision), &p.Extra.Rates[maxCoins]), base_j, &withdrawDy, &withdrawDyFee); err != nil {
			return err
		}
		dy.Set(&withdrawDy)
	}

	return nil
}
