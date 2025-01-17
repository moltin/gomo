package core

// Inventory represents a product stock object in the Moltin API: https://docs.moltin.com/api/catalog/inventory
type Inventory struct {
	ID        string `json:"id,omitempty"`
	Type      string `json:"type"`
	Total     int64  `json:"total"`
	Available int64  `json:"available"`
	Allocated int64  `json:"allocated"`
}
