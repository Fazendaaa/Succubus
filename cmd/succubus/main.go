package main

import (
	"flag"
	"fmt"

	succubus "github.com/Fazendaaa/Succubus/pkg"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [ command ] [ values ]")
		flag.PrintDefaults()
	}

	project, fail := succubus.Load(".")

	if nil != fail {
		fmt.Errorf("%w", fail)
	}
}
