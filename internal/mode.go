package internal

type Mode string

const (
	COMPETITIVE    Mode = "competitive"
	CUSTOM         Mode = "custom"
	DEATHMATCH     Mode = "deathmatch"
	ESCALATION     Mode = "escalation"
	TEAMDEATHMATCH Mode = "teamdeathmatch"
	NEWMAP         Mode = "newmap"
	REPLICATION    Mode = "replication"
	SNOWBALLFIGHT  Mode = "snowballfight"
	SPIKERUSH      Mode = "spikerush"
	SWIFTPLAY      Mode = "swiftplay"
	UNRATED        Mode = "unrated"
)
