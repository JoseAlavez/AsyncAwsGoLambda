package command

type ExampleCommand struct {
}

func NewExampleCommand() ExampleCommand {
	return ExampleCommand{}
}

func (c ExampleCommand) Support(executionContext interface{}) bool {
	return true
}

func (c ExampleCommand) Execute(executionContext interface{}, executionResults chan<- Result) NextState {

	executionResults <- Result{
		_type: Success,
		data:  "Action #1 executed successfully.",
	}

	executionResults <- Result{
		_type: Success,
		data:  "Action #2 executed successfully.",
	}

	return Stop
}
