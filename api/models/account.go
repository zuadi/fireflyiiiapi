package models

type Account struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Active    string `json:"active,omitempty"`
	Order     int    `json:"order,omitempty"`
	Name      string `json:"order" validate:"required"`
	Type      string `json:"order" validate:"required"`

	Description string `json:"description,omitempty"`

	CurrencyId          string      `json:"currency_id,omitempty"`
	CurrencyCode        string      `json:"currency_code,omitempty"`
	ForeignAmount       string      `json:"foreign_amount,omitempty"`
	ForeignCurrencyId   string      `json:"foreign_currency_id,omitempty"`
	ForeignCurrencyCode string      `json:"foreign_currency_code,omitempty"`
	BudgetId            string      `json:"budget_id,omitempty"`
	CategoryId          string      `json:"category_id,omitempty"`
	CategoryName        string      `json:"category_name,omitempty"`
	SourceId            string      `json:"source_id,omitempty"`
	SourceName          string      `json:"source_name,omitempty"`
	DestinationId       string      `json:"destination_id,omitempty"`
	DestinationName     string      `json:"destination_name,omitempty"`
	Reconciled          bool        `json:"reconciled,omitempty"`
	PiggyBankId         string      `json:"piggy_bank_id,omitempty"`
	PiggyBankName       string      `json:"piggy_bank_name,omitempty"`
	BillId              string      `json:"bill_id,omitempty"`
	BillName            string      `json:"bill_name,omitempty"`
	Tags                interface{} `json:"tags,omitempty"`
	Notes               string      `json:"notes,omitempty"`
	InternalReference   string      `json:"internal_reference,omitempty"`
	ExternalId          string      `json:"external_id,omitempty"`
	ExternalUrl         string      `json:"external_url,omitempty"`
	BunqPaymentId       string      `json:"bunq_payment_id,omitempty"`
	SepaCC              string      `json:"sepa_cc,omitempty"`
	SepaCTOP            string      `json:"sepa_ct_op,omitempty"`
	SepaCTId            string      `json:"sepa_ct_id,omitempty"`
	SepaDB              string      `json:"sepa_db,omitempty"`
	SepaCountry         string      `json:"sepa_country,omitempty"`
	SepaEP              string      `json:"sepa_ep,omitempty"`
	SepaCI              string      `json:"sepa_ci,omitempty"`
	SepaBatchId         string      `json:"sepa_batch_id,omitempty"`
	InterestDate        string      `json:"interest_date,omitempty"`
	BookDate            string      `json:"book_date,omitempty"`
	ProcessDate         string      `json:"process_date,omitempty"`
	DueDate             string      `json:"due_date,omitempty"`
	PaymentDate         string      `json:"payment_date,omitempty"`
	InvoiceDate         string      `json:"invoice_date,omitempty"`
}
