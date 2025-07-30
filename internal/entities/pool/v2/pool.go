package v2

import "github.com/holiman/uint256"

type Pool struct {
	Address string `json:"id"`
	Tokens  []struct {
		ID       string `json:"id"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"tokens"`
	NTokens           int      `json:"nTokens,omitempty"`
	Reserves          []string `json:"reserves,omitempty"`
	TotalSupply       string   `json:"totalSupply,omitempty"`
	D                 string   `json:"d"`
	MidFee            string   `json:"midFee"`
	OutFee            string   `json:"outFee"`
	FeeGamma          string   `json:"feeGamma"`
	AdminFee          string   `json:"adminFee"`
	InitialAGamma     string   `json:"initialAGamma"`
	FutureAGamma      string   `json:"futureAGamma"`
	InitialAGammaTime string   `json:"initialAGammaTime"`
	FutureAGammaTime  string   `json:"futureAGammaTime"`
	Precisions        []string `json:"precisions"`
	PriceScales       []string `json:"priceScales"`
	LpTokenAddress    string   `json:"lpTokenAddress,omitempty"`
	BlockTimestamp    int64    `json:"blockTimestamp,omitempty"`
}

type FeeInfo struct {
	MidFee, OutFee, FeeGamma, AdminFee uint256.Int
}
