package succubus

import (
	"reflect"
	"testing"
)

func TestSemanticAnalysis(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		src := Project{
			container: Container{
				dockerfile: Dockerfile{
					base: Image{
						name: "golang",
					},
				},
			},
			objectives: Objectives{
				required: map[string][]string{},
				objectives: map[string]*Objective{
					"base": &Objective{
						name:      "base",
						container: Container{},
						tasks: map[string]*Task{
							"run": &Task{
								name:      "run",
								env_file:  "",
								env:       make([]Env, 0),
								container: Container{},
								commands: Commands{
									commands: []string{
										"exit 0",
									},
								},
							},
							"test": &Task{
								name:      "test",
								env_file:  "",
								env:       make([]Env, 0),
								container: Container{},
								commands: Commands{
									commands: []string{
										"go test --verbose",
									},
								},
							},
							"add": &Task{
								name:      "add",
								env_file:  "",
								env:       make([]Env, 0),
								container: Container{},
								commands: Commands{
									commands: []string{
										"go get $$ARGV[0]",
									},
								},
							},
							"rm": &Task{
								name:      "rm",
								env_file:  "",
								env:       make([]Env, 0),
								container: Container{},
								commands: Commands{
									commands: []string{
										"go mod tidy",
									},
								},
							},
						},
					},
					"dev": &Objective{
						name:      "dev",
						container: Container{},
						tasks: map[string]*Task{
							"anal": &Task{
								name:      "anal",
								env_file:  "",
								env:       make([]Env, 0),
								container: Container{},
								commands: Commands{
									commands: []string{
										"gosec -fmt=sonarqujbe -out report.json ./..",
									},
								},
							},
							"linter": &Task{
								name:      "linter",
								env_file:  "",
								env:       make([]Env, 0),
								container: Container{},
								commands: Commands{
									commands: []string{
										"golangci-lint run",
									},
								},
							},
							"doc": &Task{
								name:      "doc",
								env_file:  "",
								env:       make([]Env, 0),
								container: Container{},
								commands: Commands{
									commands: []string{
										"go doc .",
									},
								},
							},
						},
					},
				},
			},
		}
		value, fail := SemanticAnalysis(src)

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}

		if !reflect.DeepEqual(src, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", src, value)
		}
	})
}
