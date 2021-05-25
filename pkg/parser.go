package succubus

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// checkProject looks for the config file
// It returns the file or an error
func checkProject(succPath string) (data []byte, fail error) {
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

//
func stringToRules(origin string) (rules []Command, fail error) {
	return
}

//
func interfaceToCommands(origin interface{}) (rules []Command, fail error) {
	values, ok := origin.(map[string]interface{})

	if !ok {
		return rules, fmt.Errorf("malformed command declaration")
	}

	return []Command{
		{
			command: values,
			env:     []string{""},
		},
	}
}

// commandsToRules handles cases when rules as explicid with or without
// environments variables presents.
// It returns the equivalent
func commandsToRules(origin interface{}) (rules []Command, fail error) {
	if values, ok := origin.(string); ok {
		return stringToRules(values)
	}

	return interfaceToCommands(origin)
}

// interfaceToRules handles the case when the a Command is:
// - just a command
// - a command and/or env
// - a sequence of both previous
// It returns an error when the task is not a valid one
func interfaceToRules(origin interface{}) (task Task, fail error) {
	values, ok := origin.(map[string][]string)

	if !ok {
		return task, fmt.Errorf("command malformed")
	}

	if _, ok = values["rules"]; !ok {
		return task, fmt.Errorf("rules malformed, missing 'rules' definition")
	}

	rules, fail := commandsToRules(values["rules"])

	if nil != fail {
		return task, fmt.Errorf("%w;", fail)
	}

	task.rules = rules

	if _, ok = values["env"]; ok {
		task.env = values["env"]
	}

	return task, fail
}

// stringToTasks converts the basic scenario when a command is just a simple
// string.
// It returns a Task object.
func stringToTasks(origin interface{}) (task Task, fail error) {
	command, ok := origin.(string)

	if !ok {
		return task, fmt.Errorf("could not convert command")
	}

	task.rules = []Command{
		{
			env:     []string{""},
			command: command,
		},
	}

	return task, fail
}

// interfaceToTask handles the interface to Task conversion.
// It returns a sequence of Task structures.
func interfaceToTask(origin interface{}) (task Task, fail error) {
	task, fail = stringToTasks(origin)

	if nil != fail {
		return interfaceToRules(origin)
	}

	return task, fail
}

// mapToBase just handles the base tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToBase(read map[interface{}]interface{}, succ *Project) (fail error) {
	if _, ok := read["base"]; !ok {
		return fmt.Errorf("missing base tag")
	}

	tasks, ok := read["base"].(map[interface{}]interface{})

	if !ok {
		return fmt.Errorf("%w;\n missing base run tag", fail)
	}

	succ.base.run, fail = interfaceToTask(tasks["run"])

	if nil != fail {
		return fmt.Errorf("%w;\n missing base run tag", fail)
	}

	return fail
}

// mapToDev just handles the dev tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToDev(read map[interface{}]interface{}, succ *Project) (fail error) {
	if _, ok := read["dev"]; !ok {
		return fmt.Errorf("missing dev tag")
	}

	return fail
}

// mapToExtended just handles the extended tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToExtended(read map[interface{}]interface{}, succ *Project) (fail error) {

	return fail
}

// fromFileToProject as many use cases might not need Succubus full power, but to
// handle code base is easy always to handle the "worst case scenario", this
// function does that, translate the user specific configuration to a more
// generic one.
// It returns the config file as full blown Succubus standard
func fromFileToProject(read map[interface{}]interface{}) (succ Project, fail error) {
	if image, ok := read["image"]; ok {
		succ.image = fmt.Sprintf("%v", image)
	} else {
		succ.image = ""
	}

	if dockerfile, ok := read["dockerfile"]; ok {
		succ.dockerfile = fmt.Sprintf("%v", dockerfile)
	} else {
		succ.dockerfile = ""
	}

	if fail = mapToBase(read, &succ); nil != fail {
		return succ, fmt.Errorf("%w;\n error while mapping base tag", fail)
	}

	if fail = mapToDev(read, &succ); nil != fail {
		return succ, fmt.Errorf("%w;\n error while mapping dev tag", fail)
	}

	if fail = mapToExtended(read, &succ); nil != fail {
		return succ, fmt.Errorf("%w;\n error while mapping extended tag", fail)
	}

	return succ, fail
}

// ParseProject just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func ParseProject(succPath string) (succ Project, fail error) {
	read, fail := checkProject(succPath)

	if nil != fail {
		return succ, fail
	}

	fromFile := make(map[interface{}]interface{})
	fail = yaml.Unmarshal([]byte(read), &fromFile)

	if nil != fail {
		return succ, fail
	}

	return fromFileToProject(fromFile)
}
