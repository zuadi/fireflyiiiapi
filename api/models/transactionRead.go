package models

type TransactionRead struct {
	Type       string     `json:"type,omitempty"`
	ID         string     `json:"id,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
	Links      Links      `json:"links,omitempty"`
}
