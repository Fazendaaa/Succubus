package succubus

type Container struct {
	dockerfile Dockerfile
	workdir    string
	env_file   string
}
