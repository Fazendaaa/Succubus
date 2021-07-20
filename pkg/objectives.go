package succubus

type Objective struct {
	name      string
	container Container
	tasks     map[string]*Task
}

type Objectives struct {
	required   map[string][]string
	objectives map[string]*Objective
}
