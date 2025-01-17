package core

type Transaction struct {
	ID              string         `json:"id,omitempty"`
	Type            string         `json:"type"`
	Reference       string         `json:"reference"`
	Gateway         string         `json:"gateway"`
	Amount          int32          `json:"amount"`
	Currency        string         `json:"currency"`
	TransactionType string         `json:"transaction_type"`
	Status          string         `json:"status"`
	Relationships   *Relationships `json:"relationships,omitempty"`
}
