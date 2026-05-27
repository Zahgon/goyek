package goyek

// Task represents a named task that can have action and dependencies.
type Task struct {
	// Name uniquely identifies the task.
	// It cannot be empty and should be easily representable on the CLI.
	Name string

	// Usage provides information what the task does.
	Usage string

	// Action is function that is called when the task is run.
	// A task can have only dependencies and no action to act as a pipeline.
	Action func(a *A)

	// Deps is a collection of defined tasks
	// that need to be run before this task is executed.
	Deps Deps

	// Parallel marks that this task can be run in parallel
	// with (and only with) other parallel tasks.
	Parallel bool
}

// DefinedTask represents a task that has been defined.
// It can be used as a dependency for another task.
//
// A DefinedTask is not safe for concurrent use.
type DefinedTask struct {
	name     string
	usage    string
	deps     []*DefinedTask
	action   func(a *A)
	parallel bool
	flow     *Flow
}

// Deps represents a collection of dependencies.
type Deps []*DefinedTask

// Name returns the name of the task.
func (r *DefinedTask) Name() string {
	_ = "STUB: not implemented"

	// SetName changes the name of the task.
	return ""
}

func (r *DefinedTask) SetName(s string) { _ = "STUB: not implemented"; return }

// Usage returns the description of the task.
func (r *DefinedTask) Usage() string {
	_ = "STUB: not implemented"

	// SetUsage sets the description of the task.
	return ""
}

func (r *DefinedTask) SetUsage(s string) {
	_ = "STUB: not implemented"

	// Action returns the action of the task.
	return
}

func (r *DefinedTask) Action() func(a *A) {
	_ = "STUB: not implemented"

	// SetAction changes the action of the task.
	return nil
}

func (r *DefinedTask) SetAction(fn func(a *A)) {
	_ = "STUB: not implemented"

	// Deps returns all task's dependencies.
	return
}

func (r *DefinedTask) Deps() Deps { _ = "STUB: not implemented"; return *new(Deps) }

// SetDeps sets all task's dependencies.
func (r *DefinedTask) SetDeps(deps Deps) { _ = "STUB: not implemented"; return }

func (r *DefinedTask) noCycle(deps Deps, visited map[string]bool) bool {
	_ = "STUB: not implemented"
	return false
}

// already checked this branch
