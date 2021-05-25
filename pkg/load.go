package succubus

// Load checks for:
// 1. whether or not the project as a valid image reference?
// 2. is there a Dockerfile listed in the config?
// 3. is there a Dockerfile presented in the project directory?
// It returns whether or not the config file is ok
func Load(projectPath string) (project Project, fail error) {
	lexed, fail := LexProject(projectPath)

	if nil != fail {
		return project, fail
	}

	parsed, fail := ParseProject(lexed)

	if nil != fail {
		return parsed, fail
	}

	return project, fail
}
