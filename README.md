# AsyncAwsGoLambda
An aws async go lambda that serves as template for running isolated chained commands per execution.

# Build Example on Linux
```
cd $GOPATH/src/AsyncAwsGoLambdaTemplate
env GOOS=linux GOARCH=amd64 go build -o /tmp/AsyncAwsGoLambdaTemplate
zip -j /tmp/AsyncAwsGoLambdaTemplate.zip /tmp/AsyncAwsGoLambdaTemplate
aws lambda update-function-code --function-name FUNCTION_NAME --zip-file fileb:///tmp/AsyncAwsGoLambdaTemplate.zip
```
