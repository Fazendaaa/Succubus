package controllers

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
	succubus "github.com/Fazendaaa/Succubus/pkg"
)

func Init(params []string, projectPath string) (channel chan samael.ResponseChannel, fail error) {
	fail = succubus.CreateProject(projectPath)

	if nil != fail {
		return channel, fmt.Errorf("%w;\nerror while generating project", fail)
	}

	channel = make(chan samael.ResponseChannel)

	go func() {
		channel <- samael.ResponseChannel{
			Process:  "init",
			Response: nil,
		}
	}()

	return channel, fail
}
