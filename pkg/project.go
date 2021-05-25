package succubus

type Project struct {
	name       string
	image      string
	dockerfile string
	env_file   string
	context    string
	base       Base
	dev        Dev
	extended   interface{}
}
