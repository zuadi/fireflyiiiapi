package models

type Attributes struct {
	Created_at   string        `json:"created_at,omitempty"`
	Updated_at   string        `json:"updated_at,omitempty"`
	User         string        `json:"user,omitempty"`
	GroupTitle   string        `json:"group_title,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Email        string        `json:"email,omitempty"`
	Blocked      bool          `json:"blocked,omitempty"`
	BlockedCode  string        `json:"blocked_code,omitempty"`
	Role         string        `json:"role,omitempty"`
}
