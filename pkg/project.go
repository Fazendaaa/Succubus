package succubus

type Project struct {
	name       string
	image      Image
	dockerfile Dockerfile
	env_file   string
	context    string
	base       Base
	dev        Dev
	extended   interface{}
}
