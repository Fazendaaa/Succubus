package succubus

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Command struct {
	env     []string
	command string
}

type Task struct {
	env   []string
	rules []Command
}

type Base struct {
	run  Task
	test Task
	add  Task
}

type Dev struct {
	doc    Task
	anal   Task
	linter Task
}

type Config struct {
	image      string
	dockerfile string
	base       Base
	dev        Dev
	extended   interface{}
}

// checkConfig looks for the config file
// It returns the file or an error
func checkConfig(succPath string) (data []byte, fail error) {
	file, fail := readFileOrDir(succPath)

	if nil != fail {
		return nil, fail
	}

	read, fail := ioutil.ReadFile(file)

	if nil != fail {
		return nil, fail
	}

	return read, fail
}

// interfaceToCommand handles the case when the a Command is:
// - just a command
// - a command and/or env
// - a sequence of both previous
// It returns an error when the task is not a valid one
func interfaceToCommand(origin interface{}) (task Task, fail error) {
	values, ok := origin.(map[string][]string)

	if !ok {
		return task, fmt.Errorf("command malformed")
	}

	if _, ok = values["command"]; !ok {
		return task, fmt.Errorf("command malformed")
	}

	task.command = values["command"]

	if _, ok = values["env"]; ok {
		task.env = values["env"]
	}

	return task, fail
}

// stringToCommand converts the basic scenario when a command is just a simple
// string.
// It returns a Task object.
func stringToCommand(origin interface{}) (task Task, fail error) {
	command, ok := origin.(string)

	if !ok {
		return task, fmt.Errorf("could not convert command")
	}

	task.command = []string{command}

	return task, fail
}

// interfaceToTask handles the interface to Task conversion.
// It returns a sequence of Task structures.
func interfaceToTask(key string, origin interface{}) (task Task, fail error) {
	tags, ok := origin.(map[interface{}]interface{})

	if !ok {
		return task, fmt.Errorf("command malformed")
	}

	read, ok := tags[key]

	if !ok {
		return task, fmt.Errorf("missing %s tag", key)
	}

	task, fail = stringToCommand(read)

	if nil != fail {
		return interfaceToCommand(read)
	}

	return task, fail
}

// mapToBase just handles the base tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToBase(read map[interface{}]interface{}, succ *Config) (fail error) {
	if _, ok := read["base"]; !ok {
		return fmt.Errorf("missing base tag")
	}

	succ.base.run, fail = interfaceToTask("run", read["base"])

	if nil != fail {
		return fmt.Errorf("%w; missing base run tag", fail)
	}

	return fail
}

// mapToDev just handles the dev tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToDev(read map[interface{}]interface{}, succ *Config) (fail error) {
	if _, ok := read["dev"]; !ok {
		return fmt.Errorf("missing dev tag")
	}

	return fail
}

// mapToExtended just handles the extended tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToExtended(read map[interface{}]interface{}, succ *Config) (fail error) {

	return fail
}

// fromFileToConfig as many use cases might not need Succubus full power, but to
// handle code base is easy always to handle the "worst case scenario", this
// function does that, translate the user specific configuration to a more
// generic one.
// It returns the config file as full blown Succubus standard
func fromFileToConfig(read map[interface{}]interface{}) (succ Config, fail error) {
	if image, ok := read["image"]; ok {
		succ.image = fmt.Sprintf("%v", image)
	} else {
		succ.image = ""
	}

	if fail = mapToBase(read, &succ); nil != fail {
		return succ, fmt.Errorf("%w; error while mapping base tag", fail)
	}

	if fail = mapToDev(read, &succ); nil != fail {
		return succ, fmt.Errorf("%w; error while mapping dev tag", fail)
	}

	if fail = mapToExtended(read, &succ); nil != fail {
		return succ, fmt.Errorf("%w; error while mapping extended tag", fail)
	}

	return succ, fail
}

// ParseConfig just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func ParseConfig(succPath string) (succ Config, fail error) {
	read, fail := checkConfig(succPath)

	if nil != fail {
		return succ, fail
	}

	fromFile := make(map[interface{}]interface{})
	fail = yaml.Unmarshal([]byte(read), &fromFile)

	if nil != fail {
		return succ, fail
	}

	return fromFileToConfig(fromFile)
}
