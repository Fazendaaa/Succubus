package succubus

import (
	"io/ioutil"
	"os"
	"path/filepath"

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

// exists checks whether or not a given file exists, this soluction is based in
// [this](https://stackoverflow.com/a/12527546/7092954) post
// It returns whether or not the file exists
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
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
func readFileOrDir(succPath string) (file string, err error) {
	read, fail := os.Stat(succPath)

	if nil != fail {
		return file, err
	}

	switch mode := read.Mode(); {
	case mode.IsDir():
		yaml := filepath.Join(succPath, "succ.yaml")

		if exists(yaml) {
			return yaml, err
		}

		yml := filepath.Join(succPath, "succ.yml")

		if exists(yml) {
			return yml, err
		}

		return file, err
	case mode.IsRegular():
		return succPath, err
	}

	return file, err
}

// checkConfig lookd for the config file
// It returns the file or an error
func checkConfig(succPath string) (succ Config, err error) {
	succ = Config{}
	file, fail := readFileOrDir(succPath)

	if nil != fail {
		return succ, fail
	}

	read, fail := ioutil.ReadFile(file)

	if nil != fail {
		return succ, fail
	}

	return succ, yaml.Unmarshal([]byte(read), &succ)
}

// ParseConfig just reads the given Succubus config file and checks it whether
// or not it's a valid one.
// It returns whether or not the config file is valid and any error encountered.
func ParseConfig(succPath string) (succ Config, err error) {
	succ, fail := checkConfig(succPath)

	if nil != fail {
		return succ, fail
	}

	return succ, nil
}
