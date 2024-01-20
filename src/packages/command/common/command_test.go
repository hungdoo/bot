package command

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	cmd := Command{
		Name:     "a",
		Type:     ContractCall,
		Data:     []string{"1", "2"},
		IdleTime: time.Minute,
	}
	b, err := json.MarshalIndent(cmd, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
}
