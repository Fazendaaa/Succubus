package succubus

type Usage struct {
	short string
	long  string
}

type Param struct {
	name        string
	description string
}

type CMD struct {
	name  string
	usage Usage
	args  []Param
}
