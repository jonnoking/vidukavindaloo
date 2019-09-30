package models

import ()

//https://fantasy.premierleague.com/api/entry/4719576/

// Entry tbc
type Entry struct {
	ID                         int            `json:"id"`
	JoinedTime                 string         `json:"joined_time"`
	StartedEvent               int            `json:"started_event"`
	FavouriteTeam              int            `json:"favourite_team"`
	PlayerFirstName            string         `json:"player_first_name"`
	PlayerLastName             string         `json:"player_last_name"`
	PlayerRegionID             int            `json:"player_region_id"`
	PlayerRegionName           string         `json:"player_region_name"`
	PlayerRegionISOCodeShort   string         `json:"player_region_iso_code_short"`
	PlayerRegionISOCodeLong    string         `json:"player_region_iso_code_long"`
	SummaryOverallPoints       int            `json:"summary_overall_points"`
	SummaryOverallRank         int            `json:"summary_overall_rank"`
	SummaryEventPoints         int            `json:"summary_event_points"`
	SummaryEventRank           int            `json:"summary_event_rank"`
	CurrentEvent               int            `json:"current_event"`
	Leagues                    []EntryLeagues `json:"leagues"`
	Name                       string         `json:"name"`
	Kit                        string         `json:"kit"`
	LastDeadlineBank           int            `json:"last_deadline_bank"`
	LastDeadlineValue          int            `json:"last_deadline_value"`
	LastDeadlineTotalTransfers int            `json:"last_deadline_total_transfers"`
}

// EntryLeagues tbc
type EntryLeagues struct {
	ClassicLeagues []ClassicLeague
	H2HLeagues     []H2HLeague
}

// ClassicLeague tbc
type ClassicLeague struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
	Created        string `json:"created"`
	Closed         bool   `json:"closed"`
	Rank           int    `json:"rank"`
	MaxEntries     bool   `json:"max_entries"`
	LeagueType     string `json:"league_type"`
	Scoring        string `json:"scoring"`
	AdminEntry     bool   `json:"admin_entry"`
	StartEvent     int    `json:"start_event"`
	EntryRank      int    `json:"entry_rank"`
	EntryLastRank  int    `json:"entry_last_rank"`
	EntryCanLeave  bool   `json:"entry_can_leave"`
	EntryCanAdmin  bool   `json:"entry_can_admin"`
	EntryCanInvite bool   `json:"entry_can_invite"`
}

// H2HLeague tbc
type H2HLeague struct {
}

// Kit tbc
type Kit struct {
}

// https://fantasy.premierleague.com/api/entry/1759299/history/

// EntryHistory tbc
type EntryHistory struct {
	Current []EntryEvent    `json:"current"`
	Past    []EventPastYear `json:"past"`
	Chips   []Chips         `json:"chips"`
}

// EntryEvent tbc
type EntryEvent struct {
	Event             int `json:"event"`
	Points            int `json:"points"`
	TotalPoints       int `json:"total_points"`
	Rank              int `json:"rank"`
	RankSort          int `json:"rank_sort"`
	OverallRank       int `json:"overall_rank"`
	Bank              int `json:"bank"`
	Value             int `json:"value"`
	EventTransfers    int `json:"event_transfers"`
	EventTransferCost int `json:"event_transfers_cost"`
	PointsOnBench     int `json:"points_on_bench"`
}

// EventPastYear tbc
type EventPastYear struct {
	SeasonName  string `json:"season_name"`
	TotalPoints int    `json:"total_points"`
	Rank        int    `json:"rank"`
}

// https://fantasy.premierleague.com/api/entry/1759299/event/6/picks/

// EntryPicks tbc
type EntryPicks struct {
	ActiveChip    ActiveChip      `json:"active_chip"`
	AutomaticSubs []AutomaticSubs `json:"automatic_subs"`
	EntryHistory  EntryHistory    `json:"entry_history"`
	Picks         []Pick          `json:"picks"`
}

// ActiveChip tbc
type ActiveChip struct {
}

// AutomaticSubs tbc
type AutomaticSubs struct {
	Entry      int `json:"entry"`
	ElementIn  int `json:"element_in"`
	PlayerIn   Player
	ElementOut int `json:"element_out"`
	PlayerOut  Player
	Event      int `json:"event"`
}

// EventHistory tbc
type EventHistory struct {
	Event             int `json:"event"`
	Point             int `json:"points"`
	TotalPoints       int `json:"total_points"`
	Rank              int `json:"rank"`
	RankSort          int `json:"rank_sort"`
	OverallRank       int `json:"overall_rank"`
	Bank              int `json:"bank"`
	Value             int `json:"value"`
	EventTransfers    int `json:"event_transfers"`
	EventTransferCost int `json:"event_transfers_cost"`
	PointsOnBench     int `json:"points_on_bench"`
}

// Pick tbc
type Pick struct {
	Element       int `json:"element"`
	Player        Player
	Team          Team
	PlayerType    PlayerType
	Position      int  `json:"position"`
	Multiplier    int  `json:"multiplier"`
	IsCaptain     bool `json:"is_captain"`
	IsViceCaptain bool `json:"is_vice_captain"`
}
