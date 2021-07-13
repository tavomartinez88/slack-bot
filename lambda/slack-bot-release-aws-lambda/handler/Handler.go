package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/tavomartinez88/slack-bot/dynamodb"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
	p "github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/processor"
	"time"
)

const (
	FormatDateTimecategory = "01-02-2006 15:04:05"
)

func HandleRequest(ctx context.Context, request models.Request) error {
	slackBotDb := dynamodb.GetSlackBotDb()
	processor := p.NewProcessor(slackBotDb)
	err := processor.Process(uuid.New().String(), time.Now().Format(FormatDateTimecategory), request)

	if err != nil {
		return err
	}

	return nil
}