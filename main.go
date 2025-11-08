package main

import (
	"fmt"
	"os"

	"github.com/tnyeanderson/debate-timer/cmd"
)

var version = "development"
var commit string

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
