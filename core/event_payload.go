package core

import (
	"encoding/json"
	"strings"
)

// EventPayload represents the payload delivered to a webhook URL.
// See https://docs.moltin.com/api/advanced/events/event-payload
type EventPayload struct {
	ID          string               `json:"id"`
	TriggeredBy string               `json:"triggered_by"`
	Integration Integration          `json:"integration"`
	Attempt     int                  `json:"attempt"`
	Resources   EventPayloadResource `json:"resources"`
}

type EventPayloadResource interface {
	eventPayloadMember()
}

type EventPayloadDeleted Relationship

func (_ *EventPayloadDeleted) eventPayloadMember() {}

func payloadResourceForEvent(event string) EventPayloadResource {
	bits := strings.SplitN(event, ".", 2)
	resource := bits[0]
	if len(bits) == 2 && bits[1] == "deleted" {
		resource = ""
	}
	return payloadResource(resource)
}

func (ep *EventPayload) UnmarshalJSON(b []byte) error {
	type Alias EventPayload
	payload := struct {
		*Alias
		Resources json.RawMessage `json:"value"`
	}{Alias: (*Alias)(ep)}
	err := json.Unmarshal(b, &payload)
	if err != nil {
		return err
	}
	ep.Resources = payloadResourceForEvent(ep.TriggeredBy)
	return json.Unmarshal(payload.Resources, ep.Resources)
}
