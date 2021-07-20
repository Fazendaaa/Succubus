package succubus

import (
	"fmt"
)

// checkEnv looks for the environment variables is declared in the same was
// docker-compose does, besides checks the source one is declared to avoid any
// runtimes issues.
// It returns wether or not everything is ok
func checkEnv(origin []Env) (env []Env, fail error) {
	env = make([]Env, len(origin))

	return env, fail
}

// checkEnvFile
//
func checkEnvFile(origin string) (env_file string, fail error) {
	return env_file, fail
}

// checkDockerfile
//
func checkDockerfile(origin Dockerfile) (dockerfile Dockerfile, fail error) {
	if "" != origin.base.name && !isLower(origin.base.name) {
		return dockerfile, fmt.Errorf("wrong image format name: '%s'", origin.base.name)
	}

	return dockerfile, fail
}

// checkContainer
// It returns the project container or the error presented with it
func checkContainer(origin Container) (container Container, fail error) {
	container.dockerfile, fail = checkDockerfile(origin.dockerfile)

	if nil != fail {
		return container, fmt.Errorf("%w;\nwrong container configuration", fail)
	}

	container.env_file, fail = checkEnvFile(origin.env_file)

	if nil != fail {
		return container, fmt.Errorf("%w;\nwrong container env_file configuration", fail)
	}

	return Container(origin), fail
}

// checkCommands
// It returns
func checkCommands(origin Commands) (commands Commands, fail error) {
	commands.env = origin.env
	commands.commands = origin.commands

	return commands, fail
}

// checkTask check for a Task corectness
// It returns a fixed Task or an error
func checkTask(origin Task) (task *Task, fail error) {
	task = &Task{}
	task.name = origin.name
	task.container, fail = checkContainer(origin.container)

	if nil != fail {
		return task, fmt.Errorf("%w;\ncontainer malformed in the task", fail)
	}

	task.env_file, fail = checkEnvFile(origin.env_file)

	if nil != fail {
		return task, fmt.Errorf("%w;\nenv_file malformed in the task", fail)
	}

	task.env, fail = checkEnv(origin.env)

	if nil != fail {
		return task, fmt.Errorf("%w;\nenv malformed in the task", fail)
	}

	task.commands, fail = checkCommands(origin.commands)

	if nil != fail {
		return task, fmt.Errorf("%w;\ncommands malformed in the task", fail)
	}

	return task, fail
}

// checkObjective
// It returns the read Objective or the error associated with it
func checkObjective(origin Objective, values []string) (objective *Objective, fail error) {
	objective = &Objective{}
	objective.name = origin.name
	objective.container = origin.container
	objective.tasks = make(map[string]*Task)

	for _, task := range values {
		objective.tasks[task], fail = checkTask(*(origin.tasks[task]))

		if nil != fail {
			return objective, fmt.Errorf("%w;\n'%s' rule presented task", fail, task)
		}
	}

	return objective, fail
}

// checkObjectives
// It returns all Objectives or the error associated with it
func checkObjectives(origin Objectives) (objectives Objectives, fail error) {
	objectives.objectives = make(map[string]*Objective)
	objectives.required, fail = copyMap(origin.required)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\nError while copying required objectives", fail)
	}

	for key, value := range origin.required {
		objectives.objectives[key], fail = checkObjective(*(origin.objectives[key]), value)

		if nil != fail {
			return objectives, fmt.Errorf("%w;\nrequire '%s' objective is malformed", fail, key)
		}
	}

	for key, value := range origin.objectives {
		// Don't recheck the required ones
		if nil != origin.required[key] {
			continue
		}

		values, fail := keysOf(value.tasks)

		if nil != fail {
			return objectives, fmt.Errorf("%w;\nrequire '%s' objective is malformed", fail, key)
		}

		objectives.objectives[key], fail = checkObjective(*value, values)

		if nil != fail {
			return objectives, fmt.Errorf("%w;\nextended objective malformed", fail)
		}
	}

	return objectives, fail
}

// checkProjectName
// Returns the project name or why is not allowed
func checkProjectName(origin string) (name string, fail error) {
	if "" != origin && isUpper(origin) {
		return name, fmt.Errorf("project name -- '%s' -- has a upper case name which is not allowed", origin)
	}

	return origin, fail
}

// checkProjectTag
// Returns the project tag or why is not allowed
func checkProjectTag(origin string) (tag string, fail error) {
	if "" == origin {
		return tag, fail
	}

	return origin, fail
}

// checkProjectVersion
// Returns the project version or why is not allowed
func checkProjectVersion(origin string) (version string, fail error) {
	if "" == origin {
		return version, fail
	}

	return origin, fail
}

// ParseProject goes trough the "logic" behind a Succubus' configuration file.
// It's important tho to point out that this or any other function presented
// in this package checks whether or not the given command is "plausible" just
// if all the key have a string as a value; if a image tag points to a
// non-existing image or if a command is wrongfull typed or even not exists,
// this function doesn't perform this kind of pre-processing due to the runtime
// nature of those tasks.
// It returns the processed config file
func ParseProject(origin Project) (project Project, fail error) {
	project.name, fail = checkProjectName(origin.name)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed name in project", fail)
	}

	project.tag, fail = checkProjectTag(origin.tag)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed tag in project", fail)
	}

	project.version, fail = checkProjectVersion(origin.version)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed version in project", fail)
	}

	project.container, fail = checkContainer(origin.container)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed container in project", fail)
	}

	project.objectives, fail = checkObjectives(origin.objectives)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed objectives in project", fail)
	}

	return project, fail
}
