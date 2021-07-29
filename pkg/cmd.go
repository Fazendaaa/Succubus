package succubus

type Usage struct {
	Short string
	Long  string
}

type Exec struct {
	Name     string
	Usage    Usage
	Function func(...interface{})
	Cmd      *CMD
}

type CMD struct {
	Name  string
	Usage Usage
	Flags map[string][]Exec
}
