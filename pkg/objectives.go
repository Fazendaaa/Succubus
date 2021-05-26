package succubus

type Base struct {
	run  Task
	test Task
	add  Task
	rm   Task
}

type Dev struct {
	doc    Task
	anal   Task
	linter Task
}

type Objectives struct {
	base     Base
	dev      Dev
	extended interface{}
}
