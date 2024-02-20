package tomb

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/hungdoo/bot/src/common"
	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/tombplus"
)

type TombCommand struct {
	command.Command  `json:"command" bson:"command"`
	Id               string    `json:"-" bson:"_id,unique"`
	Rpc              string    `json:"rpc" bson:"rpc"`
	Contract         string    `json:"contract" bson:"contract"`
	Up               bool      `json:"up" bson:"up"`
	PkIdx            int64     `json:"pkIdx" bson:"pkIdx"`
	Key              string    `json:"key" bson:"key"`
	SentTx           string    `json:"sent_tx" bson:"sent_tx"`
	CurrentEpoch     int64     `json:"currentEpoch" bson:"currentEpoch"`
	LastVotedEpoch   int64     `json:"lastEpoch" bson:"lastEpoch"`
	VoteEndTimestamp time.Time `json:"voteEndTimestamp" bson:"voteEndTimestamp"`
}

func (c TombCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string `json:"name"`
		Type string `json:"type"`
		// Data     []string `json:"data"`
		IdleTime         string `json:"idletime"`
		Rpc              string `json:"rpc"`
		Contract         string `json:"contract"`
		Up               bool   `json:"up"`
		PkIdx            int64  `json:"pkIdx"`
		Key              string `json:"key" bson:"key"`
		SentTx           string `json:"sent_tx" bson:"sent_tx"`
		LastVotedEpoch   int64  `json:"lastEpoch"`
		VoteEndTimestamp string `json:"voteEndTimestamp"`
		Command          string `json:"command"`
	}{
		Name: c.Name,
		Type: c.Type.String(),
		// Data:     c.Data,
		IdleTime:         c.IdleTime.String(),
		Rpc:              c.Rpc,
		Contract:         c.Contract,
		Up:               c.Up,
		PkIdx:            c.PkIdx,
		Key:              c.Key,
		SentTx:           c.SentTx,
		LastVotedEpoch:   c.LastVotedEpoch,
		VoteEndTimestamp: c.VoteEndTimestamp.String(),
		Command:          fmt.Sprintf("add tomb %s %s %s %v %v %v", c.Name, c.Rpc, c.Contract, c.Up, c.PkIdx, c.Key),
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

func (c *TombCommand) Execute(mustReport bool, subcommand string) (string, *common.ErrorWithSeverity) {
	contractAddress := ethCommon.HexToAddress(c.Contract)
	cli, err := tombplus.GetClient(c.Rpc, contractAddress)
	if err != nil {
		return "", common.NewErrorWithSeverity(common.Error, err.Error())
	}

	switch subcommand {
	case "stats":
		pk, err := LoadSecrets(int(c.PkIdx), c.Key)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		user, err := tombplus.AddressFromPriKey(pk)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		rewards := cli.GetRewards(user)

		return fmt.Sprintf("rewards: %v", rewards), nil

	case "clear":
		c.SentTx = ""
		return "", nil

	case "claim":
		pk, err := LoadSecrets(int(c.PkIdx), c.Key)
		if err != nil {
			return "", common.NewErrorWithSeverity(common.Error, err.Error())
		}
		res, err2 := cli.Claim(pk)
		if err2 != nil {
			return "", err2
		}
		// c.SentTx = res.Hash().String() // no need to record claim tx
		return res.Hash().String(), nil

	default:
		if len(c.SentTx) != 0 {
			toCheck := c.SentTx
			// Has pending tx
			if err := cli.CheckResult(toCheck); err != nil {
				return "", err
			}
			// tx successful, clear sent tx hash
			c.SentTx = ""
			c.LastVotedEpoch = c.CurrentEpoch
			c.VoteEndTimestamp = time.Time{}
			return fmt.Sprintf("tx[%s] successful", toCheck), nil
		}

		// if c.LastVotedEpoch == 0 {
		// 	lastVotedEpoch, err := cli.GetUserLastedVoteEpochId(user)
		// 	if err != nil {
		// 		return "", common.NewErrorWithSeverity(common.Error, err.Error())
		// 	}
		// 	if lastVotedEpoch.Int64() > 0 {
		// 		c.LastVotedEpoch = lastVotedEpoch.Int64()
		// 	}
		// }

		if c.VoteEndTimestamp.IsZero() { // vote end not set
			c.CurrentEpoch = cli.CurrentEpoch()

			if c.CurrentEpoch > 0 && c.CurrentEpoch > c.LastVotedEpoch {
				timestamps, err := cli.Tomb.GetEpochObservationTimestamps(&bind.CallOpts{}, big.NewInt(c.CurrentEpoch))
				if err != nil {
					return "", common.NewErrorWithSeverity(common.Info, err.Error())
				}
				c.VoteEndTimestamp = time.Unix(timestamps.ObservationStartTimestamp.Int64(), 0)
				return fmt.Sprintf("vote end timestamp set to %s", c.VoteEndTimestamp.String()), nil
			}

			if mustReport {
				return fmt.Sprintf("already voted currentEpoch[%v]/last[%v]", c.CurrentEpoch, c.LastVotedEpoch), nil
			}
			return "", nil

		} else if c.VoteEndTimestamp.After(time.Now()) && time.Until(c.VoteEndTimestamp) < 5*time.Minute { // 5min before vote end
			// determine vote side up/down
			data, err := cli.Tomb.UpcomingEpochData(&bind.CallOpts{}, big.NewInt(c.CurrentEpoch))
			if err != nil {
				return "", common.NewErrorWithSeverity(common.Info, err.Error())
			}
			var up bool
			cmp := data.TshareVotesUp.Cmp(data.TshareVotesDown)
			if cmp >= 1 {
				up = true
			} else if cmp <= -1 {
				up = false
			} else {
				up = c.Up
			}

			pk, err := LoadSecrets(int(c.PkIdx), c.Key)
			if err != nil {
				return "", common.NewErrorWithSeverity(common.Error, err.Error())
			}
			res, errWithSeverity := cli.Flip(pk, up)
			if errWithSeverity != nil {
				return "", errWithSeverity
			}
			c.SentTx = res.Hash().String()

			return fmt.Sprintf("tx[%s] sent. Voted Up:%v. Data: %+v", c.SentTx, up, data), nil
		} else if c.VoteEndTimestamp.Before(time.Now()) {
			c.VoteEndTimestamp = time.Time{}
			return "too late", nil
		}

		if mustReport {
			return fmt.Sprintf("too early till vote end. %vm left", time.Until(c.VoteEndTimestamp).Minutes()), nil
		}
		return "", nil
	}
}
