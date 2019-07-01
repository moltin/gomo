package core

import (
	"encoding/json"
)

// Modifier represents a modifier for a variation option:
type Modifier struct {
	ID           string        `json:"id,omitempty"`
	Type         string        `json:"type"`
	ModifierType string        `json:"modifier_type"`
	Value        ModifierValue `json:"value"`
}

// ModifierValue repesents a number of possible structures depending
// on the modifier type
type ModifierValue interface {
	modifierMember()
}

// ModifierValuePlain is used for modifier types: description_equals,
// description_prepend, description_append, commoditytype, name_equals,
// name_prepend, name_append, slug_equals, slug_prepend, slug_append,
// sku_equals, sku_prepend, sku_append, status
type ModifierValuePlain struct {
	Value string
}

func (_ *ModifierValuePlain) modifierMember() {}

func (v *ModifierValuePlain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

func (v *ModifierValuePlain) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &v.Value)
}

// ModifierValueBuilder is used for modifier types: sku_builder, slug_builder
type ModifierValueBuilder struct {
	Seek string `json:"seek"`
	Set  string `json:"set"`
}

func (_ *ModifierValueBuilder) modifierMember() {}

// ModifierValuePrice is used for modifier types: price_increment,
// price_decrement, price_equals
type ModifierValuePrice struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	IncludesTax bool   `json:"includes_tax"`
}

func (_ *ModifierValuePrice) modifierMember() {}

func modifierValueFor(mtype string) ModifierValue {
	switch mtype {
	case
		"description_equals",
		"description_prepend",
		"description_append",
		"commoditytype",
		"name_equals",
		"name_prepend",
		"name_append",
		"slug_equals",
		"slug_prepend",
		"slug_append",
		"sku_equals",
		"sku_prepend",
		"sku_append",
		"status":
		return &ModifierValuePlain{""}
	case
		"sku_builder",
		"slug_builder":
		return &ModifierValueBuilder{}
	case
		"price_increment",
		"price_decrement",
		"price_equals":
		return &ModifierValuePrice{}
	}
	return nil
}

// UnmarshalJSON unmarshals a Modifier from JSON
func (m *Modifier) UnmarshalJSON(b []byte) error {
	type Alias Modifier
	modifier := struct {
		*Alias
		Value json.RawMessage `json:"value"`
	}{Alias: (*Alias)(m)}
	err := json.Unmarshal(b, &modifier)
	if err != nil {
		return err
	}
	m.Value = modifierValueFor(m.ModifierType)
	return json.Unmarshal(modifier.Value, m.Value)
}
