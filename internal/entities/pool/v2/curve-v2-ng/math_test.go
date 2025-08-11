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

func Test_GeometricMean_For_TwoCryptoNG(t *testing.T) {
	var calculatedMean, expectedMean uint256.Int
	expectedMean.SetFromDecimal("29029250277969780379")

	var x = []uint256.Int{
		*uint256.MustFromDecimal("32543594247210753297"),
		*uint256.MustFromDecimal("25894416126861417166"),
	}

	calculatedMean = *GeometricMean(x)

	diff := new(uint256.Int)
	if calculatedMean.Cmp(&expectedMean) > 0 {
		diff.Sub(&calculatedMean, &expectedMean)
	} else {
		diff.Sub(&expectedMean, &calculatedMean)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mgeometricMean FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedMean.String(), expectedMean.String(), diff.String())
	} else {
		t.Logf("\033[32mgeometricMean SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedMean.String(), expectedMean.String(), diff.String())
	}
}

func Test_GeometricMean_For_TriCryptoNG(t *testing.T) {
	var calculatedMean, expectedMean uint256.Int
	expectedMean.SetFromDecimal("17671000109279242")

	var x = []uint256.Int{
		*uint256.MustFromDecimal("63624729793505614488987"),
		*uint256.MustFromDecimal("220406131330584"),
		*uint256.MustFromDecimal("393490059984"),
	}

	calculatedMean = *GeometricMean(x)

	diff := new(uint256.Int)
	if calculatedMean.Cmp(&expectedMean) > 0 {
		diff.Sub(&calculatedMean, &expectedMean)
	} else {
		diff.Sub(&expectedMean, &calculatedMean)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mgeometricMean FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedMean.String(), expectedMean.String(), diff.String())
	} else {
		t.Logf("\033[32mgeometricMean SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedMean.String(), expectedMean.String(), diff.String())
	}
}

func Test_ReductionCoefficient_For_TwoCryptoNG(t *testing.T) {
	var feeGamma, precision uint256.Int
	feeGamma.SetFromDecimal("300000000000000000")
	precision.SetFromDecimal("1000000000000000000")

	var x = []uint256.Int{
		*uint256.MustFromDecimal("32543594247210753297"),
		*uint256.MustFromDecimal("25894416126861417166"),
	}

	var calculatedRC, expectedRC uint256.Int
	expectedRC.SetFromDecimal("958630974138453631")

	ReductionCoefficient(x, &feeGamma, &calculatedRC)

	diff := new(uint256.Int)
	if calculatedRC.Cmp(&expectedRC) > 0 {
		diff.Sub(&calculatedRC, &expectedRC)
	} else {
		diff.Sub(&expectedRC, &calculatedRC)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mReductionCoefficient FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedRC.String(), expectedRC.String(), diff.String())
	} else {
		t.Logf("\033[32mReductionCoefficient SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedRC.String(), expectedRC.String(), diff.String())
	}
}

func Test_ReductionCoefficient_For_TriCryptoNG(t *testing.T) {
	var feeGamma, precision uint256.Int
	feeGamma.SetFromDecimal("300000000000000000")
	precision.SetFromDecimal("1000000000000000000")

	var x = []uint256.Int{
		*uint256.MustFromDecimal("63624729793505614488987"),
		*uint256.MustFromDecimal("220406131330584"),
		*uint256.MustFromDecimal("393490059984"),
	}

	var calculatedRC, expectedRC uint256.Int
	expectedRC.SetFromDecimal("230769230769230769")

	ReductionCoefficient(x, &feeGamma, &calculatedRC)

	diff := new(uint256.Int)
	if calculatedRC.Cmp(&expectedRC) > 0 {
		diff.Sub(&calculatedRC, &expectedRC)
	} else {
		diff.Sub(&expectedRC, &calculatedRC)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mReductionCoefficient FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedRC.String(), expectedRC.String(), diff.String())
	} else {
		t.Logf("\033[32mReductionCoefficient SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedRC.String(), expectedRC.String(), diff.String())
	}
}

