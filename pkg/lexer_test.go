package succubus

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		src := Config{
			image: "some-user/some-image",
			base: Base{
				run: Task{
					env: []string{""},
					rules: []Command{
						Command{
							env:     []string{""},
							command: "",
						},
					},
				},
				add: Task{
					env: []string{""},
					rules: []Command{
						Command{
							env:     []string{""},
							command: "",
						},
					},
				},
				test: Task{
					env: []string{""},
					rules: []Command{
						Command{
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
						Command{
							env:     []string{""},
							command: "",
						},
					},
				},
				anal: Task{
					env: []string{""},
					rules: []Command{
						Command{
							env:     []string{""},
							command: "",
						},
					},
				},
				linter: Task{
					env: []string{""},
					rules: []Command{
						Command{
							env:     []string{""},
							command: "",
						},
					},
				},
			},
			extended: nil,
		}
		value, err := LexerConfig(src)

		if nil != err {
			t.Errorf("got %v and the given error condition is: %s", value, err)
		}

		if !reflect.DeepEqual(src, value) {
			t.Errorf("got mismatching configurations: %v,  %v", src, value)
		}
	})
}
