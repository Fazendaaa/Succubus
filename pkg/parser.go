package succubus

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	image string
	base  struct {
		run  []string
		test []string
		add  []string
	}
	dev struct {
		doc    []string
		anal   []string
		linter []string
	}
	extended struct {
		any []string
	}
}

// checkConfig checks the following rule:
// 1. is there a 'succ.yaml'?
// 2. is there a 'succ.yml'?
// 3. is there a '<filenameHere>.<yml|yaml>'?
// It returns the file or an error
func checkConfig(succPath string) (data []byte, err error) {
	succ := Config{}
	data, fail := ioutil.ReadFile(succPath)

	if nil != fail {
		return nil, fail
	}

	fail = yaml.Unmarshal([]byte(data), &succ)

	if nil != fail {
		return nil, fail
	}

	dump, fail := yaml.Marshal(&succ)

	if nil != fail {
		return nil, fail
	}

	return data, fail
}

// ParseConfig just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func ParseConfig(succPath string) (valid bool, err error) {
	data, fail := checkConfig(succPath)

	if nil != fail {
		return false, fail
	}

	valid, fail := configLogic(data)

	return true, nil
}
