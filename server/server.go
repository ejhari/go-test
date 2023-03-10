package server

import "fmt"

// Build-time flags.
var Version string
var ShortCommit string

func init() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("ShortCommit: %s\n", ShortCommit)
}
