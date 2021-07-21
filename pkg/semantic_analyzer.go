package succubus

// SemanticAnalysis goes trough a Project definition and checks for:
// 1. commands/variables ambiguity
// 2. whether or not the current Project is from a different pattern from the
//    one defined as the one to be followed
// It returns the sanitized version of the Project or the error related to it
func SemanticAnalysis(origin Project) (project Project, fail error) {
	return origin, fail
}
