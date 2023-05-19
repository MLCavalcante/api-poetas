package models

type Poeta struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Historia string `json:"historia"`
}

var Poetas []Poeta