package core_test

import (
	"testing"

	"github.com/moltin/gomo/core"
)

func TestAddressType(t *testing.T) {
	var ca core.Address
	ca.SetType()

	if ca.Type != `address` {
		t.Error("failed to set type")
	}
}

func TestAdyenGatewayType(t *testing.T) {
	var ca core.AdyenGateway
	ca.SetType()

	if ca.Type != `gateway` {
		t.Error("failed to set type")
	}
}

func TestBraintreeGatewayType(t *testing.T) {
	var cb core.BraintreeGateway
	cb.SetType()

	if cb.Type != `gateway` {
		t.Error("failed to set type")
	}
}

func TestBrandType(t *testing.T) {
	var cb core.Brand
	cb.SetType()

	if cb.Type != `brand` {
		t.Error("failed to set type")
	}
}

func TestCardConnectGatewayType(t *testing.T) {
	var cc core.CardConnectGateway
	cc.SetType()

	if cc.Type != `gateway` {
		t.Error("failed to set type")
	}
}

func TestCartType(t *testing.T) {
	var cc core.Cart
	cc.SetType()

	if cc.Type != `cart` {
		t.Error("failed to set type")
	}
}

func TestCartItemType(t *testing.T) {
	var cc core.CartItem
	cc.SetType()

	if cc.Type != `cart_item` {
		t.Error("failed to set type")
	}
}

func TestCategoryType(t *testing.T) {
	var cc core.Category
	cc.SetType()

	if cc.Type != `category` {
		t.Error("failed to set type")
	}
}

func TestCollectionType(t *testing.T) {
	var cc core.Collection
	cc.SetType()

	if cc.Type != `collection` {
		t.Error("failed to set type")
	}
}

func TestCurrencyType(t *testing.T) {
	var cc core.Currency
	cc.SetType()

	if cc.Type != `currency` {
		t.Error("failed to set type")
	}
}

func TestCustomerType(t *testing.T) {
	var cc core.Customer
	cc.SetType()

	if cc.Type != `customer` {
		t.Error("failed to set type")
	}
}

func TestEventType(t *testing.T) {
	var ce core.Event
	ce.SetType()

	if ce.Type != `event` {
		t.Error("failed to set type")
	}
}

func TestFileType(t *testing.T) {
	var cf core.File
	cf.SetType()

	if cf.Type != `file` {
		t.Error("failed to set type")
	}
}

func TestFlowType(t *testing.T) {
	var cf core.Flow
	cf.SetType()

	if cf.Type != `flow` {
		t.Error("failed to set type")
	}
}

func TestFieldType(t *testing.T) {
	var cf core.Field
	cf.SetType()

	if cf.Type != `field` {
		t.Error("failed to set type")
	}
}

func TestJobType(t *testing.T) {
	var cj core.Job
	cj.SetType()

	if cj.Type != `job` {
		t.Error("failed to set type")
	}
}

func TestVariationType(t *testing.T) {
	var cv core.Variation
	cv.SetType()

	if cv.Type != `product-variation` {
		t.Error("failed to set type")
	}
}

func TestOptionType(t *testing.T) {
	var co core.Option
	co.SetType()

	if co.Type != `option` {
		t.Error("failed to set type")
	}
}

func TestOrderType(t *testing.T) {
	var co core.Order
	co.SetType()

	if co.Type != `order` {
		t.Error("failed to set type")
	}
}

func TestModifierType(t *testing.T) {
	var cm core.Modifier
	cm.SetType()

	if cm.Type != `modifier` {
		t.Error("failed to set type")
	}
}

func TestInventoryType(t *testing.T) {
	var ci core.Inventory
	ci.SetType()

	if ci.Type != `inventory` {
		t.Error("failed to set type")
	}
}

func TestIntegrationType(t *testing.T) {
	var ci core.Integration
	ci.SetType()

	if ci.Type != `integration` {
		t.Error("failed to set type")
	}
}

func TestIntegrationLogType(t *testing.T) {
	var ci core.IntegrationLog
	ci.SetType()

	if ci.Type != `integration-log` {
		t.Error("failed to set type")
	}
}

func TestProductType(t *testing.T) {
	var cp core.Product
	cp.SetType()

	if cp.Type != `product` {
		t.Error("failed to set type")
	}
}

func TestPromotionType(t *testing.T) {
	var cp core.Promotion
	cp.SetType()

	if cp.Type != `promotion` {
		t.Error("failed to set type")
	}
}

func TestTransactionType(t *testing.T) {
	var ct core.Transaction
	ct.SetType()

	if ct.Type != `transaction` {
		t.Error("failed to set type")
	}
}

func TestSettingsType(t *testing.T) {
	var cs core.Settings
	cs.SetType()

	if cs.Type != `settings` {
		t.Error("failed to set type")
	}
}

func TestStockTransactionType(t *testing.T) {
	var cs core.StockTransaction
	cs.SetType()

	if cs.Type != `stock-transaction` {
		t.Error("failed to set type")
	}
}

func TestStripeGatewayType(t *testing.T) {
	var cs core.StripeGateway
	cs.SetType()

	if cs.Type != `gateway` {
		t.Error("failed to set type")
	}
}

