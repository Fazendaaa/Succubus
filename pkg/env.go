package succubus

import (
	"fmt"
	"strings"
)

type Env struct {
	source  string
	destiny string
}

func (env *Env) parseSource(source string) (ok bool) {
	return false
}

func (env *Env) parseDestiny(source []string) (ok bool) {
	return false
}

//
func CreateEnv(origin string) (env Env, fail error) {
	splited := strings.Split(origin, "=")

	if !env.parseSource(splited[0]) {
		return env, fmt.Errorf("Error while parsing source environment variable")
	}

	if !env.parseDestiny(splited[1 : len(splited)-1]) {
		return env, fmt.Errorf("Error while parsing source environment variable")
	}

	return env, fail
}
