package tomb

import (
	"fmt"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/tombplus"
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
		lastVotedEpoch, err := cli.GetUserLastedVoteEpochId(user)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		pauseEpoch := cli.GetPauseGameAtEpoch()
		mason := cli.GetUpgradedMasonry()
		return fmt.Sprintf("cur-last-pause-mason: %v-%v-%v-%v", currentEpoch, lastVotedEpoch, pauseEpoch, mason), nil

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
		lastVotedEpoch, err := cli.GetUserLastedVoteEpochId(user)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		maxFutureFlips := cli.MaxAllowedFutureFlips()
		if maxFutureFlips <= 0 {
			return "", common.NewErrorWithSeverity(common.Error, "maxFutureFlips <= 0")
		}
		if currentEpoch == lastVotedEpoch.Int64() {
			res, errWithSeverity := cli.Flipmultiple(pk, maxFutureFlips-1, c.Up)
			if errWithSeverity != nil {
				return "", errWithSeverity
			}
			return res, nil
		} else if currentEpoch > lastVotedEpoch.Int64() {
			res, errWithSeverity := cli.Flipmultiple(pk, maxFutureFlips, c.Up)
			if errWithSeverity != nil {
				return "", errWithSeverity
			}
			return res, nil
		}
		return "", common.NewErrorWithSeverity(common.Info, "already Voted")
	}
	return "", common.NewErrorWithSeverity(common.Info, "no action")
}
