package dynamodb

import (
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
)

type IDynamoSlackBotDb interface {
	CreateRelease(model models.Request) error
}
