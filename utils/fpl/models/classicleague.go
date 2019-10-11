package models

import ()

//https://fantasy.premierleague.com/api/leagues-classic/1132753/standings/

// ClassicLeague
type ClassicLeague struct {
	ID             int                    `json:"id"`
	Name           string                 `json:"name"`
	ShortName      string                 `json:"short_name"`
	Created        string                 `json:"created"`
	Closed         bool                   `json:"closed"`
	Rank           int                    `json:"rank"`
	MaxEntries     bool                   `json:"max_entries"`
	LeagueType     string                 `json:"league_type"`
	Scoring        string                 `json:"scoring"`
	AdminEntry     bool                   `json:"admin_entry"`
	CodePrivacy    string                 `json:"code_privacy"`
	StartEvent     int                    `json:"start_event"`
	EntryRank      int                    `json:"entry_rank"`
	EntryLastRank  int                    `json:"entry_last_rank"`
	EntryCanLeave  bool                   `json:"entry_can_leave"`
	EntryCanAdmin  bool                   `json:"entry_can_admin"`
	EntryCanInvite bool                   `json:"entry_can_invite"`
	Standings      ClassicLeagueStandings `json:"standings"`
}

type ClassicLeagueNewEntries struct {
	HasNext bool                            `json:"has_next"`
	Page    int                             `json:"page"`
	Results []ClassicLeagueNewEntriesResult `json:"results"`
}

type ClassicLeagueNewEntriesResult struct {
	Entry           int    `json:"entry"`
	EntryName       string `json:"entry_name"`
	JoinedTime      string `json:"joined_time"`
	PlayerFirstName string `json:"player_first_name"`
	PlayerLastName  string `json:"player_last_name"`
}

type ClassicLeagueStandings struct {
	HasNext bool                          `json:"has_next"`
	Page    int                           `json:"page"`
	Results []ClassicLeagueStandingResult `json:"results"`
}

type ClassicLeagueStandingResult struct {
	ID         int    `json:"id"`
	EventTotal int    `json:"event_total"`
	PlayerName string `json:"player_name"`
	Rank       int    `json:"rank"`
	LastRank   int    `json:"last_rank"`
	RankSort   int    `json:"rank_sort"`
	Total      int    `json:"total"`
	Entry      int    `json:"entry"`
	EntryName  string `json:"entry_name"`
}
