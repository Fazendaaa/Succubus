package succubus

// Load checks for:
// 1. whether or not the project as a valid image reference?
// 2. is there a Dockerfile listed in the config?
// 3. is there a Dockerfile presented in the project directory?
// It returns whether or not the config file is ok
func Load(succPath string) (succ Config, err error) {
	parsed, fail := ParseConfig(succPath)

	if nil != fail {

	}

	succ, fail = LexerConfig(parsed)

	if nil != fail {
		return succ, fail
	}

	return succ, fail
}
