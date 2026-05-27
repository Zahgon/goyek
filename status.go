package goyek

// Status of a task run.
type Status uint8

// Statuses of task run.
const (
	StatusNotRun Status = iota
	StatusPassed
	StatusFailed
	StatusSkipped
)

func (s Status) String() string { _ = "STUB: not implemented"; return "" }
