package main

import (
	"fmt"
	"os"
)

var resources = map[string]string{
	"Address":     "address",
	"Brand":       "brand",
	"Cart":        "cart",
	"Category":    "category",
	"Collection":  "collection",
	"Currency":    "currency",
	"File":        "file",
	"Integration": "integrations",
	"Order":       "order",
	//"Gateway": "payment-gateway",
	"Product":          "product",
	"Settings":         "settings",
	"StockTransaction": "stock-transaction",
	"Transaction":      "transaction",
}

func main() {
	f, err := os.OpenFile(
		"event_payload_resources.go",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, `package core`)
	fmt.Fprintln(f, ``)
	for t, n := range resources {
		fmt.Fprintf(f, "type EventPayload%s struct {\n", t)
		fmt.Fprintf(f, "	Data %s `json:\"%s\"`\n", t, n)
		fmt.Fprintln(f, `}`)
		fmt.Fprintln(f, ``)
		fmt.Fprintf(f, "func (_ *EventPayload%s) eventPayloadMember() {}\n", t)
		fmt.Fprintln(f, ``)
	}
	fmt.Fprintln(f, ``)
	fmt.Fprintln(f, `func payloadResource(resource string) EventPayloadResource {`)
	fmt.Fprintln(f, `	switch resource {`)
	for t, n := range resources {
		fmt.Fprintf(f, "	case `%s`:\n", n)
		fmt.Fprintf(f, "		return &EventPayload%s{}\n", t)
	}
	fmt.Fprintln(f, `	default:`)
	fmt.Fprintln(f, `		return &EventPayloadDeleted{}`)
	fmt.Fprintln(f, `	}`)
	fmt.Fprintln(f, `}`)
}
