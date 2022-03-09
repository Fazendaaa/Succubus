package succubus

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
)

// Load checks for:
// 1. whether or not the project is has ambient issues
// 2. whether or not the project is has declaration issues
// 3. whether or not the project is has Project pattern issues
// It returns whether or not the Project is ok
func Load(projectPath string) (project Project, fail error) {
	lexed, fail := samael.LexProject("succ", projectPath, projectFunc)

	if nil != fail {
		return project, fail
	}

	casted, ok := lexed.(Project)

	if !ok {
		return project, fmt.Errorf("cannot normalize project")
	}

	parsed, fail := ParseProject(casted)

	if nil != fail {
		return parsed, fail
	}

	return SemanticAnalysis(parsed)
}
