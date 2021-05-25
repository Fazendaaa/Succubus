package succubus

type Project struct {
	image      string
	dockerfile string
	env_file   string
	context    string
	base       Base
	dev        Dev
	extended   interface{}
}
