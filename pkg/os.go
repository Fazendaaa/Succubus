package succubus

import (
	"os"
	"path/filepath"
)

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
