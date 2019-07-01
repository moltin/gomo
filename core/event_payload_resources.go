package core

//go:generate go run ../internal/generators/generate_event_payload_resources

type EventPayloadAddress struct {
	Data Address `json:"address"`
}

func (_ *EventPayloadAddress) eventPayloadMember() {}

type EventPayloadCart struct {
	Data Cart `json:"cart"`
}

func (_ *EventPayloadCart) eventPayloadMember() {}

type EventPayloadCategory struct {
	Data Category `json:"category"`
}

func (_ *EventPayloadCategory) eventPayloadMember() {}

type EventPayloadCollection struct {
	Data Collection `json:"collection"`
}

func (_ *EventPayloadCollection) eventPayloadMember() {}

type EventPayloadCurrency struct {
	Data Currency `json:"currency"`
}

func (_ *EventPayloadCurrency) eventPayloadMember() {}

type EventPayloadStockTransaction struct {
	Data StockTransaction `json:"stock-transaction"`
}

func (_ *EventPayloadStockTransaction) eventPayloadMember() {}

type EventPayloadTransaction struct {
	Data Transaction `json:"transaction"`
}

func (_ *EventPayloadTransaction) eventPayloadMember() {}

type EventPayloadBrand struct {
	Data Brand `json:"brand"`
}

func (_ *EventPayloadBrand) eventPayloadMember() {}

type EventPayloadFile struct {
	Data File `json:"file"`
}

func (_ *EventPayloadFile) eventPayloadMember() {}

type EventPayloadIntegration struct {
	Data Integration `json:"integrations"`
}

func (_ *EventPayloadIntegration) eventPayloadMember() {}

type EventPayloadOrder struct {
	Data Order `json:"order"`
}

func (_ *EventPayloadOrder) eventPayloadMember() {}

type EventPayloadProduct struct {
	Data Product `json:"product"`
}

func (_ *EventPayloadProduct) eventPayloadMember() {}

type EventPayloadSettings struct {
	Data Settings `json:"settings"`
}

func (_ *EventPayloadSettings) eventPayloadMember() {}


func payloadResource(resource string) EventPayloadResource {
	switch resource {
	case `address`:
		return &EventPayloadAddress{}
	case `cart`:
		return &EventPayloadCart{}
	case `category`:
		return &EventPayloadCategory{}
	case `collection`:
		return &EventPayloadCollection{}
	case `currency`:
		return &EventPayloadCurrency{}
	case `settings`:
		return &EventPayloadSettings{}
	case `stock-transaction`:
		return &EventPayloadStockTransaction{}
	case `transaction`:
		return &EventPayloadTransaction{}
	case `brand`:
		return &EventPayloadBrand{}
	case `file`:
		return &EventPayloadFile{}
	case `integrations`:
		return &EventPayloadIntegration{}
	case `order`:
		return &EventPayloadOrder{}
	case `product`:
		return &EventPayloadProduct{}
	default:
		return &EventPayloadDeleted{}
	}
}
