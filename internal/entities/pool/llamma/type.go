package llamma

import (
	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/maths/int256"
)

type (
	Extra struct {
		A                   *uint256.Int
		Aminus1             *uint256.Int
		MaxOracleDnPow      *uint256.Int
		LogARatio           *int256.Int
		BasePrice           *uint256.Int
		PriceOracle         *uint256.Int
		BorrowedPrecision   *uint256.Int
		CollateralPrecision *uint256.Int
		SwapFee             *uint256.Int
		AdminFee            *uint256.Int
		ActiveBand          int64
		MinBand             int64
		MaxBand             int64
		BandsX              map[int64]*uint256.Int
		BandsY              map[int64]*uint256.Int
	}

	DetailedTrade struct {
		InAmount  uint256.Int
		OutAmount uint256.Int
		N1        int64
		N2        int64
		TicksIn   []uint256.Int
		LastTickJ uint256.Int
		AdminFee  uint256.Int
	}
)
