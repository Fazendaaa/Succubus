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
func stringToRules(origin string) (rules []Command, fail error) {
	return rules, fail
}

//
func interfaceToCommands(origin interface{}) (rules []Command, fail error) {
	// values, ok := origin.(map[string]interface{})
	_, ok := origin.(map[string]interface{})

	if !ok {
		return rules, fmt.Errorf("malformed command declaration")
	}

	return rules, fail
}

// commandsToRules handles cases when rules as explicit with or without
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

// mapToBase just handles the base tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToBase(base map[interface{}]interface{}, project *Project) (fail error) {
	run, ok := base["run"]

	if !ok {
		return fmt.Errorf("%w;\nmissing base 'run' rules", fail)
	}

	project.objectives.base.run, fail = interfaceToTask(run)

	if nil != fail {
		return fmt.Errorf("%w;\nmalformed base 'run' rules", fail)
	}

	return fail
}

// mapToDev just handles the dev tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToDev(dev map[interface{}]interface{}, project *Project) (fail error) {

	return fail
}

// mapToExtended just handles the extended tag conversion.
// It returns a error if anything is wrong, like a nested structure or something
// like it.
func mapToExtended(read map[interface{}]interface{}, project *Project) (fail error) {

	return fail
}

// interfaceToObjectives
func interfaceToObjectives(read map[interface{}]interface{}, project *Project) (fail error) {
	if _, ok := read["base"]; !ok {
		return fmt.Errorf("missing 'base' objective")
	}
	if _, ok := read["dev"]; !ok {
		return fmt.Errorf("missing 'dev' objective")
	}

	base, ok := read["base"].(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("malformed 'base' objective")
	}

	dev, ok := read["dev"].(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("malformed 'dev' objective")
	}

	if fail = mapToBase(base, project); nil != fail {
		return fail
	}
	if fail = mapToDev(dev, project); nil != fail {
		return fail
	}

	if _, ok := read["extended"]; ok {
		extended, ok := read["extended"].(map[interface{}]interface{})

		if !ok {
			return fmt.Errorf("malformed 'extended' objective")
		}
		if fail = mapToExtended(extended, project); nil != fail {
			return fail
		}
	}

	return fail
}

// interfaceToProject as many use cases might not need Succubus full power, but to
// handle code base is easy always to handle the "worst case scenario", this
// function does that, translate the user specific configuration to a more
// generic one.
// It returns the config file as full blown Succubus standard
func interfaceToProject(read map[interface{}]interface{}) (project Project, fail error) {
	if _, ok := read["image"]; ok {
		project.image, fail = interfaceToImage(read["image"])

		if nil != fail {
			return project, fmt.Errorf("%w;\n'image' tag presented and malformed", fail)
		}
	}

	if _, ok := read["dockerfile"]; ok {
		project.dockerfile, fail = interfaceToDockerfile(read["dockerfile"])

		if nil != fail {
			return project, fmt.Errorf("%w;\n'dockerfile' tag presented and malformed", fail)
		}
	}

	if _, ok := read["objectives"]; !ok {
		return project, fmt.Errorf("missing 'objectives' tag")
	}

	objectives, ok := read["objectives"].(map[interface{}]interface{})

	if !ok {
		return project, fmt.Errorf("malformed 'objectives' tag")
	}

	if fail = interfaceToObjectives(objectives, &project); nil != fail {
		return project, fmt.Errorf("%w;\n'objectives' tag presented and malformed", fail)
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
