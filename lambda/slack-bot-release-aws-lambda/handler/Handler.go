package handler

import (
	"context"
	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/tavomartinez88/go-modules/logger/loggerImpl"
	"github.com/tavomartinez88/slack-bot/dynamodb"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
	"time"
)

const FormatDateTimecategory = "01-02-2006 15:04:05"
var logger = loggerImpl.GetLogger()

func HandleRequest(ctx context.Context, request models.Request) error {

	now := time.Now()
	db := dynamodb.GetDynamo()

	if !IsValidInput(request) {
		return errors.New("Invalid request, left complete fields on body")
	}

	request.Id = uuid.New().String()
	request.CreateDate = now.Format(FormatDateTimecategory)

	logger.Println("Send request to create release")
	err := db.CreateRelease(request)

	if err != nil {
		logger.Fatal("Error trying create release on dynamodb",err)
	}

	logger.Println("Send request to create release successful")
	return nil
}

func IsValidInput(request models.Request) bool {
	return request.Title != "" && request.Description != "" && request.Product != "" && request.Detail != "" && request.Team != "" && request.Status != "" && request.Owner != "" && request.Result != ""
}