package telecommands

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/hungdoo/bot/src/packages/command"
	"github.com/hungdoo/bot/src/packages/contract"
)

func TestCommandMap(t *testing.T) {
	m := CommandMap{
		"a": &contract.Command{
			Command: command.Command{
				Name:     "name_111",
				Enabled:  true,
				IdleTime: time.Second * 60,
			},
		},
		"b": &contract.Command{
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
