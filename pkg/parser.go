package succubus

import "fmt"

// checkEnv looks for the environment variables is declared in the same was
// docker-compose does, besides checks the source one is declared to avoid any
// runtimes issues.
// It returns wether or not everything is ok
func checkEnv(env string) (ok bool, fail error) {
	return false, fail
}

// checkTask check for a Task corectness
// It returns a fixed Task or an error
func checkTask(origin Task) (task Task, fail error) {
	if 0 == len(origin.rules) {
		return task, fmt.Errorf("")
	}

	return task, fail
}

// checkBase
func checkBase(origin Base) (base Base, fail error) {
	// if !checkTask(origin.add) {
	// 	return base, fail
	// }

	return base, fail
}

// checkDev
func checkDev(origin Dev) (dev Dev, fail error) {
	return dev, fail
}

// checkExtended
func checkExtended(origin interface{}) (extended []Extended, fail error) {
	return extended, fail
}

// checkContainer
func checkContainer(origin Container) (container Container, fail error) {
	if "" == origin.dockerfile.base.name {
		return container, fail
	}

	if !IsLower(origin.dockerfile.base.name) {
		return container, fmt.Errorf("wrong image format name: '%s'", origin.dockerfile.base.name)
	}

	return Container(origin), fail
}

// checkForDockerfile
func checkForDockerfile(dockerfile Dockerfile) (ok bool, fail error) {
	if "" == dockerfile.path {
		return true, fail
	}

	return true, fail
}

// checkContext
func checkContext(origin string) (context string, fail error) {
	return context, fail
}

// checkEnvFile
func checkEnvFile(origin string) (env_file string, fail error) {
	return env_file, fail
}

// checkObjetives
func checkObjetives(origin Objectives) (objectives Objectives, fail error) {
	base, fail := checkBase(origin.base)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\n'base' objective malformed", fail)
	}

	objectives.base = base
	dev, fail := checkDev(origin.dev)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\n'dev' objective malformed", fail)
	}

	objectives.dev = dev

	extended, fail := checkExtended(origin.extended)

	if nil != fail {
		return objectives, fmt.Errorf("%w;\n'extended' objective malformed", fail)
	}

	objectives.extended = extended

	return objectives, fail
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
	// container, fail := checkContainer(origin.container)

	if nil != fail {
		return project, fmt.Errorf("%w;\n'image' malformed in project", fail)
	}

	// objective, fail := checkObjetives(origin.objectives)

	if nil != fail {
		return project, fmt.Errorf("%w;\n'objectives' malformed in project", fail)
	}

	return project, fail
}
