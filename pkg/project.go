package succubus

type Project struct {
	filename  string
	version   string
	tag       string
	name      string
	interact  Commands
	init      Commands
	container Container // A Container is declared because it's easy to handle
	// reuse code instead of using the declared variables in the configuration
	// file like dockerfile or even image, as a container is formed in the by
	// the system
	objectives Objectives
	systems    Systems
}

func CreateProject(projectPath string) (fail error) {
	return fail
}
