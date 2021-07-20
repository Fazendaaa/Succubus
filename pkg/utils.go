package succubus

import "unicode"

// isUpper [this](https://stackoverflow.com/a/59293875/7092954)
// It returns whether or all characters are upper case
func isUpper(s string) (ok bool) {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// isUpper [this](https://stackoverflow.com/a/59293875/7092954)
// It returns whether or all characters are lower case
func isLower(s string) (ok bool) {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// removeEmptyStrings - Use this to remove empty string values inside an array.
// https://gist.github.com/johnpili/84c3064d30a9b041c87e43ba4bcb63a2
// This happens when allocation is bigger and empty
func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
