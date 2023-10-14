package domain

type Executor interface {
	ProgName() string
	Look() (string, error)
	Execute(args ...string) error
}
