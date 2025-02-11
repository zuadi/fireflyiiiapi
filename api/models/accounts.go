package models

type GetAccounts struct {
	Data AccountArray `json:"data" validate:"required"`
	Meta Meta         `json:"data" validate:"required"`
}
