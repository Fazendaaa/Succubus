package succubus

import "fmt"

// checkEnv looks for the environment variables is declared in the same was
// docker-compose does, besides checks the source one is declared to avoid any
// runtimes issues.
// It returns wether or not everything is ok
func checkEnv(origin []Env) (env []Env, fail error) {
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

// checkRules
func checkRules(origin Rules) (rules Rules, fail error) {
	return rules, fail
}

// checkTask check for a Task corectness
// It returns a fixed Task or an error
func checkTask(origin Task) (task Task, fail error) {
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

	task.rules, fail = checkRules(origin.rules)

	if nil != fail {
		return task, fmt.Errorf("%w;\nrules malformed in the task", fail)
	}

	return task, fail
}

// checkBase
// It returns the Base objective or the error associated with it
func checkBase(origin Base) (base Base, fail error) {
	base.add, fail = checkTask(origin.add)

	if nil != fail {
		return base, fmt.Errorf("%w;\nadd rule presented in base task", fail)
	}

	base.rm, fail = checkTask(origin.rm)

	if nil != fail {
		return base, fmt.Errorf("%w;\nrm rule presented in base task", fail)
	}

	base.test, fail = checkTask(origin.test)

	if nil != fail {
		return base, fmt.Errorf("%w;\ntest rule presented in base task", fail)
	}

	base.test, fail = checkTask(origin.test)

	if nil != fail {
		return base, fmt.Errorf("%w;\ntest rule presented in base task", fail)
	}

	return base, fail
}

// checkDev
// It returns the Dev objective or the error associated with it
func checkDev(origin Dev) (dev Dev, fail error) {
	dev.doc, fail = checkTask(origin.doc)

	if nil != fail {
		return dev, fmt.Errorf("%w;\ndoc rule presented in base task", fail)
	}

	dev.anal, fail = checkTask(origin.anal)

	if nil != fail {
		return dev, fmt.Errorf("%w;\nanal rule presented in base task", fail)
	}

	dev.linter, fail = checkTask(origin.linter)

	if nil != fail {
		return dev, fmt.Errorf("%w;\nlinter rule presented in base task", fail)
	}

	return dev, fail
}

// checkObjective
// It returns the Objective or the error associated with it
func checkObjective(origin interface{}) (objective map[interface{}]interface{}, fail error) {
	converted, ok := origin.(map[interface{}]interface{})

	if !ok {
		return objective, fmt.Errorf("invalid objective declaration")
	}

	for key, element := range converted {
		objective[key], fail = taskToObjective(element)

		if nil != fail {
			return objective, fmt.Errorf("%w;\ninvalid task -- '%s' -- in the objective", fail, key)
		}
	}

	return objective, fail
}

// checkExtended
// It returns the Extendeds objective or the error associated with it
func checkExtended(origin []Extended) (extended []Extended, fail error) {
	for userDefined := range origin {
		element, fail := checkObjective(userDefined)

		if nil != fail {
			return extended, fmt.Errorf("%w;\n", fail)
		}

		converted, ok := element.(map[string]Extended)

		if !ok {
			return extended, fmt.Errorf("")
		}

		extended = append(extended, converted)
	}

	return extended, fail
}

// checkObjectives
// It returns all Objectives or the error associated with it
func checkObjectives(origin Objectives) (objectives Objectives, fail error) {
	objectives.base, fail = checkBase(origin.base)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\nbase objective malformed", fail)
	}

	objectives.dev, fail = checkDev(origin.dev)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\ndev objective malformed", fail)
	}

	objectives.extended, fail = checkExtended(origin.extended)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\nextended objective malformed", fail)
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
		return project, fmt.Errorf("%w;\nmalformed container  in project", fail)
	}

	project.objectives, fail = checkObjectives(origin.objectives)

	if nil != fail {
		return project, fmt.Errorf("%w;\nmalformed objectives in project", fail)
	}

	return project, fail
}
