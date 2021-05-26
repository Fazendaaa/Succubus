package succubus

import (
	"fmt"
	"regexp"
	"strings"
)

type Env struct {
	source  string
	destiny string
}

// envValidator based in [this](https://stackoverflow.com/a/17564142/7092954)
// question.
// It returns whether or not the environment variable is a valid one
func envValidator(source string) (ok bool) {
	return regexp.MustCompilePOSIX("^[a-zA-Z_$][a-zA-Z_$0-9]*$").MatchString(source)
}

// parseDestiny destiny environment variables are those which are given to the
// project and probably will be used in it.
// It returns whether or not the operation has been successfull.
func (env *Env) parseDestiny(destiny string) (ok bool, fail error) {
	if !envValidator(destiny) {
		return false, fmt.Errorf("given '%s' is not a valid one", destiny)
	}

	env.source = destiny

	return true, fail
}

// parseSource source environment variables are those declared by the user to be
// injected as another -- destiny -- env var into the current project, being in
// a Command, Rule, Task, Objective or Project level. This envs can be plain
// text or even another variable.
// It returns whether or not the operation has been successfull.
func (env *Env) parseSource(source []string) (ok bool, fail error) {
	env.source = strings.Join(source, "")

	return true, fail
}

// CreateEnv Goes trough the process of checking whether or not a environment
// variable is properly declared -- even tho doesn't check whether or not is
// "valid" in a way to check whether the source value exists or the destiny
// value is required or not.
// It returns a Environment Variable object and whether or not some error occurred
func CreateEnv(origin string) (env Env, fail error) {
	splitted := strings.Split(origin, "=")

	if 2 > len(splitted) {
		return env, fmt.Errorf("given '%s' environment variable isn't a valid one -- misssing '=' operator", origin)
	}

	if ok, fail := env.parseDestiny(splitted[0]); !ok {
		return env, fmt.Errorf("%w;\nerror while destiny source environment variable", fail)
	}

	if ok, fail := env.parseSource(splitted[1 : len(splitted)-1]); !ok {
		return env, fmt.Errorf("%w;\nerror while parsing source environment variable", fail)
	}

	return env, fail
}
