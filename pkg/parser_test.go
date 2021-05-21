package succubus

import "testing"

func TestParse(t *testing.T) {
	t.Run("named", func(t *testing.T) {
		value, err := ParseConfig("../test/succ/default/")

		if false == value {
			t.Errorf("got %t want %t and the given error condition: %s", value, true, err)
		}
	})
}
