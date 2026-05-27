package middleware

import (
	"github.com/goyek/goyek/v3"
)

// ReportFlow is a middleware which reports the flow execution status.
//
// The format is based on the reports provided by the Go test runner.
func ReportFlow(next goyek.Executor) goyek.Executor {
	_ = "STUB: not implemented"
	return *new(goyek.Executor)
}
