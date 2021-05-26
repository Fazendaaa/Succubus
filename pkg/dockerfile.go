package succubus

type Dockerfile struct {
	path        string
	context     string
	base        Image
	multistages []string
	labels      []string
}
