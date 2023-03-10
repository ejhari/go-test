package main

import (
	"fmt"

	"github.com/mitchellh/colorstring"

	_ "gitub.com/ejhari/go-test/server"
)

func main() {
	msg := colorstring.Color("[green]Hello [red]World!\n")
	fmt.Printf("%s\n", msg)

	msg = colorstring.Color("[blue]Hello [yellow]World!\n")
	fmt.Printf("%s\n", msg)

	msg = colorstring.Color("[purple]Hello [cyan]World!\n")
	fmt.Printf("%s\n", msg)
}
