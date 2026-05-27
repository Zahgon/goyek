package goyek

import (
	"context"
	"errors"
	"io"
)

// Flow is the root type of the package.
// Use Register methods to register all tasks
// and Run or Main method to execute provided tasks.
//
// A Flow is not safe for concurrent use.
type Flow struct {
	output io.Writer
	usage  func()
	logger Logger

	tasks               map[string]*DefinedTask // snapshot of defined tasks
	defaultTask         *DefinedTask            // task to run when none is explicitly provided
	middlewares         []Middleware
	executorMiddlewares []ExecutorMiddleware
}

// DefaultFlow is the default flow.
// The top-level functions such as Define, Main, and so on are wrappers for the methods of Flow.
var DefaultFlow = &Flow{}

// Tasks returns all tasks sorted in lexicographical order.
func Tasks() []*DefinedTask { _ = "STUB: not implemented"; return nil }

// Tasks returns all tasks sorted in lexicographical order.
func (f *Flow) Tasks() []*DefinedTask { _ = "STUB: not implemented"; return nil }

// Define registers the task. It panics in case of any error.
func Define(task Task) *DefinedTask { _ = "STUB: not implemented"; return nil }

// Define registers the task. It panics in case of any error.
func (f *Flow) Define(task Task) *DefinedTask {
	_ = "STUB: not implemented"
	// validate
	return nil
}

// Undefine unregisters the task. It panics in case of any error.
func Undefine(task *DefinedTask) { _ = "STUB: not implemented"; return }

// Undefine unregisters the task. It panics in case of any error.
func (f *Flow) Undefine(task *DefinedTask) { _ = "STUB: not implemented"; return }

func (f *Flow) isDefined(name string, flow *Flow) bool { _ = "STUB: not implemented"; return false }

// defined in other flow

// Output returns the destination used for printing messages.
// [os.Stdout] is returned if output was not set or was set to nil.
func Output() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

// Output returns the destination used for printing messages.
// [os.Stdout] is returned if output was not set or was set to nil.
func (f *Flow) Output() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

// SetOutput sets the output destination.
func SetOutput(out io.Writer) { _ = "STUB: not implemented"; return }

// SetOutput sets the output destination.
func (f *Flow) SetOutput(out io.Writer) {
	_ = "STUB: not implemented"

	// GetLogger returns the logger used by A's logging functions
	// [CodeLineLogger] is returned if logger was not set or was set to nil.
	return
}

func GetLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

// Logger returns the logger used by A's logging functions
// [CodeLineLogger] is returned if logger was not set or was set to nil.
func (f *Flow) Logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

// SetLogger sets the logger used by A's logging functions.
//
// [A] uses following methods if implemented:
//
//	Error(w io.Writer, args ...interface{})
//	Errorf(w io.Writer, format string, args ...interface{})
//	Fatal(w io.Writer, args ...interface{})
//	Fatalf(w io.Writer, format string, args ...interface{})
//	Skip(w io.Writer, args ...interface{})
//	Skipf(w io.Writer, format string, args ...interface{})
//	Helper()
func SetLogger(logger Logger) { _ = "STUB: not implemented"; return }

// SetLogger sets the logger used by A's logging functions.
//
// [A] uses following methods if implemented:
//
//	Error(w io.Writer, args ...interface{})
//	Errorf(w io.Writer, format string, args ...interface{})
//	Fatal(w io.Writer, args ...interface{})
//	Fatalf(w io.Writer, format string, args ...interface{})
//	Skip(w io.Writer, args ...interface{})
//	Skipf(w io.Writer, format string, args ...interface{})
//	Helper()
func (f *Flow) SetLogger(logger Logger) {
	_ = "STUB: not implemented"

	// Usage returns a function that prints a usage message documenting the flow.
	// It is called when an error occurs while parsing the flow.
	// [Print] is returned if a function was not set or was set to nil.
	return
}

func Usage() func() { _ = "STUB: not implemented"; return nil }

// Usage returns a function that prints a usage message documenting the flow.
// It is called when an error occurs while parsing the flow.
// [Flow.Print] is returned if a function was not set or was set to nil.
func (f *Flow) Usage() func() { _ = "STUB: not implemented"; return nil }

// SetUsage sets the function called when an error occurs while parsing tasks.
func SetUsage(fn func()) { _ = "STUB: not implemented"; return }

