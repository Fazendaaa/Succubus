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
