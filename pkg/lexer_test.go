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
					env:     []string{""},
					command: []string{""},
				},
				add: Task{
					env:     []string{""},
					command: []string{""},
				},
				test: Task{
					env:     []string{""},
					command: []string{""},
				},
			},
			dev: Dev{
				doc: Task{
					env:     []string{""},
					command: []string{""},
				},
				anal: Task{
					env:     []string{""},
					command: []string{""},
				},
				linter: Task{
					env:     []string{""},
					command: []string{""},
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
