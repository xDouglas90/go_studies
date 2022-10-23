package models

type Personality struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

var PoaPersonalities []Personality
