package token

import "github.com/holiman/uint256"

type Token struct {
	Address     string
	Name        string
	Symbol      string
	Decimals    uint8
	TotalSupply uint256.Int
}
