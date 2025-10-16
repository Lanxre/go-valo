package accounts

import "time"

type AccountV2 struct {
	Puuid        string    `json:"puuid"`
	Region       string    `json:"region"`
	AccountLevel int       `json:"account_level"`
	Name         string    `json:"name"`
	Tag          string    `json:"tag"`
	Card         string    `json:"card"`
	Title        string    `json:"title"`
	Platforms    []string  `json:"platforms"`
	UpdatedAt    time.Time `json:"updated_at"`
}
