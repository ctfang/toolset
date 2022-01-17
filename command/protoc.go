package command

import (
	"fmt"
	"github.com/ctfang/command"
)

type ProtocCommand struct {
}

func (p ProtocCommand) Configure() command.Configure {
	return command.Configure{
		Name:        "protoc",
		Description: "生protoc文件",
	}
}

func (p ProtocCommand) Execute(input command.Input) {
	fmt.Println("123")
}
