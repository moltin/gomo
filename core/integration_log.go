package core

// IntegrationLog represents a single event log.
// See https://docs.moltin.com/api/advanced/events/logs
type IntegrationLog struct {
	ID             string              `json:"id,omitempty"`
	Type           string              `json:"type"`
	Succeeded      bool                `json:"succeeded"`
	ProcessingTime float64             `json:"processing_time"`
	Body           string              `json:"body"`
	StatusCode     int                 `json:"status_code"`
	ErrorDetail    string              `json:"error_detail"`
	Relationships  *Relationships      `json:"relationships,omitempty"`
	Meta           *IntegrationLogMeta `json:"meta,omitempty"`
}

type IntegrationLogMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

// IntegrationLogJobs is a list of jobs relating to an event.
// See https://docs.moltin.com/api/advanced/events/logs
type IntegrationLogJobs []Relationship
