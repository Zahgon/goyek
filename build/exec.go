package main

import (
	"github.com/goyek/goyek/v3"
)

// Exec runs the command in given directory.
// It calls a.Error[f] and returns false in case of any problems.
func Exec(a *goyek.A, workDir, name string, args ...string) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:gosec // it is a convenient function to run programs
