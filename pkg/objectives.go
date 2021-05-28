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
	container Container
	tasks     []Task
}

type Objectives struct {
	base     Base
	dev      Dev
	extended []Extended
}
