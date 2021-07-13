package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/slack-bot/lambda/slack-bot-release-aws-lambda/models"
	"strings"
	"testing"
)

var tests = []struct{
	Id string
	Title string
	CreateDate string
	Description string
	Product string
	Detail string
	Team string
	Status string
	Owner string
	Result string
	Observations string
}{
	{
		Title: "titulo",
		Description: "descripcion",
		Product: "producto",
		Detail: "detalle",
		Team: "equipo",
		Status: "deployado",
		Owner: "owner",
		Result: "sarass",
		Observations: "obs",
	},
	{
		Title: "titulo",
		Description: "descripcion",
		Product: "producto",
		Detail: "detalle",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXIT",
		Observations: "obs",
	},
	{
		Status: "IMPLEMENTADO",
		Result: "EXITOSO",
	},
	{
		Title: "titulo",
		Description: "descripcion",
		Product: "producto",
		Detail: "detalle",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXITOSO",
		Observations: "obs",
	},
	{
		Title: "titulo",
		Product: "producto",
		Detail: "detalle",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXITOSO",
		Observations: "obs",
	},
	{
		Title: "titulo",
		Description: "descripcion",
		Detail: "detalle",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXITOSO",
		Observations: "obs",
	},
	{
		Title: "titulo",
		Description: "descripcion",
		Product: "producto",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXITOSO",
		Observations: "obs",
	},
	{
		Title: "titulo",
		Description: "descripcion",
		Product: "producto",
		Detail: "detalle",
		Status: "IMPLEMENTADO",
		Owner: "owner",
		Result: "EXITOSO",
		Observations: "obs",
	},
	{
		Title: "titulo",
		Description: "descripcion",
		Product: "producto",
		Detail: "detalle",
		Team: "equipo",
		Status: "IMPLEMENTADO",
		Result: "EXITOSO",
		Observations: "obs",
	},
}
func TestValidInput(t *testing.T) {
	var checkError bool
	for _, request := range tests {
		checkError = false
		r := models.Request(request)
		res, err := ValidInput(r)

		if err == nil {
			assert.True(t, res)
		}

		if err != nil {
			assert.False(t, res)
			if !checkError && r.Status != StatusCanceled && r.Status != StatusImplemented && r.Status != StatusPending {
				assert.EqualValues(t, "Status invalid, Status supported are [IMPLEMENTADO, CANCELADO, PENDIENTE]", err.Error())
				assert.False(t, res)
				checkError = !checkError
			}
			if !checkError && r.Result != ResultOk && r.Result != ResultFail && r.Result != ResultPending && r.Result != ResultRollback {
				assert.EqualValues(t, "Result invalid, Result supported are [EXITOSO, PENDIENTE, ROLLBACK, FALLIDO]", err.Error())
				assert.False(t, res)
				checkError = !checkError
			}

			if !checkError {
				assert.True(t, strings.Contains(err.Error(), "is required"))
				checkError = !checkError
			}
		}
	}
}

var testsVerify = []struct{
	Field string
	Message string
}{
	{
		Field : "",
		Message: "Field is required",
	},{
		Field : "title",
		Message: "Field is required",
	},
}
func TestVerify(t *testing.T) {
	for _, d := range testsVerify {
		err := verifyField(d.Field, d.Message)
		if err != nil {
			assert.EqualValues(t, "Field is required", err.Error())
		}
	}
}

