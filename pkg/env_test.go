package succubus

import "testing"

func TestEnv(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		value, fail := CreateEnv("foo=bar")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("empty source", func(t *testing.T) {
		value, fail := CreateEnv("foo=")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("quotes", func(t *testing.T) {
		value, fail := CreateEnv("foo=\"bar\"")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("mismatched quotes in source - first", func(t *testing.T) {
		value, fail := CreateEnv("foo=\"bar")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("mismatched quotes in source - second", func(t *testing.T) {
		value, fail := CreateEnv("foo=bar\"")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("mismatched quotes in destiny - first", func(t *testing.T) {
		value, fail := CreateEnv("\"foo=\"bar\"")

		if nil == fail {
			t.Errorf("got %v in '\"foo', which isn't valid", value)
		}
	})

	t.Run("mismatched quotes in destiny - second", func(t *testing.T) {
		value, fail := CreateEnv("foo\"=\"bar\"")

		if nil == fail {
			t.Errorf("got %v in 'foo\"', which isn't valid", value)
		}
	})

	t.Run("equals in source", func(t *testing.T) {
		value, fail := CreateEnv("foo=\"bar=baz\"")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("equals in source", func(t *testing.T) {
		value, fail := CreateEnv("foo=\"bar=baz\"")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("external source", func(t *testing.T) {
		value, fail := CreateEnv("foo=${bar}")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("missconfigured external source -- missing dolar sign", func(t *testing.T) {
		value, fail := CreateEnv("foo={bar}")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("missconfigured external source -- missing close brackets", func(t *testing.T) {
		value, fail := CreateEnv("foo=${bar")

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}
	})

	t.Run("destiny not valid -- added '?' to it", func(t *testing.T) {
		value, fail := CreateEnv("fo?o=${bar}")

		if nil == fail {
			t.Errorf("got '%v' as destiny variable and this should not be valid", value)
		}
	})

	t.Run("empty destiny", func(t *testing.T) {
		value, fail := CreateEnv("=bar")

		if nil == fail {
			t.Errorf("got %v in env which isn't allowed", value)
		}
	})

	t.Run("no equals", func(t *testing.T) {
		value, fail := CreateEnv("foobar")

		if nil == fail {
			t.Errorf("got %v in env which isn't allowed", value)
		}
	})
}
