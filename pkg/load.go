package succubus

// Load checks for:
// 1. whether or not the project as a valid image reference?
// 2. is there a Dockerfile listed in the config?
// 3. is there a Dockerfile presented in the project directory?
// It returns whether or not the config file is ok
func Load(succPath string) (valid bool, err error) {
	config, fail := ParseConfig(succPath)

	if nil != fail {

	}

	commands, fail := LexerConfig(config)

	if nil != fail {

	}
}
