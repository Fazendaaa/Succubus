package succubus

type Services struct {
}

type Network struct {
}

type Volumes struct {
}

type DockerCompose struct {
	path     string
	services []Services
	network  []Network
	volumes  []Volumes
}
