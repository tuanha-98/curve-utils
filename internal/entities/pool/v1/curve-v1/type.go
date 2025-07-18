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

		SwapFee   *uint256.Int
		AdminFee  *uint256.Int
		OffPegFee *uint256.Int

		Rates      []uint256.Int
		Precisions []uint256.Int
	}
)
