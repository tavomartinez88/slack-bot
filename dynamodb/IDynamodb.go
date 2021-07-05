package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
)

type IDynamodb interface {
	GetSession() *session.Session
	GetClient(s session.Session) *dynamodb.DynamoDB
	CreateRelease(model models.Request) error
}
