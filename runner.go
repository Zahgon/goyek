package goyek

import (
	"context"
	"io"
)

// Task runner types.
type (
	// Runner represents a task runner function.
	Runner func(Input) Result

	// Input received by the task runner.
	Input struct {
		Context  context.Context
		TaskName string
		Parallel bool
		Output   io.Writer
		Logger   Logger
	}

	// Result of a task run.
	Result struct {
		Status     Status
		PanicValue interface{}
		PanicStack []byte
	}

	// Middleware represents a task runner interceptor.
	Middleware func(Runner) Runner
)

// NewRunner returns a task runner used by Flow.
//
// It can be useful for testing and debugging
// a task action or middleware.
//
// The following defaults are set for Input
// (take notice that they are different than Flow defaults):
//
//	Context = context.Background()
//	Output = io.Discard
//	Logger = FmtLogger{}
//
// It can be also used as a building block for a custom
// workflow runner if you are missing any functionalities
// provided by Flow (like concurrent dependencies execution).
func NewRunner(action func(a *A)) Runner { _ = "STUB: not implemented"; return *new(Runner) }

type taskRunner struct {
	action func(a *A)
}

// run executes the action in a separate goroutine to enable
// interuption using runtime.Goexit().
func (r taskRunner) run(in Input) Result { _ = "STUB: not implemented"; return *new(Result) }
