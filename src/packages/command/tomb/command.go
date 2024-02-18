package tomb

import (
	"encoding/json"
	"fmt"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/tombplus"
)

type TombCommand struct {
	command.Command `json:"command" bson:"command"`
	Id              string `json:"-" bson:"_id,unique"`
	Rpc             string `json:"rpc" bson:"rpc"`
	Contract        string `json:"contract" bson:"contract"`
	Up              bool   `json:"up" bson:"up"`
	PkIdx           int64  `json:"pkIdx" bson:"pkIdx"`
	Key             string `json:"key" bson:"key"`
	SentTx          string `json:"sent_tx" bson:"sent_tx"`
	LastVotedEpoch  int64  `json:"lastEpoch" bson:"lastEpoch"`
}

func (c TombCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string `json:"name"`
		Type string `json:"type"`
		// Data     []string `json:"data"`
		IdleTime       string `json:"idletime"`
		Rpc            string `json:"rpc"`
		Contract       string `json:"contract"`
		Up             bool   `json:"up"`
		PkIdx          int64  `json:"pkIdx"`
		Key            string `json:"key" bson:"key"`
		SentTx         string `json:"sent_tx" bson:"sent_tx"`
		LastVotedEpoch int64  `json:"lastEpoch"`
		Command        string `json:"command"`
	}{
		Name: c.Name,
		Type: c.Type.String(),
		// Data:     c.Data,
		IdleTime:       c.IdleTime.String(),
		Rpc:            c.Rpc,
		Contract:       c.Contract,
		Up:             c.Up,
		PkIdx:          c.PkIdx,
		Key:            c.Key,
		SentTx:         c.SentTx,
		LastVotedEpoch: c.LastVotedEpoch,
		Command:        fmt.Sprintf("add tomb %s %s %s %v %v %v %v", c.Name, c.Rpc, c.Contract, c.Up, c.PkIdx, c.Key, c.SentTx),
	})
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
	// c.Data = newValue // TODO: refractor to store into DB dynamically
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
	c.Key = newValue[4]
	return nil
}

func (c *TombCommand) Execute(_ bool, subcommand string) (string, *common.ErrorWithSeverity) {
	contractAddress := ethCommon.HexToAddress(c.Contract)
	cli, err := tombplus.GetClient(c.Rpc, contractAddress)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	pk, err := LoadSecrets(int(c.PkIdx), c.Key)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}
	user, err := tombplus.AddressFromPriKey(pk)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	switch subcommand {
	case "stats":
		// currentEpoch := cli.CurrentEpoch()
		// lastVotedEpoch, err := cli.GetUserLastedVoteEpochId(user)
		// if err != nil {
		// 	return "", common.NewErrorWithSeverity(common.Error, err.Error())
		// }
		rewards := cli.GetRewards(user)

		return fmt.Sprintf("rewards: %v", rewards), nil

	case "clear":
		c.SentTx = ""
		return "", nil

	case "claim":
		res, err := cli.Claim(pk)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Critical, err.Error())
		}
		// c.SentTx = res.Hash().String() // no need to record claim tx
		return res.Hash().String(), nil

	case "cronjob":
	default:
		if len(c.SentTx) != 0 {
			toCheck := c.SentTx
			// Has pending tx
			if err := cli.CheckResult(toCheck); err != nil {
				return "", err
			}
			// tx successful, clear sent tx hash
			c.SentTx = ""
			return fmt.Sprintf("tx[%s] successful", toCheck), nil
		}

		// Start fresh, no pending tx
		// gameStarted := cli.GameStarted()
		// if !gameStarted {
		// 	return "", common.NewErrorWithSeverity(common.Error, "game not started")
		// }
		currentEpoch := cli.CurrentEpoch()
		if c.LastVotedEpoch == 0 {
			lastVotedEpoch, err := cli.GetUserLastedVoteEpochId(user)
			if err != nil {
				return "", common.NewErrorWithSeverity(common.Error, err.Error())
			}
			c.LastVotedEpoch = lastVotedEpoch.Int64()
		}
		// maxFutureFlips := cli.MaxAllowedFutureFlips()
		// if maxFutureFlips <= 0 {
		// 	return "", common.NewErrorWithSeverity(common.Error, "maxFutureFlips <= 0")
		// }
		// Obsolete
		// if currentEpoch > 0 && currentEpoch == lastVotedEpoch.Int64() {
		// 	res, errWithSeverity := cli.Flipmultiple(pk, maxFutureFlips-1, c.Up)
		// 	if errWithSeverity != nil {
		// 		return "", errWithSeverity
		// 	}
		// 	c.SentTx = res.Hash().String()
		// 	return fmt.Sprintf("tx[%s] sent", c.SentTx), nil
		// } else
		if currentEpoch > 0 && currentEpoch > c.LastVotedEpoch {
			canFlip, err := cli.CanFlipForCurrentEpoch()
			if err != nil {
				return "", common.NewErrorWithSeverity(common.Error, err.Error())
			}
			if !canFlip {
				c.LastVotedEpoch = currentEpoch // skip this epoch
				return "", common.NewErrorWithSeverity(common.Info, "too late to vote")
			}
			res, errWithSeverity := cli.Flip(pk, c.Up)
			if errWithSeverity != nil {
				return "", errWithSeverity
			}
			c.SentTx = res.Hash().String()
			c.LastVotedEpoch = currentEpoch
			return fmt.Sprintf("tx[%s] sent", c.SentTx), nil
		}
		return "", common.NewErrorWithSeverity(common.Info, "already Voted")
	}
	return "", common.NewErrorWithSeverity(common.Info, "no action")
}
