package cmdparser

import (
	"bytes"
	"fmt"

	"github.com/urfave/cli"
)

var app *cli.App
var buf *bytes.Buffer

func GetParser() *cli.App {
	if app != nil {
		return app
	}
	app = cli.NewApp()
	app.Name = "tele"
	app.UsageText = "command [command options] [arguments...]"
	app.Email = "hungddoo@lol.com"
	buf = new(bytes.Buffer)
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintln(c.App.Writer, command, " not found")
	}
	app.Writer = buf
	return app
}

func GetOutput() string {
	out := buf.String()
	buf.Reset()
	return out
}
