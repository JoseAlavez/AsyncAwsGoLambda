package main

import (
	"context"
	"testing"
)

func Test_handle(t *testing.T) {
	type args struct {
		context    context.Context
		rawRequest interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Should normally execute.",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handle(tt.args.context, tt.args.rawRequest); (err != nil) != tt.wantErr {
				t.Errorf("handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
