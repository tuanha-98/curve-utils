package curvev1ngmeta

import (
	"github.com/holiman/uint256"
	entities "github.com/tuanha-98/curve-utils/internal/entities/pool/v1"
	curvev1 "github.com/tuanha-98/curve-utils/internal/entities/pool/v1/curve-v1"
	curvev1ng "github.com/tuanha-98/curve-utils/internal/entities/pool/v1/curve-v1-ng"

	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type Static = curvev1.Static
type Extra = curvev1.Extra

type BasePool interface {
	GetNumTokens() int
	GetFeeInfo() entities.FeeInfo
	GetDy(i, j int, dx, dy *uint256.Int) error
	XpMem(rates []uint256.Int, reserves []uint256.Int) []uint256.Int
	BaseCalculateTokenAmount(amounts []uint256.Int, deposit bool, mintAmount *uint256.Int) error
	CalculateWithdrawOneCoin(tokenAmount *uint256.Int, index int, dy *uint256.Int, fee *uint256.Int) error
	GetBasePoolType() string
}

type PoolSimulator struct {
	curvev1ng.PoolSimulator
	BasePool BasePool
}

func NewPool(
	entityPool entities.Pool,
	basePool entities.Pool,
) (*PoolSimulator, error) {
	poolSim, err := curvev1ng.NewPool(entityPool)
	if err != nil {
		return nil, err
	}
	var basePoolSim BasePool
	switch basePool.GetBasePoolType() {
	case "curvev1":
		basePoolSim, err = curvev1.NewPool(basePool)
		if err != nil {
			return nil, err
		}
	case "curvev1ng":
		basePoolSim, err = curvev1ng.NewPool(basePool)
		if err != nil {
			return nil, err
		}
	}

	pool := &PoolSimulator{
		PoolSimulator: *poolSim,
		BasePool:      basePoolSim,
	}
	return pool, nil
}

func (p *PoolSimulator) GetDyUnderlying(
	i, j int, _dx *uint256.Int,
	//out put
	dy *uint256.Int,
) error {
	var nCoins = p.NumTokens
	var maxCoins = nCoins - 1
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
		x = number.SafeAdd(&xp[i], number.Div(number.Mul(_dx, &p.Extra.Rates[i]), Precision))
	} else {
		if j == 0 {
			var base_inputs = make([]uint256.Int, baseNCoins)
			for k := 0; k < baseNCoins; k++ {
				base_inputs[k].Clear()
			}
			base_inputs[base_i].Set(_dx)
			var err = p.BasePool.BaseCalculateTokenAmount(base_inputs, true, &amountOut)
			if err != nil {
				return err
			}
			x = number.Div(number.SafeMul(&amountOut, &p.PoolSimulator.Extra.Rates[maxCoins]), Precision)
			number.SafeAddZ(x, &xp[maxCoins], x)
		} else {
			if err := p.BasePool.GetDy(base_i, base_j, _dx, dy); err != nil {
				return err
			}
			return nil
		}
	}

	err := p.PoolSimulator.GetDyByX(meta_i, meta_j, x, xp, &amountOut)
	if err != nil {
		return err
	}
	if base_j < 0 {
		dy.Set(&amountOut)
	} else {
		err = p.BasePool.CalculateWithdrawOneCoin(&amountOut, base_j, &withdrawDy, &withdrawDyFee)

		if err != nil {
			return err
		}
		dy.Set(&withdrawDy)
	}
	return nil
}
