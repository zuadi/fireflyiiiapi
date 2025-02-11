package models

type TransactionStore struct {
	ErrorIfDuplicateHash bool                    `json:"error_if_duplicate_hash,omitempty"`
	ApplyRules           bool                    `json:"apply_rules,omitempty"`
	FireWebhooks         bool                    `json:"fire_webhooks,omitempty"`
	GroupTitle           string                  `json:"group_title,omitempty"`
	Transactions         []TransactionSplitStore `json:"transactions,omitempty" validate:"required"`
}
