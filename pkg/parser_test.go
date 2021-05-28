package succubus

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		src := Project{
			image: Image{
				registry: "",
				owner:    "some-user",
				name:     "some-image",
				digest:   "",
			},
			objectives: Objectives{
				base: Base{
					run: Task{
						env: []Env{
							{
								source:  "",
								destiny: "",
							},
						},
						rules: []Command{
							{
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								command: "",
							},
						},
					},
					add: Task{
						env: []Env{
							{
								source:  "",
								destiny: "",
							},
						},
						rules: []Command{
							{
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								command: "",
							},
						},
					},
					rm: Task{
						env: []Env{
							{
								source:  "",
								destiny: "",
							},
						},
						rules: []Command{
							{
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								command: "",
							},
						},
					},
				},
				dev: Dev{
					doc: Task{
						env: []Env{
							{
								source:  "",
								destiny: "",
							},
						},
						rules: []Command{
							{
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								command: "",
							},
						},
					},
					anal: Task{
						env: []Env{
							{
								source:  "",
								destiny: "",
							},
						},
						rules: []Command{
							{
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								command: "",
							},
						},
					},
					linter: Task{
						env: []Env{
							{
								source:  "",
								destiny: "",
							},
						},
						rules: []Command{
							{
								env: []Env{
									{
										source:  "",
										destiny: "",
									},
								},
								command: "",
							},
						},
					},
				},
				extended: nil,
			},
		}
		value, fail := ParseProject(src)

		if nil != fail {
			t.Errorf("got %v and the given error condition is: %s", value, fail)
		}

		if !reflect.DeepEqual(src, value) {
			t.Errorf("got mismatching configurations:\n%v\n%v", src, value)
		}
	})
}
