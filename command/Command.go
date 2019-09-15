package command

type Command interface {
	Support(executionContext interface{}) bool
	Execute(executionContext interface{}) ([]Result, NextState)
}

type ChanneledCommand interface {
	Support(executionContext interface{}) bool
	Execute(executionContext interface{}, executionResults chan<- Result) NextState
}

type Type int

const (
	Success Type = iota
	Failure
)

func (t Type) String() string {
	return [...]string{"Success", "Failure"}[t]
}

type Result struct {
	_type Type
	data  interface{}
}

type NextState int

const (
	Continue NextState = iota
	Stop
)

func (t NextState) String() string {
	return [...]string{"Continue", "Stop"}[t]
}
