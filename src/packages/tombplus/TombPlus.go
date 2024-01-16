// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tombplus

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

// TombplusMetaData contains all meta data concerning the Tombplus contract.
var TombplusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITombPlus\",\"name\":\"_tombplus\",\"type\":\"address\"},{\"internalType\":\"contractIOTombPlus\",\"name\":\"_otombplus\",\"type\":\"address\"},{\"internalType\":\"contractITsharePlus\",\"name\":\"_tshareplus\",\"type\":\"address\"},{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochEndPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"volatilityReached\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"winnerIsUp\",\"type\":\"bool\"}],\"name\":\"EpochFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"volatilityThreshold\",\"type\":\"uint256\"}],\"name\":\"EpochStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStartPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usersUp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votesUp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usersDown\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votesDown\",\"type\":\"uint256\"}],\"name\":\"ObservationStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"flipId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"UserFlipped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochAcceptingFlips\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochEnteredPriceObservation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"epochFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochOutcome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"volatilityReached\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"winnerIsUp\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStartPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochStartTshareTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"flip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"epochIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"ups\",\"type\":\"bool[]\"}],\"name\":\"flipMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getUserFlipIdByEpochId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllowedFutureFlips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otombplus\",\"outputs\":[{\"internalType\":\"contractIOTombPlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"processUserFlips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processedRewardBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rewardBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTshareNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAllowedFutureFlips\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerTshareNumerator\",\"type\":\"uint256\"}],\"name\":\"setCfg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochDuration\",\"type\":\"uint256\"}],\"name\":\"setEpochDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flipWindow\",\"type\":\"uint256\"}],\"name\":\"setFlipWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOTombPlus\",\"name\":\"_otombplus\",\"type\":\"address\"}],\"name\":\"setOTombPlus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_volatilityThresholdBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_volatilityThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_volatilityThresholdStep\",\"type\":\"uint256\"}],\"name\":\"setVolatilityThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tombplus\",\"outputs\":[{\"internalType\":\"contractITombPlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"direct\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"renounce\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tshareplus\",\"outputs\":[{\"internalType\":\"contractITsharePlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"updateUserTsharePlusBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userFlipIdsByEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userFlips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tsharePlusBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatilityThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatilityThresholdBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatilityThresholdStep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TombplusABI is the input ABI used to generate the binding from.
// Deprecated: Use TombplusMetaData.ABI instead.
var TombplusABI = TombplusMetaData.ABI

// Tombplus is an auto generated Go binding around an Ethereum contract.
type Tombplus struct {
	TombplusCaller     // Read-only binding to the contract
	TombplusTransactor // Write-only binding to the contract
	TombplusFilterer   // Log filterer for contract events
}

