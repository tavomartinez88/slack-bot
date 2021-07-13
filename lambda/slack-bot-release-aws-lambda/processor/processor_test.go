package processor

import (
	"github.com/go-errors/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/slack-bot/dynamodb"
	"github.com/tavomartinez88/slack-bot/dynamodb/mock_dynamodb"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
	"testing"
	"time"
)

func TestNewProcessor(t *testing.T) {
	p := NewProcessor(dynamodb.GetSlackBotDb())
	assert.NotNil(t, p)
}

func TestProcessor_ProcessThrowErrorFromDb(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	errorDynamoDb := errors.New("Error trying create release on dynamodb")
	mockSlackBotDb := mock_dynamodb.NewMockIDynamoSlackBotDb(mockCtrl)
	processor := &Processor{mockSlackBotDb}
	id := "123"
	createDate := time.Now().Format(FormatDateTimecategory)

	request := models.Request{
		Id: id,
		Title: "titulo",
		CreateDate: createDate,
		Description: "descripcion",
		Product: "producto",
		Detail: "detalle",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXITOSO",
		Observations: "obs",
	}

	mockSlackBotDb.EXPECT().CreateRelease(request).Return(errorDynamoDb)

	err := processor.Process(id, createDate, request)

	if err == nil {
		t.Fail()
	}
}

func TestProcessor_ProcessWithoutStatusThrowErrorFromValidation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSlackBotDb := mock_dynamodb.NewMockIDynamoSlackBotDb(mockCtrl)
	processor := &Processor{mockSlackBotDb}
	id := "123"
	createDate := time.Now().Format(FormatDateTimecategory)

	request := models.Request{}

	err := processor.Process(id, createDate, request)

	if err == nil {
		t.Fail()
	}

	if err != nil {
		assert.EqualValues(t, "Status invalid, Status supported are [IMPLEMENTADO, CANCELADO, PENDIENTE]", err.Error())
	}
}

func TestProcessor_ProcessSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSlackBotDb := mock_dynamodb.NewMockIDynamoSlackBotDb(mockCtrl)
	processor := &Processor{mockSlackBotDb}
	id := "123"
	createDate := time.Now().Format(FormatDateTimecategory)

	request := models.Request{
		Id: id,
		Title: "titulo",
		CreateDate: createDate,
		Description: "descripcion",
		Product: "producto",
		Detail: "detalle",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXITOSO",
		Observations: "obs",
	}

	mockSlackBotDb.EXPECT().CreateRelease(request).Return(nil)

	err := processor.Process(id, createDate, request)

	assert.Nil(t, err)
}
