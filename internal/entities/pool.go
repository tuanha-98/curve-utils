package entities

import "github.com/holiman/uint256"

type Pool struct {
	Address string `json:"id"`
	Tokens  []struct {
		ID       string `json:"id"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"tokens"`
	NTokens         int      `json:"nTokens,omitempty"`
	BasePoolAddress string   `json:"basePoolAddress,omitempty"`
	Reserves        []string `json:"reserves,omitempty"`
	TotalSupply     string   `json:"totalSupply,omitempty"`
	SwapFee         string   `json:"swapFee"`
	AdminFee        string   `json:"adminFee"`
	OffPegFee       string   `json:"offPegFee"`
	InitialA        string   `json:"initialA"`
	FutureA         string   `json:"futureA"`
	InitialATime    string   `json:"initialATime"`
	FutureATime     string   `json:"futureATime"`
	APrecision      string   `json:"APrecision"`
	Precisions      []string `json:"precisions"`
	Rates           []string `json:"rates"`
	Kind            string   `json:"kind,omitempty"`
	LpTokenAddress  string   `json:"lpTokenAddress,omitempty"`
	BasePoolType    string   `json:"basePoolType,omitempty"`
	BlockTimestamp  int64    `json:"blockTimestamp,omitempty"`
}

type FeeInfo struct {
	SwapFee, AdminFee, OffPegFee uint256.Int
}

func (p *Pool) IsMeta() bool {
	return len(p.BasePoolAddress) > 0
}

func (p *Pool) GetBasePoolType() string {
	return p.BasePoolType
}
