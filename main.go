package main

import (
	"JoseAlavez/AwsGoLambda/command"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"runtime/debug"
)

func main() {
	lambda.Start(handle)
}

func handle(context context.Context, rawRequest interface{}) (err error) {

	defer func() {
		if recover := recover(); recover != nil {
			log.Printf("Exit with error: %+v.\n%s", recover, string(debug.Stack()))
		}
		err = nil
	}()

	command.NewChainCommand(
		command.NewExampleCommand(),
	).Execute(rawRequest)

	return nil
}