func Test_NewtonY_For_TwoCryptoNG(t *testing.T) {
	var i = 0
	var A, gamma, precision, D, dx, expectedY uint256.Int
	A.SetFromDecimal("400000")
	gamma.SetFromDecimal("145000000000000")
	precision.SetFromDecimal("1000000000000000000")
	D.SetFromDecimal("58104810078143560570")
	dx.SetFromDecimal("12345000000")
	expectedY.SetFromDecimal("32543594247210753303")

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

	xp[i].Add(&xp[i], &dx)
	xp[0].Mul(&xp[0], &precisions[0])
	xp[1].Div(
		number.Mul(number.Mul(&xp[1], &precisions[1]), &priceScale[0]),
		number.TenPow(18),
	)

	xp = SortArray(xp)

	lim_mul := number.TenPow(20)

	if gamma.Cmp(MaxGammaTwoSmall) > 0 {
		lim_mul = number.Div(
			number.Mul(
				lim_mul,
				MaxGammaTwoSmall,
			),
			&gamma,
		)
	}

	calculatedY := new(uint256.Int)
	err := Newton_y(&A, &gamma, xp, &D, i, lim_mul, calculatedY)
	if err != nil {
		t.Fatalf("Newton_y failed: %v", err)
	}

	diff := new(uint256.Int)
	if calculatedY.Cmp(&expectedY) > 0 {
		diff.Sub(calculatedY, &expectedY)
	} else {
		diff.Sub(&expectedY, calculatedY)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mnewton_y FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedY.String(), expectedY.String(), diff.String())
	} else {
		t.Logf("\033[32mnewton_y SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedY.String(), expectedY.String(), diff.String())
	}
}

func Test_NewtonY_For_TriCryptoNG(t *testing.T) {
	var i = 0
	var A, gamma, precision, D, dx, expectedY uint256.Int
	A.SetFromDecimal("1707629")
	gamma.SetFromDecimal("11809167828997")
	precision.SetFromDecimal("1000000000000000000")
	D.SetFromDecimal("659319789163351247966687332")
	dx.SetFromDecimal("12345000000")
	expectedY.SetFromDecimal("221734358807933451995707231")

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

	xp[i].Add(&xp[i], &dx)
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

	calculatedY := new(uint256.Int)
	err := Newton_y(&A, &gamma, xp, &D, i, nil, calculatedY)
	if err != nil {
		t.Fatalf("Newton_y failed: %v", err)
	}

	diff := new(uint256.Int)
	if calculatedY.Cmp(&expectedY) > 0 {
		diff.Sub(calculatedY, &expectedY)
	} else {
		diff.Sub(&expectedY, calculatedY)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mnewton_y FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedY.String(), expectedY.String(), diff.String())
	} else {
		t.Logf("\033[32mnewton_y SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedY.String(), expectedY.String(), diff.String())
	}
}

func Test_SnekmateLog2_For_TriCryptoNG(t *testing.T) {
	var x, expectedLog2 uint256.Int
	x.SetFromDecimal("1000000000000000000")
	expectedLog2.SetFromDecimal("60")

	calculatedLog2 := SnekmateLog2(&x, true)

	diff := new(uint256.Int)
	if calculatedLog2.Cmp(&expectedLog2) > 0 {
		diff.Sub(calculatedLog2, &expectedLog2)
	} else {
		diff.Sub(&expectedLog2, calculatedLog2)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31msnekmate_log_2 FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedLog2.String(), expectedLog2.String(), diff.String())
	} else {
		t.Logf("\033[32msnekmate_log_2 SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedLog2.String(), expectedLog2.String(), diff.String())
	}
}

func Test_Cbrt_For_TriCryptoNG(t *testing.T) {
	var x, expectedCbrt uint256.Int
	x.SetFromDecimal("100")
	expectedCbrt.SetFromDecimal("4641588833612")

	calculatedCbrt := Cbrt(&x)

	diff := new(uint256.Int)
	if calculatedCbrt.Cmp(&expectedCbrt) > 0 {
		diff.Sub(calculatedCbrt, &expectedCbrt)
	} else {
		diff.Sub(&expectedCbrt, calculatedCbrt)
	}

	maxAllowedDiff := uint256.NewInt(2)
	if diff.Cmp(maxAllowedDiff) > 0 {
		t.Errorf("\033[31mcbrt FAILED calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedCbrt.String(), expectedCbrt.String(), diff.String())
	} else {
		t.Logf("\033[32mcbrt SUCCESS: calculated %s, expected %s (diff: %s wei)\033[0m",
			calculatedCbrt.String(), expectedCbrt.String(), diff.String())
	}
}
