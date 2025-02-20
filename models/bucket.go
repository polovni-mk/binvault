package models

type Bucket struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Public bool   `json:"public"`
}
