package succubus

import "testing"

func TestParse(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		value, fail := ParseProject("../test/default/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("extended", func(t *testing.T) {
		value, fail := ParseProject("../test/extended/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("missing", func(t *testing.T) {
		value, fail := ParseProject("../test/missing/")

		if nil == fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("named", func(t *testing.T) {
		value, fail := ParseProject("../test/named/foo.yaml")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("yml", func(t *testing.T) {
		value, fail := ParseProject("../test/yml/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("many commands", func(t *testing.T) {
		value, fail := ParseProject("../test/manyCommands/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})

	t.Run("nested env", func(t *testing.T) {
		value, fail := ParseProject("../test/nestedEnv/")

		if nil != fail {
			t.Errorf("got %v and the given error condition: %s", value, fail)
		}
	})
}
