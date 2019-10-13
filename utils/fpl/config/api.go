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

func (p *FPLAPI) GetBoostrapAPI() string {
	return p.Bootstrap
}

func (p *FPLAPI) GetEntryAPI(entryID int) string {
	return str.Replace(p.Entry, ENTRY_ID, strc.Itoa(entryID), 1)
}

func (p *FPLAPI) GetEntryHistoryAPI(entryID int) string {
	return str.Replace(p.EntryHistory, ENTRY_ID, strc.Itoa(entryID), 1)
}

func (p *FPLAPI) GetEntryTransfersAPI(entryID int) string {
	return str.Replace(p.EntryTransfers, ENTRY_ID, strc.Itoa(entryID), 1)
}

func (p *FPLAPI) GetEntryGameweekAPI(entryID int, eventID int) string {
	return str.Replace(str.Replace(p.EntryGameweek, ENTRY_ID, strc.Itoa(entryID), 1), EVENT_ID, strc.Itoa(eventID), 1)
}

func (p *FPLAPI) GetMyTeamAPI(entryID int) string {
	return str.Replace(p.MyTeam, ENTRY_ID, strc.Itoa(entryID), 1)
}

func (p *FPLAPI) GetElementAPI(elementID int) string {
	return str.Replace(p.Element, ELEMENT_ID, strc.Itoa(elementID), 1)
}

func (p *FPLAPI) GetClassicLeagueAPI(leagueID int) string {
	return str.Replace(p.ClassicLeague, LEAGUE_ID, strc.Itoa(leagueID), 1)
}

func (p *FPLAPI) GetH2HLeagueAPI(leagueID int) string {
	return str.Replace(p.H2HLeague, LEAGUE_ID, strc.Itoa(leagueID), 1)
}

func (p *FPLAPI) GetGameWeekAPI(eventID int) string {
	return str.Replace(p.GameWeek, EVENT_ID, strc.Itoa(eventID), 1)
}
