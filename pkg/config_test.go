package config

import "testing"

func TestParse(t *testing.T) {
	value, err := ParseConfig("../test/succ/default/succ.yaml")

	if false == value {
		t.Errorf("got %t want %t and the given error condition: %s", value, true, err)
	}
}
