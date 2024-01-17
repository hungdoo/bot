package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/packages/command/contract"
	"github.com/hungdoo/bot/src/packages/db"
	"github.com/hungdoo/bot/src/services/telecommands"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFactory(t *testing.T) {
	f := telecommands.NewCommadFactory()
	jobs, err := f.GetJobs()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Before")
	for _, j := range jobs {
		_prev, err := j.GetPrev()
		if err != nil {
			fmt.Printf("[%s] before GetPrev failed: [%s]", j.GetName(), err)
		}
		fmt.Println(_prev)
		filter := bson.M{"name": j.GetName()}
		update := bson.M{"$set": bson.M{"prev": decimal.NewFromInt(111111).String()}}
		if err := db.GetDb().Update("commands", filter, update); err != nil {
			fmt.Printf("Job [%s] update db failed: [%s]", j.GetName(), err)
			continue
		}
	}

	fmt.Println("After")
	jobs, err = f.GetJobs()
	if err != nil {
		fmt.Print(err)
	}
	for _, j := range jobs {
		_prev, err := j.GetPrev()
		if err != nil {
			fmt.Printf("[%s] before GetPrev failed: [%s]", j.GetName(), err)
		}
		fmt.Println(_prev)
	}

	// res, err := json.MarshalIndent(j, "", "")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Print(string(res))
}
func TestCommandMap(t *testing.T) {
	m := telecommands.CommandMap{
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
