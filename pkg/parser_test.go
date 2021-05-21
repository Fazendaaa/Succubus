package succubus

import "testing"

func TestParse(t *testing.T) {
	t.Run("named", func(t *testing.T) {
		value, err := ParseConfig("../test/default/")

		if nil != err {
			t.Errorf("got %v and the given error condition: %s", value, err)
		}
	})
}
