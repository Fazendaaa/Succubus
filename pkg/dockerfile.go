package succubus

type Dockerfile struct {
	path        string
	context     string
	base        Image
	args        []Env
	multistages []string
	labels      []string
}
