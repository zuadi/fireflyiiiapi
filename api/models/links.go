package models

import "encoding/json"

type Links struct {
	Self  string                `json:"self,omitempty"`
	Other map[string]NestedLink `json:"-"`
}

type NestedLink struct {
	Rel string `json:"rel"`
	Uri string `json:"uri"`
}

func (l *Links) UnmarshalJSON(data []byte) error {
	// Create a temporary map to handle unknown keys
	temp := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Extract "self" field if present
	if self, found := temp["self"]; found {
		json.Unmarshal(self, &l.Self)
		delete(temp, "self")
	}

	// Store remaining keys dynamically in Other
	l.Other = make(map[string]NestedLink)
	for key, value := range temp {
		var link NestedLink
		if err := json.Unmarshal(value, &link); err != nil {
			return err
		}
		l.Other[key] = link
	}
	return nil
}
