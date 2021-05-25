package succubus

import "testing"

func TestLex(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		value, fail := LexProject("../test/default/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("extended", func(t *testing.T) {
		value, fail := LexProject("../test/extended/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("missing", func(t *testing.T) {
		value, fail := LexProject("../test/missing/")

		if nil == fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("named", func(t *testing.T) {
		value, fail := LexProject("../test/named/foo.yaml")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("yml", func(t *testing.T) {
		value, fail := LexProject("../test/yml/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("many commands", func(t *testing.T) {
		value, fail := LexProject("../test/manyCommands/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("nested env", func(t *testing.T) {
		value, fail := LexProject("../test/nestedEnv/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})
}
