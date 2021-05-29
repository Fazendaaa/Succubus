package succubus

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		src := Project{
			tag:     "",
			name:    "",
			version: "",
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
				base: Base{
					run: Task{
						rules: Rules{
							commands: []Command{
								{
									env: []Env{
										{
											source:  "",
											destiny: "",
										},
									},
									command: "",
								},
							},
						},
					},
					add: Task{
						rules: Rules{
							commands: []Command{
								{
									env: []Env{
										{
											source:  "",
											destiny: "",
										},
									},
									command: "",
								},
							},
						},
					},
					rm: Task{
						rules: Rules{
							commands: []Command{
								{
									env: []Env{
										{
											source:  "",
											destiny: "",
										},
									},
									command: "",
								},
							},
						},
					},
				},
				dev: Dev{
					doc: Task{
						rules: Rules{
							commands: []Command{
								{
									env: []Env{
										{
											source:  "",
											destiny: "",
										},
									},
									command: "",
								},
							},
						},
					},
					anal: Task{
						rules: Rules{
							commands: []Command{
								{
									env: []Env{
										{
											source:  "",
											destiny: "",
										},
									},
									command: "",
								},
							},
						},
					},
					linter: Task{
						rules: Rules{
							commands: []Command{
								{
									env: []Env{
										{
											source:  "",
											destiny: "",
										},
									},
									command: "",
								},
							},
						},
					},
				},
				extended: nil,
			},
		}
		value, fail := ParseProject(src)

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}

		if !reflect.DeepEqual(src, value) {
			t.Errorf("got mismatching configurations:\n%v\n%v", src, value)
		}
	})
}
