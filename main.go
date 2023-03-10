package main

import (
	"fmt"

	"github.com/mitchellh/colorstring"
)

func main() {
	msg := colorstring.Color("[green]Hello [red]World!\n")
	fmt.Printf("%s\n", msg)
}
