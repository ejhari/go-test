package server

import (
	"fmt"
	"runtime"
)

// Build-time flags.
var Version string
var ShortCommit string

func init() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("ShortCommit: %s\n", ShortCommit)
	fmt.Printf("Go Version: %s\n", runtime.Version())
}
