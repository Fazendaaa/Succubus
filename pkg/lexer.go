package succubus

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// checkProject looks for the config file
// It returns the file or an error
func checkProject(projectPath string) (data []byte, fail error) {
	file, fail := readFileOrDir(projectPath)

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
func interfaceToImage(origin interface{}) (image Image, fail error) {
	return image, fail
}

//
func interfaceToDockerfile(origin interface{}) (dockerfile Dockerfile, fail error) {
	return dockerfile, fail
}

//
func simpleCommandsToRules(origin string) (rules []Command, fail error) {
	return rules, fail
}

// interfaceToContainer
func interfaceToContainer(origin interface{}) (container Container, fail error) {
	return container, fail
}

//
func complexCommandsToRules(origin interface{}) (commands []Command, fail error) {
	values, ok := origin.(map[interface{}]interface{})

	if !ok {
		return commands, fmt.Errorf("malformed command declaration")
	}

	rules, ok := values["rules"]

	if !ok {
		return commands, fmt.Errorf("malformed rules declaration")
	}

	if env, ok := values["env"]; ok {
		return commands, fmt.Errorf("malformed rules declaration")
	}

	return commands, fail
}

// commandsToRules handles cases when rules as explicit with or without
// environments variables presents.
// It returns the equivalent
func commandsToRules(origin interface{}) (rules []Command, fail error) {
	if values, ok := origin.(string); ok {
		return simpleCommandsToRules(values)
	}

	return complexCommandsToRules(origin)
}

// interfaceToRules handles the case when the a Command is:
// - just a command
// - a command and/or env
// - a sequence of both previous
// It returns an error when the task is not a valid one
func interfaceToRules(origin interface{}) (task Task, fail error) {
	values, ok := origin.(map[interface{}]interface{})

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
		// task.env = values["env"]
	}

	return task, fail
}

// stringToTasks converts the basic scenario when a command is just a simple
// string.
// It returns a Task object.
func stringToTasks(origin interface{}) (task Task, fail error) {
	// command, ok := origin.(string)
	_, ok := origin.(string)

	if !ok {
		return task, fmt.Errorf("could not convert command")
	}

	// task.rules = []Command{
	// 	{
	// 		env:     []string{""},
	// 		command: command,
	// 	},
	// }

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

// baseToObjective just handles the base tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func baseToObjective(origin map[interface{}]interface{}) (base Base, fail error) {
	run, ok := origin["run"]

	if !ok {
		return base, fmt.Errorf("%w;\nmissing 'run' rules", fail)
	}

	base.run, fail = interfaceToTask(run)

	if nil != fail {
		return base, fmt.Errorf("%w;\nmalformed 'run' rules", fail)
	}

	add, ok := origin["add"]

	if !ok {
		return base, fmt.Errorf("%w;\nmissing 'add' rules", fail)
	}

	base.run, fail = interfaceToTask(add)

	if nil != fail {
		return base, fmt.Errorf("%w;\nmalformed 'add' rules", fail)
	}

	rm, ok := origin["rm"]

	if !ok {
		return base, fmt.Errorf("%w;\nmissing 'rm' rules", fail)
	}

	base.run, fail = interfaceToTask(rm)

	if nil != fail {
		return base, fmt.Errorf("%w;\nmalformed 'rm' rules", fail)
	}

	test, ok := origin["test"]

	if !ok {
		return base, fmt.Errorf("%w;\nmissing 'test' rules", fail)
	}

	base.run, fail = interfaceToTask(test)

	if nil != fail {
		return base, fmt.Errorf("%w;\nmalformed 'test' rules", fail)
	}

	return base, fail
}

// devToObjective just handles the dev tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func devToObjective(origin map[interface{}]interface{}) (dev Dev, fail error) {
	return dev, fail
}

// extendedToObjective just handles the extended tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func extendedToObjective(origin map[interface{}]interface{}) (extendeds []Extended, fail error) {
	return extendeds, fail
}

// tasksToObjectives
// It return
func tasksToObjectives(origin map[interface{}]interface{}) (objectives Objectives, fail error) {
	if _, ok := origin["base"]; !ok {
		return objectives, fmt.Errorf("missing 'base' objective")
	}
	if _, ok := origin["dev"]; !ok {
		return objectives, fmt.Errorf("missing 'dev' objective")
	}

	base, ok := origin["base"].(map[interface{}]interface{})

	if !ok {
		return objectives, fmt.Errorf("malformed 'base' objective")
	}

	dev, ok := origin["dev"].(map[interface{}]interface{})

	if !ok {
		return objectives, fmt.Errorf("malformed 'dev' objective")
	}

	if objectives.base, fail = baseToObjective(base); nil != fail {
		return objectives, fail
	}
	if objectives.dev, fail = devToObjective(dev); nil != fail {
		return objectives, fail
	}
	if objectives.extended, fail = extendedToObjective(origin); nil != fail {
		return objectives, fail
	}

	return objectives, fail
}

// containerToProject
func containerToProject(origin map[interface{}]interface{}) (container Container, fail error) {
	if _, ok := origin["image"]; ok {
		container, fail = interfaceToContainer(origin["image"])

		if nil != fail {
			return container, fmt.Errorf("%w;\nimage presented and malformed", fail)
		}
	}

	if _, ok := origin["dockerfile"]; ok {
		container, fail = interfaceToContainer(origin["dockerfile"])

		if nil != fail {
			return container, fmt.Errorf("%w;\ndockerfile presented and malformed", fail)
		}
	}

	return container, fail
}

// objetivesToProject
// It returns all the declared objectives in the project
func objetivesToProject(origin map[interface{}]interface{}) (objectives Objectives, fail error) {
	if _, ok := origin["objectives"]; !ok {
		return objectives, fmt.Errorf("missing objectives")
	}

	read, ok := origin["objectives"].(map[interface{}]interface{})

	if !ok {
		return objectives, fmt.Errorf("malformed objectives")
	}

	objectives, fail = tasksToObjectives(read)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\nobjectives presented and malformed", fail)
	}

	return objectives, fail
}

// versionToProject
func versionToProject(read map[interface{}]interface{}) (value string, fail error) {
	return value, fail
}

// tagToProject
func tagToProject(read map[interface{}]interface{}) (value string, fail error) {
	return value, fail
}

// nameToProject
func nameToProject(read map[interface{}]interface{}) (value string, fail error) {
	return value, fail
}

// interfaceToProject as many use cases might not need Succubus full power, but to
// handle code base is easy always to handle the "worst case scenario", this
// function does that, translate the user specific configuration to a more
// generic one.
// It returns the config file as full blown Succubus standard
func interfaceToProject(read map[interface{}]interface{}) (project Project, fail error) {
	project.container, fail = containerToProject(read)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed project container", fail)
	}

	project.objectives, fail = objetivesToProject(read)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed project objectives", fail)
	}

	project.version, fail = versionToProject(read)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed project version", fail)
	}

	project.tag, fail = tagToProject(read)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed project tag", fail)
	}

	project.name, fail = nameToProject(read)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed project name", fail)
	}

	return project, fail
}

// LexProject just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func LexProject(projectPath string) (project Project, fail error) {
	data, fail := checkProject(projectPath)

	if nil != fail {
		return project, fail
	}

	read := make(map[interface{}]interface{})
	fail = yaml.Unmarshal([]byte(data), &read)

	if nil != fail {
		return project, fail
	}

	return interfaceToProject(read)
}
