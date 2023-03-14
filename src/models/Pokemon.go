package models

type Pokemon struct {
	ID     int    `json:"id"`
	Image  string `json:"image"`
	Moves  string `json:"moves"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
	Stats  []Stat `json:"stats"`
}

type Stat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
