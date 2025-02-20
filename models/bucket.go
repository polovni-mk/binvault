package models

type Bucket struct {
	Id        string  `json:"id"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedAt string  `json:"createdAt"`
	Access    string  `json:"access"`
}
