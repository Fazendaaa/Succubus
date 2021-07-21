package succubus

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// checkConfig looks for the config file
// It returns the file or an error
func checkConfig(projectPath string) (data []byte, fail error) {
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
	read, ok := origin.(string)

	if !ok {
		return image, fmt.Errorf("image declaration is malformed")
	}

	split := strings.Split(read, "/")

	if 1 <= len(split) {
		image.name = split[0]
	}

	return image, fail
}

//
func interfaceToDockerfile(origin interface{}) (dockerfile Dockerfile, fail error) {
	return dockerfile, fail
}

// containerToObjective
// It return
func containerToObjective(origin interface{}) (container Container, fail error) {
	return container, fail
}

// commandsToTask
// It return
func commandToTask(origin interface{}) (commands Commands, fail error) {
	command, ok := origin.(string)

	if !ok {
		return commands, fmt.Errorf("malformed '%s' task", command)
	}

	commands.commands = removeEmptyStrings(strings.Split(command, "\n"))

	return commands, fail
}

// interfaceToEnv
//
func interfaceToEnv(origin interface{}) (env []Env, fail error) {
	read, ok := origin.([]interface{})

	if !ok {
		return env, fmt.Errorf("error while reading env declaration")
	}

	env = make([]Env, len(read))

	for index, value := range read {
		declared, ok := value.(string)

		if !ok {
			return env, fmt.Errorf("malformed env declaration")
		}

		environments := strings.Split(declared, "=")

		if 2 != len(environments) {
			return env, fmt.Errorf("malformed env declaration, env -- %s -- has more than one declaration", read)
		}

		env[index].source = environments[0]
		env[index].destiny = environments[1]
	}

	return env, fail
}

// interfaceToEnvFile
//
func interfaceToEnvFile(origin interface{}) (env string, fail error) {
	return env, fail
}

// interfaceToContainer
//
func interfaceToContainer(origin interface{}) (container Container, fail error) {
	return container, fail
}

// complexTask
//
func complexTask(origin map[interface{}]interface{}) (task *Task, fail error) {
	task = &Task{}

	if nil != origin["container"] {
		task.container, fail = interfaceToContainer(origin["container"])
	}
	if nil != fail {
		return task, fmt.Errorf("%w;\nError while getting task container decalaration values", fail)
	}

	if nil != origin["env"] {
		task.env, fail = interfaceToEnv(origin["env"])
	}
	if nil != fail {
		return task, fmt.Errorf("%w;\nError while getting task env decalaration values", fail)
	}

	if nil != origin["env_file"] {
		task.env_file, fail = interfaceToEnvFile(origin["env_file"])
	}
	if nil != fail {
		return task, fmt.Errorf("%w;\nError while getting task env_file decalaration values", fail)
	}

	if nil != origin["commands"] {
		task.commands, fail = commandToTask(origin["commands"])
	}
	if nil != fail {
		return task, fmt.Errorf("%w;\nError while getting task commands decalaration values", fail)
	}

	return task, fail
}

// tasksToObjective
// It return
func tasksToObjective(origin map[interface{}]interface{}) (tasks map[string]*Task, fail error) {
	tasks = make(map[string]*Task)

	for key, value := range origin {
		name, ok := key.(string)

		if !ok {
			return tasks, fmt.Errorf("malformed '%s' task", name)
		}

		complex, ok := value.(map[interface{}]interface{})

		if ok {
			tasks[name], fail = complexTask(complex)
		} else {
			tasks[name] = &Task{}
			tasks[name].commands, fail = commandToTask(value)
			tasks[name].env = make([]Env, 0)
		}

		if nil != fail {
			return tasks, fmt.Errorf("%w;\nmalformed '%s' complex task", fail, name)
		}

		tasks[name].name = name
	}

	return tasks, fail
}

// objectiveToObjectives handles the reading process of Objective to Objetives,
// as the required values cannot be declared in the manifest, it won't be
// populated
// It return the Objectives
func objectiveToObjectives(origin map[interface{}]interface{}) (objectives Objectives, fail error) {
	objectives.objectives = make(map[string]*Objective)

	for key, value := range origin {
		name, ok := key.(string)

		if !ok {
			return objectives, fmt.Errorf("malformed '%s' objective", name)
		}

		objectives.objectives[name] = &Objective{}
		objectives.objectives[name].name = name
		read, ok := value.(map[interface{}]interface{})

		if !ok {
			return objectives, fmt.Errorf("malformed '%s' objective", name)
		}

		if nil != read["container"] {
			objectives.objectives[name].container, fail = containerToObjective(read["container"])
		}
		if nil != fail {
			return objectives, fmt.Errorf("%w;\nmalformed container declaration in '%s' objective", fail, name)
		}

		objectives.objectives[name].tasks, fail = tasksToObjective(read)

		if nil != fail {
			return objectives, fmt.Errorf("%w;\nmalformed tasks in '%s' objective", fail, name)
		}
	}

	return objectives, fail
}

// containerToProject
func containerToProject(origin map[interface{}]interface{}) (container Container, fail error) {
	if _, ok := origin["image"]; ok {
		container.dockerfile.base, fail = interfaceToImage(origin["image"])
	}
	if nil != fail {
		return container, fmt.Errorf("%w;\nimage presented and malformed", fail)
	}

	if _, ok := origin["dockerfile"]; ok {
		container.dockerfile, fail = interfaceToDockerfile(origin["dockerfile"])
	}
	if nil != fail {
		return container, fmt.Errorf("%w;\ndockerfile presented and malformed", fail)
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

	objectives, fail = objectiveToObjectives(read)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\nobjectives presented and malformed", fail)
	}

	objectives.required = make(map[string][]string)

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

// interactToProject
func interactToProject(read map[interface{}]interface{}) (interact Commands, fail error) {
	return interact, fail
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

	project.interact, fail = interactToProject(read)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed project name", fail)
	}

	return project, fail
}

// LexProject just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func LexProject(projectPath string) (project Project, fail error) {
	data, fail := checkConfig(projectPath)

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
