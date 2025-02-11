package models

type TransactionArray struct {
	Data  []TransactionRead `json:"data,omitempty"`
	Meta  interface{}       `json:"meta,omitempty"`
	Links PageLink          `json:"link,omitempty"`
}
