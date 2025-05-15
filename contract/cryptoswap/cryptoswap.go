// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[2]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_index\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampAgamma\",\"inputs\":[{\"name\":\"initial_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"current_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ClaimAdminFee\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"admin_fee_receiver\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"gamma\",\"type\":\"uint256\"},{\"name\":\"mid_fee\",\"type\":\"uint256\"},{\"name\":\"out_fee\",\"type\":\"uint256\"},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"fee_gamma\",\"type\":\"uint256\"},{\"name\":\"adjustment_step\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"ma_half_time\",\"type\":\"uint256\"},{\"name\":\"initial_price\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_coins\",\"type\":\"address[2]\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[2]\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A_gamma\",\"inputs\":[{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"name\":\"future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A_gamma\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_parameters\",\"inputs\":[{\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"name\":\"_new_admin_fee\",\"type\":\"uint256\"},{\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"name\":\"_new_ma_half_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_new_parameters\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_transfer_ownership\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_transfer_ownership\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"kill_me\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"unkill_me\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin_fee_receiver\",\"inputs\":[{\"name\":\"_admin_fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lp_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_scale\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit_a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"kill_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"transfer_ownership_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_actions_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Contract *ContractCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Contract *ContractSession) A() (*big.Int, error) {
	return _Contract.Contract.A(&_Contract.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Contract *ContractCallerSession) A() (*big.Int, error) {
	return _Contract.Contract.A(&_Contract.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Contract *ContractCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Contract *ContractSession) D() (*big.Int, error) {
	return _Contract.Contract.D(&_Contract.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Contract *ContractCallerSession) D() (*big.Int, error) {
	return _Contract.Contract.D(&_Contract.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Contract *ContractCaller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Contract *ContractSession) AdjustmentStep() (*big.Int, error) {
	return _Contract.Contract.AdjustmentStep(&_Contract.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Contract *ContractCallerSession) AdjustmentStep() (*big.Int, error) {
	return _Contract.Contract.AdjustmentStep(&_Contract.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Contract *ContractCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Contract *ContractSession) AdminActionsDeadline() (*big.Int, error) {
	return _Contract.Contract.AdminActionsDeadline(&_Contract.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Contract *ContractCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Contract.Contract.AdminActionsDeadline(&_Contract.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Contract *ContractCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Contract *ContractSession) AdminFee() (*big.Int, error) {
	return _Contract.Contract.AdminFee(&_Contract.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Contract *ContractCallerSession) AdminFee() (*big.Int, error) {
	return _Contract.Contract.AdminFee(&_Contract.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_Contract *ContractCaller) AdminFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "admin_fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_Contract *ContractSession) AdminFeeReceiver() (common.Address, error) {
	return _Contract.Contract.AdminFeeReceiver(&_Contract.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_Contract *ContractCallerSession) AdminFeeReceiver() (common.Address, error) {
	return _Contract.Contract.AdminFeeReceiver(&_Contract.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Contract *ContractCaller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Contract *ContractSession) AllowedExtraProfit() (*big.Int, error) {
	return _Contract.Contract.AllowedExtraProfit(&_Contract.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Contract *ContractCallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _Contract.Contract.AllowedExtraProfit(&_Contract.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Contract *ContractSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Contract *ContractCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Contract *ContractCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [2]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calc_token_amount", amounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Contract *ContractSession) CalcTokenAmount(amounts [2]*big.Int) (*big.Int, error) {
	return _Contract.Contract.CalcTokenAmount(&_Contract.CallOpts, amounts)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x8d8ea727.
//
// Solidity: function calc_token_amount(uint256[2] amounts) view returns(uint256)
func (_Contract *ContractCallerSession) CalcTokenAmount(amounts [2]*big.Int) (*big.Int, error) {
	return _Contract.Contract.CalcTokenAmount(&_Contract.CallOpts, amounts)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Contract *ContractCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Contract *ContractSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Contract.Contract.CalcWithdrawOneCoin(&_Contract.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Contract *ContractCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Contract.Contract.CalcWithdrawOneCoin(&_Contract.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_Contract *ContractCaller) Coins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_Contract *ContractSession) Coins(i *big.Int) (common.Address, error) {
	return _Contract.Contract.Coins(&_Contract.CallOpts, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_Contract *ContractCallerSession) Coins(i *big.Int) (common.Address, error) {
	return _Contract.Contract.Coins(&_Contract.CallOpts, i)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Contract *ContractCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Contract *ContractSession) Fee() (*big.Int, error) {
	return _Contract.Contract.Fee(&_Contract.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Contract *ContractCallerSession) Fee() (*big.Int, error) {
	return _Contract.Contract.Fee(&_Contract.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Contract *ContractCaller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Contract *ContractSession) FeeGamma() (*big.Int, error) {
	return _Contract.Contract.FeeGamma(&_Contract.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Contract *ContractCallerSession) FeeGamma() (*big.Int, error) {
	return _Contract.Contract.FeeGamma(&_Contract.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Contract *ContractCaller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Contract *ContractSession) FutureAGamma() (*big.Int, error) {
	return _Contract.Contract.FutureAGamma(&_Contract.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Contract *ContractCallerSession) FutureAGamma() (*big.Int, error) {
	return _Contract.Contract.FutureAGamma(&_Contract.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Contract *ContractCaller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Contract *ContractSession) FutureAGammaTime() (*big.Int, error) {
	return _Contract.Contract.FutureAGammaTime(&_Contract.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Contract *ContractCallerSession) FutureAGammaTime() (*big.Int, error) {
	return _Contract.Contract.FutureAGammaTime(&_Contract.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Contract *ContractCaller) FutureAdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Contract *ContractSession) FutureAdjustmentStep() (*big.Int, error) {
	return _Contract.Contract.FutureAdjustmentStep(&_Contract.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Contract *ContractCallerSession) FutureAdjustmentStep() (*big.Int, error) {
	return _Contract.Contract.FutureAdjustmentStep(&_Contract.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Contract *ContractCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Contract *ContractSession) FutureAdminFee() (*big.Int, error) {
	return _Contract.Contract.FutureAdminFee(&_Contract.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Contract *ContractCallerSession) FutureAdminFee() (*big.Int, error) {
	return _Contract.Contract.FutureAdminFee(&_Contract.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Contract *ContractCaller) FutureAllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Contract *ContractSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Contract.Contract.FutureAllowedExtraProfit(&_Contract.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Contract *ContractCallerSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Contract.Contract.FutureAllowedExtraProfit(&_Contract.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Contract *ContractCaller) FutureFeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Contract *ContractSession) FutureFeeGamma() (*big.Int, error) {
	return _Contract.Contract.FutureFeeGamma(&_Contract.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Contract *ContractCallerSession) FutureFeeGamma() (*big.Int, error) {
	return _Contract.Contract.FutureFeeGamma(&_Contract.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Contract *ContractCaller) FutureMaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Contract *ContractSession) FutureMaHalfTime() (*big.Int, error) {
	return _Contract.Contract.FutureMaHalfTime(&_Contract.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Contract *ContractCallerSession) FutureMaHalfTime() (*big.Int, error) {
	return _Contract.Contract.FutureMaHalfTime(&_Contract.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Contract *ContractCaller) FutureMidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Contract *ContractSession) FutureMidFee() (*big.Int, error) {
	return _Contract.Contract.FutureMidFee(&_Contract.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Contract *ContractCallerSession) FutureMidFee() (*big.Int, error) {
	return _Contract.Contract.FutureMidFee(&_Contract.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Contract *ContractCaller) FutureOutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Contract *ContractSession) FutureOutFee() (*big.Int, error) {
	return _Contract.Contract.FutureOutFee(&_Contract.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Contract *ContractCallerSession) FutureOutFee() (*big.Int, error) {
	return _Contract.Contract.FutureOutFee(&_Contract.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Contract *ContractCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Contract *ContractSession) FutureOwner() (common.Address, error) {
	return _Contract.Contract.FutureOwner(&_Contract.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Contract *ContractCallerSession) FutureOwner() (common.Address, error) {
	return _Contract.Contract.FutureOwner(&_Contract.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Contract *ContractCaller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Contract *ContractSession) Gamma() (*big.Int, error) {
	return _Contract.Contract.Gamma(&_Contract.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Contract *ContractCallerSession) Gamma() (*big.Int, error) {
	return _Contract.Contract.Gamma(&_Contract.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Contract *ContractCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Contract *ContractSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDy(&_Contract.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Contract *ContractCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDy(&_Contract.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Contract *ContractCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Contract *ContractSession) GetVirtualPrice() (*big.Int, error) {
	return _Contract.Contract.GetVirtualPrice(&_Contract.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Contract *ContractCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Contract.Contract.GetVirtualPrice(&_Contract.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Contract *ContractCaller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Contract *ContractSession) InitialAGamma() (*big.Int, error) {
	return _Contract.Contract.InitialAGamma(&_Contract.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Contract *ContractCallerSession) InitialAGamma() (*big.Int, error) {
	return _Contract.Contract.InitialAGamma(&_Contract.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Contract *ContractCaller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Contract *ContractSession) InitialAGammaTime() (*big.Int, error) {
	return _Contract.Contract.InitialAGammaTime(&_Contract.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Contract *ContractCallerSession) InitialAGammaTime() (*big.Int, error) {
	return _Contract.Contract.InitialAGammaTime(&_Contract.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Contract *ContractCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Contract *ContractSession) IsKilled() (bool, error) {
	return _Contract.Contract.IsKilled(&_Contract.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Contract *ContractCallerSession) IsKilled() (bool, error) {
	return _Contract.Contract.IsKilled(&_Contract.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_Contract *ContractCaller) KillDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "kill_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_Contract *ContractSession) KillDeadline() (*big.Int, error) {
	return _Contract.Contract.KillDeadline(&_Contract.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_Contract *ContractCallerSession) KillDeadline() (*big.Int, error) {
	return _Contract.Contract.KillDeadline(&_Contract.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Contract *ContractCaller) LastPrices(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "last_prices")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Contract *ContractSession) LastPrices() (*big.Int, error) {
	return _Contract.Contract.LastPrices(&_Contract.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0xc146bf94.
//
// Solidity: function last_prices() view returns(uint256)
func (_Contract *ContractCallerSession) LastPrices() (*big.Int, error) {
	return _Contract.Contract.LastPrices(&_Contract.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Contract *ContractCaller) LastPricesTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "last_prices_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Contract *ContractSession) LastPricesTimestamp() (*big.Int, error) {
	return _Contract.Contract.LastPricesTimestamp(&_Contract.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Contract *ContractCallerSession) LastPricesTimestamp() (*big.Int, error) {
	return _Contract.Contract.LastPricesTimestamp(&_Contract.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Contract *ContractCaller) LpPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lp_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Contract *ContractSession) LpPrice() (*big.Int, error) {
	return _Contract.Contract.LpPrice(&_Contract.CallOpts)
}

// LpPrice is a free data retrieval call binding the contract method 0x54f0f7d5.
//
// Solidity: function lp_price() view returns(uint256)
func (_Contract *ContractCallerSession) LpPrice() (*big.Int, error) {
	return _Contract.Contract.LpPrice(&_Contract.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Contract *ContractCaller) MaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Contract *ContractSession) MaHalfTime() (*big.Int, error) {
	return _Contract.Contract.MaHalfTime(&_Contract.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Contract *ContractCallerSession) MaHalfTime() (*big.Int, error) {
	return _Contract.Contract.MaHalfTime(&_Contract.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Contract *ContractCaller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Contract *ContractSession) MidFee() (*big.Int, error) {
	return _Contract.Contract.MidFee(&_Contract.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Contract *ContractCallerSession) MidFee() (*big.Int, error) {
	return _Contract.Contract.MidFee(&_Contract.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Contract *ContractCaller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Contract *ContractSession) OutFee() (*big.Int, error) {
	return _Contract.Contract.OutFee(&_Contract.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Contract *ContractCallerSession) OutFee() (*big.Int, error) {
	return _Contract.Contract.OutFee(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Contract *ContractCaller) PriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "price_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Contract *ContractSession) PriceOracle() (*big.Int, error) {
	return _Contract.Contract.PriceOracle(&_Contract.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x86fc88d3.
//
// Solidity: function price_oracle() view returns(uint256)
func (_Contract *ContractCallerSession) PriceOracle() (*big.Int, error) {
	return _Contract.Contract.PriceOracle(&_Contract.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Contract *ContractCaller) PriceScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "price_scale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Contract *ContractSession) PriceScale() (*big.Int, error) {
	return _Contract.Contract.PriceScale(&_Contract.CallOpts)
}

// PriceScale is a free data retrieval call binding the contract method 0xb9e8c9fd.
//
// Solidity: function price_scale() view returns(uint256)
func (_Contract *ContractCallerSession) PriceScale() (*big.Int, error) {
	return _Contract.Contract.PriceScale(&_Contract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractSession) Token() (common.Address, error) {
	return _Contract.Contract.Token(&_Contract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contract *ContractCallerSession) Token() (common.Address, error) {
	return _Contract.Contract.Token(&_Contract.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Contract *ContractCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Contract *ContractSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Contract.Contract.TransferOwnershipDeadline(&_Contract.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Contract *ContractCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Contract.Contract.TransferOwnershipDeadline(&_Contract.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Contract *ContractCaller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Contract *ContractSession) VirtualPrice() (*big.Int, error) {
	return _Contract.Contract.VirtualPrice(&_Contract.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Contract *ContractCallerSession) VirtualPrice() (*big.Int, error) {
	return _Contract.Contract.VirtualPrice(&_Contract.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Contract *ContractCaller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Contract *ContractSession) XcpProfit() (*big.Int, error) {
	return _Contract.Contract.XcpProfit(&_Contract.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Contract *ContractCallerSession) XcpProfit() (*big.Int, error) {
	return _Contract.Contract.XcpProfit(&_Contract.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Contract *ContractCaller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Contract *ContractSession) XcpProfitA() (*big.Int, error) {
	return _Contract.Contract.XcpProfitA(&_Contract.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Contract *ContractCallerSession) XcpProfitA() (*big.Int, error) {
	return _Contract.Contract.XcpProfitA(&_Contract.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Contract *ContractTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Contract *ContractSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity(&_Contract.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) payable returns(uint256)
func (_Contract *ContractTransactorSession) AddLiquidity(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity(&_Contract.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Contract *ContractTransactor) AddLiquidity0(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add_liquidity0", amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Contract *ContractSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity0(&_Contract.TransactOpts, amounts, min_mint_amount, use_eth)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xee22be23.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount, bool use_eth) payable returns(uint256)
func (_Contract *ContractTransactorSession) AddLiquidity0(amounts [2]*big.Int, min_mint_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity0(&_Contract.TransactOpts, amounts, min_mint_amount, use_eth)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Contract *ContractTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Contract *ContractSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Contract.Contract.ApplyNewParameters(&_Contract.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Contract *ContractTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Contract.Contract.ApplyNewParameters(&_Contract.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Contract *ContractTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Contract *ContractSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Contract.Contract.ApplyTransferOwnership(&_Contract.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Contract *ContractTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Contract.Contract.ApplyTransferOwnership(&_Contract.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Contract *ContractTransactor) ClaimAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claim_admin_fees")
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Contract *ContractSession) ClaimAdminFees() (*types.Transaction, error) {
	return _Contract.Contract.ClaimAdminFees(&_Contract.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Contract *ContractTransactorSession) ClaimAdminFees() (*types.Transaction, error) {
	return _Contract.Contract.ClaimAdminFees(&_Contract.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Contract *ContractTransactor) CommitNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "commit_new_parameters", _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Contract *ContractSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CommitNewParameters(&_Contract.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Contract *ContractTransactorSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CommitNewParameters(&_Contract.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Contract *ContractTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Contract *ContractSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CommitTransferOwnership(&_Contract.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Contract *ContractTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CommitTransferOwnership(&_Contract.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Contract *ContractTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Contract *ContractSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Exchange(&_Contract.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Contract *ContractTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Exchange(&_Contract.TransactOpts, i, j, dx, min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Contract *ContractTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchange0", i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Contract *ContractSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.Exchange0(&_Contract.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns(uint256)
func (_Contract *ContractTransactorSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.Exchange0(&_Contract.TransactOpts, i, j, dx, min_dy, use_eth)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Contract *ContractTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchange_underlying", i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Contract *ContractSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeUnderlying(&_Contract.TransactOpts, i, j, dx, min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns(uint256)
func (_Contract *ContractTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeUnderlying(&_Contract.TransactOpts, i, j, dx, min_dy)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Contract *ContractTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Contract *ContractSession) KillMe() (*types.Transaction, error) {
	return _Contract.Contract.KillMe(&_Contract.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Contract *ContractTransactorSession) KillMe() (*types.Transaction, error) {
	return _Contract.Contract.KillMe(&_Contract.TransactOpts)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Contract *ContractTransactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Contract *ContractSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RampAGamma(&_Contract.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Contract *ContractTransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RampAGamma(&_Contract.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Contract *ContractTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Contract *ContractSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity(&_Contract.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts) returns()
func (_Contract *ContractTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity(&_Contract.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Contract *ContractTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity0", _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Contract *ContractSession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity0(&_Contract.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x269b5581.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[2] min_amounts, bool use_eth) returns()
func (_Contract *ContractTransactorSession) RemoveLiquidity0(_amount *big.Int, min_amounts [2]*big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity0(&_Contract.TransactOpts, _amount, min_amounts, use_eth)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Contract *ContractTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Contract *ContractSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin(&_Contract.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns(uint256)
func (_Contract *ContractTransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin(&_Contract.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Contract *ContractTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity_one_coin0", token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Contract *ContractSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin0(&_Contract.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x8f15b6b5.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount, bool use_eth) returns(uint256)
func (_Contract *ContractTransactorSession) RemoveLiquidityOneCoin0(token_amount *big.Int, i *big.Int, min_amount *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin0(&_Contract.TransactOpts, token_amount, i, min_amount, use_eth)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Contract *ContractTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Contract *ContractSession) RevertNewParameters() (*types.Transaction, error) {
	return _Contract.Contract.RevertNewParameters(&_Contract.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Contract *ContractTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Contract.Contract.RevertNewParameters(&_Contract.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Contract *ContractTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Contract *ContractSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RevertTransferOwnership(&_Contract.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Contract *ContractTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RevertTransferOwnership(&_Contract.TransactOpts)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_Contract *ContractTransactor) SetAdminFeeReceiver(opts *bind.TransactOpts, _admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_admin_fee_receiver", _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_Contract *ContractSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAdminFeeReceiver(&_Contract.TransactOpts, _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_Contract *ContractTransactorSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAdminFeeReceiver(&_Contract.TransactOpts, _admin_fee_receiver)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Contract *ContractTransactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Contract *ContractSession) StopRampAGamma() (*types.Transaction, error) {
	return _Contract.Contract.StopRampAGamma(&_Contract.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Contract *ContractTransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _Contract.Contract.StopRampAGamma(&_Contract.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Contract *ContractTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Contract *ContractSession) UnkillMe() (*types.Transaction, error) {
	return _Contract.Contract.UnkillMe(&_Contract.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Contract *ContractTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Contract.Contract.UnkillMe(&_Contract.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contract.Contract.Fallback(&_Contract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contract.Contract.Fallback(&_Contract.TransactOpts, calldata)
}

// ContractAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Contract contract.
type ContractAddLiquidityIterator struct {
	Event *ContractAddLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAddLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddLiquidity represents a AddLiquidity event raised by the Contract contract.
type ContractAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fee          *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Contract *ContractFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*ContractAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddLiquidityIterator{contract: _Contract.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Contract *ContractFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *ContractAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddLiquidity)
				if err := _Contract.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddLiquidity is a log parse operation binding the contract event 0x540ab385f9b5d450a27404172caade516b3ba3f4be88239ac56a2ad1de2a1f5a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256 fee, uint256 token_supply)
func (_Contract *ContractFilterer) ParseAddLiquidity(log types.Log) (*ContractAddLiquidity, error) {
	event := new(ContractAddLiquidity)
	if err := _Contract.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimAdminFeeIterator is returned from FilterClaimAdminFee and is used to iterate over the raw logs and unpacked data for ClaimAdminFee events raised by the Contract contract.
type ContractClaimAdminFeeIterator struct {
	Event *ContractClaimAdminFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractClaimAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimAdminFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClaimAdminFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractClaimAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimAdminFee represents a ClaimAdminFee event raised by the Contract contract.
type ContractClaimAdminFee struct {
	Admin  common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimAdminFee is a free log retrieval operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Contract *ContractFilterer) FilterClaimAdminFee(opts *bind.FilterOpts, admin []common.Address) (*ContractClaimAdminFeeIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimAdminFeeIterator{contract: _Contract.contract, event: "ClaimAdminFee", logs: logs, sub: sub}, nil
}

// WatchClaimAdminFee is a free log subscription operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Contract *ContractFilterer) WatchClaimAdminFee(opts *bind.WatchOpts, sink chan<- *ContractClaimAdminFee, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimAdminFee)
				if err := _Contract.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimAdminFee is a log parse operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Contract *ContractFilterer) ParseClaimAdminFee(log types.Log) (*ContractClaimAdminFee, error) {
	event := new(ContractClaimAdminFee)
	if err := _Contract.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Contract contract.
type ContractCommitNewAdminIterator struct {
	Event *ContractCommitNewAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCommitNewAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCommitNewAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCommitNewAdmin represents a CommitNewAdmin event raised by the Contract contract.
type ContractCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Contract *ContractFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*ContractCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &ContractCommitNewAdminIterator{contract: _Contract.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Contract *ContractFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *ContractCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCommitNewAdmin)
				if err := _Contract.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Contract *ContractFilterer) ParseCommitNewAdmin(log types.Log) (*ContractCommitNewAdmin, error) {
	event := new(ContractCommitNewAdmin)
	if err := _Contract.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Contract contract.
type ContractCommitNewParametersIterator struct {
	Event *ContractCommitNewParameters // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCommitNewParameters)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCommitNewParameters)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCommitNewParameters represents a CommitNewParameters event raised by the Contract contract.
type ContractCommitNewParameters struct {
	Deadline           *big.Int
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Contract *ContractFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*ContractCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &ContractCommitNewParametersIterator{contract: _Contract.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Contract *ContractFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *ContractCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCommitNewParameters)
				if err := _Contract.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitNewParameters is a log parse operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Contract *ContractFilterer) ParseCommitNewParameters(log types.Log) (*ContractCommitNewParameters, error) {
	event := new(ContractCommitNewParameters)
	if err := _Contract.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Contract contract.
type ContractNewAdminIterator struct {
	Event *ContractNewAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractNewAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewAdmin represents a NewAdmin event raised by the Contract contract.
type ContractNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Contract *ContractFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*ContractNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewAdminIterator{contract: _Contract.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Contract *ContractFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *ContractNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewAdmin)
				if err := _Contract.contract.UnpackLog(event, "NewAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Contract *ContractFilterer) ParseNewAdmin(log types.Log) (*ContractNewAdmin, error) {
	event := new(ContractNewAdmin)
	if err := _Contract.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Contract contract.
type ContractNewParametersIterator struct {
	Event *ContractNewParameters // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewParameters)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractNewParameters)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewParameters represents a NewParameters event raised by the Contract contract.
type ContractNewParameters struct {
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Contract *ContractFilterer) FilterNewParameters(opts *bind.FilterOpts) (*ContractNewParametersIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &ContractNewParametersIterator{contract: _Contract.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Contract *ContractFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *ContractNewParameters) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewParameters)
				if err := _Contract.contract.UnpackLog(event, "NewParameters", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewParameters is a log parse operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Contract *ContractFilterer) ParseNewParameters(log types.Log) (*ContractNewParameters, error) {
	event := new(ContractNewParameters)
	if err := _Contract.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRampAgammaIterator is returned from FilterRampAgamma and is used to iterate over the raw logs and unpacked data for RampAgamma events raised by the Contract contract.
type ContractRampAgammaIterator struct {
	Event *ContractRampAgamma // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRampAgammaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRampAgamma)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRampAgamma)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRampAgammaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRampAgammaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRampAgamma represents a RampAgamma event raised by the Contract contract.
type ContractRampAgamma struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialGamma *big.Int
	FutureGamma  *big.Int
	InitialTime  *big.Int
	FutureTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRampAgamma is a free log retrieval operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Contract *ContractFilterer) FilterRampAgamma(opts *bind.FilterOpts) (*ContractRampAgammaIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return &ContractRampAgammaIterator{contract: _Contract.contract, event: "RampAgamma", logs: logs, sub: sub}, nil
}

// WatchRampAgamma is a free log subscription operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Contract *ContractFilterer) WatchRampAgamma(opts *bind.WatchOpts, sink chan<- *ContractRampAgamma) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRampAgamma)
				if err := _Contract.contract.UnpackLog(event, "RampAgamma", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRampAgamma is a log parse operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Contract *ContractFilterer) ParseRampAgamma(log types.Log) (*ContractRampAgamma, error) {
	event := new(ContractRampAgamma)
	if err := _Contract.contract.UnpackLog(event, "RampAgamma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Contract contract.
type ContractRemoveLiquidityIterator struct {
	Event *ContractRemoveLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRemoveLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRemoveLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRemoveLiquidity represents a RemoveLiquidity event raised by the Contract contract.
type ContractRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Contract *ContractFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*ContractRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRemoveLiquidityIterator{contract: _Contract.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Contract *ContractFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *ContractRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRemoveLiquidity)
				if err := _Contract.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256 token_supply)
func (_Contract *ContractFilterer) ParseRemoveLiquidity(log types.Log) (*ContractRemoveLiquidity, error) {
	event := new(ContractRemoveLiquidity)
	if err := _Contract.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Contract contract.
type ContractRemoveLiquidityOneIterator struct {
	Event *ContractRemoveLiquidityOne // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRemoveLiquidityOne)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRemoveLiquidityOne)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Contract contract.
type ContractRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinIndex   *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Contract *ContractFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*ContractRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRemoveLiquidityOneIterator{contract: _Contract.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Contract *ContractFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *ContractRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRemoveLiquidityOne)
				if err := _Contract.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Contract *ContractFilterer) ParseRemoveLiquidityOne(log types.Log) (*ContractRemoveLiquidityOne, error) {
	event := new(ContractRemoveLiquidityOne)
	if err := _Contract.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Contract contract.
type ContractStopRampAIterator struct {
	Event *ContractStopRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStopRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractStopRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStopRampA represents a StopRampA event raised by the Contract contract.
type ContractStopRampA struct {
	CurrentA     *big.Int
	CurrentGamma *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Contract *ContractFilterer) FilterStopRampA(opts *bind.FilterOpts) (*ContractStopRampAIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &ContractStopRampAIterator{contract: _Contract.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Contract *ContractFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *ContractStopRampA) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStopRampA)
				if err := _Contract.contract.UnpackLog(event, "StopRampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStopRampA is a log parse operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Contract *ContractFilterer) ParseStopRampA(log types.Log) (*ContractStopRampA, error) {
	event := new(ContractStopRampA)
	if err := _Contract.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Contract contract.
type ContractTokenExchangeIterator struct {
	Event *ContractTokenExchange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTokenExchange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTokenExchange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTokenExchange represents a TokenExchange event raised by the Contract contract.
type ContractTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Contract *ContractFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*ContractTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTokenExchangeIterator{contract: _Contract.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Contract *ContractFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *ContractTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTokenExchange)
				if err := _Contract.contract.UnpackLog(event, "TokenExchange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenExchange is a log parse operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Contract *ContractFilterer) ParseTokenExchange(log types.Log) (*ContractTokenExchange, error) {
	event := new(ContractTokenExchange)
	if err := _Contract.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
