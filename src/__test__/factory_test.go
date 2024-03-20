package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/command/contract"
	"github.com/hungdoo/bot/src/services/telecommands"
)

func TestFactoryAddJob(t *testing.T) {
	f := telecommands.NewCommandFactory()
	ret := f.Add(command.Balance, []string{"Balance", "url", "wallet1", "wallet2"})
	t.Log(ret)
	ret = f.Add(command.ContractCall, []string{"ContractCall", "url", "address", "funcSelect", "params", "0", "1", "18"})
	t.Log(ret)
	ret = f.Add(command.Tomb, []string{"Tomb", "rpc", "contract", "false", "0", "key"})
	t.Log(ret)
	// jobs, err := f.GetJobs()
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// t.Log(jobs)
}
func TestFactoryGetJobs(t *testing.T) {
	f := telecommands.NewCommandFactory()
	jobs, err := f.GetJobs()
	if err != nil {
		fmt.Print(err)
	}

	res, err := json.MarshalIndent(jobs, "", "")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(res))
}
func TestCommandMap(t *testing.T) {
	m := telecommands.CommandMap{
		"a": &contract.ContractCommand{
			Command: command.Command{
				Name:     "name_111",
				Enabled:  true,
				IdleTime: time.Second * 60,
			},
		},
		"b": &contract.ContractCommand{
			Command: command.Command{
				Name:     "name_222",
				Enabled:  false,
				IdleTime: time.Second * 60,
			},
		},
	}

	testSearchs := []string{"11", "22", "name"}
	for _, s := range testSearchs {
		res, err := json.MarshalIndent(m.Search(s), "", "")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("Test [%v]", s)
		fmt.Print(string(res))
	}
}
