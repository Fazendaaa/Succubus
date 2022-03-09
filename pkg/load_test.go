package succubus

import (
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		value, fail := Load("../test/config/default/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("dockerfile", func(t *testing.T) {
		value, fail := Load("../test/config/dockerfile/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("env_file/objective", func(t *testing.T) {
		value, fail := Load("../test/config/env_file/objective.yaml")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("env_file/rule", func(t *testing.T) {
		value, fail := Load("../test/config/env_file/rule.yaml")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("env_file/task", func(t *testing.T) {
		value, fail := Load("../test/config/env_file/task.yaml")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("exec", func(t *testing.T) {
		value, fail := Load("../test/config/exec/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("extended", func(t *testing.T) {
		value, fail := Load("../test/config/extended/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("many_commands", func(t *testing.T) {
		value, fail := Load("../test/config/many_commands/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("missing", func(t *testing.T) {
		_, fail := Load("../test/config/missing/")

		if nil == fail {
			t.Errorf("Expected error and got:\n%v\n", fail)
		}
	})

	t.Run("named", func(t *testing.T) {
		value, fail := Load("../test/config/named/foo.yaml")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("project_name", func(t *testing.T) {
		value, fail := Load("../test/config/project_name/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("simple", func(t *testing.T) {
		value, fail := Load("../test/config/simple/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})

	t.Run("yml", func(t *testing.T) {
		value, fail := Load("../test/config/yml/")

		if nil != fail {
			t.Errorf("got:\n%v\nand the given error condition:\n%s", value, fail)
		}
	})
}
