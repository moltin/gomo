package core

type EventPayloadCart struct {
	Data Cart `json:"cart"`
}

func (_ *EventPayloadCart) eventPayloadMember() {}

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

type EventPayloadStockTransaction struct {
	Data StockTransaction `json:"stock-transaction"`
}

func (_ *EventPayloadStockTransaction) eventPayloadMember() {}

type EventPayloadBrand struct {
	Data Brand `json:"brand"`
}

func (_ *EventPayloadBrand) eventPayloadMember() {}

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

type EventPayloadFile struct {
	Data File `json:"file"`
}

func (_ *EventPayloadFile) eventPayloadMember() {}

type EventPayloadTransaction struct {
	Data Transaction `json:"transaction"`
}

func (_ *EventPayloadTransaction) eventPayloadMember() {}

type EventPayloadAddress struct {
	Data Address `json:"address"`
}

func (_ *EventPayloadAddress) eventPayloadMember() {}


func payloadResource(resource string) EventPayloadResource {
	switch resource {
	case `file`:
		return &EventPayloadFile{}
	case `transaction`:
		return &EventPayloadTransaction{}
	case `address`:
		return &EventPayloadAddress{}
	case `category`:
		return &EventPayloadCategory{}
	case `collection`:
		return &EventPayloadCollection{}
	case `currency`:
		return &EventPayloadCurrency{}
	case `product`:
		return &EventPayloadProduct{}
	case `settings`:
		return &EventPayloadSettings{}
	case `stock-transaction`:
		return &EventPayloadStockTransaction{}
	case `brand`:
		return &EventPayloadBrand{}
	case `cart`:
		return &EventPayloadCart{}
	case `integrations`:
		return &EventPayloadIntegration{}
	case `order`:
		return &EventPayloadOrder{}
	default:
		return &EventPayloadDeleted{}
	}
}
