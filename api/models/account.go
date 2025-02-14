package models

type Account struct {
	CreatedAt                   string  `json:"created_at,omitempty"`
	UpdatedAt                   string  `json:"updated_at,omitempty"`
	Active                      bool    `json:"active,omitempty"`
	Order                       int     `json:"order,omitempty"`
	Name                        string  `json:"name" validate:"required"`
	Type                        string  `json:"type" validate:"required"`
	AccountRole                 string  `json:"account_role"`
	CurrencyId                  string  `json:"currency_id,omitempty"`
	CurrencyCode                string  `json:"currency_code,omitempty"`
	CurrencySymbol              string  `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces       int     `json:"currency_decimal_places,omitempty"`
	NativeCurrencyId            string  `json:"native_currency_id,omitempty"`
	NativeCurrencyCode          string  `json:"native_currency_code,omitempty"`
	NativeCurrencySymbol        string  `json:"native_currency_symbol,omitempty"`
	NativeCurrencyDecimalPlaces int     `json:"native_currency_decimal_places,omitempty"`
	CurrentBalance              string  `json:"current_balance,omitempty"`
	NativeCurrentBalance        string  `json:"native_current_balance,omitempty"`
	CurrentBalanceDate          string  `json:"current_balance_date,omitempty"`
	Notes                       string  `json:"notes,omitempty"`
	MonthlyPaymentDate          string  `json:"monthly_payment_date,omitempty"`
	CreditCardType              string  `json:"credit_card_type,omitempty"`
	AccountNumber               string  `json:"account_number,omitempty"`
	IBAN                        string  `json:"iban,omitempty"`
	BIC                         string  `json:"bic,omitempty"`
	VirtualBalance              string  `json:"virtual_balance,omitempty"`
	OpeningBalance              string  `json:"opening_balance,omitempty"`
	OpeningBalanceDate          string  `json:"opening_balance_date,omitempty"`
	LiabilityType               string  `json:"liability_type,omitempty"`
	LiabilityDirection          string  `json:"liability_direction,omitempty"`
	Interest                    string  `json:"interest,omitempty"`
	InterestPeriod              string  `json:"interest_period,omitempty"`
	CurrentDebt                 string  `json:"current_debt,omitempty"`
	IncludeNetWorth             bool    `json:"include_net_worth,omitempty"`
	Longitude                   float32 `json:"longitude,omitempty"`
	Latitude                    float32 `json:"latitude,omitempty"`
	ZoomLevel                   int     `json:"zoom_level,omitempty"`
}
