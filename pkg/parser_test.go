package succubus

import "testing"

func TestParse(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		value, err := ParseConfig("../test/default/")

		if nil != err {
			t.Errorf("got %v and the given error condition: %s", value, err)
		}
	})

	t.Run("extended", func(t *testing.T) {
		value, err := ParseConfig("../test/extended/")

		if nil != err {
			t.Errorf("got %v and the given error condition: %s", value, err)
		}

	})

	t.Run("missing", func(t *testing.T) {
		value, err := ParseConfig("../test/missing/")

		if nil == err {
			t.Errorf("got %v and the given error condition: %s", value, err)
		}

	})

	t.Run("named", func(t *testing.T) {
		value, err := ParseConfig("../test/named/foo.yaml")

		if nil != err {
			t.Errorf("got %v and the given error condition: %s", value, err)
		}

	})

	t.Run("yml", func(t *testing.T) {
		value, err := ParseConfig("../test/yml/")

		if nil != err {
			t.Errorf("got %v and the given error condition: %s", value, err)
		}

	})

	t.Run("many commands", func(t *testing.T) {
		value, err := ParseConfig("../test/manyCommands/")

		if nil != err {
			t.Errorf("got %v and the given error condition: %s", value, err)
		}

	})
}
