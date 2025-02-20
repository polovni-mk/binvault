package models

type Bucket struct {
	Id          string `json:"id"`
	CreatedBy   string `json:"createdBy"`
	CreatedAt   string `json:"createdAt"`
	AccessLevel string `json:"accessLevel"`
}
