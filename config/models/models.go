package models

type UBS struct {
	CSV CSV `yaml:"csv,omitempty"`
}

type CSV struct {
	Seperator               string `yaml:"seperator,omitempty"`
	FieldLimit              int    `yaml:"fieldLimit,omitempty"`
	FirstItemForBankaccount string `yaml:"firstItemForBankaccount,omitempty"`
	FirstItemForCreditCard  string `yaml:"firstItemForCreditCard,omitempty"`
}
