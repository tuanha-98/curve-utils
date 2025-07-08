package entities

type Pool struct {
	Address       string `json:"id"`
	Amplification struct {
		Initial     string `json:"initial"`
		Future      string `json:"future"`
		InitialTime string `json:"initialTime"`
		FutureTime  string `json:"futureTime"`
	} `json:"amplification"`
	BasePool string `json:"basePool,omitempty"`
	Fee      struct {
		SwapFee             string `json:"swapFee"`
		AdminFee            string `json:"adminFee"`
		OffPegFeeMultiplier string `json:"offPegFeeMultipliers"`
	} `json:"fee"`
	Type        string `json:"kind,omitempty"`
	LpToken     string `json:"lpToken,omitempty"`
	Multipliers struct {
		APrecision           string   `json:"aPrecision"`
		PrecisionMultipliers []string `json:"precisions"`
		RateMultipliers      []string `json:"rateMultipliers"`
	} `json:"multipliers"`
	Ncoins   int      `json:"ncoins,omitempty"`
	Reserves []string `json:"reserves,omitempty"`
	Tokens   []struct {
		ID       string `json:"id"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"tokens"`
	TotalSupply string `json:"totalSupply,omitempty"`
}

func (p *Pool) IsMeta() bool {
	return len(p.BasePool) > 0
}
