package succubus

// Load checks for:
// 1. whether or not the project as a valid image reference?
// 2. is there a Dockerfile listed in the config?
// 3. is there a Dockerfile presented in the project directory?
// It returns whether or not the config file is ok
func Load(succPath string) (succ Project, err error) {
	parsed, fail := ParseProject(succPath)

	if nil != fail {
		return succ, err
	}

	succ, fail = LexerProject(parsed)

	if nil != fail {
		return succ, fail
	}

	return succ, fail
}
