package succubus

type Commands struct {
	name     string
	env_file string
	env      []Env
	commands []string
}
