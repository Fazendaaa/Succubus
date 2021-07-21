package succubus

// Load checks for:
// 1. whether or not the project is has ambient issues
// 2. whether or not the project is has declaration issues
// 3. whether or not the project is has Project pattern issues
// It returns whether or not the Project is ok
func Load(projectPath string) (project Project, fail error) {
	lexed, fail := LexProject(projectPath)

	if nil != fail {
		return project, fail
	}

	parsed, fail := ParseProject(lexed)

	if nil != fail {
		return parsed, fail
	}

	return SemanticAnalysis(parsed)
}
