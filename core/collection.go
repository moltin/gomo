package core

// Collection is a Moltin Collection - https://docs.moltin.com/catalog/collections
type Collection struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Meta        *struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

// CollectionIncludes is possible includes for a Collection
type CollectionIncludes struct {
	Products []Product `json:"products"`
}
