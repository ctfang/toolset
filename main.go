package main

import (
	"github.com/ctfang/command"
	command2 "github.com/go-home-admin/toolset/command"
)

func main() {
	app := command.New()
	app.AddCommand(command2.ProtocCommand{})
	app.Run()
}
