package succubus

// Stores the description related to given command
type Usage struct {
	Short string
	Long  string
}

// Stores the given command
type CMD struct {
	Name     string
	Usage    Usage
	Params   []string // Just the parameters name
	Function func(...interface{})
}
