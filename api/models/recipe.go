package models

type Recipe struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Author      string       `json:"author"`
	Submitter   string       `json:"submitter"`
	Course      string       `json:"course"`
	Tags        []string     `json:"tags"`
	Ingredients []Ingredient `json:"ingredients"`
	CanDouble   bool         `json:"canDouble"`
}
