package test

import (
	"testing"
	"time"

	command "github.com/hungdoo/bot/src/packages/command/common"
	"github.com/hungdoo/bot/src/services/telecommands"
)

func TestServiceLock(t *testing.T) {
	s := telecommands.NewService()
	s.Factory.Add(command.Balance, []string{"Balance1", "url", "wallet1", "wallet2"})
	s.Factory.Add(command.Tomb, []string{"Tomb1", "rpc", "contract", "false", "0", "key"})

	s2 := telecommands.NewService()

	// Check locking
	go s.Work()
	go s2.Work()

	time.Sleep(25 * time.Second)
}
