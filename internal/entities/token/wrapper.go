package token

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	token, err := NewContractWrapper(client, address)

	name, err := token.Name(nil)
	if err != nil {
		return nil, err
	}
	symbol, err := token.Symbol(nil)
	if err != nil {
		return nil, err
	}
	decimals, err := token.Decimals(nil)
	if err != nil {
		return nil, err
	}

	return &Token{
		Address:  address.String(),
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}, nil
}
