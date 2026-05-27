package goyek

import (
	"context"
	"io"
)

type (
	// Executor represents a flow execution function.
	Executor func(ExecuteInput) error

	// ExecuteInput received by the flow executor.
	ExecuteInput struct {
		Context   context.Context
		Tasks     []string
		SkipTasks []string
		NoDeps    bool
		Output    io.Writer
		Logger    Logger
	}

	// ExecutorMiddleware represents a flow execution interceptor.
	ExecutorMiddleware func(Executor) Executor

	executor struct {
		defined     map[string]*DefinedTask
		middlewares []Middleware
		defaultTask *DefinedTask
	}
)

// Execute runs provided tasks and all their dependencies.
// Each task is executed at most once.
//
//nolint:gocyclo // Contains graph traversal logic.
func (r *executor) Execute(in ExecuteInput) error {
	_ = "STUB: not implemented"
	// Handle default task.
	return nil
}

// Add dependencies to be run first.

// Run task sychronously.

// Find all parallel tasks that have not been run
// and have no dependencies.

// Parallel task has none not-executed dependencies so we can run it.

// Run parallel tasks.

func (r *executor) validate(in ExecuteInput) error { _ = "STUB: not implemented"; return nil }

func (r *executor) canRunTask(task *DefinedTask, visited map[string]bool, noDeps bool) bool {
	_ = "STUB: not implemented"
	return false
}

// We cannot run a non-parallel task in parallel.

// Dependencies are not honored so we can just run the task.

// The task has a dependency which is not executed yet.

func (r *executor) runParallelTasks(ctx context.Context, tasks []*DefinedTask, output io.Writer, logger Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *executor) runTask(ctx context.Context, task *DefinedTask, output io.Writer, logger Logger) error {
	_ = "STUB: not implemented"
	// prepare runner
	return nil
}

// apply defined middlewares

// run action
