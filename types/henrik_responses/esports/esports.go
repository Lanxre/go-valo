package esports

import "time"

type ScheduleEsports []struct {
	Date   time.Time `json:"date"`
	State  string    `json:"state"`
	Type   string    `json:"type"`
	Vod    string    `json:"vod"`
	League struct {
		Name       string `json:"name"`
		Identifier string `json:"identifier"`
		Icon       string `json:"icon"`
		Region     string `json:"region"`
	} `json:"league"`
	Tournament struct {
		Name   string `json:"name"`
		Season string `json:"season"`
	} `json:"tournament"`
	Match struct {
		ID       string `json:"id"`
		GameType struct {
			Type  string `json:"type"`
			Count int    `json:"count"`
		} `json:"game_type"`
		Teams []struct {
			Name     string `json:"name"`
			Code     string `json:"code"`
			Icon     string `json:"icon"`
			HasWon   bool   `json:"has_won"`
			GameWins int    `json:"game_wins"`
			Record   struct {
				Wins   int `json:"wins"`
				Losses int `json:"losses"`
			} `json:"record"`
		} `json:"teams"`
	} `json:"match"`
}