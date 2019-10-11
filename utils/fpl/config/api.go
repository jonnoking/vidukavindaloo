package config

import (
	strc "strconv"
	str "strings"
)

type FPLAPI struct {
	Bootstrap      string
	Fixtures       string
	Element        string
	ClassicLeague  string
	H2HLeague      string
	Entry          string
	EntryHistory   string
	EntryGameweek  string
	EntryTransfers string
	MyTeam         string
	GameWeek       string
}

func GetAPI() *FPLAPI {
	return &FPLAPI{
		Bootstrap:      "https://fantasy.premierleague.com/api/bootstrap-static/",
		Fixtures:       "https://fantasy.premierleague.com/api/fixtures/",
		Element:        "https://fantasy.premierleague.com/api/element-summary/{element_id}/",
		ClassicLeague:  "https://fantasy.premierleague.com/api/leagues-classic/{league_id}/standings/",
		H2HLeague:      "https://fantasy.premierleague.com/api/leagues-h2h/{league_id}/standings/",
		Entry:          "https://fantasy.premierleague.com/api/entry/{entry_id}/",
		EntryHistory:   "https://fantasy.premierleague.com/api/entry/{entry_id}/history/",
		EntryGameweek:  "https://fantasy.premierleague.com/api/entry/{entry_id}/event/{event_id}/picks/",
		EntryTransfers: "https://fantasy.premierleague.com/api/entry/{entry_id}/transfers/",
		MyTeam:         "https://fantasy.premierleague.com/api/my-team/{entry_id}/",
		GameWeek:       "https://fantasy.premierleague.com/api/event/{entry_id}/live/",
	}
}

func GetBoostrapAPI() string {
	return GetAPI().Bootstrap
}

func GetEntryAPI(entryID int) string {
	return str.Replace(GetAPI().Entry, ENTRY_ID, strc.Itoa(entryID), 1)
}

func GetEntryHistoryAPI(entryID int) string {
	return str.Replace(GetAPI().EntryHistory, ENTRY_ID, strc.Itoa(entryID), 1)
}

func GetEntryTransfersAPI(entryID int) string {
	return str.Replace(GetAPI().EntryTransfers, ENTRY_ID, strc.Itoa(entryID), 1)
}

func GetEntryGameweekAPI(entryID int, eventID int) string {
	return str.Replace(str.Replace(GetAPI().EntryGameweek, ENTRY_ID, strc.Itoa(entryID), 1), EVENT_ID, strc.Itoa(eventID), 1)
}

func GetMyTeamAPI(entryID int) string {
	return str.Replace(GetAPI().MyTeam, ENTRY_ID, strc.Itoa(entryID), 1)
}

func GetElementAPI(elementID int) string {
	return str.Replace(GetAPI().Element, ELEMENT_ID, strc.Itoa(elementID), 1)
}

func GetClassicLeagueAPI(leagueID int) string {
	return str.Replace(GetAPI().ClassicLeague, LEAGUE_ID, strc.Itoa(leagueID), 1)
}

func GetH2HLeagueAPI(leagueID int) string {
	return str.Replace(GetAPI().H2HLeague, LEAGUE_ID, strc.Itoa(leagueID), 1)
}

func GetGameWeekAPI(eventID int) string {
	return str.Replace(GetAPI().GameWeek, EVENT_ID, strc.Itoa(eventID), 1)
}
