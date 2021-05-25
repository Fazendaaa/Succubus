package succubus

type Command struct {
	env_file string
	env      []Env
	command  string
}
