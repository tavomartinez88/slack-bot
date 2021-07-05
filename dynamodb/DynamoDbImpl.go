package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
)

var tableName = "Releases"

type dynamo struct{}

func GetDynamo() *dynamo {
	return &dynamo{}
}

func (d *dynamo) GetSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func (d *dynamo) GetClient(s *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(s)
}

func (d *dynamo) CreateRelease(model models.Request) error {
	item, err := dynamodbattribute.MarshalMap(model)

	if err != nil {
		return err
	}

	doc := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(tableName),
	}

	_, err = d.GetClient(d.GetSession()).PutItem(doc)

	if err != nil {
		return err
	}

	return nil
}

