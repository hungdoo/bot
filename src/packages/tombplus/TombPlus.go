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

// MasonryPlusFlip is an auto generated low-level Go binding around an user-defined struct.
type MasonryPlusFlip struct {
	EpochId           *big.Int
	Up                bool
	TsharePlusBalance *big.Int
}

// MasonryPlusHistoricEpochData is an auto generated low-level Go binding around an user-defined struct.
type MasonryPlusHistoricEpochData struct {
	VolatilityReached             bool
	WinnerIsUp                    bool
	RewardPerWinningVoteNumerator *big.Int
}

// MasonryPlusIndexedHistoricEpochData is an auto generated low-level Go binding around an user-defined struct.
type MasonryPlusIndexedHistoricEpochData struct {
	EpochId           *big.Int
	HistoricEpochData MasonryPlusHistoricEpochData
}

// TombplusMetaData contains all meta data concerning the Tombplus contract.
var TombplusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITombPlus\",\"name\":\"_tombplus\",\"type\":\"address\"},{\"internalType\":\"contractIOTombPlus\",\"name\":\"_otombplus\",\"type\":\"address\"},{\"internalType\":\"contractITsharePlus\",\"name\":\"_tshareplus\",\"type\":\"address\"},{\"internalType\":\"contractIPyth\",\"name\":\"_pyth\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_pythPriceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gameStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_distribRate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStartPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochEndPrice\",\"type\":\"uint256\"}],\"name\":\"EmergencyFinishEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStartPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochEndPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"volatilityReached\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"winnerIsUp\",\"type\":\"bool\"}],\"name\":\"EpochFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"flipId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"UserFlipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UserWon\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numsOfFlips\",\"type\":\"uint256[]\"}],\"name\":\"batchSetMaxAllowedFutureFlipsOverride\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"batchUnsetMaxAllowedFutureFlipsOverride\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canFlipForCurrentEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPrice\",\"type\":\"uint256\"}],\"name\":\"emergencyFinishEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"startPriceData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"endPriceData\",\"type\":\"bytes[]\"}],\"name\":\"finishEpoch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"startPriceData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"endPriceData\",\"type\":\"bytes[]\"}],\"name\":\"finishEpochPermissionless\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"flip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[]\",\"name\":\"ups\",\"type\":\"bool[]\"}],\"name\":\"flipMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStartTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"getEpochIdFromTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getEpochObservationTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"observationStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"observationEndTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfEpochs\",\"type\":\"uint256\"}],\"name\":\"getLastFinishedEpochs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"volatilityReached\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"winnerIsUp\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerWinningVoteNumerator\",\"type\":\"uint256\"}],\"internalType\":\"structMasonryPlus.HistoricEpochData\",\"name\":\"historicEpochData\",\"type\":\"tuple\"}],\"internalType\":\"structMasonryPlus.IndexedHistoricEpochData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"getTimeInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"getUserFlipIdByEpochId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserFlips\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tsharePlusBalance\",\"type\":\"uint256\"}],\"internalType\":\"structMasonryPlus.Flip[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"historicEpochData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"volatilityReached\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"winnerIsUp\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerWinningVoteNumerator\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinishedEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lottery\",\"outputs\":[{\"internalType\":\"contractILottery\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllowedFutureFlips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxAllowedFutureFlipsOverride\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otombplus\",\"outputs\":[{\"internalType\":\"contractIOTombPlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"pauseGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGameAtEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGameImmediately\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"processUserFlips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processedRewardBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pyth\",\"outputs\":[{\"internalType\":\"contractIPyth\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pythPriceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeFlips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rewardBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTshareNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_distribRate\",\"type\":\"uint256\"}],\"name\":\"setDistribCfg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flipWindow\",\"type\":\"uint256\"}],\"name\":\"setFlipWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameStartTimestamp\",\"type\":\"uint256\"}],\"name\":\"setGameStartTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILottery\",\"name\":\"_lottery\",\"type\":\"address\"}],\"name\":\"setLottery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAllowedFutureFlips\",\"type\":\"uint256\"}],\"name\":\"setMaxAllowedFutureFlips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numOfFlips\",\"type\":\"uint256\"}],\"name\":\"setMaxAllowedFutureFlipsOverride\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOTombPlus\",\"name\":\"_otombplus\",\"type\":\"address\"}],\"name\":\"setOTombPlus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPyth\",\"name\":\"_pyth\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_pythPriceID\",\"type\":\"bytes32\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_pythAllowedTimeOvershoot\",\"type\":\"uint64\"}],\"name\":\"setPythAllowedTimeOvershoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardPerTshareNumerator\",\"type\":\"uint256\"}],\"name\":\"setRewardPerTshareNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUpgradedMasonry\",\"name\":\"_upgradedMasonry\",\"type\":\"address\"}],\"name\":\"setUpgradedMasonry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_volatilityThresholdBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_volatilityThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_volatilityThresholdStep\",\"type\":\"uint256\"}],\"name\":\"setVolatilityThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tombplus\",\"outputs\":[{\"internalType\":\"contractITombPlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"direct\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"renounce\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tshareplus\",\"outputs\":[{\"internalType\":\"contractITsharePlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"unsetMaxAllowedFutureFlipsOverride\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"upcomingEpochData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tshareVotesUp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tshareVotesDown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tshareUsersUp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tshareUsersDown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tshareTotalSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"updateUserTsharePlusBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradedMasonry\",\"outputs\":[{\"internalType\":\"contractIUpgradedMasonry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userCanPlaceOneMoreFlip\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userFlipIdsByEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userFlips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tsharePlusBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatilityThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatilityThresholdBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatilityThresholdStep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// CanFlipForCurrentEpoch is a free data retrieval call binding the contract method 0x7d766406.
//
// Solidity: function canFlipForCurrentEpoch() view returns(bool)
func (_Tombplus *TombplusCaller) CanFlipForCurrentEpoch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "canFlipForCurrentEpoch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanFlipForCurrentEpoch is a free data retrieval call binding the contract method 0x7d766406.
//
// Solidity: function canFlipForCurrentEpoch() view returns(bool)
func (_Tombplus *TombplusSession) CanFlipForCurrentEpoch() (bool, error) {
	return _Tombplus.Contract.CanFlipForCurrentEpoch(&_Tombplus.CallOpts)
}

// CanFlipForCurrentEpoch is a free data retrieval call binding the contract method 0x7d766406.
//
// Solidity: function canFlipForCurrentEpoch() view returns(bool)
func (_Tombplus *TombplusCallerSession) CanFlipForCurrentEpoch() (bool, error) {
	return _Tombplus.Contract.CanFlipForCurrentEpoch(&_Tombplus.CallOpts)
}

// CurrentEpochId is a free data retrieval call binding the contract method 0xeacdc5ff.
//
// Solidity: function currentEpochId() view returns(uint256)
func (_Tombplus *TombplusCaller) CurrentEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "currentEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpochId is a free data retrieval call binding the contract method 0xeacdc5ff.
//
// Solidity: function currentEpochId() view returns(uint256)
func (_Tombplus *TombplusSession) CurrentEpochId() (*big.Int, error) {
	return _Tombplus.Contract.CurrentEpochId(&_Tombplus.CallOpts)
}

