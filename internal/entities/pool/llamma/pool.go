package llamma

type Band struct {
	Index int64  `json:"index"`
	BandX string `json:"bandX"`
	BandY string `json:"bandY"`
}

type Pool struct {
	Address string `json:"id"`
	Tokens  []struct {
		ID       string `json:"id"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"tokens"`
	NTokens             int      `json:"nTokens,omitempty"`
	Reserves            []string `json:"reserves,omitempty"`
	A                   string   `json:"A"`
	BasePrice           string   `json:"basePrice"`
	PriceOracle         string   `json:"priceOracle"`
	SwapFee             string   `json:"swapFee"`
	AdminFee            string   `json:"adminFee"`
	ActiveBand          int64    `json:"activeBand"`
	MinBand             int64    `json:"minBand"`
	MaxBand             int64    `json:"maxBand"`
	Bands               []Band   `json:"bands"`
	BorrowedPrecision   string   `json:"borrowedPrecision"`
	CollateralPrecision string   `json:"collateralPrecision"`
	BlockTimestamp      int64    `json:"blockTimestamp,omitempty"`
	UseDynamicFee       bool     `json:"useDynamicFee,omitempty"`
}
