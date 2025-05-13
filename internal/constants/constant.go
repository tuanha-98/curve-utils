package constants

import m "github.com/tuanha-98/curve-utils/internal/utils/maths"

var (
	Precision      = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(18))
	FeeDenominator = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(10))

	Wei9  = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(9))
	Wei14 = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(14))
	Wei15 = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(15))
	Wei16 = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(16))
	Wei17 = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(17))
	Wei18 = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(18))
	Wei20 = new(m.Uint256).Exp(new(m.Uint256).SetUint64(10), new(m.Uint256).SetUint64(20))
)
