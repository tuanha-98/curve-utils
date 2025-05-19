package token

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	erc20 "github.com/tuanha-98/curve-utils/contract/erc20"
)

func NewContractWrapper(client *ethclient.Client, tokenAddress common.Address) (*erc20.ContractCaller, error) {
	tokenContract, err := erc20.NewContractCaller(tokenAddress, client)
	if err != nil {
		return nil, err
	}
	return tokenContract, nil
}

func NewToken(address common.Address) (*Token, error) {
	client, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	if err != nil {
		return nil, err
	}

	if address == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
		return &Token{
			Address:  address.String(),
			Name:     "Ether",
			Symbol:   "ETH",
			Decimals: 18,
		}, nil
	} else {
		tokenContract, err := NewContractWrapper(client, address)
		if err != nil {
			return nil, err
		}

		name, err := tokenContract.Name(nil)
		if err != nil {
			return nil, err
		}
		symbol, err := tokenContract.Symbol(nil)
		if err != nil {
			return nil, err
		}
		decimals, err := tokenContract.Decimals(nil)
		if err != nil {
			return nil, err
		}
		totalSupply, err := tokenContract.TotalSupply(nil)
		if err != nil {
			return nil, err
		}

		return &Token{
			Address:     address.String(),
			Name:        name,
			Symbol:      symbol,
			Decimals:    decimals,
			TotalSupply: *uint256.MustFromBig(totalSupply),
		}, nil
	}
}
