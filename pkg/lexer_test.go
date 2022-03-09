package succubus

import (
	"reflect"
	"testing"

	samael "github.com/Fazendaaa/Samael/pkg"
)

func TestLex(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		value, fail := samael.LexProject("succ", "../test/config/default/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("extended", func(t *testing.T) {
		value, fail := samael.LexProject("succ", "../test/config/extended/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("missing", func(t *testing.T) {
		value, fail := samael.LexProject("succ", "../test/config/missing/", projectFunc)

		if nil == fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("named", func(t *testing.T) {
		value, fail := samael.LexProject("succ", "../test/config/named/foo.yaml", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("yml", func(t *testing.T) {
		value, fail := samael.LexProject("succ", "../test/config/yml/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("many commands", func(t *testing.T) {
		value, fail := samael.LexProject("succ", "../test/config/many_commands/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("simple", func(t *testing.T) {
		filename := "../test/config/simple/"
		value := Project{
			filename: filename + "succ.yml",
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
		src, fail := samael.LexProject("succ", filename, projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}

		if !reflect.DeepEqual(src, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", src, value)
		}
	})
}