// TombplusCaller is an auto generated read-only Go binding around an Ethereum contract.
type TombplusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TombplusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TombplusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TombplusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TombplusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TombplusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TombplusSession struct {
	Contract     *Tombplus         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TombplusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TombplusCallerSession struct {
	Contract *TombplusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TombplusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TombplusTransactorSession struct {
	Contract     *TombplusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TombplusRaw is an auto generated low-level Go binding around an Ethereum contract.
type TombplusRaw struct {
	Contract *Tombplus // Generic contract binding to access the raw methods on
}

// TombplusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TombplusCallerRaw struct {
	Contract *TombplusCaller // Generic read-only contract binding to access the raw methods on
}

// TombplusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TombplusTransactorRaw struct {
	Contract *TombplusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTombplus creates a new instance of Tombplus, bound to a specific deployed contract.
func NewTombplus(address common.Address, backend bind.ContractBackend) (*Tombplus, error) {
	contract, err := bindTombplus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tombplus{TombplusCaller: TombplusCaller{contract: contract}, TombplusTransactor: TombplusTransactor{contract: contract}, TombplusFilterer: TombplusFilterer{contract: contract}}, nil
}

// NewTombplusCaller creates a new read-only instance of Tombplus, bound to a specific deployed contract.
func NewTombplusCaller(address common.Address, caller bind.ContractCaller) (*TombplusCaller, error) {
	contract, err := bindTombplus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TombplusCaller{contract: contract}, nil
}

// NewTombplusTransactor creates a new write-only instance of Tombplus, bound to a specific deployed contract.
func NewTombplusTransactor(address common.Address, transactor bind.ContractTransactor) (*TombplusTransactor, error) {
	contract, err := bindTombplus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TombplusTransactor{contract: contract}, nil
}

// NewTombplusFilterer creates a new log filterer instance of Tombplus, bound to a specific deployed contract.
func NewTombplusFilterer(address common.Address, filterer bind.ContractFilterer) (*TombplusFilterer, error) {
	contract, err := bindTombplus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TombplusFilterer{contract: contract}, nil
}

// bindTombplus binds a generic wrapper to an already deployed contract.
func bindTombplus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TombplusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tombplus *TombplusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tombplus.Contract.TombplusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tombplus *TombplusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.Contract.TombplusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tombplus *TombplusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tombplus.Contract.TombplusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tombplus *TombplusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tombplus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tombplus *TombplusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tombplus *TombplusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tombplus.Contract.contract.Transact(opts, method, params...)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_Tombplus *TombplusCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_Tombplus *TombplusSession) Denominator() (*big.Int, error) {
	return _Tombplus.Contract.Denominator(&_Tombplus.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_Tombplus *TombplusCallerSession) Denominator() (*big.Int, error) {
	return _Tombplus.Contract.Denominator(&_Tombplus.CallOpts)
}

// EpochAcceptingFlips is a free data retrieval call binding the contract method 0xa5cbc898.
//
// Solidity: function epochAcceptingFlips() view returns(bool)
func (_Tombplus *TombplusCaller) EpochAcceptingFlips(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochAcceptingFlips")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochAcceptingFlips is a free data retrieval call binding the contract method 0xa5cbc898.
//
// Solidity: function epochAcceptingFlips() view returns(bool)
func (_Tombplus *TombplusSession) EpochAcceptingFlips() (bool, error) {
	return _Tombplus.Contract.EpochAcceptingFlips(&_Tombplus.CallOpts)
}

// EpochAcceptingFlips is a free data retrieval call binding the contract method 0xa5cbc898.
//
// Solidity: function epochAcceptingFlips() view returns(bool)
func (_Tombplus *TombplusCallerSession) EpochAcceptingFlips() (bool, error) {
	return _Tombplus.Contract.EpochAcceptingFlips(&_Tombplus.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Tombplus *TombplusCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Tombplus *TombplusSession) EpochDuration() (*big.Int, error) {
	return _Tombplus.Contract.EpochDuration(&_Tombplus.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint256)
func (_Tombplus *TombplusCallerSession) EpochDuration() (*big.Int, error) {
	return _Tombplus.Contract.EpochDuration(&_Tombplus.CallOpts)
}

// EpochEnteredPriceObservation is a free data retrieval call binding the contract method 0x1e03f37b.
//
// Solidity: function epochEnteredPriceObservation() view returns(bool)
func (_Tombplus *TombplusCaller) EpochEnteredPriceObservation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochEnteredPriceObservation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochEnteredPriceObservation is a free data retrieval call binding the contract method 0x1e03f37b.
//
// Solidity: function epochEnteredPriceObservation() view returns(bool)
func (_Tombplus *TombplusSession) EpochEnteredPriceObservation() (bool, error) {
	return _Tombplus.Contract.EpochEnteredPriceObservation(&_Tombplus.CallOpts)
}

// EpochEnteredPriceObservation is a free data retrieval call binding the contract method 0x1e03f37b.
//
// Solidity: function epochEnteredPriceObservation() view returns(bool)
func (_Tombplus *TombplusCallerSession) EpochEnteredPriceObservation() (bool, error) {
	return _Tombplus.Contract.EpochEnteredPriceObservation(&_Tombplus.CallOpts)
}

// EpochFinished is a free data retrieval call binding the contract method 0xa5888a27.
//
// Solidity: function epochFinished(uint256 epochId) view returns(bool)
func (_Tombplus *TombplusCaller) EpochFinished(opts *bind.CallOpts, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochFinished", epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochFinished is a free data retrieval call binding the contract method 0xa5888a27.
//
// Solidity: function epochFinished(uint256 epochId) view returns(bool)
func (_Tombplus *TombplusSession) EpochFinished(epochId *big.Int) (bool, error) {
	return _Tombplus.Contract.EpochFinished(&_Tombplus.CallOpts, epochId)
}

// EpochFinished is a free data retrieval call binding the contract method 0xa5888a27.
//
// Solidity: function epochFinished(uint256 epochId) view returns(bool)
func (_Tombplus *TombplusCallerSession) EpochFinished(epochId *big.Int) (bool, error) {
	return _Tombplus.Contract.EpochFinished(&_Tombplus.CallOpts, epochId)
}

// EpochNumber is a free data retrieval call binding the contract method 0xf4145a83.
//
// Solidity: function epochNumber() view returns(uint256)
func (_Tombplus *TombplusCaller) EpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochNumber is a free data retrieval call binding the contract method 0xf4145a83.
//
// Solidity: function epochNumber() view returns(uint256)
func (_Tombplus *TombplusSession) EpochNumber() (*big.Int, error) {
	return _Tombplus.Contract.EpochNumber(&_Tombplus.CallOpts)
}

// EpochNumber is a free data retrieval call binding the contract method 0xf4145a83.
//
// Solidity: function epochNumber() view returns(uint256)
func (_Tombplus *TombplusCallerSession) EpochNumber() (*big.Int, error) {
	return _Tombplus.Contract.EpochNumber(&_Tombplus.CallOpts)
}

// EpochOutcome is a free data retrieval call binding the contract method 0x4f764e48.
//
// Solidity: function epochOutcome() view returns(uint256 price, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusCaller) EpochOutcome(opts *bind.CallOpts) (struct {
	Price             *big.Int
	VolatilityReached bool
	WinnerIsUp        bool
}, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochOutcome")

	outstruct := new(struct {
		Price             *big.Int
		VolatilityReached bool
		WinnerIsUp        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VolatilityReached = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.WinnerIsUp = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// EpochOutcome is a free data retrieval call binding the contract method 0x4f764e48.
//
// Solidity: function epochOutcome() view returns(uint256 price, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusSession) EpochOutcome() (struct {
	Price             *big.Int
	VolatilityReached bool
	WinnerIsUp        bool
}, error) {
	return _Tombplus.Contract.EpochOutcome(&_Tombplus.CallOpts)
}

// EpochOutcome is a free data retrieval call binding the contract method 0x4f764e48.
//
// Solidity: function epochOutcome() view returns(uint256 price, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusCallerSession) EpochOutcome() (struct {
	Price             *big.Int
	VolatilityReached bool
	WinnerIsUp        bool
}, error) {
	return _Tombplus.Contract.EpochOutcome(&_Tombplus.CallOpts)
}

// EpochStartPrice is a free data retrieval call binding the contract method 0x84be9a83.
//
// Solidity: function epochStartPrice() view returns(uint256)
func (_Tombplus *TombplusCaller) EpochStartPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochStartPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStartPrice is a free data retrieval call binding the contract method 0x84be9a83.
//
// Solidity: function epochStartPrice() view returns(uint256)
func (_Tombplus *TombplusSession) EpochStartPrice() (*big.Int, error) {
	return _Tombplus.Contract.EpochStartPrice(&_Tombplus.CallOpts)
}

// EpochStartPrice is a free data retrieval call binding the contract method 0x84be9a83.
//
// Solidity: function epochStartPrice() view returns(uint256)
func (_Tombplus *TombplusCallerSession) EpochStartPrice() (*big.Int, error) {
	return _Tombplus.Contract.EpochStartPrice(&_Tombplus.CallOpts)
}

// EpochStartTshareTotalSupply is a free data retrieval call binding the contract method 0x3a8b4c8a.
//
// Solidity: function epochStartTshareTotalSupply() view returns(uint256)
func (_Tombplus *TombplusCaller) EpochStartTshareTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "epochStartTshareTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochStartTshareTotalSupply is a free data retrieval call binding the contract method 0x3a8b4c8a.
//
// Solidity: function epochStartTshareTotalSupply() view returns(uint256)
func (_Tombplus *TombplusSession) EpochStartTshareTotalSupply() (*big.Int, error) {
	return _Tombplus.Contract.EpochStartTshareTotalSupply(&_Tombplus.CallOpts)
}

// EpochStartTshareTotalSupply is a free data retrieval call binding the contract method 0x3a8b4c8a.
//
// Solidity: function epochStartTshareTotalSupply() view returns(uint256)
func (_Tombplus *TombplusCallerSession) EpochStartTshareTotalSupply() (*big.Int, error) {
	return _Tombplus.Contract.EpochStartTshareTotalSupply(&_Tombplus.CallOpts)
}

// FlipWindow is a free data retrieval call binding the contract method 0x43f6ab38.
//
// Solidity: function flipWindow() view returns(uint256)
func (_Tombplus *TombplusCaller) FlipWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "flipWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlipWindow is a free data retrieval call binding the contract method 0x43f6ab38.
//
// Solidity: function flipWindow() view returns(uint256)
func (_Tombplus *TombplusSession) FlipWindow() (*big.Int, error) {
	return _Tombplus.Contract.FlipWindow(&_Tombplus.CallOpts)
}

// FlipWindow is a free data retrieval call binding the contract method 0x43f6ab38.
//
// Solidity: function flipWindow() view returns(uint256)
func (_Tombplus *TombplusCallerSession) FlipWindow() (*big.Int, error) {
	return _Tombplus.Contract.FlipWindow(&_Tombplus.CallOpts)
}

// GameStarted is a free data retrieval call binding the contract method 0x5e123ce4.
//
// Solidity: function gameStarted() view returns(bool)
func (_Tombplus *TombplusCaller) GameStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "gameStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GameStarted is a free data retrieval call binding the contract method 0x5e123ce4.
//
// Solidity: function gameStarted() view returns(bool)
func (_Tombplus *TombplusSession) GameStarted() (bool, error) {
	return _Tombplus.Contract.GameStarted(&_Tombplus.CallOpts)
}

// GameStarted is a free data retrieval call binding the contract method 0x5e123ce4.
//
// Solidity: function gameStarted() view returns(bool)
func (_Tombplus *TombplusCallerSession) GameStarted() (bool, error) {
	return _Tombplus.Contract.GameStarted(&_Tombplus.CallOpts)
}

// GetUserFlipIdByEpochId is a free data retrieval call binding the contract method 0x968277ea.
//
// Solidity: function getUserFlipIdByEpochId(address user, uint256 epochId) view returns(bool found, uint256 id)
func (_Tombplus *TombplusCaller) GetUserFlipIdByEpochId(opts *bind.CallOpts, user common.Address, epochId *big.Int) (struct {
	Found bool
	Id    *big.Int
}, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "getUserFlipIdByEpochId", user, epochId)

	outstruct := new(struct {
		Found bool
		Id    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Found = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserFlipIdByEpochId is a free data retrieval call binding the contract method 0x968277ea.
//
// Solidity: function getUserFlipIdByEpochId(address user, uint256 epochId) view returns(bool found, uint256 id)
func (_Tombplus *TombplusSession) GetUserFlipIdByEpochId(user common.Address, epochId *big.Int) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _Tombplus.Contract.GetUserFlipIdByEpochId(&_Tombplus.CallOpts, user, epochId)
}

// GetUserFlipIdByEpochId is a free data retrieval call binding the contract method 0x968277ea.
//
// Solidity: function getUserFlipIdByEpochId(address user, uint256 epochId) view returns(bool found, uint256 id)
func (_Tombplus *TombplusCallerSession) GetUserFlipIdByEpochId(user common.Address, epochId *big.Int) (struct {
	Found bool
	Id    *big.Int
}, error) {
	return _Tombplus.Contract.GetUserFlipIdByEpochId(&_Tombplus.CallOpts, user, epochId)
}

// MaxAllowedFutureFlips is a free data retrieval call binding the contract method 0x71bc2606.
//
// Solidity: function maxAllowedFutureFlips() view returns(uint256)
func (_Tombplus *TombplusCaller) MaxAllowedFutureFlips(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "maxAllowedFutureFlips")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAllowedFutureFlips is a free data retrieval call binding the contract method 0x71bc2606.
//
// Solidity: function maxAllowedFutureFlips() view returns(uint256)
func (_Tombplus *TombplusSession) MaxAllowedFutureFlips() (*big.Int, error) {
	return _Tombplus.Contract.MaxAllowedFutureFlips(&_Tombplus.CallOpts)
}

// MaxAllowedFutureFlips is a free data retrieval call binding the contract method 0x71bc2606.
//
// Solidity: function maxAllowedFutureFlips() view returns(uint256)
func (_Tombplus *TombplusCallerSession) MaxAllowedFutureFlips() (*big.Int, error) {
	return _Tombplus.Contract.MaxAllowedFutureFlips(&_Tombplus.CallOpts)
}

// Otombplus is a free data retrieval call binding the contract method 0x5ef5d76c.
//
// Solidity: function otombplus() view returns(address)
func (_Tombplus *TombplusCaller) Otombplus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "otombplus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Otombplus is a free data retrieval call binding the contract method 0x5ef5d76c.
//
// Solidity: function otombplus() view returns(address)
func (_Tombplus *TombplusSession) Otombplus() (common.Address, error) {
	return _Tombplus.Contract.Otombplus(&_Tombplus.CallOpts)
}

// Otombplus is a free data retrieval call binding the contract method 0x5ef5d76c.
//
// Solidity: function otombplus() view returns(address)
func (_Tombplus *TombplusCallerSession) Otombplus() (common.Address, error) {
	return _Tombplus.Contract.Otombplus(&_Tombplus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tombplus *TombplusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tombplus *TombplusSession) Owner() (common.Address, error) {
	return _Tombplus.Contract.Owner(&_Tombplus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tombplus *TombplusCallerSession) Owner() (common.Address, error) {
	return _Tombplus.Contract.Owner(&_Tombplus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Tombplus *TombplusCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Tombplus *TombplusSession) PendingOwner() (common.Address, error) {
	return _Tombplus.Contract.PendingOwner(&_Tombplus.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Tombplus *TombplusCallerSession) PendingOwner() (common.Address, error) {
	return _Tombplus.Contract.PendingOwner(&_Tombplus.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Tombplus *TombplusCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Tombplus *TombplusSession) PriceFeed() (common.Address, error) {
	return _Tombplus.Contract.PriceFeed(&_Tombplus.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Tombplus *TombplusCallerSession) PriceFeed() (common.Address, error) {
	return _Tombplus.Contract.PriceFeed(&_Tombplus.CallOpts)
}

// ProcessedRewardBalances is a free data retrieval call binding the contract method 0xf04badf6.
//
// Solidity: function processedRewardBalances(address ) view returns(uint256)
func (_Tombplus *TombplusCaller) ProcessedRewardBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "processedRewardBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProcessedRewardBalances is a free data retrieval call binding the contract method 0xf04badf6.
//
// Solidity: function processedRewardBalances(address ) view returns(uint256)
func (_Tombplus *TombplusSession) ProcessedRewardBalances(arg0 common.Address) (*big.Int, error) {
	return _Tombplus.Contract.ProcessedRewardBalances(&_Tombplus.CallOpts, arg0)
}

// ProcessedRewardBalances is a free data retrieval call binding the contract method 0xf04badf6.
//
// Solidity: function processedRewardBalances(address ) view returns(uint256)
func (_Tombplus *TombplusCallerSession) ProcessedRewardBalances(arg0 common.Address) (*big.Int, error) {
	return _Tombplus.Contract.ProcessedRewardBalances(&_Tombplus.CallOpts, arg0)
}

// RewardBalance is a free data retrieval call binding the contract method 0x67b40cf7.
//
// Solidity: function rewardBalance(address user) view returns(uint256 balance)
func (_Tombplus *TombplusCaller) RewardBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "rewardBalance", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardBalance is a free data retrieval call binding the contract method 0x67b40cf7.
//
// Solidity: function rewardBalance(address user) view returns(uint256 balance)
func (_Tombplus *TombplusSession) RewardBalance(user common.Address) (*big.Int, error) {
	return _Tombplus.Contract.RewardBalance(&_Tombplus.CallOpts, user)
}

// RewardBalance is a free data retrieval call binding the contract method 0x67b40cf7.
//
// Solidity: function rewardBalance(address user) view returns(uint256 balance)
func (_Tombplus *TombplusCallerSession) RewardBalance(user common.Address) (*big.Int, error) {
	return _Tombplus.Contract.RewardBalance(&_Tombplus.CallOpts, user)
}

// RewardPerTshareNumerator is a free data retrieval call binding the contract method 0xe4188980.
//
// Solidity: function rewardPerTshareNumerator() view returns(uint256)
func (_Tombplus *TombplusCaller) RewardPerTshareNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "rewardPerTshareNumerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTshareNumerator is a free data retrieval call binding the contract method 0xe4188980.
//
// Solidity: function rewardPerTshareNumerator() view returns(uint256)
func (_Tombplus *TombplusSession) RewardPerTshareNumerator() (*big.Int, error) {
	return _Tombplus.Contract.RewardPerTshareNumerator(&_Tombplus.CallOpts)
}

// RewardPerTshareNumerator is a free data retrieval call binding the contract method 0xe4188980.
//
// Solidity: function rewardPerTshareNumerator() view returns(uint256)
func (_Tombplus *TombplusCallerSession) RewardPerTshareNumerator() (*big.Int, error) {
	return _Tombplus.Contract.RewardPerTshareNumerator(&_Tombplus.CallOpts)
}

// Tombplus is a free data retrieval call binding the contract method 0x779c94df.
//
// Solidity: function tombplus() view returns(address)
func (_Tombplus *TombplusCaller) Tombplus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "tombplus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tombplus is a free data retrieval call binding the contract method 0x779c94df.
//
// Solidity: function tombplus() view returns(address)
func (_Tombplus *TombplusSession) Tombplus() (common.Address, error) {
	return _Tombplus.Contract.Tombplus(&_Tombplus.CallOpts)
}

// Tombplus is a free data retrieval call binding the contract method 0x779c94df.
//
// Solidity: function tombplus() view returns(address)
func (_Tombplus *TombplusCallerSession) Tombplus() (common.Address, error) {
	return _Tombplus.Contract.Tombplus(&_Tombplus.CallOpts)
}

// Tshareplus is a free data retrieval call binding the contract method 0x6200b680.
//
// Solidity: function tshareplus() view returns(address)
func (_Tombplus *TombplusCaller) Tshareplus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "tshareplus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tshareplus is a free data retrieval call binding the contract method 0x6200b680.
//
// Solidity: function tshareplus() view returns(address)
func (_Tombplus *TombplusSession) Tshareplus() (common.Address, error) {
	return _Tombplus.Contract.Tshareplus(&_Tombplus.CallOpts)
}

// Tshareplus is a free data retrieval call binding the contract method 0x6200b680.
//
// Solidity: function tshareplus() view returns(address)
func (_Tombplus *TombplusCallerSession) Tshareplus() (common.Address, error) {
	return _Tombplus.Contract.Tshareplus(&_Tombplus.CallOpts)
}

// UserFlipIdsByEpochId is a free data retrieval call binding the contract method 0x5fd94ef1.
//
// Solidity: function userFlipIdsByEpochId(address , uint256 ) view returns(uint256)
func (_Tombplus *TombplusCaller) UserFlipIdsByEpochId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "userFlipIdsByEpochId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFlipIdsByEpochId is a free data retrieval call binding the contract method 0x5fd94ef1.
//
// Solidity: function userFlipIdsByEpochId(address , uint256 ) view returns(uint256)
func (_Tombplus *TombplusSession) UserFlipIdsByEpochId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tombplus.Contract.UserFlipIdsByEpochId(&_Tombplus.CallOpts, arg0, arg1)
}

// UserFlipIdsByEpochId is a free data retrieval call binding the contract method 0x5fd94ef1.
//
// Solidity: function userFlipIdsByEpochId(address , uint256 ) view returns(uint256)
func (_Tombplus *TombplusCallerSession) UserFlipIdsByEpochId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Tombplus.Contract.UserFlipIdsByEpochId(&_Tombplus.CallOpts, arg0, arg1)
}

// UserFlips is a free data retrieval call binding the contract method 0x3c9e8907.
//
// Solidity: function userFlips(address , uint256 ) view returns(uint256 epochId, bool up, uint256 tsharePlusBalance)
func (_Tombplus *TombplusCaller) UserFlips(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	EpochId           *big.Int
	Up                bool
	TsharePlusBalance *big.Int
}, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "userFlips", arg0, arg1)

	outstruct := new(struct {
		EpochId           *big.Int
		Up                bool
		TsharePlusBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EpochId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Up = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TsharePlusBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserFlips is a free data retrieval call binding the contract method 0x3c9e8907.
//
// Solidity: function userFlips(address , uint256 ) view returns(uint256 epochId, bool up, uint256 tsharePlusBalance)
func (_Tombplus *TombplusSession) UserFlips(arg0 common.Address, arg1 *big.Int) (struct {
	EpochId           *big.Int
	Up                bool
	TsharePlusBalance *big.Int
}, error) {
	return _Tombplus.Contract.UserFlips(&_Tombplus.CallOpts, arg0, arg1)
}

// UserFlips is a free data retrieval call binding the contract method 0x3c9e8907.
//
// Solidity: function userFlips(address , uint256 ) view returns(uint256 epochId, bool up, uint256 tsharePlusBalance)
func (_Tombplus *TombplusCallerSession) UserFlips(arg0 common.Address, arg1 *big.Int) (struct {
	EpochId           *big.Int
	Up                bool
	TsharePlusBalance *big.Int
}, error) {
	return _Tombplus.Contract.UserFlips(&_Tombplus.CallOpts, arg0, arg1)
}

// VolatilityThreshold is a free data retrieval call binding the contract method 0x9fe88c84.
//
// Solidity: function volatilityThreshold() view returns(uint256)
func (_Tombplus *TombplusCaller) VolatilityThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "volatilityThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VolatilityThreshold is a free data retrieval call binding the contract method 0x9fe88c84.
//
// Solidity: function volatilityThreshold() view returns(uint256)
func (_Tombplus *TombplusSession) VolatilityThreshold() (*big.Int, error) {
	return _Tombplus.Contract.VolatilityThreshold(&_Tombplus.CallOpts)
}

// VolatilityThreshold is a free data retrieval call binding the contract method 0x9fe88c84.
//
// Solidity: function volatilityThreshold() view returns(uint256)
func (_Tombplus *TombplusCallerSession) VolatilityThreshold() (*big.Int, error) {
	return _Tombplus.Contract.VolatilityThreshold(&_Tombplus.CallOpts)
}

// VolatilityThresholdBase is a free data retrieval call binding the contract method 0x8b0b77b7.
//
// Solidity: function volatilityThresholdBase() view returns(uint256)
func (_Tombplus *TombplusCaller) VolatilityThresholdBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "volatilityThresholdBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VolatilityThresholdBase is a free data retrieval call binding the contract method 0x8b0b77b7.
//
// Solidity: function volatilityThresholdBase() view returns(uint256)
func (_Tombplus *TombplusSession) VolatilityThresholdBase() (*big.Int, error) {
	return _Tombplus.Contract.VolatilityThresholdBase(&_Tombplus.CallOpts)
}

// VolatilityThresholdBase is a free data retrieval call binding the contract method 0x8b0b77b7.
//
// Solidity: function volatilityThresholdBase() view returns(uint256)
func (_Tombplus *TombplusCallerSession) VolatilityThresholdBase() (*big.Int, error) {
	return _Tombplus.Contract.VolatilityThresholdBase(&_Tombplus.CallOpts)
}

// VolatilityThresholdStep is a free data retrieval call binding the contract method 0xda172a0c.
//
// Solidity: function volatilityThresholdStep() view returns(uint256)
func (_Tombplus *TombplusCaller) VolatilityThresholdStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "volatilityThresholdStep")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VolatilityThresholdStep is a free data retrieval call binding the contract method 0xda172a0c.
//
// Solidity: function volatilityThresholdStep() view returns(uint256)
func (_Tombplus *TombplusSession) VolatilityThresholdStep() (*big.Int, error) {
	return _Tombplus.Contract.VolatilityThresholdStep(&_Tombplus.CallOpts)
}

// VolatilityThresholdStep is a free data retrieval call binding the contract method 0xda172a0c.
//
// Solidity: function volatilityThresholdStep() view returns(uint256)
func (_Tombplus *TombplusCallerSession) VolatilityThresholdStep() (*big.Int, error) {
	return _Tombplus.Contract.VolatilityThresholdStep(&_Tombplus.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Tombplus *TombplusTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Tombplus *TombplusSession) Claim() (*types.Transaction, error) {
	return _Tombplus.Contract.Claim(&_Tombplus.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Tombplus *TombplusTransactorSession) Claim() (*types.Transaction, error) {
	return _Tombplus.Contract.Claim(&_Tombplus.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tombplus *TombplusTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tombplus *TombplusSession) ClaimOwnership() (*types.Transaction, error) {
	return _Tombplus.Contract.ClaimOwnership(&_Tombplus.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tombplus *TombplusTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Tombplus.Contract.ClaimOwnership(&_Tombplus.TransactOpts)
}

// Flip is a paid mutator transaction binding the contract method 0x105c2738.
//
// Solidity: function flip(uint256 epochId, bool up) returns()
func (_Tombplus *TombplusTransactor) Flip(opts *bind.TransactOpts, epochId *big.Int, up bool) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "flip", epochId, up)
}

// Flip is a paid mutator transaction binding the contract method 0x105c2738.
//
// Solidity: function flip(uint256 epochId, bool up) returns()
func (_Tombplus *TombplusSession) Flip(epochId *big.Int, up bool) (*types.Transaction, error) {
	return _Tombplus.Contract.Flip(&_Tombplus.TransactOpts, epochId, up)
}

// Flip is a paid mutator transaction binding the contract method 0x105c2738.
//
// Solidity: function flip(uint256 epochId, bool up) returns()
func (_Tombplus *TombplusTransactorSession) Flip(epochId *big.Int, up bool) (*types.Transaction, error) {
	return _Tombplus.Contract.Flip(&_Tombplus.TransactOpts, epochId, up)
}

// FlipMultiple is a paid mutator transaction binding the contract method 0x85b716d5.
//
// Solidity: function flipMultiple(uint256[] epochIds, bool[] ups) returns()
func (_Tombplus *TombplusTransactor) FlipMultiple(opts *bind.TransactOpts, epochIds []*big.Int, ups []bool) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "flipMultiple", epochIds, ups)
}

// FlipMultiple is a paid mutator transaction binding the contract method 0x85b716d5.
//
// Solidity: function flipMultiple(uint256[] epochIds, bool[] ups) returns()
func (_Tombplus *TombplusSession) FlipMultiple(epochIds []*big.Int, ups []bool) (*types.Transaction, error) {
	return _Tombplus.Contract.FlipMultiple(&_Tombplus.TransactOpts, epochIds, ups)
}

// FlipMultiple is a paid mutator transaction binding the contract method 0x85b716d5.
//
// Solidity: function flipMultiple(uint256[] epochIds, bool[] ups) returns()
func (_Tombplus *TombplusTransactorSession) FlipMultiple(epochIds []*big.Int, ups []bool) (*types.Transaction, error) {
	return _Tombplus.Contract.FlipMultiple(&_Tombplus.TransactOpts, epochIds, ups)
}

// FlipState is a paid mutator transaction binding the contract method 0x8e920351.
//
// Solidity: function flipState() returns()
func (_Tombplus *TombplusTransactor) FlipState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "flipState")
}

// FlipState is a paid mutator transaction binding the contract method 0x8e920351.
//
// Solidity: function flipState() returns()
func (_Tombplus *TombplusSession) FlipState() (*types.Transaction, error) {
	return _Tombplus.Contract.FlipState(&_Tombplus.TransactOpts)
}

// FlipState is a paid mutator transaction binding the contract method 0x8e920351.
//
// Solidity: function flipState() returns()
func (_Tombplus *TombplusTransactorSession) FlipState() (*types.Transaction, error) {
	return _Tombplus.Contract.FlipState(&_Tombplus.TransactOpts)
}

// ProcessUserFlips is a paid mutator transaction binding the contract method 0xbd39ee05.
//
// Solidity: function processUserFlips(address user) returns()
func (_Tombplus *TombplusTransactor) ProcessUserFlips(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "processUserFlips", user)
}

// ProcessUserFlips is a paid mutator transaction binding the contract method 0xbd39ee05.
//
// Solidity: function processUserFlips(address user) returns()
func (_Tombplus *TombplusSession) ProcessUserFlips(user common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.ProcessUserFlips(&_Tombplus.TransactOpts, user)
}

// ProcessUserFlips is a paid mutator transaction binding the contract method 0xbd39ee05.
//
// Solidity: function processUserFlips(address user) returns()
func (_Tombplus *TombplusTransactorSession) ProcessUserFlips(user common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.ProcessUserFlips(&_Tombplus.TransactOpts, user)
}

// SetCfg is a paid mutator transaction binding the contract method 0xdceda7b2.
//
// Solidity: function setCfg(uint256 _maxAllowedFutureFlips, uint256 _rewardPerTshareNumerator) returns()
func (_Tombplus *TombplusTransactor) SetCfg(opts *bind.TransactOpts, _maxAllowedFutureFlips *big.Int, _rewardPerTshareNumerator *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setCfg", _maxAllowedFutureFlips, _rewardPerTshareNumerator)
}

// SetCfg is a paid mutator transaction binding the contract method 0xdceda7b2.
//
// Solidity: function setCfg(uint256 _maxAllowedFutureFlips, uint256 _rewardPerTshareNumerator) returns()
func (_Tombplus *TombplusSession) SetCfg(_maxAllowedFutureFlips *big.Int, _rewardPerTshareNumerator *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetCfg(&_Tombplus.TransactOpts, _maxAllowedFutureFlips, _rewardPerTshareNumerator)
}

// SetCfg is a paid mutator transaction binding the contract method 0xdceda7b2.
//
// Solidity: function setCfg(uint256 _maxAllowedFutureFlips, uint256 _rewardPerTshareNumerator) returns()
func (_Tombplus *TombplusTransactorSession) SetCfg(_maxAllowedFutureFlips *big.Int, _rewardPerTshareNumerator *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetCfg(&_Tombplus.TransactOpts, _maxAllowedFutureFlips, _rewardPerTshareNumerator)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_Tombplus *TombplusTransactor) SetEpochDuration(opts *bind.TransactOpts, _epochDuration *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setEpochDuration", _epochDuration)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_Tombplus *TombplusSession) SetEpochDuration(_epochDuration *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetEpochDuration(&_Tombplus.TransactOpts, _epochDuration)
}

// SetEpochDuration is a paid mutator transaction binding the contract method 0x30024dfe.
//
// Solidity: function setEpochDuration(uint256 _epochDuration) returns()
func (_Tombplus *TombplusTransactorSession) SetEpochDuration(_epochDuration *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetEpochDuration(&_Tombplus.TransactOpts, _epochDuration)
}

// SetFlipWindow is a paid mutator transaction binding the contract method 0x81a640be.
//
// Solidity: function setFlipWindow(uint256 _flipWindow) returns()
func (_Tombplus *TombplusTransactor) SetFlipWindow(opts *bind.TransactOpts, _flipWindow *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setFlipWindow", _flipWindow)
}

// SetFlipWindow is a paid mutator transaction binding the contract method 0x81a640be.
//
// Solidity: function setFlipWindow(uint256 _flipWindow) returns()
func (_Tombplus *TombplusSession) SetFlipWindow(_flipWindow *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetFlipWindow(&_Tombplus.TransactOpts, _flipWindow)
}

// SetFlipWindow is a paid mutator transaction binding the contract method 0x81a640be.
//
// Solidity: function setFlipWindow(uint256 _flipWindow) returns()
func (_Tombplus *TombplusTransactorSession) SetFlipWindow(_flipWindow *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetFlipWindow(&_Tombplus.TransactOpts, _flipWindow)
}

// SetOTombPlus is a paid mutator transaction binding the contract method 0xa22779b6.
//
// Solidity: function setOTombPlus(address _otombplus) returns()
func (_Tombplus *TombplusTransactor) SetOTombPlus(opts *bind.TransactOpts, _otombplus common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setOTombPlus", _otombplus)
}

// SetOTombPlus is a paid mutator transaction binding the contract method 0xa22779b6.
//
// Solidity: function setOTombPlus(address _otombplus) returns()
func (_Tombplus *TombplusSession) SetOTombPlus(_otombplus common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetOTombPlus(&_Tombplus.TransactOpts, _otombplus)
}

// SetOTombPlus is a paid mutator transaction binding the contract method 0xa22779b6.
//
// Solidity: function setOTombPlus(address _otombplus) returns()
func (_Tombplus *TombplusTransactorSession) SetOTombPlus(_otombplus common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetOTombPlus(&_Tombplus.TransactOpts, _otombplus)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Tombplus *TombplusTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Tombplus *TombplusSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetPriceFeed(&_Tombplus.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Tombplus *TombplusTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetPriceFeed(&_Tombplus.TransactOpts, _priceFeed)
}

// SetVolatilityThreshold is a paid mutator transaction binding the contract method 0x0b0eebb0.
//
// Solidity: function setVolatilityThreshold(uint256 _volatilityThresholdBase, uint256 _volatilityThreshold, uint256 _volatilityThresholdStep) returns()
func (_Tombplus *TombplusTransactor) SetVolatilityThreshold(opts *bind.TransactOpts, _volatilityThresholdBase *big.Int, _volatilityThreshold *big.Int, _volatilityThresholdStep *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setVolatilityThreshold", _volatilityThresholdBase, _volatilityThreshold, _volatilityThresholdStep)
}

// SetVolatilityThreshold is a paid mutator transaction binding the contract method 0x0b0eebb0.
//
// Solidity: function setVolatilityThreshold(uint256 _volatilityThresholdBase, uint256 _volatilityThreshold, uint256 _volatilityThresholdStep) returns()
func (_Tombplus *TombplusSession) SetVolatilityThreshold(_volatilityThresholdBase *big.Int, _volatilityThreshold *big.Int, _volatilityThresholdStep *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetVolatilityThreshold(&_Tombplus.TransactOpts, _volatilityThresholdBase, _volatilityThreshold, _volatilityThresholdStep)
}

// SetVolatilityThreshold is a paid mutator transaction binding the contract method 0x0b0eebb0.
//
// Solidity: function setVolatilityThreshold(uint256 _volatilityThresholdBase, uint256 _volatilityThreshold, uint256 _volatilityThresholdStep) returns()
func (_Tombplus *TombplusTransactorSession) SetVolatilityThreshold(_volatilityThresholdBase *big.Int, _volatilityThreshold *big.Int, _volatilityThresholdStep *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetVolatilityThreshold(&_Tombplus.TransactOpts, _volatilityThresholdBase, _volatilityThreshold, _volatilityThresholdStep)
}

// StartGame is a paid mutator transaction binding the contract method 0xd65ab5f2.
//
// Solidity: function startGame() returns()
func (_Tombplus *TombplusTransactor) StartGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "startGame")
}

// StartGame is a paid mutator transaction binding the contract method 0xd65ab5f2.
//
// Solidity: function startGame() returns()
func (_Tombplus *TombplusSession) StartGame() (*types.Transaction, error) {
	return _Tombplus.Contract.StartGame(&_Tombplus.TransactOpts)
}

// StartGame is a paid mutator transaction binding the contract method 0xd65ab5f2.
//
// Solidity: function startGame() returns()
func (_Tombplus *TombplusTransactorSession) StartGame() (*types.Transaction, error) {
	return _Tombplus.Contract.StartGame(&_Tombplus.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_Tombplus *TombplusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "transferOwnership", newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_Tombplus *TombplusSession) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _Tombplus.Contract.TransferOwnership(&_Tombplus.TransactOpts, newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_Tombplus *TombplusTransactorSession) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _Tombplus.Contract.TransferOwnership(&_Tombplus.TransactOpts, newOwner, direct, renounce)
}

// UpdateUserTsharePlusBalance is a paid mutator transaction binding the contract method 0x2e8da1e5.
//
// Solidity: function updateUserTsharePlusBalance(address user) returns()
func (_Tombplus *TombplusTransactor) UpdateUserTsharePlusBalance(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "updateUserTsharePlusBalance", user)
}

// UpdateUserTsharePlusBalance is a paid mutator transaction binding the contract method 0x2e8da1e5.
//
// Solidity: function updateUserTsharePlusBalance(address user) returns()
func (_Tombplus *TombplusSession) UpdateUserTsharePlusBalance(user common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.UpdateUserTsharePlusBalance(&_Tombplus.TransactOpts, user)
}

// UpdateUserTsharePlusBalance is a paid mutator transaction binding the contract method 0x2e8da1e5.
//
// Solidity: function updateUserTsharePlusBalance(address user) returns()
func (_Tombplus *TombplusTransactorSession) UpdateUserTsharePlusBalance(user common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.UpdateUserTsharePlusBalance(&_Tombplus.TransactOpts, user)
}

// TombplusEpochFinishedIterator is returned from FilterEpochFinished and is used to iterate over the raw logs and unpacked data for EpochFinished events raised by the Tombplus contract.
type TombplusEpochFinishedIterator struct {
	Event *TombplusEpochFinished // Event containing the contract specifics and raw log

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
func (it *TombplusEpochFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TombplusEpochFinished)
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
		it.Event = new(TombplusEpochFinished)
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
func (it *TombplusEpochFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TombplusEpochFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TombplusEpochFinished represents a EpochFinished event raised by the Tombplus contract.
type TombplusEpochFinished struct {
	EpochNumber       *big.Int
	EpochEndPrice     *big.Int
	VolatilityReached bool
	WinnerIsUp        bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEpochFinished is a free log retrieval operation binding the contract event 0x3b5e82289135e7d610beee7f392f9bce4057fdab01229f22e31cb2b23e5c1df1.
//
// Solidity: event EpochFinished(uint256 epochNumber, uint256 epochEndPrice, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusFilterer) FilterEpochFinished(opts *bind.FilterOpts) (*TombplusEpochFinishedIterator, error) {

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "EpochFinished")
	if err != nil {
		return nil, err
	}
	return &TombplusEpochFinishedIterator{contract: _Tombplus.contract, event: "EpochFinished", logs: logs, sub: sub}, nil
}

// WatchEpochFinished is a free log subscription operation binding the contract event 0x3b5e82289135e7d610beee7f392f9bce4057fdab01229f22e31cb2b23e5c1df1.
//
// Solidity: event EpochFinished(uint256 epochNumber, uint256 epochEndPrice, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusFilterer) WatchEpochFinished(opts *bind.WatchOpts, sink chan<- *TombplusEpochFinished) (event.Subscription, error) {

	logs, sub, err := _Tombplus.contract.WatchLogs(opts, "EpochFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TombplusEpochFinished)
				if err := _Tombplus.contract.UnpackLog(event, "EpochFinished", log); err != nil {
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

// ParseEpochFinished is a log parse operation binding the contract event 0x3b5e82289135e7d610beee7f392f9bce4057fdab01229f22e31cb2b23e5c1df1.
//
// Solidity: event EpochFinished(uint256 epochNumber, uint256 epochEndPrice, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusFilterer) ParseEpochFinished(log types.Log) (*TombplusEpochFinished, error) {
	event := new(TombplusEpochFinished)
	if err := _Tombplus.contract.UnpackLog(event, "EpochFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TombplusEpochStartedIterator is returned from FilterEpochStarted and is used to iterate over the raw logs and unpacked data for EpochStarted events raised by the Tombplus contract.
type TombplusEpochStartedIterator struct {
	Event *TombplusEpochStarted // Event containing the contract specifics and raw log

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
func (it *TombplusEpochStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TombplusEpochStarted)
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
		it.Event = new(TombplusEpochStarted)
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
func (it *TombplusEpochStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TombplusEpochStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TombplusEpochStarted represents a EpochStarted event raised by the Tombplus contract.
type TombplusEpochStarted struct {
	EpochNumber         *big.Int
	VolatilityThreshold *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterEpochStarted is a free log retrieval operation binding the contract event 0x41787f1277821474072e18df95f0bd9ed9f117003aa97732ebbd737225b32b02.
//
// Solidity: event EpochStarted(uint256 epochNumber, uint256 volatilityThreshold)
func (_Tombplus *TombplusFilterer) FilterEpochStarted(opts *bind.FilterOpts) (*TombplusEpochStartedIterator, error) {

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "EpochStarted")
	if err != nil {
		return nil, err
	}
	return &TombplusEpochStartedIterator{contract: _Tombplus.contract, event: "EpochStarted", logs: logs, sub: sub}, nil
}

// WatchEpochStarted is a free log subscription operation binding the contract event 0x41787f1277821474072e18df95f0bd9ed9f117003aa97732ebbd737225b32b02.
//
// Solidity: event EpochStarted(uint256 epochNumber, uint256 volatilityThreshold)
func (_Tombplus *TombplusFilterer) WatchEpochStarted(opts *bind.WatchOpts, sink chan<- *TombplusEpochStarted) (event.Subscription, error) {

	logs, sub, err := _Tombplus.contract.WatchLogs(opts, "EpochStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TombplusEpochStarted)
				if err := _Tombplus.contract.UnpackLog(event, "EpochStarted", log); err != nil {
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

// ParseEpochStarted is a log parse operation binding the contract event 0x41787f1277821474072e18df95f0bd9ed9f117003aa97732ebbd737225b32b02.
//
// Solidity: event EpochStarted(uint256 epochNumber, uint256 volatilityThreshold)
func (_Tombplus *TombplusFilterer) ParseEpochStarted(log types.Log) (*TombplusEpochStarted, error) {
	event := new(TombplusEpochStarted)
	if err := _Tombplus.contract.UnpackLog(event, "EpochStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TombplusObservationStartedIterator is returned from FilterObservationStarted and is used to iterate over the raw logs and unpacked data for ObservationStarted events raised by the Tombplus contract.
type TombplusObservationStartedIterator struct {
	Event *TombplusObservationStarted // Event containing the contract specifics and raw log

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
func (it *TombplusObservationStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TombplusObservationStarted)
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
		it.Event = new(TombplusObservationStarted)
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
func (it *TombplusObservationStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TombplusObservationStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TombplusObservationStarted represents a ObservationStarted event raised by the Tombplus contract.
type TombplusObservationStarted struct {
	EpochNumber     *big.Int
	EpochStartPrice *big.Int
	UsersUp         *big.Int
	VotesUp         *big.Int
	UsersDown       *big.Int
	VotesDown       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterObservationStarted is a free log retrieval operation binding the contract event 0x3101bbcbd5369a33e197b1769e4c1e6825472b375cac2f9831fc2fd426574641.
//
// Solidity: event ObservationStarted(uint256 epochNumber, uint256 epochStartPrice, uint256 usersUp, uint256 votesUp, uint256 usersDown, uint256 votesDown)
func (_Tombplus *TombplusFilterer) FilterObservationStarted(opts *bind.FilterOpts) (*TombplusObservationStartedIterator, error) {

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "ObservationStarted")
	if err != nil {
		return nil, err
	}
	return &TombplusObservationStartedIterator{contract: _Tombplus.contract, event: "ObservationStarted", logs: logs, sub: sub}, nil
}

// WatchObservationStarted is a free log subscription operation binding the contract event 0x3101bbcbd5369a33e197b1769e4c1e6825472b375cac2f9831fc2fd426574641.
//
// Solidity: event ObservationStarted(uint256 epochNumber, uint256 epochStartPrice, uint256 usersUp, uint256 votesUp, uint256 usersDown, uint256 votesDown)
func (_Tombplus *TombplusFilterer) WatchObservationStarted(opts *bind.WatchOpts, sink chan<- *TombplusObservationStarted) (event.Subscription, error) {

	logs, sub, err := _Tombplus.contract.WatchLogs(opts, "ObservationStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TombplusObservationStarted)
				if err := _Tombplus.contract.UnpackLog(event, "ObservationStarted", log); err != nil {
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

// ParseObservationStarted is a log parse operation binding the contract event 0x3101bbcbd5369a33e197b1769e4c1e6825472b375cac2f9831fc2fd426574641.
//
// Solidity: event ObservationStarted(uint256 epochNumber, uint256 epochStartPrice, uint256 usersUp, uint256 votesUp, uint256 usersDown, uint256 votesDown)
func (_Tombplus *TombplusFilterer) ParseObservationStarted(log types.Log) (*TombplusObservationStarted, error) {
	event := new(TombplusObservationStarted)
	if err := _Tombplus.contract.UnpackLog(event, "ObservationStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TombplusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tombplus contract.
type TombplusOwnershipTransferredIterator struct {
	Event *TombplusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TombplusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TombplusOwnershipTransferred)
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
		it.Event = new(TombplusOwnershipTransferred)
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
func (it *TombplusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TombplusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TombplusOwnershipTransferred represents a OwnershipTransferred event raised by the Tombplus contract.
type TombplusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tombplus *TombplusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TombplusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TombplusOwnershipTransferredIterator{contract: _Tombplus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tombplus *TombplusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TombplusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tombplus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TombplusOwnershipTransferred)
				if err := _Tombplus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tombplus *TombplusFilterer) ParseOwnershipTransferred(log types.Log) (*TombplusOwnershipTransferred, error) {
	event := new(TombplusOwnershipTransferred)
	if err := _Tombplus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TombplusUserFlippedIterator is returned from FilterUserFlipped and is used to iterate over the raw logs and unpacked data for UserFlipped events raised by the Tombplus contract.
type TombplusUserFlippedIterator struct {
	Event *TombplusUserFlipped // Event containing the contract specifics and raw log

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
func (it *TombplusUserFlippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TombplusUserFlipped)
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
		it.Event = new(TombplusUserFlipped)
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
func (it *TombplusUserFlippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TombplusUserFlippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TombplusUserFlipped represents a UserFlipped event raised by the Tombplus contract.
type TombplusUserFlipped struct {
	User    common.Address
	FlipId  *big.Int
	EpochId *big.Int
	Up      bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserFlipped is a free log retrieval operation binding the contract event 0xf251991d6027899592edc87c4426dc7649cfa3d68d09c9e537465680bf30dc9e.
//
// Solidity: event UserFlipped(address user, uint256 flipId, uint256 epochId, bool up)
func (_Tombplus *TombplusFilterer) FilterUserFlipped(opts *bind.FilterOpts) (*TombplusUserFlippedIterator, error) {

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "UserFlipped")
	if err != nil {
		return nil, err
	}
	return &TombplusUserFlippedIterator{contract: _Tombplus.contract, event: "UserFlipped", logs: logs, sub: sub}, nil
}

// WatchUserFlipped is a free log subscription operation binding the contract event 0xf251991d6027899592edc87c4426dc7649cfa3d68d09c9e537465680bf30dc9e.
//
// Solidity: event UserFlipped(address user, uint256 flipId, uint256 epochId, bool up)
func (_Tombplus *TombplusFilterer) WatchUserFlipped(opts *bind.WatchOpts, sink chan<- *TombplusUserFlipped) (event.Subscription, error) {

	logs, sub, err := _Tombplus.contract.WatchLogs(opts, "UserFlipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TombplusUserFlipped)
				if err := _Tombplus.contract.UnpackLog(event, "UserFlipped", log); err != nil {
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

// ParseUserFlipped is a log parse operation binding the contract event 0xf251991d6027899592edc87c4426dc7649cfa3d68d09c9e537465680bf30dc9e.
//
// Solidity: event UserFlipped(address user, uint256 flipId, uint256 epochId, bool up)
func (_Tombplus *TombplusFilterer) ParseUserFlipped(log types.Log) (*TombplusUserFlipped, error) {
	event := new(TombplusUserFlipped)
	if err := _Tombplus.contract.UnpackLog(event, "UserFlipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
