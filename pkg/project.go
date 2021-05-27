package succubus

type Project struct {
	vertion    string
	name       string
	image      Image
	dockerfile Dockerfile
	env_file   string
	context    string
	objectives Objectives
}
