package command

import (
	"reflect"
	"testing"
)

type ChanneledCommandMock struct {
	support   bool
	nextState NextState
	results   []Result
}

func (c ChanneledCommandMock) Support(executionContext interface{}) bool {
	return c.support
}

func (c ChanneledCommandMock) Execute(executionContext interface{}, executionResults chan<- Result) NextState {
	for _, result := range c.results {
		executionResults <- result
	}
	return c.nextState
}

func TestCommandChain_Execute(t *testing.T) {
	type fields struct {
		commands []ChanneledCommand
	}
	type args struct {
		executionContext interface{}
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantResult    []Result
		wantNextState NextState
	}{
		{
			name: "Should execute chain and return results.",
			fields: fields{
				commands: []ChanneledCommand{
					ChanneledCommandMock{
						support:   false,
						nextState: Continue,
						results: []Result{
							{_type: Failure, data: "Should skip failure."},
							{_type: Success, data: "Should skip success."},
						},
					},
					ChanneledCommandMock{
						support:   true,
						nextState: Stop,
						results: []Result{
							{_type: Failure, data: "A failure 1."},
							{_type: Success, data: "A success 1."},
						},
					},
					ChanneledCommandMock{
						support:   true,
						nextState: Continue,
						results: []Result{
							{_type: Failure, data: "A failure 2."},
							{_type: Success, data: "A success 2."},
						},
					},
				},
			},
			args: args{
				executionContext: nil,
			},
			wantResult: []Result{
				{_type: Failure, data: "A failure 1."},
				{_type: Success, data: "A success 1."},
			},
			wantNextState: Stop,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChainCommand{
				commands: tt.fields.commands,
			}
			gotResult, gotNextState := c.Execute(tt.args.executionContext)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Execute() gotResult = %v, wantResult %v", gotResult, tt.wantResult)
			}
			if gotNextState != tt.wantNextState {
				t.Errorf("Execute() gotNextState = %v, wantNextState %v", gotNextState, tt.wantNextState)
			}
		})
	}
}

func TestCommandChain_Support(t *testing.T) {
	type fields struct {
		commands []ChanneledCommand
	}
	type args struct {
		executionContext interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should always support the execution.",
			fields: fields{
				commands: nil,
			},
			args: args{
				executionContext: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChainCommand{
				commands: tt.fields.commands,
			}
			if got := c.Support(tt.args.executionContext); got != tt.want {
				t.Errorf("Support() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCommandChain(t *testing.T) {
	type args struct {
		commands []ChanneledCommand
	}
	tests := []struct {
		name string
		args args
		want ChainCommand
	}{
		{
			name: "Should always create a new instance.",
			args: args{
				commands: []ChanneledCommand{ChanneledCommandMock{}},
			},
			want: ChainCommand{
				commands: []ChanneledCommand{ChanneledCommandMock{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChainCommand(tt.args.commands...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChainCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
