package middleware

import (
	"time"

	"github.com/goyek/goyek/v3"
)

// ReportLongRun is a middleware which reports the task when it is long running.
func ReportLongRun(d time.Duration) func(next goyek.Runner) goyek.Runner {
	_ = "STUB: not implemented"
	return nil
}

// run
