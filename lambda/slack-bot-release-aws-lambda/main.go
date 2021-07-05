package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
