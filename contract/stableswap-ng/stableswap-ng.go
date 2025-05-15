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
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TokenExchangeUnderlying\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_id\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"fees\",\"type\":\"uint256[]\",\"indexed\":false},{\"name\":\"invariant\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampA\",\"inputs\":[{\"name\":\"old_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"new_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"t\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyNewFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"offpeg_fee_multiplier\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetNewMATime\",\"inputs\":[{\"name\":\"ma_exp_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"D_ma_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_offpeg_fee_multiplier\",\"type\":\"uint256\"},{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_coins\",\"type\":\"address[]\"},{\"name\":\"_rate_multipliers\",\"type\":\"uint256[]\"},{\"name\":\"_asset_types\",\"type\":\"uint8[]\"},{\"name\":\"_method_ids\",\"type\":\"bytes4[]\"},{\"name\":\"_oracles\",\"type\":\"address[]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_received\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"_min_received\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_imbalance\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_max_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_claim_admin_fees\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw_admin_fees\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ema_price\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_p\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dx\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dy\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A_precise\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"stored_rates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"dynamic_fee\",\"inputs\":[{\"name\":\"i\",\"type\":\"int128\"},{\"name\":\"j\",\"type\":\"int128\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A\",\"inputs\":[{\"name\":\"_future_A\",\"type\":\"uint256\"},{\"name\":\"_future_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_new_fee\",\"inputs\":[{\"name\":\"_new_fee\",\"type\":\"uint256\"},{\"name\":\"_new_offpeg_fee_multiplier\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_ma_exp_time\",\"inputs\":[{\"name\":\"_ma_exp_time\",\"type\":\"uint256\"},{\"name\":\"_D_ma_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"N_COINS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"offpeg_fee_multiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_exp_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D_ma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_last_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"salt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}]}]",
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

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_Contract *ContractCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_Contract *ContractSession) APrecise() (*big.Int, error) {
	return _Contract.Contract.APrecise(&_Contract.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_Contract *ContractCallerSession) APrecise() (*big.Int, error) {
	return _Contract.Contract.APrecise(&_Contract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_Contract *ContractCaller) DMaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "D_ma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_Contract *ContractSession) DMaTime() (*big.Int, error) {
	return _Contract.Contract.DMaTime(&_Contract.CallOpts)
}

// DMaTime is a free data retrieval call binding the contract method 0x9c4258c4.
//
// Solidity: function D_ma_time() view returns(uint256)
func (_Contract *ContractCallerSession) DMaTime() (*big.Int, error) {
	return _Contract.Contract.DMaTime(&_Contract.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_Contract *ContractCaller) DOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "D_oracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_Contract *ContractSession) DOracle() (*big.Int, error) {
	return _Contract.Contract.DOracle(&_Contract.CallOpts)
}

// DOracle is a free data retrieval call binding the contract method 0x907a016b.
//
// Solidity: function D_oracle() view returns(uint256)
func (_Contract *ContractCallerSession) DOracle() (*big.Int, error) {
	return _Contract.Contract.DOracle(&_Contract.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_Contract *ContractCaller) NCOINS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "N_COINS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_Contract *ContractSession) NCOINS() (*big.Int, error) {
	return _Contract.Contract.NCOINS(&_Contract.CallOpts)
}

// NCOINS is a free data retrieval call binding the contract method 0x29357750.
//
// Solidity: function N_COINS() view returns(uint256)
func (_Contract *ContractCallerSession) NCOINS() (*big.Int, error) {
	return _Contract.Contract.NCOINS(&_Contract.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_Contract *ContractCaller) AdminBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "admin_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_Contract *ContractSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.AdminBalances(&_Contract.CallOpts, arg0)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 arg0) view returns(uint256)
func (_Contract *ContractCallerSession) AdminBalances(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.AdminBalances(&_Contract.CallOpts, arg0)
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Contract *ContractSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Contract *ContractCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_Contract *ContractSession) Balances(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_Contract *ContractCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_Contract *ContractCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_Contract *ContractSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _Contract.Contract.CalcTokenAmount(&_Contract.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3db06dd8.
//
// Solidity: function calc_token_amount(uint256[] _amounts, bool _is_deposit) view returns(uint256)
func (_Contract *ContractCallerSession) CalcTokenAmount(_amounts []*big.Int, _is_deposit bool) (*big.Int, error) {
	return _Contract.Contract.CalcTokenAmount(&_Contract.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_Contract *ContractCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_Contract *ContractSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Contract.Contract.CalcWithdrawOneCoin(&_Contract.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_Contract *ContractCallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Contract.Contract.CalcWithdrawOneCoin(&_Contract.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Contract *ContractCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Contract *ContractSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Coins(&_Contract.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Contract *ContractCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Coins(&_Contract.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCallerSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_Contract *ContractCaller) DynamicFee(opts *bind.CallOpts, i *big.Int, j *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dynamic_fee", i, j)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_Contract *ContractSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _Contract.Contract.DynamicFee(&_Contract.CallOpts, i, j)
}

// DynamicFee is a free data retrieval call binding the contract method 0x76a9cd3e.
//
// Solidity: function dynamic_fee(int128 i, int128 j) view returns(uint256)
func (_Contract *ContractCallerSession) DynamicFee(i *big.Int, j *big.Int) (*big.Int, error) {
	return _Contract.Contract.DynamicFee(&_Contract.CallOpts, i, j)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_Contract *ContractCaller) EmaPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ema_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_Contract *ContractSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.EmaPrice(&_Contract.CallOpts, i)
}

// EmaPrice is a free data retrieval call binding the contract method 0x90d20837.
//
// Solidity: function ema_price(uint256 i) view returns(uint256)
func (_Contract *ContractCallerSession) EmaPrice(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.EmaPrice(&_Contract.CallOpts, i)
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

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Contract *ContractCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Contract *ContractSession) FutureA() (*big.Int, error) {
	return _Contract.Contract.FutureA(&_Contract.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_Contract *ContractCallerSession) FutureA() (*big.Int, error) {
	return _Contract.Contract.FutureA(&_Contract.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Contract *ContractCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Contract *ContractSession) FutureATime() (*big.Int, error) {
	return _Contract.Contract.FutureATime(&_Contract.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_Contract *ContractCallerSession) FutureATime() (*big.Int, error) {
	return _Contract.Contract.FutureATime(&_Contract.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_Contract *ContractCaller) GetBalances(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_Contract *ContractSession) GetBalances() ([]*big.Int, error) {
	return _Contract.Contract.GetBalances(&_Contract.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[])
func (_Contract *ContractCallerSession) GetBalances() ([]*big.Int, error) {
	return _Contract.Contract.GetBalances(&_Contract.CallOpts)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_Contract *ContractCaller) GetDx(opts *bind.CallOpts, i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_dx", i, j, dy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_Contract *ContractSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDx(&_Contract.CallOpts, i, j, dy)
}

// GetDx is a free data retrieval call binding the contract method 0x67df02ca.
//
// Solidity: function get_dx(int128 i, int128 j, uint256 dy) view returns(uint256)
func (_Contract *ContractCallerSession) GetDx(i *big.Int, j *big.Int, dy *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDx(&_Contract.CallOpts, i, j, dy)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Contract *ContractCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Contract *ContractSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDy(&_Contract.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_Contract *ContractCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetDy(&_Contract.CallOpts, i, j, dx)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_Contract *ContractCaller) GetP(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get_p", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_Contract *ContractSession) GetP(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetP(&_Contract.CallOpts, i)
}

// GetP is a free data retrieval call binding the contract method 0xec023862.
//
// Solidity: function get_p(uint256 i) view returns(uint256)
func (_Contract *ContractCallerSession) GetP(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetP(&_Contract.CallOpts, i)
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

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Contract *ContractCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Contract *ContractSession) InitialA() (*big.Int, error) {
	return _Contract.Contract.InitialA(&_Contract.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_Contract *ContractCallerSession) InitialA() (*big.Int, error) {
	return _Contract.Contract.InitialA(&_Contract.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Contract *ContractCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Contract *ContractSession) InitialATime() (*big.Int, error) {
	return _Contract.Contract.InitialATime(&_Contract.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_Contract *ContractCallerSession) InitialATime() (*big.Int, error) {
	return _Contract.Contract.InitialATime(&_Contract.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_Contract *ContractCaller) LastPrice(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "last_price", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_Contract *ContractSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.LastPrice(&_Contract.CallOpts, i)
}

// LastPrice is a free data retrieval call binding the contract method 0x3931ab52.
//
// Solidity: function last_price(uint256 i) view returns(uint256)
func (_Contract *ContractCallerSession) LastPrice(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.LastPrice(&_Contract.CallOpts, i)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_Contract *ContractCaller) MaExpTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ma_exp_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_Contract *ContractSession) MaExpTime() (*big.Int, error) {
	return _Contract.Contract.MaExpTime(&_Contract.CallOpts)
}

// MaExpTime is a free data retrieval call binding the contract method 0x1be913a5.
//
// Solidity: function ma_exp_time() view returns(uint256)
func (_Contract *ContractCallerSession) MaExpTime() (*big.Int, error) {
	return _Contract.Contract.MaExpTime(&_Contract.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_Contract *ContractCaller) MaLastTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ma_last_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_Contract *ContractSession) MaLastTime() (*big.Int, error) {
	return _Contract.Contract.MaLastTime(&_Contract.CallOpts)
}

// MaLastTime is a free data retrieval call binding the contract method 0x1ddc3b01.
//
// Solidity: function ma_last_time() view returns(uint256)
func (_Contract *ContractCallerSession) MaLastTime() (*big.Int, error) {
	return _Contract.Contract.MaLastTime(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Contract *ContractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Contract *ContractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonces(&_Contract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_Contract *ContractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonces(&_Contract.CallOpts, arg0)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_Contract *ContractCaller) OffpegFeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "offpeg_fee_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_Contract *ContractSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _Contract.Contract.OffpegFeeMultiplier(&_Contract.CallOpts)
}

// OffpegFeeMultiplier is a free data retrieval call binding the contract method 0x8edfdd5f.
//
// Solidity: function offpeg_fee_multiplier() view returns(uint256)
func (_Contract *ContractCallerSession) OffpegFeeMultiplier() (*big.Int, error) {
	return _Contract.Contract.OffpegFeeMultiplier(&_Contract.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_Contract *ContractCaller) PriceOracle(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "price_oracle", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_Contract *ContractSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.PriceOracle(&_Contract.CallOpts, i)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 i) view returns(uint256)
func (_Contract *ContractCallerSession) PriceOracle(i *big.Int) (*big.Int, error) {
	return _Contract.Contract.PriceOracle(&_Contract.CallOpts, i)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Contract *ContractCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Contract *ContractSession) Salt() ([32]byte, error) {
	return _Contract.Contract.Salt(&_Contract.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Contract *ContractCallerSession) Salt() ([32]byte, error) {
	return _Contract.Contract.Salt(&_Contract.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_Contract *ContractCaller) StoredRates(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stored_rates")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_Contract *ContractSession) StoredRates() ([]*big.Int, error) {
	return _Contract.Contract.StoredRates(&_Contract.CallOpts)
}

// StoredRates is a free data retrieval call binding the contract method 0xfd0684b1.
//
// Solidity: function stored_rates() view returns(uint256[])
func (_Contract *ContractCallerSession) StoredRates() ([]*big.Int, error) {
	return _Contract.Contract.StoredRates(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contract *ContractSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Contract *ContractCallerSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_Contract *ContractTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_Contract *ContractSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity(&_Contract.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xb72df5de.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount) returns(uint256)
func (_Contract *ContractTransactorSession) AddLiquidity(_amounts []*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity(&_Contract.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_Contract *ContractTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_Contract *ContractSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity0(&_Contract.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xa7256d09.
//
// Solidity: function add_liquidity(uint256[] _amounts, uint256 _min_mint_amount, address _receiver) returns(uint256)
func (_Contract *ContractTransactorSession) AddLiquidity0(_amounts []*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddLiquidity0(&_Contract.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Contract *ContractSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Contract *ContractTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _spender, _value)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Contract *ContractTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchange", i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Contract *ContractSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Exchange(&_Contract.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Contract *ContractTransactorSession) Exchange(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Exchange(&_Contract.TransactOpts, i, j, _dx, _min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Contract *ContractTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchange0", i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Contract *ContractSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Exchange0(&_Contract.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Exchange0 is a paid mutator transaction binding the contract method 0xddc1f59d.
//
// Solidity: function exchange(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Contract *ContractTransactorSession) Exchange0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Exchange0(&_Contract.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Contract *ContractTransactor) ExchangeReceived(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchange_received", i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Contract *ContractSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeReceived(&_Contract.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived is a paid mutator transaction binding the contract method 0x7e3db030.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy) returns(uint256)
func (_Contract *ContractTransactorSession) ExchangeReceived(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeReceived(&_Contract.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Contract *ContractTransactor) ExchangeReceived0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exchange_received0", i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Contract *ContractSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeReceived0(&_Contract.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeReceived0 is a paid mutator transaction binding the contract method 0xafb43012.
//
// Solidity: function exchange_received(int128 i, int128 j, uint256 _dx, uint256 _min_dy, address _receiver) returns(uint256)
func (_Contract *ContractTransactorSession) ExchangeReceived0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ExchangeReceived0(&_Contract.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_Contract *ContractTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_Contract *ContractSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Permit(&_Contract.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns(bool)
func (_Contract *ContractTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Permit(&_Contract.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Contract *ContractTransactor) RampA(opts *bind.TransactOpts, _future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "ramp_A", _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Contract *ContractSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RampA(&_Contract.TransactOpts, _future_A, _future_time)
}

// RampA is a paid mutator transaction binding the contract method 0x3c157e64.
//
// Solidity: function ramp_A(uint256 _future_A, uint256 _future_time) returns()
func (_Contract *ContractTransactorSession) RampA(_future_A *big.Int, _future_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RampA(&_Contract.TransactOpts, _future_A, _future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_Contract *ContractTransactor) RemoveLiquidity(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity", _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_Contract *ContractSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity(&_Contract.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd40ddb8c.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts) returns(uint256[])
func (_Contract *ContractTransactorSession) RemoveLiquidity(_burn_amount *big.Int, _min_amounts []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity(&_Contract.TransactOpts, _burn_amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_Contract *ContractTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity0", _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_Contract *ContractSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity0(&_Contract.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x5e604cd2.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver) returns(uint256[])
func (_Contract *ContractTransactorSession) RemoveLiquidity0(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity0(&_Contract.TransactOpts, _burn_amount, _min_amounts, _receiver)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_Contract *ContractTransactor) RemoveLiquidity1(opts *bind.TransactOpts, _burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity1", _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_Contract *ContractSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity1(&_Contract.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x2969e04a.
//
// Solidity: function remove_liquidity(uint256 _burn_amount, uint256[] _min_amounts, address _receiver, bool _claim_admin_fees) returns(uint256[])
func (_Contract *ContractTransactorSession) RemoveLiquidity1(_burn_amount *big.Int, _min_amounts []*big.Int, _receiver common.Address, _claim_admin_fees bool) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidity1(&_Contract.TransactOpts, _burn_amount, _min_amounts, _receiver, _claim_admin_fees)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_Contract *ContractTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity_imbalance", _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_Contract *ContractSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityImbalance(&_Contract.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x7706db75.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount) returns(uint256)
func (_Contract *ContractTransactorSession) RemoveLiquidityImbalance(_amounts []*big.Int, _max_burn_amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityImbalance(&_Contract.TransactOpts, _amounts, _max_burn_amount)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_Contract *ContractTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, _amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity_imbalance0", _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_Contract *ContractSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityImbalance0(&_Contract.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x4a6e32c6.
//
// Solidity: function remove_liquidity_imbalance(uint256[] _amounts, uint256 _max_burn_amount, address _receiver) returns(uint256)
func (_Contract *ContractTransactorSession) RemoveLiquidityImbalance0(_amounts []*big.Int, _max_burn_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityImbalance0(&_Contract.TransactOpts, _amounts, _max_burn_amount, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_Contract *ContractTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity_one_coin", _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_Contract *ContractSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin(&_Contract.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received) returns(uint256)
func (_Contract *ContractTransactorSession) RemoveLiquidityOneCoin(_burn_amount *big.Int, i *big.Int, _min_received *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin(&_Contract.TransactOpts, _burn_amount, i, _min_received)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_Contract *ContractTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "remove_liquidity_one_coin0", _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_Contract *ContractSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin0(&_Contract.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x081579a5.
//
// Solidity: function remove_liquidity_one_coin(uint256 _burn_amount, int128 i, uint256 _min_received, address _receiver) returns(uint256)
func (_Contract *ContractTransactorSession) RemoveLiquidityOneCoin0(_burn_amount *big.Int, i *big.Int, _min_received *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveLiquidityOneCoin0(&_Contract.TransactOpts, _burn_amount, i, _min_received, _receiver)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_Contract *ContractTransactor) SetMaExpTime(opts *bind.TransactOpts, _ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_ma_exp_time", _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_Contract *ContractSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaExpTime(&_Contract.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetMaExpTime is a paid mutator transaction binding the contract method 0x65bbea6b.
//
// Solidity: function set_ma_exp_time(uint256 _ma_exp_time, uint256 _D_ma_time) returns()
func (_Contract *ContractTransactorSession) SetMaExpTime(_ma_exp_time *big.Int, _D_ma_time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaExpTime(&_Contract.TransactOpts, _ma_exp_time, _D_ma_time)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_Contract *ContractTransactor) SetNewFee(opts *bind.TransactOpts, _new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set_new_fee", _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_Contract *ContractSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetNewFee(&_Contract.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// SetNewFee is a paid mutator transaction binding the contract method 0x015c2838.
//
// Solidity: function set_new_fee(uint256 _new_fee, uint256 _new_offpeg_fee_multiplier) returns()
func (_Contract *ContractTransactorSession) SetNewFee(_new_fee *big.Int, _new_offpeg_fee_multiplier *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetNewFee(&_Contract.TransactOpts, _new_fee, _new_offpeg_fee_multiplier)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Contract *ContractTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stop_ramp_A")
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Contract *ContractSession) StopRampA() (*types.Transaction, error) {
	return _Contract.Contract.StopRampA(&_Contract.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0x551a6588.
//
// Solidity: function stop_ramp_A() returns()
func (_Contract *ContractTransactorSession) StopRampA() (*types.Transaction, error) {
	return _Contract.Contract.StopRampA(&_Contract.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Contract *ContractSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Contract *ContractTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Contract *ContractSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, _from, _to, _value)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Contract *ContractTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw_admin_fees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Contract *ContractSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAdminFees(&_Contract.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x30c54085.
//
// Solidity: function withdraw_admin_fees() returns()
func (_Contract *ContractTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAdminFees(&_Contract.TransactOpts)
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
	TokenAmounts []*big.Int
	Fees         []*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
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

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_Contract *ContractFilterer) ParseAddLiquidity(log types.Log) (*ContractAddLiquidity, error) {
	event := new(ContractAddLiquidity)
	if err := _Contract.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApplyNewFeeIterator is returned from FilterApplyNewFee and is used to iterate over the raw logs and unpacked data for ApplyNewFee events raised by the Contract contract.
type ContractApplyNewFeeIterator struct {
	Event *ContractApplyNewFee // Event containing the contract specifics and raw log

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
func (it *ContractApplyNewFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApplyNewFee)
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
		it.Event = new(ContractApplyNewFee)
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
func (it *ContractApplyNewFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApplyNewFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApplyNewFee represents a ApplyNewFee event raised by the Contract contract.
type ContractApplyNewFee struct {
	Fee                 *big.Int
	OffpegFeeMultiplier *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterApplyNewFee is a free log retrieval operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_Contract *ContractFilterer) FilterApplyNewFee(opts *bind.FilterOpts) (*ContractApplyNewFeeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return &ContractApplyNewFeeIterator{contract: _Contract.contract, event: "ApplyNewFee", logs: logs, sub: sub}, nil
}

// WatchApplyNewFee is a free log subscription operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_Contract *ContractFilterer) WatchApplyNewFee(opts *bind.WatchOpts, sink chan<- *ContractApplyNewFee) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApplyNewFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApplyNewFee)
				if err := _Contract.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
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

// ParseApplyNewFee is a log parse operation binding the contract event 0x750d10a7f37466ce785ee6bcb604aac543358db42afbcc332a3c12a49c80bf6d.
//
// Solidity: event ApplyNewFee(uint256 fee, uint256 offpeg_fee_multiplier)
func (_Contract *ContractFilterer) ParseApplyNewFee(log types.Log) (*ContractApplyNewFee, error) {
	event := new(ContractApplyNewFee)
	if err := _Contract.contract.UnpackLog(event, "ApplyNewFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the Contract contract.
type ContractRampAIterator struct {
	Event *ContractRampA // Event containing the contract specifics and raw log

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
func (it *ContractRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRampA)
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
		it.Event = new(ContractRampA)
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
func (it *ContractRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRampA represents a RampA event raised by the Contract contract.
type ContractRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Contract *ContractFilterer) FilterRampA(opts *bind.FilterOpts) (*ContractRampAIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &ContractRampAIterator{contract: _Contract.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Contract *ContractFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *ContractRampA) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRampA)
				if err := _Contract.contract.UnpackLog(event, "RampA", log); err != nil {
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

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 old_A, uint256 new_A, uint256 initial_time, uint256 future_time)
func (_Contract *ContractFilterer) ParseRampA(log types.Log) (*ContractRampA, error) {
	event := new(ContractRampA)
	if err := _Contract.contract.UnpackLog(event, "RampA", log); err != nil {
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
	TokenAmounts []*big.Int
	Fees         []*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
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

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x347ad828e58cbe534d8f6b67985d791360756b18f0d95fd9f197a66cc46480ea.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 token_supply)
func (_Contract *ContractFilterer) ParseRemoveLiquidity(log types.Log) (*ContractRemoveLiquidity, error) {
	event := new(ContractRemoveLiquidity)
	if err := _Contract.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Contract contract.
type ContractRemoveLiquidityImbalanceIterator struct {
	Event *ContractRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *ContractRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRemoveLiquidityImbalance)
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
		it.Event = new(ContractRemoveLiquidityImbalance)
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
func (it *ContractRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Contract contract.
type ContractRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts []*big.Int
	Fees         []*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_Contract *ContractFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*ContractRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRemoveLiquidityImbalanceIterator{contract: _Contract.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_Contract *ContractFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *ContractRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRemoveLiquidityImbalance)
				if err := _Contract.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] token_amounts, uint256[] fees, uint256 invariant, uint256 token_supply)
func (_Contract *ContractFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*ContractRemoveLiquidityImbalance, error) {
	event := new(ContractRemoveLiquidityImbalance)
	if err := _Contract.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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
	TokenId     *big.Int
	TokenAmount *big.Int
	CoinAmount  *big.Int
	TokenSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
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

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x6f48129db1f37ccb9cc5dd7e119cb32750cabdf75b48375d730d26ce3659bbe1.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, int128 token_id, uint256 token_amount, uint256 coin_amount, uint256 token_supply)
func (_Contract *ContractFilterer) ParseRemoveLiquidityOne(log types.Log) (*ContractRemoveLiquidityOne, error) {
	event := new(ContractRemoveLiquidityOne)
	if err := _Contract.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetNewMATimeIterator is returned from FilterSetNewMATime and is used to iterate over the raw logs and unpacked data for SetNewMATime events raised by the Contract contract.
type ContractSetNewMATimeIterator struct {
	Event *ContractSetNewMATime // Event containing the contract specifics and raw log

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
func (it *ContractSetNewMATimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetNewMATime)
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
		it.Event = new(ContractSetNewMATime)
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
func (it *ContractSetNewMATimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetNewMATimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetNewMATime represents a SetNewMATime event raised by the Contract contract.
type ContractSetNewMATime struct {
	MaExpTime *big.Int
	DMaTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetNewMATime is a free log retrieval operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_Contract *ContractFilterer) FilterSetNewMATime(opts *bind.FilterOpts) (*ContractSetNewMATimeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return &ContractSetNewMATimeIterator{contract: _Contract.contract, event: "SetNewMATime", logs: logs, sub: sub}, nil
}

// WatchSetNewMATime is a free log subscription operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_Contract *ContractFilterer) WatchSetNewMATime(opts *bind.WatchOpts, sink chan<- *ContractSetNewMATime) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetNewMATime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetNewMATime)
				if err := _Contract.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
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

// ParseSetNewMATime is a log parse operation binding the contract event 0x68dc4e067dff1862b896b7a0faf55f97df1a60d0aaa79481b69d675f2026a28c.
//
// Solidity: event SetNewMATime(uint256 ma_exp_time, uint256 D_ma_time)
func (_Contract *ContractFilterer) ParseSetNewMATime(log types.Log) (*ContractSetNewMATime, error) {
	event := new(ContractSetNewMATime)
	if err := _Contract.contract.UnpackLog(event, "SetNewMATime", log); err != nil {
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
	A   *big.Int
	T   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
func (_Contract *ContractFilterer) FilterStopRampA(opts *bind.FilterOpts) (*ContractStopRampAIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &ContractStopRampAIterator{contract: _Contract.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
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

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 A, uint256 t)
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

// FilterTokenExchange is a free log retrieval operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
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

// WatchTokenExchange is a free log subscription operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
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

// ParseTokenExchange is a log parse operation binding the contract event 0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140.
//
// Solidity: event TokenExchange(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Contract *ContractFilterer) ParseTokenExchange(log types.Log) (*ContractTokenExchange, error) {
	event := new(ContractTokenExchange)
	if err := _Contract.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTokenExchangeUnderlyingIterator is returned from FilterTokenExchangeUnderlying and is used to iterate over the raw logs and unpacked data for TokenExchangeUnderlying events raised by the Contract contract.
type ContractTokenExchangeUnderlyingIterator struct {
	Event *ContractTokenExchangeUnderlying // Event containing the contract specifics and raw log

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
func (it *ContractTokenExchangeUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTokenExchangeUnderlying)
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
		it.Event = new(ContractTokenExchangeUnderlying)
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
func (it *ContractTokenExchangeUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTokenExchangeUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTokenExchangeUnderlying represents a TokenExchangeUnderlying event raised by the Contract contract.
type ContractTokenExchangeUnderlying struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangeUnderlying is a free log retrieval operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Contract *ContractFilterer) FilterTokenExchangeUnderlying(opts *bind.FilterOpts, buyer []common.Address) (*ContractTokenExchangeUnderlyingIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTokenExchangeUnderlyingIterator{contract: _Contract.contract, event: "TokenExchangeUnderlying", logs: logs, sub: sub}, nil
}

// WatchTokenExchangeUnderlying is a free log subscription operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Contract *ContractFilterer) WatchTokenExchangeUnderlying(opts *bind.WatchOpts, sink chan<- *ContractTokenExchangeUnderlying, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TokenExchangeUnderlying", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTokenExchangeUnderlying)
				if err := _Contract.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
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

// ParseTokenExchangeUnderlying is a log parse operation binding the contract event 0xd013ca23e77a65003c2c659c5442c00c805371b7fc1ebd4c206c41d1536bd90b.
//
// Solidity: event TokenExchangeUnderlying(address indexed buyer, int128 sold_id, uint256 tokens_sold, int128 bought_id, uint256 tokens_bought)
func (_Contract *ContractFilterer) ParseTokenExchangeUnderlying(log types.Log) (*ContractTokenExchangeUnderlying, error) {
	event := new(ContractTokenExchangeUnderlying)
	if err := _Contract.contract.UnpackLog(event, "TokenExchangeUnderlying", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ContractTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
