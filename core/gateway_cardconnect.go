package core

// CardConnectGateway is a Moltin CardConnectGateway - https://docs.moltin.com/payments/gateways/configure-cardconnect
type CardConnectGateway struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Enabled    bool   `json:"enabled"`
	Test       bool   `json:"test"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	MerchantID string `json:"merchant_id"`
}
