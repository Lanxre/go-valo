package accounts

type AccountV1 struct {
	PUUID         string `json:"puuid"`
	Region        string `json:"region"`
	AccountLevel  int    `json:"account_level"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Card          Card   `json:"card"`
	LastUpdate    string `json:"last_update"`
	LastUpdateRaw int    `json:"last_update_raw"`
}

type Card struct {
	Small string `json:"small"`
	Large string `json:"large"`
	Wide  string `json:"wide"`
	ID    string `json:"id"`
}
