package processor

import (
	"github.com/tavomartinez88/go-modules/logger/loggerImpl"
	"github.com/tavomartinez88/slack-bot/dynamodb"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/utils"
)

const (
	FormatDateTimecategory = "01-02-2006 15:04:05"
)

var logger = loggerImpl.GetLogger()


type Processor struct {
	SlackBotDb dynamodb.IDynamoSlackBotDb
}

func NewProcessor(i dynamodb.IDynamoSlackBotDb) *Processor {
	return &Processor{i}
}

func (receiver *Processor) Process(id string, createdDate string, request models.Request) error {
	logger.Info("Start Process")

	_, err := utils.ValidInput(request)

	if err != nil {
		logger.Error("Invalid request", err)
		return err
	}

	request.Id = id
	request.CreateDate = createdDate

	logger.Println("Send request to create release")

	err = receiver.SlackBotDb.CreateRelease(request)

	if err != nil {
		logger.Error("Error trying create release on dynamodb", err)
		logger.Info("Finished Process")
		return err
	}

	logger.Println("Send request to create release successful")
	logger.Info("Finished Process")
	return nil
}