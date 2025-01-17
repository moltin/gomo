package core

type Job struct {
	ID      string `json:"id,omitempty"`
	Type    string `json:"type"`
	JobType string `json:"job_type"`
	Link    *struct {
		Href string `json:"href"`
	} `json:"link,omitempty"`
	Status     string      `json:"status,omitempty"`
	Timestamps *Timestamps `json:"timestamps,omitempty"`
}
