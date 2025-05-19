package stablemetang

import (
	"github.com/holiman/uint256"
	stable "github.com/tuanha-98/curve-utils/internal/entities/pool/stable"
	stableng "github.com/tuanha-98/curve-utils/internal/entities/pool/stable-ng"

	token "github.com/tuanha-98/curve-utils/internal/entities/token"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

type BasePool interface {
	GetTokens() []token.Token
	XpMem(rates []uint256.Int, reserves []uint256.Int) []uint256.Int
	CalculateTokenAmount(amounts []uint256.Int, deposit bool, mintAmount *uint256.Int, fees []uint256.Int) error
	CalculateWithdrawOneCoin(tokenAmount *uint256.Int, index int, dy *uint256.Int, fee *uint256.Int) error
}

type Pool struct {
	stableng.Pool
	basePool BasePool
}

func NewPool(
	address, exchange, baseAddress, baseExchange string,
	reserves, baseReserves []uint256.Int,
	tokens, baseTokens []token.Token,
	a_precision, off_peg_fee_multipler, initial_a, future_a, swap_fee, admin_fee, total_supply uint256.Int,
	base_a_precision, base_off_peg_fee_multipler, base_initial_a, base_future_a, base_swap_fee, base_admin_fee, base_total_supply uint256.Int,
	rate_multipliers, base_rate_multipliers []uint256.Int,
	initial_a_time, future_a_time, base_initial_a_time, base_future_a_time int64,
) *Pool {
	numtokens := len(tokens)

	pool := &Pool{
		Pool: stableng.Pool{
			Address:       address,
			Exchange:      exchange,
			Reserves:      reserves,
			LpSupply:      total_supply,
			NumTokens:     numtokens,
			NumTokensU256: *number.SetUint64(uint64(numtokens)),
			Extra: stableng.Extra{
				APrecision:          &a_precision,
				OffPegFeeMultiplier: &off_peg_fee_multipler,
				InitialA:            &initial_a,
				FutureA:             &future_a,
				InitialATime:        initial_a_time,
				FutureATime:         future_a_time,
				SwapFee:             &swap_fee,
				AdminFee:            &admin_fee,
				RateMultipliers:     rate_multipliers,
			},
		},
	}

	if base_off_peg_fee_multipler.IsZero() {
		pool.basePool = stable.NewPool(
			baseAddress,
			baseExchange,
			base_rate_multipliers,
			baseReserves,
			baseTokens,
			base_a_precision,
			base_initial_a,
			base_future_a,
			base_swap_fee,
			base_admin_fee,
			base_total_supply,
			base_initial_a_time,
			base_future_a_time,
		)
	} else {
		pool.basePool = stableng.NewPool(
			baseAddress,
			baseExchange,
			baseReserves,
			baseTokens,
			base_a_precision,
			base_off_peg_fee_multipler,
			base_initial_a,
			base_future_a,
			base_swap_fee,
			base_admin_fee,
			base_total_supply,
			base_rate_multipliers,
			base_initial_a_time,
			base_future_a_time,
		)
	}

	return pool
}
