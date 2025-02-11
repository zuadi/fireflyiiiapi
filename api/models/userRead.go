package models

type UserRead struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
	Links      Links      `json:"links"`
}
