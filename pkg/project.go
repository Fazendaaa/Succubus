package succubus

type Project struct {
	version   string
	tag       string
	name      string
	container Container // A Container is declared because it's easy to handle
	// reuse code instead of using the declared variables in the configuration
	// file like dockerfile or even image, as a container is formed in the by
	// the system
	objectives Objectives
	system     DockerCompose
}
