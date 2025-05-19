package stablemetang

import (
	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

func (p *Pool) GetDyUnderlying(
	i, j int, _dx *uint256.Int,
	//out put
	dy *uint256.Int,
) error {

	var baseNCoins = len(p.basePool.GetTokens())
	xp := p.basePool.XpMem(p.Extra.RateMultipliers, p.Reserves)

	var base_i = i - MAX_METAPOOL_COIN_INDEX
	var base_j = j - MAX_METAPOOL_COIN_INDEX

	input_is_base_coin := base_i >= 0
	output_is_base_coin := base_j >= 0
	if input_is_base_coin && output_is_base_coin {
		// should be rejected at the outer level already
		return ErrAllBasePoolTokens
	}
	if !input_is_base_coin && !output_is_base_coin {
		// all meta coins, should not happen (should be redirected to GetDy instead)
		return ErrAllMetaPoolTokens
	}

	var tokenInIndex, tokenOutIndex int

	if output_is_base_coin {
		tokenInIndex = i
		tokenOutIndex = MAX_METAPOOL_COIN_INDEX
	} else {
		tokenInIndex = MAX_METAPOOL_COIN_INDEX
		tokenOutIndex = j
	}

	var x *uint256.Int
	var amountOut, adminFee, withdrawDy, withdrawDyFee uint256.Int
	if output_is_base_coin {
		x = number.SafeAdd(&xp[i], number.SafeMul(_dx, number.Div(&p.Extra.RateMultipliers[i], Precision)))
	} else {
		addLiquidityAmounts := make([]uint256.Int, baseNCoins)
		feeAmounts := make([]uint256.Int, baseNCoins)
		for k := 0; k < baseNCoins; k++ {
			addLiquidityAmounts[k].Clear()
		}
		addLiquidityAmounts[base_i].Set(_dx)

		var mintAmount uint256.Int

		if err := p.basePool.CalculateTokenAmount(addLiquidityAmounts[:baseNCoins], true, &mintAmount, feeAmounts[:baseNCoins]); err != nil {
			return err
		}

		x = number.Div(number.SafeMul(&mintAmount, &p.Extra.RateMultipliers[MAX_METAPOOL_COIN_INDEX]), Precision)
		number.SafeAddZ(x, &xp[MAX_METAPOOL_COIN_INDEX], x)
	}

	err := p.GetDyByX(tokenInIndex, tokenOutIndex, x, xp, &amountOut, &adminFee)
	if err != nil {
		return err
	}
	if output_is_base_coin {
		err = p.basePool.CalculateWithdrawOneCoin(&amountOut, base_j, &withdrawDy, &withdrawDyFee)
		if err != nil {
			return err
		}
		dy.Set(&withdrawDy)
	} else {
		dy.Set(&amountOut)
	}
	return nil
}
