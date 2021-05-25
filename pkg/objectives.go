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
