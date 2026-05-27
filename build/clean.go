package main

import (
	"github.com/goyek/goyek/v3"
)

var _ = goyek.Define(goyek.Task{
	Name:  "clean",
	Usage: "remove files created during build pipeline",
	Action: func(a *goyek.A) {
		remove(a, "coverage.out")
		remove(a, "coverage.html")
	},
})

func remove(a *goyek.A, path string) { _ = "STUB: not implemented"; return }
