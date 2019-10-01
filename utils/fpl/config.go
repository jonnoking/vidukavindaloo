package fpl

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
		Element:        "https://fantasy.premierleague.com/api/element-summary/{element_id}/",
		LeagueClassic:  "https://fantasy.premierleague.com/api/leagues-classic/{league_id}/standings/",
		LeagueH2H:      "https://fantasy.premierleague.com/api/leagues-h2h/{league_id}/standings/",
		Entry:          "https://fantasy.premierleague.com/api/entry/{entry_id}/",
		EntryHistory:   "https://fantasy.premierleague.com/api/entry/{entry_id}/history/",
		EntryGameweek:  "https://fantasy.premierleague.com/api/entry/{entry_id}/event/{event_id}/",
		EntryTransfers: "https://fantasy.premierleague.com/api/entry/{entry_id}/transfers/",
		MyTeam:         "https://fantasy.premierleague.com/api/my-team/{entry_id}/",
		GameWeek:       "https://fantasy.premierleague.com/api/event/{event_id}/live/",
	}
}

const entry_id = "{entry_id}"
const element_id = "{element_id}"
const league_id = "{league_id}"
const event_id = "{event_id}"

func GetEntryAPI(entryID int) string {
	return str.Replace(GetAPI().Entry, entry_id, strc.Itoa(entryID), 1)
}

func GetEntryHistoryAPI(entryID int) string {
	return str.Replace(GetAPI().EntryHistory, entry_id, strc.Itoa(entryID), 1)
}

func GetEntryTransfersAPI(entryID int) string {
	return str.Replace(GetAPI().EntryTransfers, entry_id, strc.Itoa(entryID), 1)
}

func GetEntryGameweekAPI(entryID int, eventID int) string {
	return str.Replace(str.Replace(GetAPI().EntryGameweek, entry_id, strc.Itoa(entryID), 1), event_id, strc.Itoa(eventID), 1)
}

func GetMyTeamAPI(entryID int) string {
	return str.Replace(GetAPI().MyTeam, entry_id, strc.Itoa(entryID), 1)
}

func GetElementAPI(elementID int) string {
	return str.Replace(GetAPI().Element, element_id, strc.Itoa(elementID), 1)
}

func GetLeagueClassicAPI(leagueID int) string {
	return str.Replace(GetAPI().LeagueClassic, league_id, strc.Itoa(leagueID), 1)
}

func GetLeagueH2HAPI(leagueID int) string {
	return str.Replace(GetAPI().LeagueH2H, league_id, strc.Itoa(leagueID), 1)
}

func GetGameWeekAPI(eventID int) string {
	return str.Replace(GetAPI().GameWeek, event_id, strc.Itoa(eventID), 1)
}
