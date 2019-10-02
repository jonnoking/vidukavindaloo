package config

import (
	strc "strconv"
	str "strings"
)

type FPLAPI struct {
	Bootstrap      string
	Fixtures       string
	Element        string
	LeagueClassic  string
	LeagueH2H      string
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
		Element:        "https://fantasy.premierleague.com/api/element-summary/{ELEMENT_ID}/",
		LeagueClassic:  "https://fantasy.premierleague.com/api/leagues-classic/{LEAGUE_ID}/standings/",
		LeagueH2H:      "https://fantasy.premierleague.com/api/leagues-h2h/{LEAGUE_ID}/standings/",
		Entry:          "https://fantasy.premierleague.com/api/entry/{ENTRY_ID}/",
		EntryHistory:   "https://fantasy.premierleague.com/api/entry/{ENTRY_ID}/history/",
		EntryGameweek:  "https://fantasy.premierleague.com/api/entry/{ENTRY_ID}/event/{EVENT_ID}/",
		EntryTransfers: "https://fantasy.premierleague.com/api/entry/{ENTRY_ID}/transfers/",
		MyTeam:         "https://fantasy.premierleague.com/api/my-team/{ENTRY_ID}/",
		GameWeek:       "https://fantasy.premierleague.com/api/event/{EVENT_ID}/live/",
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

func GetLeagueClassicAPI(leagueID int) string {
	return str.Replace(GetAPI().LeagueClassic, LEAGUE_ID, strc.Itoa(leagueID), 1)
}

func GetLeagueH2HAPI(leagueID int) string {
	return str.Replace(GetAPI().LeagueH2H, LEAGUE_ID, strc.Itoa(leagueID), 1)
}

func GetGameWeekAPI(eventID int) string {
	return str.Replace(GetAPI().GameWeek, EVENT_ID, strc.Itoa(eventID), 1)
}
