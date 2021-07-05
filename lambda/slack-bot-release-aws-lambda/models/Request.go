package models

type Request struct {
	Id string `json:"id"`
	Title string `json:"title" validate:"required"`
	CreateDate string `json:"create_date"`
	Description string `json:"description" validate:"required"`
	Product string `json:"product" validate:"required"`
	Detail string `json:"detail" validate:"required"`
	Team string `json:"team" validate:"required"`
	Status string `json:"status" validate:"required"`
	Owner string `json:"owner" validate:"required"`
	Result string `json:"result"`
	Observations string `json:"observations"`
}