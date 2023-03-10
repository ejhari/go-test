package main

import (
	"fmt"

	"github.com/mitchellh/colorstring"
)

// Build-time flags.
var Version string
var ShortCommit string

func main() {
	fmt.Printf("%s\n", Version)
	fmt.Printf("%s\n", ShortCommit)

	msg := colorstring.Color("[green]Hello [red]World!\n")
	fmt.Printf("%s\n", msg)
}
