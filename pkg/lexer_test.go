package succubus

import "testing"

func TestLexer(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		value, err := LexerConfig(Config{
			image: "some-user/some-image",
			base: {
				run:  "",
				add:  "",
				test: "",
			},
			dev: {
				doc:    "",
				anal:   "",
				linter: "",
			},
			extended: []Extended{
				key: "",
			},
		})

		if nil != err {
			t.Errorf("got %v and the given error condition is: %s", value, err)
		}
	})
}
