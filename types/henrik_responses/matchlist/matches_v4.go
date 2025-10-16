package matchlist

import "time"

type MatchesV4 []struct {
	Metadata struct {
		MatchID string `json:"match_id"`
		Map     struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"map"`
		GameVersion    string    `json:"game_version"`
		GameLengthInMs int       `json:"game_length_in_ms"`
		StartedAt      time.Time `json:"started_at"`
		IsCompleted    bool      `json:"is_completed"`
		Queue          struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			ModeType string `json:"mode_type"`
		} `json:"queue"`
		Season struct {
			ID    string `json:"id"`
			Short string `json:"short"`
		} `json:"season"`
		Platform string `json:"platform"`
		Premier  struct {
		} `json:"premier"`
		PartyRrPenaltys []struct {
			PartyID string `json:"party_id"`
			Penalty int    `json:"penalty"`
		} `json:"party_rr_penaltys"`
		Region  string `json:"region"`
		Cluster string `json:"cluster"`
	} `json:"metadata"`
	Players []struct {
		Puuid    string `json:"puuid"`
		Name     string `json:"name"`
		Tag      string `json:"tag"`
		TeamID   string `json:"team_id"`
		Platform string `json:"platform"`
		PartyID  string `json:"party_id"`
		Agent    struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"agent"`
		Stats struct {
			Score     int `json:"score"`
			Kills     int `json:"kills"`
			Deaths    int `json:"deaths"`
			Assists   int `json:"assists"`
			Headshots int `json:"headshots"`
			Legshots  int `json:"legshots"`
			Bodyshots int `json:"bodyshots"`
			Damage    struct {
				Dealt    int `json:"dealt"`
				Received int `json:"received"`
			} `json:"damage"`
		} `json:"stats"`
		AbilityCasts struct {
			Grenade  int `json:"grenade"`
			Ability1 int `json:"ability_1"`
			Ability2 int `json:"ability_2"`
			Ultimate int `json:"ultimate"`
		} `json:"ability_casts"`
		Tier struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
		CardID              string `json:"card_id"`
		TitleID             string `json:"title_id"`
		PreferedLevelBorder string `json:"prefered_level_border"`
		AccountLevel        int    `json:"account_level"`
		SessionPlaytimeInMs int    `json:"session_playtime_in_ms"`
		Behavior            struct {
			AfkRounds    int `json:"afk_rounds"`
			FriendlyFire struct {
				Incoming int `json:"incoming"`
				Outgoing int `json:"outgoing"`
			} `json:"friendly_fire"`
			RoundsInSpawn int `json:"rounds_in_spawn"`
		} `json:"behavior"`
		Economy struct {
			Spent struct {
				Overall int `json:"overall"`
				Average int `json:"average"`
			} `json:"spent"`
			LoadoutValue struct {
				Overall int `json:"overall"`
				Average int `json:"average"`
			} `json:"loadout_value"`
		} `json:"economy"`
	} `json:"players"`
	Observers []struct {
		Puuid               string `json:"puuid"`
		Name                string `json:"name"`
		Tag                 string `json:"tag"`
		AccountLevel        int    `json:"account_level"`
		SessionPlaytimeInMs int    `json:"session_playtime_in_ms"`
		CardID              string `json:"card_id"`
		TitleID             string `json:"title_id"`
		PartyID             string `json:"party_id"`
	} `json:"observers"`
	Coaches []struct {
		Puuid  string `json:"puuid"`
		TeamID string `json:"team_id"`
	} `json:"coaches"`
	Teams []struct {
		TeamID string `json:"team_id"`
		Rounds struct {
			Won  int `json:"won"`
			Lost int `json:"lost"`
		} `json:"rounds"`
		Won           bool `json:"won"`
		PremierRoster struct {
			ID            string   `json:"id"`
			Name          string   `json:"name"`
			Tag           string   `json:"tag"`
			Members       []string `json:"members"`
			Customization struct {
				Icon           string `json:"icon"`
				Image          string `json:"image"`
				PrimaryColor   string `json:"primary_color"`
				SecondaryColor string `json:"secondary_color"`
				TertiaryColor  string `json:"tertiary_color"`
			} `json:"customization"`
		} `json:"premier_roster"`
	} `json:"teams"`
	Rounds []struct {
		ID          int    `json:"id"`
		Result      string `json:"result"`
		Ceremony    string `json:"ceremony"`
		WinningTeam string `json:"winning_team"`
		Plant       struct {
			RoundTimeInMs int    `json:"round_time_in_ms"`
			Site          string `json:"site"`
			Location      struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
			Player struct {
				Puuid string `json:"puuid"`
				Name  string `json:"name"`
				Tag   string `json:"tag"`
				Team  string `json:"team"`
			} `json:"player"`
			PlayerLocations []struct {
				Puuid       string `json:"puuid"`
				Name        string `json:"name"`
				Tag         string `json:"tag"`
				Team        string `json:"team"`
				ViewRadians int    `json:"view_radians"`
				Location    struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"location"`
			} `json:"player_locations"`
		} `json:"plant"`
		Defuse struct {
			RoundTimeInMs int `json:"round_time_in_ms"`
			Location      struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
			Player struct {
				Puuid string `json:"puuid"`
				Name  string `json:"name"`
				Tag   string `json:"tag"`
				Team  string `json:"team"`
			} `json:"player"`
			PlayerLocations []struct {
				Puuid       string `json:"puuid"`
				Name        string `json:"name"`
				Tag         string `json:"tag"`
				Team        string `json:"team"`
				ViewRadians int    `json:"view_radians"`
				Location    struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"location"`
			} `json:"player_locations"`
		} `json:"defuse"`
		Stats []struct {
			AbilityCasts struct {
				Grenade  int `json:"grenade"`
				Ability1 int `json:"ability_1"`
				Ability2 int `json:"ability_2"`
				Ultimate int `json:"ultimate"`
			} `json:"ability_casts"`
			Player struct {
				Puuid string `json:"puuid"`
				Name  string `json:"name"`
				Tag   string `json:"tag"`
				Team  string `json:"team"`
			} `json:"player"`
			DamageEvents []struct {
				Puuid     string `json:"puuid"`
				Name      string `json:"name"`
				Tag       string `json:"tag"`
				Team      string `json:"team"`
				Bodyshots int    `json:"bodyshots"`
				Headshots int    `json:"headshots"`
				Legshots  int    `json:"legshots"`
				Damage    int    `json:"damage"`
			} `json:"damage_events"`
			Stats struct {
				Bodyshots int `json:"bodyshots"`
				Headshots int `json:"headshots"`
				Legshots  int `json:"legshots"`
				Damage    int `json:"damage"`
				Kills     int `json:"kills"`
				Assists   int `json:"assists"`
				Score     int `json:"score"`
			} `json:"stats"`
			Economy struct {
				LoadoutValue int `json:"loadout_value"`
				Remaining    int `json:"remaining"`
				Weapon       struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"weapon"`
				Armor struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"armor"`
			} `json:"economy"`
			WasAfk          bool `json:"was_afk"`
			ReceivedPenalty bool `json:"received_penalty"`
			StayedInSpawn   bool `json:"stayed_in_spawn"`
		} `json:"stats"`
	} `json:"rounds"`
	Kills []struct {
		Round           int `json:"round"`
		TimeInRoundInMs int `json:"time_in_round_in_ms"`
		TimeInMatchInMs int `json:"time_in_match_in_ms"`
		Killer          struct {
			Puuid string `json:"puuid"`
			Name  string `json:"name"`
			Tag   string `json:"tag"`
			Team  string `json:"team"`
		} `json:"killer"`
		Victim struct {
			Puuid string `json:"puuid"`
			Name  string `json:"name"`
			Tag   string `json:"tag"`
			Team  string `json:"team"`
		} `json:"victim"`
		Assistants []struct {
			Puuid string `json:"puuid"`
			Name  string `json:"name"`
			Tag   string `json:"tag"`
			Team  string `json:"team"`
		} `json:"assistants"`
		Location struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"location"`
		Weapon struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"weapon"`
		SecondaryFireMode bool `json:"secondary_fire_mode"`
		PlayerLocations   []struct {
			Puuid       string `json:"puuid"`
			Name        string `json:"name"`
			Tag         string `json:"tag"`
			Team        string `json:"team"`
			ViewRadians int    `json:"view_radians"`
			Location    struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"player_locations"`
	} `json:"kills"`
}
