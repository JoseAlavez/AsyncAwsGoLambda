package command

import (
	"reflect"
	"testing"
)

func TestExampleCommand_Execute(t *testing.T) {
	results := make(chan Result, 2)
	defer close(results)

	type args struct {
		executionContext interface{}
		executionResults chan<- Result
	}
	tests := []struct {
		name    string
		args    args
		want    NextState
		result1 Result
		result2 Result
	}{
		{
			name: "Should always execute.",
			args: args{
				executionContext: nil,
				executionResults: results,
			},
			want: Stop,
			result1: Result{
				_type: Success,
				data:  "Action #1 executed successfully.",
			},
			result2: Result{
				_type: Success,
				data:  "Action #2 executed successfully.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ExampleCommand{}
			if got := c.Execute(tt.args.executionContext, tt.args.executionResults); got != tt.want {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}

			if got := <-results; got != tt.result1 {
				t.Errorf("Execute() = %v, want %v", got, tt.result1)
			}

			if got := <-results; got != tt.result2 {
				t.Errorf("Execute() = %v, want %v", got, tt.result1)
			}
		})
	}
}

func TestExampleCommand_Support(t *testing.T) {
	type args struct {
		executionContext interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should always support the execution.",
			args: args{
				executionContext: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ExampleCommand{}
			if got := c.Support(tt.args.executionContext); got != tt.want {
				t.Errorf("Support() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewExampleCommand(t *testing.T) {
	tests := []struct {
		name string
		want ExampleCommand
	}{
		{
			name: "Should always create a new instance.",
			want: ExampleCommand{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExampleCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExampleCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
