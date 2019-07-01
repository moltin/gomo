package core

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestModifierMarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		modifier Modifier
		expected string
	}{
		{
			"name_equals",
			Modifier{
				Type:         "modifier",
				ModifierType: "name_equals",
				Value:        &ModifierValuePlain{"foo"},
			},
			`{"type":"modifier","modifier_type":"name_equals","value":"foo"}`,
		},
		{
			"sku_builder",
			Modifier{
				Type:         "modifier",
				ModifierType: "sku_builder",
				Value: &ModifierValueBuilder{
					Seek: "foo",
					Set:  "bar",
				},
			},
			`{"type":"modifier","modifier_type":"sku_builder","value":{"seek":"foo","set":"bar"}}`,
		},
		{
			"price_equals",
			Modifier{
				Type:         "modifier",
				ModifierType: "price_equals",
				Value: &ModifierValuePrice{
					Amount:      100,
					Currency:    "USD",
					IncludesTax: true,
				},
			},
			`{"type":"modifier","modifier_type":"price_equals","value":{"amount":100,"currency":"USD","includes_tax":true}}`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.modifier)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != test.expected {
				t.Errorf(
					"\nexpected: %s\ngot:       %s\n",
					test.expected,
					string(data),
				)
			}
		})
	}
}

func TestModiferUnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		json     string
		expected Modifier
	}{
		{
			"name_equals",
			`{"type":"modifier","modifier_type":"name_equals","value":"foo"}`,
			Modifier{
				Type:         "modifier",
				ModifierType: "name_equals",
				Value:        &ModifierValuePlain{"foo"},
			},
		},
		{
			"sku_builder",
			`{"type":"modifier","modifier_type":"sku_builder","value":{"seek":"foo","set":"bar"}}`,
			Modifier{
				Type:         "modifier",
				ModifierType: "sku_builder",
				Value: &ModifierValueBuilder{
					Seek: "foo",
					Set:  "bar",
				},
			},
		},
		{
			"price_equals",
			`{"type":"modifier","modifier_type":"price_equals","value":{"amount":100,"currency":"USD","includes_tax":true}}`,
			Modifier{
				Type:         "modifier",
				ModifierType: "price_equals",
				Value: &ModifierValuePrice{
					Amount:      100,
					Currency:    "USD",
					IncludesTax: true,
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var modifier Modifier
			err := json.Unmarshal([]byte(test.json), &modifier)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(modifier, test.expected) {
				t.Errorf(
					"\nexpected: %s\ngot:       %s\n",
					spew.Sdump(test.expected),
					spew.Sdump(modifier),
				)
			}
		})
	}
}