// CurrentEpochId is a free data retrieval call binding the contract method 0xeacdc5ff.
//
// Solidity: function currentEpochId() view returns(uint256)
func (_Tombplus *TombplusCallerSession) CurrentEpochId() (*big.Int, error) {
	return _Tombplus.Contract.CurrentEpochId(&_Tombplus.CallOpts)
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

// DistribRate is a free data retrieval call binding the contract method 0xb4bce820.
//
// Solidity: function distribRate() view returns(uint256)
func (_Tombplus *TombplusCaller) DistribRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "distribRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistribRate is a free data retrieval call binding the contract method 0xb4bce820.
//
// Solidity: function distribRate() view returns(uint256)
func (_Tombplus *TombplusSession) DistribRate() (*big.Int, error) {
	return _Tombplus.Contract.DistribRate(&_Tombplus.CallOpts)
}

// DistribRate is a free data retrieval call binding the contract method 0xb4bce820.
//
// Solidity: function distribRate() view returns(uint256)
func (_Tombplus *TombplusCallerSession) DistribRate() (*big.Int, error) {
	return _Tombplus.Contract.DistribRate(&_Tombplus.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Tombplus *TombplusCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Tombplus *TombplusSession) Distributor() (common.Address, error) {
	return _Tombplus.Contract.Distributor(&_Tombplus.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Tombplus *TombplusCallerSession) Distributor() (common.Address, error) {
	return _Tombplus.Contract.Distributor(&_Tombplus.CallOpts)
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

// GameStartTimestamp is a free data retrieval call binding the contract method 0x5bfe6e00.
//
// Solidity: function gameStartTimestamp() view returns(uint256)
func (_Tombplus *TombplusCaller) GameStartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "gameStartTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GameStartTimestamp is a free data retrieval call binding the contract method 0x5bfe6e00.
//
// Solidity: function gameStartTimestamp() view returns(uint256)
func (_Tombplus *TombplusSession) GameStartTimestamp() (*big.Int, error) {
	return _Tombplus.Contract.GameStartTimestamp(&_Tombplus.CallOpts)
}

// GameStartTimestamp is a free data retrieval call binding the contract method 0x5bfe6e00.
//
// Solidity: function gameStartTimestamp() view returns(uint256)
func (_Tombplus *TombplusCallerSession) GameStartTimestamp() (*big.Int, error) {
	return _Tombplus.Contract.GameStartTimestamp(&_Tombplus.CallOpts)
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

// GetEpochIdFromTimestamp is a free data retrieval call binding the contract method 0x741c90b2.
//
// Solidity: function getEpochIdFromTimestamp(uint256 t) view returns(uint256)
func (_Tombplus *TombplusCaller) GetEpochIdFromTimestamp(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "getEpochIdFromTimestamp", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochIdFromTimestamp is a free data retrieval call binding the contract method 0x741c90b2.
//
// Solidity: function getEpochIdFromTimestamp(uint256 t) view returns(uint256)
func (_Tombplus *TombplusSession) GetEpochIdFromTimestamp(t *big.Int) (*big.Int, error) {
	return _Tombplus.Contract.GetEpochIdFromTimestamp(&_Tombplus.CallOpts, t)
}

// GetEpochIdFromTimestamp is a free data retrieval call binding the contract method 0x741c90b2.
//
// Solidity: function getEpochIdFromTimestamp(uint256 t) view returns(uint256)
func (_Tombplus *TombplusCallerSession) GetEpochIdFromTimestamp(t *big.Int) (*big.Int, error) {
	return _Tombplus.Contract.GetEpochIdFromTimestamp(&_Tombplus.CallOpts, t)
}

// GetEpochObservationTimestamps is a free data retrieval call binding the contract method 0x8eda0878.
//
// Solidity: function getEpochObservationTimestamps(uint256 epochId) view returns(uint256 observationStartTimestamp, uint256 observationEndTimestamp)
func (_Tombplus *TombplusCaller) GetEpochObservationTimestamps(opts *bind.CallOpts, epochId *big.Int) (struct {
	ObservationStartTimestamp *big.Int
	ObservationEndTimestamp   *big.Int
}, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "getEpochObservationTimestamps", epochId)

	outstruct := new(struct {
		ObservationStartTimestamp *big.Int
		ObservationEndTimestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ObservationStartTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ObservationEndTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEpochObservationTimestamps is a free data retrieval call binding the contract method 0x8eda0878.
//
// Solidity: function getEpochObservationTimestamps(uint256 epochId) view returns(uint256 observationStartTimestamp, uint256 observationEndTimestamp)
func (_Tombplus *TombplusSession) GetEpochObservationTimestamps(epochId *big.Int) (struct {
	ObservationStartTimestamp *big.Int
	ObservationEndTimestamp   *big.Int
}, error) {
	return _Tombplus.Contract.GetEpochObservationTimestamps(&_Tombplus.CallOpts, epochId)
}

// GetEpochObservationTimestamps is a free data retrieval call binding the contract method 0x8eda0878.
//
// Solidity: function getEpochObservationTimestamps(uint256 epochId) view returns(uint256 observationStartTimestamp, uint256 observationEndTimestamp)
func (_Tombplus *TombplusCallerSession) GetEpochObservationTimestamps(epochId *big.Int) (struct {
	ObservationStartTimestamp *big.Int
	ObservationEndTimestamp   *big.Int
}, error) {
	return _Tombplus.Contract.GetEpochObservationTimestamps(&_Tombplus.CallOpts, epochId)
}

// GetLastFinishedEpochs is a free data retrieval call binding the contract method 0xa5586593.
//
// Solidity: function getLastFinishedEpochs(uint256 numOfEpochs) view returns((uint256,(bool,bool,uint256))[])
func (_Tombplus *TombplusCaller) GetLastFinishedEpochs(opts *bind.CallOpts, numOfEpochs *big.Int) ([]MasonryPlusIndexedHistoricEpochData, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "getLastFinishedEpochs", numOfEpochs)

	if err != nil {
		return *new([]MasonryPlusIndexedHistoricEpochData), err
	}

	out0 := *abi.ConvertType(out[0], new([]MasonryPlusIndexedHistoricEpochData)).(*[]MasonryPlusIndexedHistoricEpochData)

	return out0, err

}

// GetLastFinishedEpochs is a free data retrieval call binding the contract method 0xa5586593.
//
// Solidity: function getLastFinishedEpochs(uint256 numOfEpochs) view returns((uint256,(bool,bool,uint256))[])
func (_Tombplus *TombplusSession) GetLastFinishedEpochs(numOfEpochs *big.Int) ([]MasonryPlusIndexedHistoricEpochData, error) {
	return _Tombplus.Contract.GetLastFinishedEpochs(&_Tombplus.CallOpts, numOfEpochs)
}

// GetLastFinishedEpochs is a free data retrieval call binding the contract method 0xa5586593.
//
// Solidity: function getLastFinishedEpochs(uint256 numOfEpochs) view returns((uint256,(bool,bool,uint256))[])
func (_Tombplus *TombplusCallerSession) GetLastFinishedEpochs(numOfEpochs *big.Int) ([]MasonryPlusIndexedHistoricEpochData, error) {
	return _Tombplus.Contract.GetLastFinishedEpochs(&_Tombplus.CallOpts, numOfEpochs)
}

// GetTimeInEpoch is a free data retrieval call binding the contract method 0x3e8add35.
//
// Solidity: function getTimeInEpoch(uint256 t) view returns(uint256)
func (_Tombplus *TombplusCaller) GetTimeInEpoch(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "getTimeInEpoch", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeInEpoch is a free data retrieval call binding the contract method 0x3e8add35.
//
// Solidity: function getTimeInEpoch(uint256 t) view returns(uint256)
func (_Tombplus *TombplusSession) GetTimeInEpoch(t *big.Int) (*big.Int, error) {
	return _Tombplus.Contract.GetTimeInEpoch(&_Tombplus.CallOpts, t)
}

// GetTimeInEpoch is a free data retrieval call binding the contract method 0x3e8add35.
//
// Solidity: function getTimeInEpoch(uint256 t) view returns(uint256)
func (_Tombplus *TombplusCallerSession) GetTimeInEpoch(t *big.Int) (*big.Int, error) {
	return _Tombplus.Contract.GetTimeInEpoch(&_Tombplus.CallOpts, t)
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

// GetUserFlips is a free data retrieval call binding the contract method 0x75271cb4.
//
// Solidity: function getUserFlips(address user) view returns((uint256,bool,uint256)[])
func (_Tombplus *TombplusCaller) GetUserFlips(opts *bind.CallOpts, user common.Address) ([]MasonryPlusFlip, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "getUserFlips", user)

	if err != nil {
		return *new([]MasonryPlusFlip), err
	}

	out0 := *abi.ConvertType(out[0], new([]MasonryPlusFlip)).(*[]MasonryPlusFlip)

	return out0, err

}

// GetUserFlips is a free data retrieval call binding the contract method 0x75271cb4.
//
// Solidity: function getUserFlips(address user) view returns((uint256,bool,uint256)[])
func (_Tombplus *TombplusSession) GetUserFlips(user common.Address) ([]MasonryPlusFlip, error) {
	return _Tombplus.Contract.GetUserFlips(&_Tombplus.CallOpts, user)
}

// GetUserFlips is a free data retrieval call binding the contract method 0x75271cb4.
//
// Solidity: function getUserFlips(address user) view returns((uint256,bool,uint256)[])
func (_Tombplus *TombplusCallerSession) GetUserFlips(user common.Address) ([]MasonryPlusFlip, error) {
	return _Tombplus.Contract.GetUserFlips(&_Tombplus.CallOpts, user)
}

// HistoricEpochData is a free data retrieval call binding the contract method 0x256bd5cf.
//
// Solidity: function historicEpochData(uint256 ) view returns(bool volatilityReached, bool winnerIsUp, uint256 rewardPerWinningVoteNumerator)
func (_Tombplus *TombplusCaller) HistoricEpochData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	VolatilityReached             bool
	WinnerIsUp                    bool
	RewardPerWinningVoteNumerator *big.Int
}, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "historicEpochData", arg0)

	outstruct := new(struct {
		VolatilityReached             bool
		WinnerIsUp                    bool
		RewardPerWinningVoteNumerator *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VolatilityReached = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.WinnerIsUp = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RewardPerWinningVoteNumerator = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HistoricEpochData is a free data retrieval call binding the contract method 0x256bd5cf.
//
// Solidity: function historicEpochData(uint256 ) view returns(bool volatilityReached, bool winnerIsUp, uint256 rewardPerWinningVoteNumerator)
func (_Tombplus *TombplusSession) HistoricEpochData(arg0 *big.Int) (struct {
	VolatilityReached             bool
	WinnerIsUp                    bool
	RewardPerWinningVoteNumerator *big.Int
}, error) {
	return _Tombplus.Contract.HistoricEpochData(&_Tombplus.CallOpts, arg0)
}

// HistoricEpochData is a free data retrieval call binding the contract method 0x256bd5cf.
//
// Solidity: function historicEpochData(uint256 ) view returns(bool volatilityReached, bool winnerIsUp, uint256 rewardPerWinningVoteNumerator)
func (_Tombplus *TombplusCallerSession) HistoricEpochData(arg0 *big.Int) (struct {
	VolatilityReached             bool
	WinnerIsUp                    bool
	RewardPerWinningVoteNumerator *big.Int
}, error) {
	return _Tombplus.Contract.HistoricEpochData(&_Tombplus.CallOpts, arg0)
}

// LastFinishedEpochId is a free data retrieval call binding the contract method 0x9d656aad.
//
// Solidity: function lastFinishedEpochId() view returns(uint256)
func (_Tombplus *TombplusCaller) LastFinishedEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "lastFinishedEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinishedEpochId is a free data retrieval call binding the contract method 0x9d656aad.
//
// Solidity: function lastFinishedEpochId() view returns(uint256)
func (_Tombplus *TombplusSession) LastFinishedEpochId() (*big.Int, error) {
	return _Tombplus.Contract.LastFinishedEpochId(&_Tombplus.CallOpts)
}

// LastFinishedEpochId is a free data retrieval call binding the contract method 0x9d656aad.
//
// Solidity: function lastFinishedEpochId() view returns(uint256)
func (_Tombplus *TombplusCallerSession) LastFinishedEpochId() (*big.Int, error) {
	return _Tombplus.Contract.LastFinishedEpochId(&_Tombplus.CallOpts)
}

// Lottery is a free data retrieval call binding the contract method 0xba13a572.
//
// Solidity: function lottery() view returns(address)
func (_Tombplus *TombplusCaller) Lottery(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "lottery")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lottery is a free data retrieval call binding the contract method 0xba13a572.
//
// Solidity: function lottery() view returns(address)
func (_Tombplus *TombplusSession) Lottery() (common.Address, error) {
	return _Tombplus.Contract.Lottery(&_Tombplus.CallOpts)
}

// Lottery is a free data retrieval call binding the contract method 0xba13a572.
//
// Solidity: function lottery() view returns(address)
func (_Tombplus *TombplusCallerSession) Lottery() (common.Address, error) {
	return _Tombplus.Contract.Lottery(&_Tombplus.CallOpts)
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

// MaxAllowedFutureFlipsOverride is a free data retrieval call binding the contract method 0xc814e77b.
//
// Solidity: function maxAllowedFutureFlipsOverride(address ) view returns(uint256)
func (_Tombplus *TombplusCaller) MaxAllowedFutureFlipsOverride(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "maxAllowedFutureFlipsOverride", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAllowedFutureFlipsOverride is a free data retrieval call binding the contract method 0xc814e77b.
//
// Solidity: function maxAllowedFutureFlipsOverride(address ) view returns(uint256)
func (_Tombplus *TombplusSession) MaxAllowedFutureFlipsOverride(arg0 common.Address) (*big.Int, error) {
	return _Tombplus.Contract.MaxAllowedFutureFlipsOverride(&_Tombplus.CallOpts, arg0)
}

// MaxAllowedFutureFlipsOverride is a free data retrieval call binding the contract method 0xc814e77b.
//
// Solidity: function maxAllowedFutureFlipsOverride(address ) view returns(uint256)
func (_Tombplus *TombplusCallerSession) MaxAllowedFutureFlipsOverride(arg0 common.Address) (*big.Int, error) {
	return _Tombplus.Contract.MaxAllowedFutureFlipsOverride(&_Tombplus.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Tombplus *TombplusCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Tombplus *TombplusSession) Operator() (common.Address, error) {
	return _Tombplus.Contract.Operator(&_Tombplus.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Tombplus *TombplusCallerSession) Operator() (common.Address, error) {
	return _Tombplus.Contract.Operator(&_Tombplus.CallOpts)
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

// PauseGameAtEpoch is a free data retrieval call binding the contract method 0x9432fd84.
//
// Solidity: function pauseGameAtEpoch() view returns(uint256)
func (_Tombplus *TombplusCaller) PauseGameAtEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "pauseGameAtEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PauseGameAtEpoch is a free data retrieval call binding the contract method 0x9432fd84.
//
// Solidity: function pauseGameAtEpoch() view returns(uint256)
func (_Tombplus *TombplusSession) PauseGameAtEpoch() (*big.Int, error) {
	return _Tombplus.Contract.PauseGameAtEpoch(&_Tombplus.CallOpts)
}

// PauseGameAtEpoch is a free data retrieval call binding the contract method 0x9432fd84.
//
// Solidity: function pauseGameAtEpoch() view returns(uint256)
func (_Tombplus *TombplusCallerSession) PauseGameAtEpoch() (*big.Int, error) {
	return _Tombplus.Contract.PauseGameAtEpoch(&_Tombplus.CallOpts)
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

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_Tombplus *TombplusCaller) Pyth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "pyth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_Tombplus *TombplusSession) Pyth() (common.Address, error) {
	return _Tombplus.Contract.Pyth(&_Tombplus.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_Tombplus *TombplusCallerSession) Pyth() (common.Address, error) {
	return _Tombplus.Contract.Pyth(&_Tombplus.CallOpts)
}

// PythPriceID is a free data retrieval call binding the contract method 0xa2673ff4.
//
// Solidity: function pythPriceID() view returns(bytes32)
func (_Tombplus *TombplusCaller) PythPriceID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "pythPriceID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PythPriceID is a free data retrieval call binding the contract method 0xa2673ff4.
//
// Solidity: function pythPriceID() view returns(bytes32)
func (_Tombplus *TombplusSession) PythPriceID() ([32]byte, error) {
	return _Tombplus.Contract.PythPriceID(&_Tombplus.CallOpts)
}

// PythPriceID is a free data retrieval call binding the contract method 0xa2673ff4.
//
// Solidity: function pythPriceID() view returns(bytes32)
func (_Tombplus *TombplusCallerSession) PythPriceID() ([32]byte, error) {
	return _Tombplus.Contract.PythPriceID(&_Tombplus.CallOpts)
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

// UpcomingEpochData is a free data retrieval call binding the contract method 0x2ee31206.
//
// Solidity: function upcomingEpochData(uint256 ) view returns(uint256 tshareVotesUp, uint256 tshareVotesDown, uint256 tshareUsersUp, uint256 tshareUsersDown, uint256 tshareTotalSupply)
func (_Tombplus *TombplusCaller) UpcomingEpochData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TshareVotesUp     *big.Int
	TshareVotesDown   *big.Int
	TshareUsersUp     *big.Int
	TshareUsersDown   *big.Int
	TshareTotalSupply *big.Int
}, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "upcomingEpochData", arg0)

	outstruct := new(struct {
		TshareVotesUp     *big.Int
		TshareVotesDown   *big.Int
		TshareUsersUp     *big.Int
		TshareUsersDown   *big.Int
		TshareTotalSupply *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TshareVotesUp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TshareVotesDown = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TshareUsersUp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TshareUsersDown = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TshareTotalSupply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UpcomingEpochData is a free data retrieval call binding the contract method 0x2ee31206.
//
// Solidity: function upcomingEpochData(uint256 ) view returns(uint256 tshareVotesUp, uint256 tshareVotesDown, uint256 tshareUsersUp, uint256 tshareUsersDown, uint256 tshareTotalSupply)
func (_Tombplus *TombplusSession) UpcomingEpochData(arg0 *big.Int) (struct {
	TshareVotesUp     *big.Int
	TshareVotesDown   *big.Int
	TshareUsersUp     *big.Int
	TshareUsersDown   *big.Int
	TshareTotalSupply *big.Int
}, error) {
	return _Tombplus.Contract.UpcomingEpochData(&_Tombplus.CallOpts, arg0)
}

// UpcomingEpochData is a free data retrieval call binding the contract method 0x2ee31206.
//
// Solidity: function upcomingEpochData(uint256 ) view returns(uint256 tshareVotesUp, uint256 tshareVotesDown, uint256 tshareUsersUp, uint256 tshareUsersDown, uint256 tshareTotalSupply)
func (_Tombplus *TombplusCallerSession) UpcomingEpochData(arg0 *big.Int) (struct {
	TshareVotesUp     *big.Int
	TshareVotesDown   *big.Int
	TshareUsersUp     *big.Int
	TshareUsersDown   *big.Int
	TshareTotalSupply *big.Int
}, error) {
	return _Tombplus.Contract.UpcomingEpochData(&_Tombplus.CallOpts, arg0)
}

// UpgradedMasonry is a free data retrieval call binding the contract method 0x902d2684.
//
// Solidity: function upgradedMasonry() view returns(address)
func (_Tombplus *TombplusCaller) UpgradedMasonry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "upgradedMasonry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradedMasonry is a free data retrieval call binding the contract method 0x902d2684.
//
// Solidity: function upgradedMasonry() view returns(address)
func (_Tombplus *TombplusSession) UpgradedMasonry() (common.Address, error) {
	return _Tombplus.Contract.UpgradedMasonry(&_Tombplus.CallOpts)
}

// UpgradedMasonry is a free data retrieval call binding the contract method 0x902d2684.
//
// Solidity: function upgradedMasonry() view returns(address)
func (_Tombplus *TombplusCallerSession) UpgradedMasonry() (common.Address, error) {
	return _Tombplus.Contract.UpgradedMasonry(&_Tombplus.CallOpts)
}

// UserCanPlaceOneMoreFlip is a free data retrieval call binding the contract method 0xca4d56ce.
//
// Solidity: function userCanPlaceOneMoreFlip(address user) view returns(bool)
func (_Tombplus *TombplusCaller) UserCanPlaceOneMoreFlip(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Tombplus.contract.Call(opts, &out, "userCanPlaceOneMoreFlip", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserCanPlaceOneMoreFlip is a free data retrieval call binding the contract method 0xca4d56ce.
//
// Solidity: function userCanPlaceOneMoreFlip(address user) view returns(bool)
func (_Tombplus *TombplusSession) UserCanPlaceOneMoreFlip(user common.Address) (bool, error) {
	return _Tombplus.Contract.UserCanPlaceOneMoreFlip(&_Tombplus.CallOpts, user)
}

// UserCanPlaceOneMoreFlip is a free data retrieval call binding the contract method 0xca4d56ce.
//
// Solidity: function userCanPlaceOneMoreFlip(address user) view returns(bool)
func (_Tombplus *TombplusCallerSession) UserCanPlaceOneMoreFlip(user common.Address) (bool, error) {
	return _Tombplus.Contract.UserCanPlaceOneMoreFlip(&_Tombplus.CallOpts, user)
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

// BatchSetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xf3771d90.
//
// Solidity: function batchSetMaxAllowedFutureFlipsOverride(address[] users, uint256[] numsOfFlips) returns()
func (_Tombplus *TombplusTransactor) BatchSetMaxAllowedFutureFlipsOverride(opts *bind.TransactOpts, users []common.Address, numsOfFlips []*big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "batchSetMaxAllowedFutureFlipsOverride", users, numsOfFlips)
}

// BatchSetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xf3771d90.
//
// Solidity: function batchSetMaxAllowedFutureFlipsOverride(address[] users, uint256[] numsOfFlips) returns()
func (_Tombplus *TombplusSession) BatchSetMaxAllowedFutureFlipsOverride(users []common.Address, numsOfFlips []*big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.BatchSetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, users, numsOfFlips)
}

// BatchSetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xf3771d90.
//
// Solidity: function batchSetMaxAllowedFutureFlipsOverride(address[] users, uint256[] numsOfFlips) returns()
func (_Tombplus *TombplusTransactorSession) BatchSetMaxAllowedFutureFlipsOverride(users []common.Address, numsOfFlips []*big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.BatchSetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, users, numsOfFlips)
}

// BatchUnsetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0x3bc8a212.
//
// Solidity: function batchUnsetMaxAllowedFutureFlipsOverride(address[] users) returns()
func (_Tombplus *TombplusTransactor) BatchUnsetMaxAllowedFutureFlipsOverride(opts *bind.TransactOpts, users []common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "batchUnsetMaxAllowedFutureFlipsOverride", users)
}

// BatchUnsetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0x3bc8a212.
//
// Solidity: function batchUnsetMaxAllowedFutureFlipsOverride(address[] users) returns()
func (_Tombplus *TombplusSession) BatchUnsetMaxAllowedFutureFlipsOverride(users []common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.BatchUnsetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, users)
}

// BatchUnsetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0x3bc8a212.
//
// Solidity: function batchUnsetMaxAllowedFutureFlipsOverride(address[] users) returns()
func (_Tombplus *TombplusTransactorSession) BatchUnsetMaxAllowedFutureFlipsOverride(users []common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.BatchUnsetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, users)
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

// EmergencyFinishEpoch is a paid mutator transaction binding the contract method 0x88b9a220.
//
// Solidity: function emergencyFinishEpoch(uint256 startTimestamp, uint256 endTimestamp, uint256 startPrice, uint256 endPrice) returns()
func (_Tombplus *TombplusTransactor) EmergencyFinishEpoch(opts *bind.TransactOpts, startTimestamp *big.Int, endTimestamp *big.Int, startPrice *big.Int, endPrice *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "emergencyFinishEpoch", startTimestamp, endTimestamp, startPrice, endPrice)
}

// EmergencyFinishEpoch is a paid mutator transaction binding the contract method 0x88b9a220.
//
// Solidity: function emergencyFinishEpoch(uint256 startTimestamp, uint256 endTimestamp, uint256 startPrice, uint256 endPrice) returns()
func (_Tombplus *TombplusSession) EmergencyFinishEpoch(startTimestamp *big.Int, endTimestamp *big.Int, startPrice *big.Int, endPrice *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.EmergencyFinishEpoch(&_Tombplus.TransactOpts, startTimestamp, endTimestamp, startPrice, endPrice)
}

// EmergencyFinishEpoch is a paid mutator transaction binding the contract method 0x88b9a220.
//
// Solidity: function emergencyFinishEpoch(uint256 startTimestamp, uint256 endTimestamp, uint256 startPrice, uint256 endPrice) returns()
func (_Tombplus *TombplusTransactorSession) EmergencyFinishEpoch(startTimestamp *big.Int, endTimestamp *big.Int, startPrice *big.Int, endPrice *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.EmergencyFinishEpoch(&_Tombplus.TransactOpts, startTimestamp, endTimestamp, startPrice, endPrice)
}

// FinishEpoch is a paid mutator transaction binding the contract method 0x8e0ddafa.
//
// Solidity: function finishEpoch(uint64 startTimestamp, uint64 endTimestamp, bytes[] startPriceData, bytes[] endPriceData) payable returns()
func (_Tombplus *TombplusTransactor) FinishEpoch(opts *bind.TransactOpts, startTimestamp uint64, endTimestamp uint64, startPriceData [][]byte, endPriceData [][]byte) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "finishEpoch", startTimestamp, endTimestamp, startPriceData, endPriceData)
}

// FinishEpoch is a paid mutator transaction binding the contract method 0x8e0ddafa.
//
// Solidity: function finishEpoch(uint64 startTimestamp, uint64 endTimestamp, bytes[] startPriceData, bytes[] endPriceData) payable returns()
func (_Tombplus *TombplusSession) FinishEpoch(startTimestamp uint64, endTimestamp uint64, startPriceData [][]byte, endPriceData [][]byte) (*types.Transaction, error) {
	return _Tombplus.Contract.FinishEpoch(&_Tombplus.TransactOpts, startTimestamp, endTimestamp, startPriceData, endPriceData)
}

// FinishEpoch is a paid mutator transaction binding the contract method 0x8e0ddafa.
//
// Solidity: function finishEpoch(uint64 startTimestamp, uint64 endTimestamp, bytes[] startPriceData, bytes[] endPriceData) payable returns()
func (_Tombplus *TombplusTransactorSession) FinishEpoch(startTimestamp uint64, endTimestamp uint64, startPriceData [][]byte, endPriceData [][]byte) (*types.Transaction, error) {
	return _Tombplus.Contract.FinishEpoch(&_Tombplus.TransactOpts, startTimestamp, endTimestamp, startPriceData, endPriceData)
}

// FinishEpochPermissionless is a paid mutator transaction binding the contract method 0x93ce912c.
//
// Solidity: function finishEpochPermissionless(uint64 startTimestamp, uint64 endTimestamp, bytes[] startPriceData, bytes[] endPriceData) payable returns()
func (_Tombplus *TombplusTransactor) FinishEpochPermissionless(opts *bind.TransactOpts, startTimestamp uint64, endTimestamp uint64, startPriceData [][]byte, endPriceData [][]byte) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "finishEpochPermissionless", startTimestamp, endTimestamp, startPriceData, endPriceData)
}

// FinishEpochPermissionless is a paid mutator transaction binding the contract method 0x93ce912c.
//
// Solidity: function finishEpochPermissionless(uint64 startTimestamp, uint64 endTimestamp, bytes[] startPriceData, bytes[] endPriceData) payable returns()
func (_Tombplus *TombplusSession) FinishEpochPermissionless(startTimestamp uint64, endTimestamp uint64, startPriceData [][]byte, endPriceData [][]byte) (*types.Transaction, error) {
	return _Tombplus.Contract.FinishEpochPermissionless(&_Tombplus.TransactOpts, startTimestamp, endTimestamp, startPriceData, endPriceData)
}

// FinishEpochPermissionless is a paid mutator transaction binding the contract method 0x93ce912c.
//
// Solidity: function finishEpochPermissionless(uint64 startTimestamp, uint64 endTimestamp, bytes[] startPriceData, bytes[] endPriceData) payable returns()
func (_Tombplus *TombplusTransactorSession) FinishEpochPermissionless(startTimestamp uint64, endTimestamp uint64, startPriceData [][]byte, endPriceData [][]byte) (*types.Transaction, error) {
	return _Tombplus.Contract.FinishEpochPermissionless(&_Tombplus.TransactOpts, startTimestamp, endTimestamp, startPriceData, endPriceData)
}

// Flip is a paid mutator transaction binding the contract method 0x1d263f67.
//
// Solidity: function flip(bool up) returns()
func (_Tombplus *TombplusTransactor) Flip(opts *bind.TransactOpts, up bool) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "flip", up)
}

// Flip is a paid mutator transaction binding the contract method 0x1d263f67.
//
// Solidity: function flip(bool up) returns()
func (_Tombplus *TombplusSession) Flip(up bool) (*types.Transaction, error) {
	return _Tombplus.Contract.Flip(&_Tombplus.TransactOpts, up)
}

// Flip is a paid mutator transaction binding the contract method 0x1d263f67.
//
// Solidity: function flip(bool up) returns()
func (_Tombplus *TombplusTransactorSession) Flip(up bool) (*types.Transaction, error) {
	return _Tombplus.Contract.Flip(&_Tombplus.TransactOpts, up)
}

// FlipMultiple is a paid mutator transaction binding the contract method 0x52add018.
//
// Solidity: function flipMultiple(bool[] ups) returns()
func (_Tombplus *TombplusTransactor) FlipMultiple(opts *bind.TransactOpts, ups []bool) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "flipMultiple", ups)
}

// FlipMultiple is a paid mutator transaction binding the contract method 0x52add018.
//
// Solidity: function flipMultiple(bool[] ups) returns()
func (_Tombplus *TombplusSession) FlipMultiple(ups []bool) (*types.Transaction, error) {
	return _Tombplus.Contract.FlipMultiple(&_Tombplus.TransactOpts, ups)
}

// FlipMultiple is a paid mutator transaction binding the contract method 0x52add018.
//
// Solidity: function flipMultiple(bool[] ups) returns()
func (_Tombplus *TombplusTransactorSession) FlipMultiple(ups []bool) (*types.Transaction, error) {
	return _Tombplus.Contract.FlipMultiple(&_Tombplus.TransactOpts, ups)
}

// PauseGame is a paid mutator transaction binding the contract method 0x1214e31e.
//
// Solidity: function pauseGame(uint256 epochId) returns()
func (_Tombplus *TombplusTransactor) PauseGame(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "pauseGame", epochId)
}

// PauseGame is a paid mutator transaction binding the contract method 0x1214e31e.
//
// Solidity: function pauseGame(uint256 epochId) returns()
func (_Tombplus *TombplusSession) PauseGame(epochId *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.PauseGame(&_Tombplus.TransactOpts, epochId)
}

// PauseGame is a paid mutator transaction binding the contract method 0x1214e31e.
//
// Solidity: function pauseGame(uint256 epochId) returns()
func (_Tombplus *TombplusTransactorSession) PauseGame(epochId *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.PauseGame(&_Tombplus.TransactOpts, epochId)
}

// PauseGameImmediately is a paid mutator transaction binding the contract method 0x7f463865.
//
// Solidity: function pauseGameImmediately() returns()
func (_Tombplus *TombplusTransactor) PauseGameImmediately(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "pauseGameImmediately")
}

// PauseGameImmediately is a paid mutator transaction binding the contract method 0x7f463865.
//
// Solidity: function pauseGameImmediately() returns()
func (_Tombplus *TombplusSession) PauseGameImmediately() (*types.Transaction, error) {
	return _Tombplus.Contract.PauseGameImmediately(&_Tombplus.TransactOpts)
}

// PauseGameImmediately is a paid mutator transaction binding the contract method 0x7f463865.
//
// Solidity: function pauseGameImmediately() returns()
func (_Tombplus *TombplusTransactorSession) PauseGameImmediately() (*types.Transaction, error) {
	return _Tombplus.Contract.PauseGameImmediately(&_Tombplus.TransactOpts)
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

// RemoveFlips is a paid mutator transaction binding the contract method 0x96b36b48.
//
// Solidity: function removeFlips(uint256 amount) returns()
func (_Tombplus *TombplusTransactor) RemoveFlips(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "removeFlips", amount)
}

// RemoveFlips is a paid mutator transaction binding the contract method 0x96b36b48.
//
// Solidity: function removeFlips(uint256 amount) returns()
func (_Tombplus *TombplusSession) RemoveFlips(amount *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.RemoveFlips(&_Tombplus.TransactOpts, amount)
}

// RemoveFlips is a paid mutator transaction binding the contract method 0x96b36b48.
//
// Solidity: function removeFlips(uint256 amount) returns()
func (_Tombplus *TombplusTransactorSession) RemoveFlips(amount *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.RemoveFlips(&_Tombplus.TransactOpts, amount)
}

// SetDistribCfg is a paid mutator transaction binding the contract method 0xb7c7f9e4.
//
// Solidity: function setDistribCfg(address _distributor, uint256 _distribRate) returns()
func (_Tombplus *TombplusTransactor) SetDistribCfg(opts *bind.TransactOpts, _distributor common.Address, _distribRate *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setDistribCfg", _distributor, _distribRate)
}

// SetDistribCfg is a paid mutator transaction binding the contract method 0xb7c7f9e4.
//
// Solidity: function setDistribCfg(address _distributor, uint256 _distribRate) returns()
func (_Tombplus *TombplusSession) SetDistribCfg(_distributor common.Address, _distribRate *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetDistribCfg(&_Tombplus.TransactOpts, _distributor, _distribRate)
}

// SetDistribCfg is a paid mutator transaction binding the contract method 0xb7c7f9e4.
//
// Solidity: function setDistribCfg(address _distributor, uint256 _distribRate) returns()
func (_Tombplus *TombplusTransactorSession) SetDistribCfg(_distributor common.Address, _distribRate *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetDistribCfg(&_Tombplus.TransactOpts, _distributor, _distribRate)
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

// SetGameStartTimestamp is a paid mutator transaction binding the contract method 0xaa8cc03a.
//
// Solidity: function setGameStartTimestamp(uint256 _gameStartTimestamp) returns()
func (_Tombplus *TombplusTransactor) SetGameStartTimestamp(opts *bind.TransactOpts, _gameStartTimestamp *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setGameStartTimestamp", _gameStartTimestamp)
}

// SetGameStartTimestamp is a paid mutator transaction binding the contract method 0xaa8cc03a.
//
// Solidity: function setGameStartTimestamp(uint256 _gameStartTimestamp) returns()
func (_Tombplus *TombplusSession) SetGameStartTimestamp(_gameStartTimestamp *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetGameStartTimestamp(&_Tombplus.TransactOpts, _gameStartTimestamp)
}

// SetGameStartTimestamp is a paid mutator transaction binding the contract method 0xaa8cc03a.
//
// Solidity: function setGameStartTimestamp(uint256 _gameStartTimestamp) returns()
func (_Tombplus *TombplusTransactorSession) SetGameStartTimestamp(_gameStartTimestamp *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetGameStartTimestamp(&_Tombplus.TransactOpts, _gameStartTimestamp)
}

// SetLottery is a paid mutator transaction binding the contract method 0xf298083b.
//
// Solidity: function setLottery(address _lottery) returns()
func (_Tombplus *TombplusTransactor) SetLottery(opts *bind.TransactOpts, _lottery common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setLottery", _lottery)
}

// SetLottery is a paid mutator transaction binding the contract method 0xf298083b.
//
// Solidity: function setLottery(address _lottery) returns()
func (_Tombplus *TombplusSession) SetLottery(_lottery common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetLottery(&_Tombplus.TransactOpts, _lottery)
}

// SetLottery is a paid mutator transaction binding the contract method 0xf298083b.
//
// Solidity: function setLottery(address _lottery) returns()
func (_Tombplus *TombplusTransactorSession) SetLottery(_lottery common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetLottery(&_Tombplus.TransactOpts, _lottery)
}

// SetMaxAllowedFutureFlips is a paid mutator transaction binding the contract method 0xd1f9574c.
//
// Solidity: function setMaxAllowedFutureFlips(uint256 _maxAllowedFutureFlips) returns()
func (_Tombplus *TombplusTransactor) SetMaxAllowedFutureFlips(opts *bind.TransactOpts, _maxAllowedFutureFlips *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setMaxAllowedFutureFlips", _maxAllowedFutureFlips)
}

// SetMaxAllowedFutureFlips is a paid mutator transaction binding the contract method 0xd1f9574c.
//
// Solidity: function setMaxAllowedFutureFlips(uint256 _maxAllowedFutureFlips) returns()
func (_Tombplus *TombplusSession) SetMaxAllowedFutureFlips(_maxAllowedFutureFlips *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetMaxAllowedFutureFlips(&_Tombplus.TransactOpts, _maxAllowedFutureFlips)
}

// SetMaxAllowedFutureFlips is a paid mutator transaction binding the contract method 0xd1f9574c.
//
// Solidity: function setMaxAllowedFutureFlips(uint256 _maxAllowedFutureFlips) returns()
func (_Tombplus *TombplusTransactorSession) SetMaxAllowedFutureFlips(_maxAllowedFutureFlips *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetMaxAllowedFutureFlips(&_Tombplus.TransactOpts, _maxAllowedFutureFlips)
}

// SetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xaa3ba9b4.
//
// Solidity: function setMaxAllowedFutureFlipsOverride(address user, uint256 numOfFlips) returns()
func (_Tombplus *TombplusTransactor) SetMaxAllowedFutureFlipsOverride(opts *bind.TransactOpts, user common.Address, numOfFlips *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setMaxAllowedFutureFlipsOverride", user, numOfFlips)
}

// SetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xaa3ba9b4.
//
// Solidity: function setMaxAllowedFutureFlipsOverride(address user, uint256 numOfFlips) returns()
func (_Tombplus *TombplusSession) SetMaxAllowedFutureFlipsOverride(user common.Address, numOfFlips *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, user, numOfFlips)
}

// SetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xaa3ba9b4.
//
// Solidity: function setMaxAllowedFutureFlipsOverride(address user, uint256 numOfFlips) returns()
func (_Tombplus *TombplusTransactorSession) SetMaxAllowedFutureFlipsOverride(user common.Address, numOfFlips *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, user, numOfFlips)
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

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Tombplus *TombplusTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Tombplus *TombplusSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetOperator(&_Tombplus.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Tombplus *TombplusTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetOperator(&_Tombplus.TransactOpts, _operator)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0xadcb085a.
//
// Solidity: function setPriceFeed(address _pyth, bytes32 _pythPriceID) returns()
func (_Tombplus *TombplusTransactor) SetPriceFeed(opts *bind.TransactOpts, _pyth common.Address, _pythPriceID [32]byte) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setPriceFeed", _pyth, _pythPriceID)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0xadcb085a.
//
// Solidity: function setPriceFeed(address _pyth, bytes32 _pythPriceID) returns()
func (_Tombplus *TombplusSession) SetPriceFeed(_pyth common.Address, _pythPriceID [32]byte) (*types.Transaction, error) {
	return _Tombplus.Contract.SetPriceFeed(&_Tombplus.TransactOpts, _pyth, _pythPriceID)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0xadcb085a.
//
// Solidity: function setPriceFeed(address _pyth, bytes32 _pythPriceID) returns()
func (_Tombplus *TombplusTransactorSession) SetPriceFeed(_pyth common.Address, _pythPriceID [32]byte) (*types.Transaction, error) {
	return _Tombplus.Contract.SetPriceFeed(&_Tombplus.TransactOpts, _pyth, _pythPriceID)
}

// SetPythAllowedTimeOvershoot is a paid mutator transaction binding the contract method 0xce4125d8.
//
// Solidity: function setPythAllowedTimeOvershoot(uint64 _pythAllowedTimeOvershoot) returns()
func (_Tombplus *TombplusTransactor) SetPythAllowedTimeOvershoot(opts *bind.TransactOpts, _pythAllowedTimeOvershoot uint64) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setPythAllowedTimeOvershoot", _pythAllowedTimeOvershoot)
}

// SetPythAllowedTimeOvershoot is a paid mutator transaction binding the contract method 0xce4125d8.
//
// Solidity: function setPythAllowedTimeOvershoot(uint64 _pythAllowedTimeOvershoot) returns()
func (_Tombplus *TombplusSession) SetPythAllowedTimeOvershoot(_pythAllowedTimeOvershoot uint64) (*types.Transaction, error) {
	return _Tombplus.Contract.SetPythAllowedTimeOvershoot(&_Tombplus.TransactOpts, _pythAllowedTimeOvershoot)
}

// SetPythAllowedTimeOvershoot is a paid mutator transaction binding the contract method 0xce4125d8.
//
// Solidity: function setPythAllowedTimeOvershoot(uint64 _pythAllowedTimeOvershoot) returns()
func (_Tombplus *TombplusTransactorSession) SetPythAllowedTimeOvershoot(_pythAllowedTimeOvershoot uint64) (*types.Transaction, error) {
	return _Tombplus.Contract.SetPythAllowedTimeOvershoot(&_Tombplus.TransactOpts, _pythAllowedTimeOvershoot)
}

// SetRewardPerTshareNumerator is a paid mutator transaction binding the contract method 0x694d3b2b.
//
// Solidity: function setRewardPerTshareNumerator(uint256 _rewardPerTshareNumerator) returns()
func (_Tombplus *TombplusTransactor) SetRewardPerTshareNumerator(opts *bind.TransactOpts, _rewardPerTshareNumerator *big.Int) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setRewardPerTshareNumerator", _rewardPerTshareNumerator)
}

// SetRewardPerTshareNumerator is a paid mutator transaction binding the contract method 0x694d3b2b.
//
// Solidity: function setRewardPerTshareNumerator(uint256 _rewardPerTshareNumerator) returns()
func (_Tombplus *TombplusSession) SetRewardPerTshareNumerator(_rewardPerTshareNumerator *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetRewardPerTshareNumerator(&_Tombplus.TransactOpts, _rewardPerTshareNumerator)
}

// SetRewardPerTshareNumerator is a paid mutator transaction binding the contract method 0x694d3b2b.
//
// Solidity: function setRewardPerTshareNumerator(uint256 _rewardPerTshareNumerator) returns()
func (_Tombplus *TombplusTransactorSession) SetRewardPerTshareNumerator(_rewardPerTshareNumerator *big.Int) (*types.Transaction, error) {
	return _Tombplus.Contract.SetRewardPerTshareNumerator(&_Tombplus.TransactOpts, _rewardPerTshareNumerator)
}

// SetUpgradedMasonry is a paid mutator transaction binding the contract method 0x5b8024a6.
//
// Solidity: function setUpgradedMasonry(address _upgradedMasonry) returns()
func (_Tombplus *TombplusTransactor) SetUpgradedMasonry(opts *bind.TransactOpts, _upgradedMasonry common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "setUpgradedMasonry", _upgradedMasonry)
}

