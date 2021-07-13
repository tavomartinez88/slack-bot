package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
)

var tableName = "Releases"

type slackBotDb struct{}

func GetSlackBotDb() *slackBotDb {
	return &slackBotDb{}
}

func (d *slackBotDb) CreateRelease(model models.Request) error {
	client :=  dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))

	item, err := dynamodbattribute.MarshalMap(model)

	if err != nil {
		return err
	}

	doc := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(doc)

	if err != nil {
		return err
	}

	return nil
}

