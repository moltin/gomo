package core

// StripeGateway is a Moltin StripeGateway - https://docs.moltin.com/payments/gateways/configure-stripe
type StripeGateway struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Enabled bool   `json:"enabled"`
	Login   string `json:"login"`
}
