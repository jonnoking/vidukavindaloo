package config

import (
	"fmt"
)

func GetBootstrapFilename() string {
	return fmt.Sprintf("%s/bootstrap.json", FOLDER)
}

func GetEntryFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry.json", FOLDER, entryID)
}

func GetEntryHistoryFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-history.json", FOLDER, entryID)
}

func GetEntryTransfersFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-transfers.json", FOLDER, entryID)
}

func GetEntryFullFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-full.json", FOLDER, entryID)
}

func GetEntryGameWeekFilename(entryID int, eventID int) string {
	return fmt.Sprintf("%s/%d-%d-entry-picks.json", FOLDER, entryID, eventID)
}

func GetEntryGameWeekAllFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-picks-full.json", FOLDER, entryID)
}

func GetMyTeamFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-my-team.json", FOLDER, entryID)
}

func GetClassicLeagueFilename(leagueID int) string {
	return fmt.Sprintf("%s/%d-my-league-classic.json", FOLDER, leagueID)
}
