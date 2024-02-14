package bybitido

import (
	"testing"
)

func TestIdoCommand(t *testing.T) {
	cmd := &IdoCommand{}

	p, _ := cmd.Execute(true, "")
	t.Log(p)
}

func TestIdoCommand_FetchAll(t *testing.T) {
	cmd := &IdoCommand{}

	a, _ := cmd.Execute(true, "all")
	t.Log(a)
}
