package command

import "sync"

type ChainCommand struct {
	commands []ChanneledCommand
}

func NewChainCommand(commands ...ChanneledCommand) ChainCommand {

	if commands == nil || len(commands) == 0 {
		panic("Commands cannot be empty.")
	}

	return ChainCommand{commands: commands}
}

func (c ChainCommand) Support(executionContext interface{}) bool {
	return true
}

func (c ChainCommand) Execute(executionContext interface{}) ([]Result, NextState) {
	executionResults := make(chan Result)

	go c.ProduceResults(executionContext, executionResults)
	results := c.ConsumeResults(executionResults)

	return results, Stop
}

func (c ChainCommand) ConsumeResults(executionResults <-chan Result) []Result {
	results := make([]Result, 0)

	var workgroup sync.WaitGroup
	workgroup.Add(1)
	go func() {
		for result := range executionResults {
			results = append(results, result)
		}
		workgroup.Done()
	}()
	workgroup.Wait()

	return results
}

func (c ChainCommand) ProduceResults(executionContext interface{}, executionResults chan<- Result) {
	defer close(executionResults)

	for _, command := range c.commands {
		if command.Support(executionContext) {
			if next := command.Execute(executionContext, executionResults); next == Stop {
				break
			}
		}
	}

}
