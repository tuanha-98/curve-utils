package maths

import (
	"math/big"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		a, b     *big.Int
		expected *big.Int
	}{
		{big.NewInt(1), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(-1), big.NewInt(-2), big.NewInt(-1)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
	}

	for _, test := range tests {
		if got := Max(test.a, test.b); got.Cmp(test.expected) != 0 {
			t.Errorf("Max(%v, %v) = %v, expected %v", test.a, test.b, got, test.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a, b     *big.Int
		expected *big.Int
	}{
		{big.NewInt(1), big.NewInt(2), big.NewInt(1)},
		{big.NewInt(-1), big.NewInt(-2), big.NewInt(-2)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
	}

	for _, test := range tests {
		if got := Min(test.a, test.b); got.Cmp(test.expected) != 0 {
			t.Errorf("Min(%v, %v) = %v, expected %v", test.a, test.b, got, test.expected)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    *big.Int
		expected *big.Int
	}{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(-1), big.NewInt(1)},
		{big.NewInt(0), big.NewInt(0)},
	}

	for _, test := range tests {
		if got := Abs(test.input); got.Cmp(test.expected) != 0 {
			t.Errorf("Abs(%v) = %v, expected %v", test.input, got, test.expected)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		input    *big.Int
		expected *big.Int
	}{
		{big.NewInt(4), big.NewInt(2)},
		{big.NewInt(9), big.NewInt(3)},
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(-1), nil},
	}

	for _, test := range tests {
		if got := Sqrt(test.input); (got == nil && test.expected != nil) || (got != nil && test.expected != nil && got.Cmp(test.expected) != 0) {
			t.Errorf("Sqrt(%v) = %v, expected %v", test.input, got, test.expected)
		}
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		x, y     *big.Int
		expected *big.Int
	}{
		{big.NewInt(2), big.NewInt(3), big.NewInt(8)},
		{big.NewInt(3), big.NewInt(2), big.NewInt(9)},
		{big.NewInt(2), big.NewInt(0), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(1), big.NewInt(2)},
	}

	for _, test := range tests {
		if got := Pow(test.x, test.y); got.Cmp(test.expected) != 0 {
			t.Errorf("Pow(%v, %v) = %v, expected %v", test.x, test.y, got, test.expected)
		}
	}
}

func TestToUint256(t *testing.T) {
	max := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	tests := []struct {
		input    *big.Int
		expected *big.Int
	}{
		{big.NewInt(1), big.NewInt(1)},
		{new(big.Int).Add(max, big.NewInt(1)), big.NewInt(0)},
		{max, max},
		{big.NewInt(-1), max}, // -1 should wrap around to max uint256
	}

	for _, test := range tests {
		if got := ToUint256(test.input); got.Cmp(test.expected) != 0 {
			t.Errorf("ToUint256(%v) = %v, expected %v", test.input, got, test.expected)
		}
	}
}

func TestToInt256(t *testing.T) {
	maxInt256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 255), big.NewInt(1))
	maxUint256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
	tests := []struct {
		input    *big.Int
		expected *big.Int
	}{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(-1), big.NewInt(-1)},
		{maxInt256, maxInt256},
		{new(big.Int).Add(maxInt256, big.NewInt(1)), new(big.Int).Sub(new(big.Int).Add(maxInt256, big.NewInt(1)), new(big.Int).Lsh(big.NewInt(1), 256))},
		{maxUint256, big.NewInt(-1)}, // Max uint256 should wrap around to -1 in int256
	}

	for _, test := range tests {
		if got := ToInt256(test.input); got.Cmp(test.expected) != 0 {
			t.Errorf("ToInt256(%v) = %v, expected %v", test.input, got, test.expected)
		}
	}
}
