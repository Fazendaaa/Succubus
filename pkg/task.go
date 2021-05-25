package succubus

type Task struct {
	env_file string
	env      []Env
	rules    []Command
}
