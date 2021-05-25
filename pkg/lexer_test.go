package succubus

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		src := Project{
			image: "some-user/some-image",
			base: Base{
				run: Task{
					env: []string{""},
					rules: []Command{
						{
							env:     []string{""},
							command: "",
						},
					},
				},
				add: Task{
					env: []string{""},
					rules: []Command{
						{
							env:     []string{""},
							command: "",
						},
					},
				},
				test: Task{
					env: []string{""},
					rules: []Command{
						{
							env:     []string{""},
							command: "",
						},
					},
				},
			},
			dev: Dev{
				doc: Task{
					env: []string{""},
					rules: []Command{
						{
							env:     []string{""},
							command: "",
						},
					},
				},
				anal: Task{
					env: []string{""},
					rules: []Command{
						{
							env:     []string{""},
							command: "",
						},
					},
				},
				linter: Task{
					env: []string{""},
					rules: []Command{
						{
							env:     []string{""},
							command: "",
						},
					},
				},
			},
			extended: nil,
		}
		value, fail := LexerProject(src)

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}

		if !reflect.DeepEqual(src, value) {
			t.Errorf("got mismatching configurations: %v,  %v", src, value)
		}
	})
}
