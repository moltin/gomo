package core_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/moltin/gomo/core"
)

func ExampleEventPayload() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			return
		}

		var payload core.EventPayload
		err = json.Unmarshal(b, &payload)
		if err != nil {
			log.Print(err)
			return
		}

		if payload.TriggeredBy != "order.created" {
			log.Printf("unexpected event: %s\n", payload.TriggeredBy)
			return
		}

		order := payload.Resources.(*core.EventPayloadOrder).Data
		log.Printf("processing new order: %s\n", order.ID)
		// process order...
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
