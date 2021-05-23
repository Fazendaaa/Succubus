package succubus

// checkEnv looks for the environement variables is declared in the same was
// docker-compose does, besides checks the source one is declared to avoid any
// runtimes issues.
// It returns wether or not everything is ok
func checkEnv(env string) (ok bool) {
	return false
}

//
func checkTask(task Task) (ok bool) {
	return false
}

func checkBase(base Base) (ok bool) {
	if !checkTask(base.add) {
		return false
	}

	return true
}

func checkForDockerfile() {

}

func checkImage(image string) (ok bool) {
	if "" == image {
		return checkForDockerfile()
	}

	return true
}

// LexerConfig goes trough the "logic" behind a Succubus' configuration file.
// It's important tho to point out that this or any other function presented
// in this package checks whether or not the given command is "plausible" just
// if all the key have a string as a value; if a image tag points to a
// non-existing image or if a command is wrongfull typed or even not exists,
// this function doesn't perform this kind of pre-processing due to the runtime
// nature of those tasks.
// It returns the processed config file
func LexerConfig(parsed Config) (succ Config, err error) {

	if !checkBase(parsed.base) {

	}

	return succ, nil
}
