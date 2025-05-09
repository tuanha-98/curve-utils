package constants

import m "github.com/tuanha-98/curve-utils/internal/utils/maths"

var (
	Precision      = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(18))
	FeeDenominator = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(10))
)
