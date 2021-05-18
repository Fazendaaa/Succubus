// Package config handles all Succubus config file tasks
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	base struct {
		run  []string
		test []string
		add  []string
		doc  []string
	}
	dev struct {
		anal   []string
		linter []string
	}
}

// ParseConfig just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func ParseConfig(succPath string) (valid bool, err error) {
	data, fail := ioutil.ReadFile(succPath)

	if nil != fail {
		return false, fail
	}

	succ := Config{}
	fail = yaml.Unmarshal([]byte(data), &succ)

	if nil != fail {
		return false, fail
	}

	return true, nil
}
