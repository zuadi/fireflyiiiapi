package models

type AccountRead struct {
	Type       string  `json:"type" validate:"required"`
	Id         string  `json:"id" validate:"required"`
	Attributes Account `json:"attributes" validate:"required"`
}
