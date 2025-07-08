package curvev1

import "github.com/holiman/uint256"

type (
	Static struct {
		PoolType   string
		APrecision *uint256.Int
	}

	Extra struct {
		InitialA     *uint256.Int
		FutureA      *uint256.Int
		InitialATime int64
		FutureATime  int64

		SwapFee             *uint256.Int
		AdminFee            *uint256.Int
		OffPegFeeMultiplier *uint256.Int

		RateMultipliers      []uint256.Int
		PrecisionMultipliers []uint256.Int
	}
)
