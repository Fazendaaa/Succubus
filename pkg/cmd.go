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
	Args     int
	Function func(...interface{})
	Flags    map[string][]CMD
}
