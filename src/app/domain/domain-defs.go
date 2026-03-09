package domain

// Executor models a component capable of describing itself and executing
// a command with the given arguments.
type Executor interface {
	ProgName() string
	Look() (string, error)
	Execute(args ...string) error
}