// SetUpgradedMasonry is a paid mutator transaction binding the contract method 0x5b8024a6.
//
// Solidity: function setUpgradedMasonry(address _upgradedMasonry) returns()
func (_Tombplus *TombplusSession) SetUpgradedMasonry(_upgradedMasonry common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetUpgradedMasonry(&_Tombplus.TransactOpts, _upgradedMasonry)
}

// SetUpgradedMasonry is a paid mutator transaction binding the contract method 0x5b8024a6.
//
// Solidity: function setUpgradedMasonry(address _upgradedMasonry) returns()
func (_Tombplus *TombplusTransactorSession) SetUpgradedMasonry(_upgradedMasonry common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.SetUpgradedMasonry(&_Tombplus.TransactOpts, _upgradedMasonry)
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

// UnpauseGame is a paid mutator transaction binding the contract method 0x305243d9.
//
// Solidity: function unpauseGame() returns()
func (_Tombplus *TombplusTransactor) UnpauseGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "unpauseGame")
}

// UnpauseGame is a paid mutator transaction binding the contract method 0x305243d9.
//
// Solidity: function unpauseGame() returns()
func (_Tombplus *TombplusSession) UnpauseGame() (*types.Transaction, error) {
	return _Tombplus.Contract.UnpauseGame(&_Tombplus.TransactOpts)
}

