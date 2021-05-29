package succubus

type Base struct {
	container Container
	run       Task
	test      Task
	add       Task
	rm        Task
}

type Dev struct {
	container Container
	doc       Task
	anal      Task
	linter    Task
}

type Extended struct {
	name      string
	container Container
	tasks     map[string]Task
}

type Objectives struct {
	base     Base
	dev      Dev
	extended map[string]Extended
}
