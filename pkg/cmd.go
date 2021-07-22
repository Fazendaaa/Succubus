package succubus

type Usage struct {
	short string
	long  string
}

type Exec struct {
	name     string
	usage    Usage
	function func(...interface{})
	cmd      *CMD
}

type CMD struct {
	name  string
	usage Usage
	flags map[string][]Exec
}
