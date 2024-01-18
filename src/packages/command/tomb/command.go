package tomb

import (
	"fmt"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/tombplus"
	"github.com/shopspring/decimal"
)

type TombCommand struct {
	command.Command
	Rpc      string `bson:"rpc"`
	Contract string `bson:"contract"`
	Up       bool   `bson:"up"`
	PkIdx    int64  `bson:"pkIdx"`
	key      string `bson:"-"`
}

func (c *TombCommand) Validate(data []string) error { // rpc, contract, up, pkIdx, k
	if len(data) < 5 {
		return fmt.Errorf("invalid params: rpc, contract, up, pkIdx, k")
	}
	return nil
}

func (c *TombCommand) SetData(newValue []string) (err error) {
	if err = c.Validate(newValue); err != nil {
		return err
	}
	c.Data = newValue // TODO: refractor to store into DB dynamically
	c.Rpc = newValue[0]
	c.Contract = newValue[1]
	c.Up, err = strconv.ParseBool(newValue[2])
	if err != nil {
		return err
	}
	c.PkIdx, err = strconv.ParseInt(newValue[3], 10, 64)
	if err != nil {
		return err
	}
	c.key = newValue[4]
	return nil
}

func (c *TombCommand) Execute(_ bool, subcommand string) (string, *common.ErrorWithSeverity) {
	contractAddress := ethCommon.HexToAddress(c.Contract)
	cli, err := tombplus.GetClient(c.Rpc, contractAddress)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	pk, err := LoadSecrets(int(c.PkIdx), c.key)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	user, err := tombplus.AddressFromPriKey(pk)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	switch subcommand {
	case "stats":
		currentEpoch := cli.CurrentEpoch()
		lastVotedEpoch := c.Prev
		isVoted, err := cli.IsVotedAtEpoch(user, currentEpoch)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		max := cli.MaxAllowedFutureFlips()
		return fmt.Sprintf("cur-isVoted-lastVoted-maxFuture: %v-%v-%v-%v", currentEpoch, isVoted, lastVotedEpoch, max), nil

	case "claim":
		res, err := cli.Claim(pk)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Critical, err.Error())
		}
		return res, nil

	case "cronjob":
	default:
		gameStarted := cli.GameStarted()
		if !gameStarted {
			return "", common.NewErrorWithSeverity(common.Error, "game not started")
		}
		currentEpoch := cli.CurrentEpoch()
		voted, err := cli.IsVotedAtEpoch(user, currentEpoch)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		if voted {
			return "", common.NewErrorWithSeverity(common.Error, "already Voted")
		}
		maxFutureFlips := cli.MaxAllowedFutureFlips()
		if maxFutureFlips <= 0 {
			return "", common.NewErrorWithSeverity(common.Error, "maxFutureFlips <= 0")
		}
		res, errWithSeverity := cli.Flipmultiple(pk, currentEpoch, maxFutureFlips-1, c.Up)
		if errWithSeverity != nil {
			return "", errWithSeverity
		}
		// last voted epoch (could be in the future)
		// for reporting only
		c.SetPrev(decimal.NewFromInt(currentEpoch + maxFutureFlips - 1))
		return res, nil
	}
	return "", common.NewErrorWithSeverity(common.Info, "no action")
}
