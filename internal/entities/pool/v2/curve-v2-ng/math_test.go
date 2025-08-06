package curvev2ng

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/tuanha-98/curve-utils/internal/utils/toolkit/number"
)

func Test_NewtonD_For_TwoCryptoNG(t *testing.T) {
	var A, gamma, precision, expectedD uint256.Int
	A.SetFromDecimal("400000")
	gamma.SetFromDecimal("145000000000000")
	precision.SetFromDecimal("1000000000000000000")
	expectedD.SetFromDecimal("58104810064367150682")

	var precisions = []uint256.Int{
		*uint256.MustFromDecimal("1"),
		*uint256.MustFromDecimal("1"),
	}

	var priceScale = []uint256.Int{
		*uint256.MustFromDecimal("56538026505692"),
	}
	var balances = []uint256.Int{
		*uint256.MustFromDecimal("25894416114516417166"),
		*uint256.MustFromDecimal("575605415656565360862706"),
	}

	var xp = make([]uint256.Int, len(balances))
	for k := range balances {
		xp[k].Set(&balances[k])
	}

	xp[0].Mul(&xp[0], &precisions[0])
	xp[1].Div(
		number.Mul(number.Mul(&xp[1], &precisions[1]), &priceScale[0]),
		number.TenPow(18),
	)

	xp = SortArray(xp)

	calculatedD := new(uint256.Int)
	err := Newton_D(&A, &gamma, xp, number.Zero, calculatedD)
	if err != nil {
		t.Fatalf("Newton_D failed: %v", err)
	}

	diff := new(uint256.Int)
	if calculatedD.Cmp(&expectedD) > 0 {
		diff.Sub(calculatedD, &expectedD)
	} else {
		diff.Sub(&expectedD, calculatedD)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mnewton_D FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedD.String(), expectedD.String(), diff.String())
	} else {
		t.Logf("\033[32mnewton_D SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedD.String(), expectedD.String(), diff.String())
	}
}

func Test_NewtonD_For_TriCryptoNG(t *testing.T) {
	var A, gamma, precision, expectedD uint256.Int
	A.SetFromDecimal("1707629")
	gamma.SetFromDecimal("11809167828997")
	precision.SetFromDecimal("1000000000000000000")
	expectedD.SetFromDecimal("659307468228931998601036430")

	var precisions = []uint256.Int{
		*uint256.MustFromDecimal("1000000000000"),
		*uint256.MustFromDecimal("10000000000"),
		*uint256.MustFromDecimal("1"),
	}

	var priceScale = []uint256.Int{
		*uint256.MustFromDecimal("55192676963173615208913"),
		*uint256.MustFromDecimal("3485034192326999988769"),
	}
	var balances = []uint256.Int{
		*uint256.MustFromDecimal("220406131330584"),
		*uint256.MustFromDecimal("393490059984"),
		*uint256.MustFromDecimal("63624729793505614488987"),
	}

	var xp = make([]uint256.Int, len(balances))
	for k := range balances {
		xp[k].Set(&balances[k])
	}

	xp[0].Mul(&xp[0], &precisions[0])

	for k := 0; k < len(balances)-1; k++ {
		xp[k+1].Div(
			number.SafeMul(
				number.SafeMul(
					&xp[k+1],
					&priceScale[k],
				),
				&precisions[k+1],
			),
			&precision,
		)
	}

	xp = SortArray(xp)

	calculatedD := new(uint256.Int)
	err := Newton_D(&A, &gamma, xp, number.Zero, calculatedD)
	if err != nil {
		t.Fatalf("Newton_D failed: %v", err)
	}

	diff := new(uint256.Int)
	if calculatedD.Cmp(&expectedD) > 0 {
		diff.Sub(calculatedD, &expectedD)
	} else {
		diff.Sub(&expectedD, calculatedD)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mnewton_D FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedD.String(), expectedD.String(), diff.String())
	} else {
		t.Logf("\033[32mnewton_D SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedD.String(), expectedD.String(), diff.String())
	}
}
