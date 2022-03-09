package succubus

import (
	"regexp"
	"unicode"
)

// hasUpper
// It returns whether container upper case or not
func hasUpper(s string) (ok bool) {
	for _, r := range s {
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			return true
		}
	}

	return false
}

// hasLower
// It returns whether container lower case or not
func hasLower(s string) (ok bool) {
	for _, r := range s {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			return true
		}
	}

	return false
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

// copyMap
// It returns
func copyMap(origin map[string][]string) (dest map[string][]string, fail error) {
	dest = make(map[string][]string)

	for key, value := range origin {
		dest[key] = value
	}

	return dest, fail
}

// keysOf
// It returns
func keysOf(origin map[string]*Task) (keys []string, fail error) {
	keys = make([]string, len(origin))
	index := 0

	for key, _ := range origin {
		keys[index] = key
		index += 1
	}

	return keys, fail
}

// isSemanticVersion
// It returns whether or not is Semantic Version
func isSemanticVersion(origin string) (ok bool) {
	condition, _ := regexp.Compile(`\d+.\d+.\d*`)

	return condition.MatchString(origin)
}

// isCalendarVersion
// It returns whether or not is Semantic Version
func isCalendarVersion(origin string) (ok bool) {
	condition, _ := regexp.Compile(`\d{2}.\d{2}.\d*`)

	return condition.MatchString(origin)
}
