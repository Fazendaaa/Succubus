package succubus

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Command struct {
	env     []string
	command []string
}

type Extended struct {
	Command
}

type Config struct {
	image string
	base  struct {
		run  []Command
		test []Command
		add  []Command
	}
	dev struct {
		doc    []Command
		anal   []Command
		linter []Command
	}
	extended struct {
		key []Extended
	}
}

// exists checks whether or not a given file exists, this soluction is based in
// [this](https://stackoverflow.com/a/12527546/7092954) post
// It returns whether or not the file exists
func exists(name string) bool {
	if _, fail := os.Stat(name); fail != nil {
		if os.IsNotExist(fail) {
			return false
		}
	}
	return true
}

// readFileOrDir is based in [this](https://stackoverflow.com/a/8824952/7092954)
// code and reads the default YAML file or the given one:
// 1. is there a 'succ.yaml'?
// 2. is there a 'succ.yml'?
// 3. is there a '<filenameHere>.<yml|yaml>'?
// Returns the filename
func readFileOrDir(succPath string) (file string, fail error) {
	read, fail := os.Stat(succPath)

	if nil != fail {
		return file, fail
	}

	switch mode := read.Mode(); {
	case mode.IsDir():
		yaml := filepath.Join(succPath, "succ.yaml")

		if exists(yaml) {
			return yaml, fail
		}

		yml := filepath.Join(succPath, "succ.yml")

		if exists(yml) {
			return yml, fail
		}

		return file, fail
	case mode.IsRegular():
		return succPath, fail
	}

	return file, fail
}

// checkConfig looks for the config file
// It returns the file or an error
func checkConfig(succPath string) (data []byte, fail error) {
	file, fail := readFileOrDir(succPath)

	if nil != fail {
		return nil, fail
	}

	read, fail := ioutil.ReadFile(file)

	if nil != fail {
		return nil, fail
	}

	return read, fail
}

// fromFileToConfig as many use cases might not need Succubus full power, but to
// handle code base is easy always to handle the "worst case scenario", this
// function does that, translate the user specific configuration to a more
// generic one.
// It returns the config file as full blown Succubus standard
func fromFileToConfig(read map[interface{}]interface{}) (succ Config, fail error) {

}

// ParseConfig just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func ParseConfig(succPath string) (succ Config, fail error) {
	read, fail := checkConfig(succPath)

	if nil != fail {
		return succ, fail
	}

	fromFile := make(map[interface{}]interface{})
	fail = yaml.Unmarshal([]byte(read), &fromFile)

	if nil != fail {
		return succ, fail
	}

	return fromFileToConfig(fromFile)
}
