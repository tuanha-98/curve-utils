package curvev1

import (
	"encoding/json"
	"io"
	"os"
	"testing"
	"time"

	"github.com/holiman/uint256"
	entities "github.com/tuanha-98/curve-utils/internal/entities/pool/v1"
)

type PoolJSON []struct {
	Pool      entities.Pool
	TestCases []struct {
		In                string `json:"in"`
		IndexIn           int    `json:"indexIn"`
		Out               string `json:"out"`
		IndexOut          int    `json:"indexOut"`
		AmountIn          string `json:"amountIn"`
		ExpectedAmountOut string `json:"expectedAmountOut"`
		Swappable         bool   `json:"swappable"`
		BlockTimestamp    int64  `json:"blockTimestamp"`
	}
}

func TestGetDYStablePool(t *testing.T) {

	jsonFile, err := os.Open("data/curvev1_pools_with_testcases.json")
	if err != nil {
		t.Fatalf("Failed to open curve.json: %v", err)
	}

	defer jsonFile.Close()

	var result PoolJSON
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if err := json.Unmarshal(byteValue, &result); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	for _, poolResult := range result {
		if poolResult.Pool.Kind == PoolTypeMeta {
			t.Logf("\033[33mSkipping META pool %s\033[0m", poolResult.Pool.Address)
			continue
		}

		NowFunc = func() time.Time {
			return time.Unix(poolResult.Pool.BlockTimestamp, 0)
		}

		pool, err := NewPool(poolResult.Pool)
		if err != nil {
			t.Fatalf("Failed to construct pool: %s", poolResult.Pool.Address)
		}

		for _, testCase := range poolResult.TestCases {
			if testCase.Swappable {
				var amountIn, amountOut uint256.Int
				amountIn.SetFromDecimal(testCase.AmountIn)
				err := pool.GetDy(testCase.IndexIn, testCase.IndexOut, &amountIn, &amountOut)
				if err != nil {
					t.Errorf("Failed to calculate GetDy for pool %s: %v", poolResult.Pool.Address, err)
					continue
				}
				expectAmountOut := uint256.MustFromDecimal(testCase.ExpectedAmountOut)

				diff := new(uint256.Int)
				if amountOut.Cmp(expectAmountOut) > 0 {
					diff.Sub(&amountOut, expectAmountOut)
				} else {
					diff.Sub(expectAmountOut, &amountOut)
				}

				maxAllowedDiff := uint256.NewInt(2)
				if diff.Cmp(maxAllowedDiff) > 0 {
					t.Errorf("\033[31mGetDy FAILED for pool %s: calculated %s, expected %s (diff: %s wei)\033[0m",
						poolResult.Pool.Address, amountOut.String(), expectAmountOut.String(), diff.String())
				} else {
					t.Logf("\033[32mGetDy SUCCESS for pool %s: calculated %s, expected %s\033[0m",
						poolResult.Pool.Address, amountOut.String(), expectAmountOut.String())
				}
			}
		}
	}
}
