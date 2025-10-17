package storedmatches

type StoredMatchesV1 []struct {
	Meta struct {
		ID  string `json:"id"`
		Map struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"map"`
		Version   string `json:"version"`
		Mode      string `json:"mode"`
		StartedAt string `json:"started_at"`
		Season    struct {
			ID    string `json:"id"`
			Short string `json:"short"`
		} `json:"season"`
		Region  string `json:"region"`
		Cluster string `json:"cluster"`
	} `json:"meta"`
	Stats struct {
		Puuid     string `json:"puuid"`
		Team      string `json:"team"`
		Level     int    `json:"level"`
		Character struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"character"`
		Tier    int `json:"tier"`
		Score   int `json:"score"`
		Kills   int `json:"kills"`
		Deaths  int `json:"deaths"`
		Assists int `json:"assists"`
		Shots   struct {
			Head int `json:"head"`
			Body int `json:"body"`
			Leg  int `json:"leg"`
		} `json:"shots"`
		Damage struct {
			Dealt    int `json:"dealt"`
			Received int `json:"received"`
		} `json:"damage"`
	} `json:"stats"`
	Teams struct {
		Red  int `json:"red"`
		Blue int `json:"blue"`
	} `json:"teams"`
}
