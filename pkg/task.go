package succubus

type Task struct {
	container Container
	env_file  string
	env       []Env
	rules     Rules
}