// UnpauseGame is a paid mutator transaction binding the contract method 0x305243d9.
//
// Solidity: function unpauseGame() returns()
func (_Tombplus *TombplusTransactorSession) UnpauseGame() (*types.Transaction, error) {
	return _Tombplus.Contract.UnpauseGame(&_Tombplus.TransactOpts)
}

// UnsetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xebcb109a.
//
// Solidity: function unsetMaxAllowedFutureFlipsOverride(address user) returns()
func (_Tombplus *TombplusTransactor) UnsetMaxAllowedFutureFlipsOverride(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Tombplus.contract.Transact(opts, "unsetMaxAllowedFutureFlipsOverride", user)
}

// UnsetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xebcb109a.
//
// Solidity: function unsetMaxAllowedFutureFlipsOverride(address user) returns()
func (_Tombplus *TombplusSession) UnsetMaxAllowedFutureFlipsOverride(user common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.UnsetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, user)
}

// UnsetMaxAllowedFutureFlipsOverride is a paid mutator transaction binding the contract method 0xebcb109a.
//
// Solidity: function unsetMaxAllowedFutureFlipsOverride(address user) returns()
func (_Tombplus *TombplusTransactorSession) UnsetMaxAllowedFutureFlipsOverride(user common.Address) (*types.Transaction, error) {
	return _Tombplus.Contract.UnsetMaxAllowedFutureFlipsOverride(&_Tombplus.TransactOpts, user)
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

// TombplusEmergencyFinishEpochIterator is returned from FilterEmergencyFinishEpoch and is used to iterate over the raw logs and unpacked data for EmergencyFinishEpoch events raised by the Tombplus contract.
type TombplusEmergencyFinishEpochIterator struct {
	Event *TombplusEmergencyFinishEpoch // Event containing the contract specifics and raw log

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
func (it *TombplusEmergencyFinishEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TombplusEmergencyFinishEpoch)
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
		it.Event = new(TombplusEmergencyFinishEpoch)
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
func (it *TombplusEmergencyFinishEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TombplusEmergencyFinishEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TombplusEmergencyFinishEpoch represents a EmergencyFinishEpoch event raised by the Tombplus contract.
type TombplusEmergencyFinishEpoch struct {
	StartTimestamp  *big.Int
	EndTimestamp    *big.Int
	EpochStartPrice *big.Int
	EpochEndPrice   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEmergencyFinishEpoch is a free log retrieval operation binding the contract event 0x01a916bca8c820ab1793d96564f9bb764d1d0b2358365acdcf513cb37e08e689.
//
// Solidity: event EmergencyFinishEpoch(uint256 startTimestamp, uint256 endTimestamp, uint256 epochStartPrice, uint256 epochEndPrice)
func (_Tombplus *TombplusFilterer) FilterEmergencyFinishEpoch(opts *bind.FilterOpts) (*TombplusEmergencyFinishEpochIterator, error) {

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "EmergencyFinishEpoch")
	if err != nil {
		return nil, err
	}
	return &TombplusEmergencyFinishEpochIterator{contract: _Tombplus.contract, event: "EmergencyFinishEpoch", logs: logs, sub: sub}, nil
}

// WatchEmergencyFinishEpoch is a free log subscription operation binding the contract event 0x01a916bca8c820ab1793d96564f9bb764d1d0b2358365acdcf513cb37e08e689.
//
// Solidity: event EmergencyFinishEpoch(uint256 startTimestamp, uint256 endTimestamp, uint256 epochStartPrice, uint256 epochEndPrice)
func (_Tombplus *TombplusFilterer) WatchEmergencyFinishEpoch(opts *bind.WatchOpts, sink chan<- *TombplusEmergencyFinishEpoch) (event.Subscription, error) {

	logs, sub, err := _Tombplus.contract.WatchLogs(opts, "EmergencyFinishEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TombplusEmergencyFinishEpoch)
				if err := _Tombplus.contract.UnpackLog(event, "EmergencyFinishEpoch", log); err != nil {
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

// ParseEmergencyFinishEpoch is a log parse operation binding the contract event 0x01a916bca8c820ab1793d96564f9bb764d1d0b2358365acdcf513cb37e08e689.
//
// Solidity: event EmergencyFinishEpoch(uint256 startTimestamp, uint256 endTimestamp, uint256 epochStartPrice, uint256 epochEndPrice)
func (_Tombplus *TombplusFilterer) ParseEmergencyFinishEpoch(log types.Log) (*TombplusEmergencyFinishEpoch, error) {
	event := new(TombplusEmergencyFinishEpoch)
	if err := _Tombplus.contract.UnpackLog(event, "EmergencyFinishEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	EpochStartPrice   *big.Int
	EpochEndPrice     *big.Int
	VolatilityReached bool
	WinnerIsUp        bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEpochFinished is a free log retrieval operation binding the contract event 0x7a3750c95021179860b0428af6e156f6d4cbbaae1a0adbe556d4bdc7efba97c6.
//
// Solidity: event EpochFinished(uint256 epochNumber, uint256 epochStartPrice, uint256 epochEndPrice, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusFilterer) FilterEpochFinished(opts *bind.FilterOpts) (*TombplusEpochFinishedIterator, error) {

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "EpochFinished")
	if err != nil {
		return nil, err
	}
	return &TombplusEpochFinishedIterator{contract: _Tombplus.contract, event: "EpochFinished", logs: logs, sub: sub}, nil
}

// WatchEpochFinished is a free log subscription operation binding the contract event 0x7a3750c95021179860b0428af6e156f6d4cbbaae1a0adbe556d4bdc7efba97c6.
//
// Solidity: event EpochFinished(uint256 epochNumber, uint256 epochStartPrice, uint256 epochEndPrice, bool volatilityReached, bool winnerIsUp)
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

// ParseEpochFinished is a log parse operation binding the contract event 0x7a3750c95021179860b0428af6e156f6d4cbbaae1a0adbe556d4bdc7efba97c6.
//
// Solidity: event EpochFinished(uint256 epochNumber, uint256 epochStartPrice, uint256 epochEndPrice, bool volatilityReached, bool winnerIsUp)
func (_Tombplus *TombplusFilterer) ParseEpochFinished(log types.Log) (*TombplusEpochFinished, error) {
	event := new(TombplusEpochFinished)
	if err := _Tombplus.contract.UnpackLog(event, "EpochFinished", log); err != nil {
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

// TombplusUserWonIterator is returned from FilterUserWon and is used to iterate over the raw logs and unpacked data for UserWon events raised by the Tombplus contract.
type TombplusUserWonIterator struct {
	Event *TombplusUserWon // Event containing the contract specifics and raw log

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
func (it *TombplusUserWonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TombplusUserWon)
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
		it.Event = new(TombplusUserWon)
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
func (it *TombplusUserWonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TombplusUserWonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TombplusUserWon represents a UserWon event raised by the Tombplus contract.
type TombplusUserWon struct {
	User    common.Address
	EpochId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserWon is a free log retrieval operation binding the contract event 0x5d65a272414584c001aa531e37121612a5b4b32aefddfe422f936093de33b10f.
//
// Solidity: event UserWon(address user, uint256 epochId, uint256 amount)
func (_Tombplus *TombplusFilterer) FilterUserWon(opts *bind.FilterOpts) (*TombplusUserWonIterator, error) {

	logs, sub, err := _Tombplus.contract.FilterLogs(opts, "UserWon")
	if err != nil {
		return nil, err
	}
	return &TombplusUserWonIterator{contract: _Tombplus.contract, event: "UserWon", logs: logs, sub: sub}, nil
}

// WatchUserWon is a free log subscription operation binding the contract event 0x5d65a272414584c001aa531e37121612a5b4b32aefddfe422f936093de33b10f.
//
// Solidity: event UserWon(address user, uint256 epochId, uint256 amount)
func (_Tombplus *TombplusFilterer) WatchUserWon(opts *bind.WatchOpts, sink chan<- *TombplusUserWon) (event.Subscription, error) {

	logs, sub, err := _Tombplus.contract.WatchLogs(opts, "UserWon")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TombplusUserWon)
				if err := _Tombplus.contract.UnpackLog(event, "UserWon", log); err != nil {
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

// ParseUserWon is a log parse operation binding the contract event 0x5d65a272414584c001aa531e37121612a5b4b32aefddfe422f936093de33b10f.
//
// Solidity: event UserWon(address user, uint256 epochId, uint256 amount)
func (_Tombplus *TombplusFilterer) ParseUserWon(log types.Log) (*TombplusUserWon, error) {
	event := new(TombplusUserWon)
	if err := _Tombplus.contract.UnpackLog(event, "UserWon", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
