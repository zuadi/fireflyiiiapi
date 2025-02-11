package models

type AccountArray struct {
	Type       string  `json:"type" validate:"required"`
	Id         string  `json:"id" validate:"required"`
	attributes Account `json:"attributes" validate:"required"`
}
