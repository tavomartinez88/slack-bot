package utils

import (
	"errors"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
	"strings"
)

const (
	StatusImplemented = "IMPLEMENTADO"
	StatusCanceled = "CANCELADO"
	StatusPending = "PENDIENTE"
	ResultOk = "EXITOSO"
	ResultPending = "PENDIENTE"
	ResultRollback = "ROLLBACK"
	ResultFail = "FALLIDO"
)

func ValidInput(request models.Request) (bool, error) {
	request.Status = strings.ToUpper(request.Status)
	request.Result = strings.ToUpper(request.Result)

	if request.Status != StatusCanceled && request.Status != StatusImplemented && request.Status != StatusPending {
		return false, errors.New("Status invalid, Status supported are [IMPLEMENTADO, CANCELADO, PENDIENTE]")
	}

	if request.Result != ResultOk && request.Result != ResultFail && request.Result != ResultPending && request.Result != ResultRollback {
		return false, errors.New("Result invalid, Result supported are [EXITOSO, PENDIENTE, ROLLBACK, FALLIDO]")
	}

	return verifyFields(request)
}

func verifyFields(request models.Request) (bool, error) {
	err := verifyField(request.Title, "Title is required")

	if err != nil {
		return false, err
	}

	err = verifyField(request.Description, "Description is required")

	if err != nil {
		return false, err
	}

	err = verifyField(request.Product, "Product is required")

	if err != nil {
		return false, err
	}

	err = verifyField(request.Detail, "Detail is required")

	if err != nil {
		return false, err
	}

	err = verifyField(request.Team, "Team is required")

	if err != nil {
		return false, err
	}

	err = verifyField(request.Status, "Status is required")

	if err != nil {
		return false, err
	}

	err = verifyField(request.Owner, "Owner is required")

	if err != nil {
		return false, err
	}

	err = verifyField(request.Result, "Result is required")

	if err != nil {
		return false, err
	}

	return true, nil
}

func verifyField(field string, description string) error {
	if field == "" {
		return errors.New(description)
	}

	return nil
}
