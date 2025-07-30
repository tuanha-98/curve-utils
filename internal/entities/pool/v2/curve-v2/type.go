package curvev2

import "github.com/holiman/uint256"

type (
	Extra struct {
		InitialAGamma     *uint256.Int
		FutureAGamma      *uint256.Int
		InitialAGammaTime int64
		FutureAGammaTime  int64

		D *uint256.Int

		MidFee   *uint256.Int
		OutFee   *uint256.Int
		FeeGamma *uint256.Int
		AdminFee *uint256.Int

		PriceScales []uint256.Int
		Precisions  []uint256.Int
	}
)