// SetUsage sets the function called when an error occurs while parsing tasks.
func (f *Flow) SetUsage(fn func()) {
	_ = "STUB: not implemented"

	// Default returns the default task.
	// nil is returned if default was not set.
	return
}

func Default() *DefinedTask { _ = "STUB: not implemented"; return nil }

// Default returns the default task.
// nil is returned if default was not set.
func (f *Flow) Default() *DefinedTask { _ = "STUB: not implemented"; return nil }

// SetDefault sets a task to run when none is explicitly provided.
// It panics in case of any error.
func SetDefault(task *DefinedTask) { _ = "STUB: not implemented"; return }

// SetDefault sets a task to run when none is explicitly provided.
// Passing nil clears the default task.
// It panics in case of any error.
func (f *Flow) SetDefault(task *DefinedTask) { _ = "STUB: not implemented"; return }

// Use adds task runner middlewares (interceptors).
func Use(middlewares ...Middleware) { _ = "STUB: not implemented"; return }

// Use adds task runner middlewares (interceptors).
func (f *Flow) Use(middlewares ...Middleware) { _ = "STUB: not implemented"; return }

// UseExecutor adds flow executor middlewares (interceptors).
func UseExecutor(middlewares ...ExecutorMiddleware) { _ = "STUB: not implemented"; return }

// UseExecutor adds flow executor middlewares (interceptors).
func (f *Flow) UseExecutor(middlewares ...ExecutorMiddleware) { _ = "STUB: not implemented"; return }

// Option configures the flow execution.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (fn optionFunc) apply(cfg *config) { _ = "STUB: not implemented"; return }

type config struct {
	noDeps    bool
	skipTasks []string
}

// NoDeps is an option to skip processing of all dependencies.
func NoDeps() Option { _ = "STUB: not implemented"; return *new(Option) }

// Skip is an option to skip processing of given tasks.
func Skip(tasks ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// FailError pointer is returned by [Flow.Execute] when a task failed.
type FailError struct {
	Task string
}

func (err *FailError) Error() string { _ = "STUB: not implemented"; return "" }

// Execute runs provided tasks and all their dependencies.
// Each task is executed at most once.
// Returns nil if no task has failed,
// [*FailError] if a task failed,
// other errors in case of invalid input or context error.
func Execute(ctx context.Context, tasks []string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Execute runs provided tasks and all their dependencies.
// Each task is executed at most once.
// Returns nil if no task has failed,
// [*FailError] if a task failed,
// other errors in case of invalid input or context error.
func (f *Flow) Execute(ctx context.Context, tasks []string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// prepare runner

// apply defined executor middlewares

const (
	exitCodePass    = 0
	exitCodeFail    = 1
	exitCodeInvalid = 2
)

// Main runs provided tasks and all their dependencies.
// Each task is executed at most once.
// It exits the current program when after the run is finished
// or SIGINT interrupted the execution.
//   - 0 exit code means that non of the tasks failed.
//   - 1 exit code means that a task has failed or the execution was interrupted.
//   - 2 exit code means that the input was invalid.
//
// Calls [Usage] when invalid args are provided.
func Main(args []string, opts ...Option) { _ = "STUB: not implemented"; return }

// Main runs provided tasks and all their dependencies.
// Each task is executed at most once.
// It exits the current program when after the run is finished
// or SIGINT interrupted the execution.
//   - 0 exit code means that non of the tasks failed.
//   - 1 exit code means that a task has failed or the execution was interrupted.
//   - 2 exit code means that the input was invalid.
//
// Calls [Usage] when invalid args are provided.
func (f *Flow) Main(args []string, opts ...Option) {
	_ = "STUB: not implemented"

	// trap Ctrl+C and call cancel on the context
	return
}

// first signal, cancel context

// second signal, hard exit

func (f *Flow) main(ctx context.Context, args []string, opts ...Option) int {
	err := f.Execute(ctx, args, opts...)
	var ferr *FailError
	if errors.As(err, &ferr) {
		return exitCodeFail
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return exitCodeFail
	}
	if err != nil {
		f.Usage()()
		return exitCodeInvalid
	}
	return exitCodePass
}

// Print prints the information about the registered tasks.
// Tasks with empty [Task.Usage] are not printed.
func Print() { _ = "STUB: not implemented"; return }

// Print prints the information about the registered tasks.
// Tasks with empty [Task.Usage] are not printed.
func (f *Flow) Print() { _ = "STUB: not implemented"; return }
