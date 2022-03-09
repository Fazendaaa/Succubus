package succubus

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		src := Project{
			tag:     "hom",
			name:    "foo",
			version: "21.07",
			interact: Commands{
				commands: []string{
					"foo",
				},
			},
			container: Container{
				env_file: "",
				dockerfile: Dockerfile{
					path:    "",
					context: "",
					args: []Env{
						{
							source:  "",
							destiny: "",
						},
					},
					multistages: []string{
						"",
					},
					base: Image{
						registry: "some-registry",
						owner:    "some-user",
						name:     "some-image",
						digest:   "some-digest",
					},
				},
			},
			objectives: Objectives{
				required: map[string][]string{
					"base": []string{
						"run",
						"test",
						"add",
						"rm",
					},
					"dev": []string{
						"anal",
						"linter",
						"doc",
					},
				},
				objectives: map[string]*Objective{
					"base": &Objective{
						name:      "base",
						container: Container{},
						tasks: map[string]*Task{
							"run": &Task{
								name:      "run",
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "",
												destiny: "",
											},
										},
										commands: []string{
											"exit 0",
										},
									},
								},
							},
							"test": &Task{
								name:      "test",
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "",
												destiny: "",
											},
										},
										commands: []string{
											"go test --verbose",
										},
									},
								},
							},
							"add": &Task{
								name:      "add",
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "ENV",
												destiny: "ENV",
											},
										},
										commands: []string{
											"go get $ARGV[0]",
										},
									},
								},
							},
							"rm": &Task{
								name:      "rm",
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "",
												destiny: "",
											},
										},
										commands: []string{
											"go mod tidy",
										},
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
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "",
												destiny: "",
											},
										},
										commands: []string{
											"gosec -fmt=sonarqujbe -out report.json ./..",
										},
									},
								},
							},
							"linter": &Task{
								name:      "linter",
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "",
												destiny: "",
											},
										},
										commands: []string{
											"golangci-lint run",
										},
									},
								},
							},
							"doc": &Task{
								name:      "doc",
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "",
												destiny: "",
											},
										},
										commands: []string{
											"go doc .",
										},
									},
								},
							},
						},
					},
					"extended": &Objective{
						name:      "extended",
						container: Container{},
						tasks: map[string]*Task{
							"coverage": &Task{
								name:      "coverage",
								container: Container{},
								env_file:  "",
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								commands: []Commands{
									Commands{
										env: []Env{
											{
												source:  "",
												destiny: "",
											},
										},
										commands: []string{
											"go test -cover ./...",
										},
									},
								},
							},
						},
					},
				},
			},
		}
		value, fail := ParseProject(src)

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}

		if !reflect.DeepEqual(src, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", src, value)
		}
	})
}
