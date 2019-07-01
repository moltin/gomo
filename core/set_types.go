package core

//go:generate go run ../internal/generators/generate_set_types/

func (a *Address) SetType() {
	a.Type = `address`
}

func (a *AdyenGateway) SetType() {
	a.Type = `gateway`
}

func (b *BraintreeGateway) SetType() {
	b.Type = `gateway`
}

func (b *Brand) SetType() {
	b.Type = `brand`
}

func (c *CardConnectGateway) SetType() {
	c.Type = `gateway`
}

func (c *Cart) SetType() {
	c.Type = `cart`
}

func (c *CartItem) SetType() {
	c.Type = `cart_item`
}

func (c *Category) SetType() {
	c.Type = `category`
}

func (c *Collection) SetType() {
	c.Type = `collection`
}

func (c *Currency) SetType() {
	c.Type = `currency`
}

func (c *Customer) SetType() {
	c.Type = `customer`
}

func (e *Event) SetType() {
	e.Type = `event`
}

func (f *File) SetType() {
	f.Type = `file`
}

func (f *Flow) SetType() {
	f.Type = `flow`
}

func (f *Field) SetType() {
	f.Type = `field`
}

func (j *Job) SetType() {
	j.Type = `job`
}

func (v *Variation) SetType() {
	v.Type = `product-variation`
}

func (o *Option) SetType() {
	o.Type = `option`
}

func (o *Order) SetType() {
	o.Type = `order`
}

func (m *Modifier) SetType() {
	m.Type = `modifier`
}

func (i *Inventory) SetType() {
	i.Type = `inventory`
}

func (i *Integration) SetType() {
	i.Type = `integration`
}

func (i *IntegrationLog) SetType() {
	i.Type = `integration-log`
}

func (p *Product) SetType() {
	p.Type = `product`
}

func (p *Promotion) SetType() {
	p.Type = `promotion`
}

func (t *Transaction) SetType() {
	t.Type = `transaction`
}

func (s *Settings) SetType() {
	s.Type = `settings`
}

func (s *StockTransaction) SetType() {
	s.Type = `stock-transaction`
}

func (s *StripeGateway) SetType() {
	s.Type = `gateway`
}

