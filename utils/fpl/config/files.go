package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type FPLFile struct {
	Folder string
	Shirt  string
	Player string
}

func (f *FPLFile) GetBootstrapFilename() string {
	return fmt.Sprintf("%s/bootstrap.json", f.Folder)
}

func (f *FPLFile) GetEntryFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry.json", f.Folder, entryID)
}

func (f *FPLFile) GetEntryHistoryFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-history.json", f.Folder, entryID)
}

func (f *FPLFile) GetEntryTransfersFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-transfers.json", f.Folder, entryID)
}

func (f *FPLFile) GetEntryFullFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-full.json", f.Folder, entryID)
}

func (f *FPLFile) GetEntryGameWeekFilename(entryID int, eventID int) string {
	return fmt.Sprintf("%s/%d-%d-entry-picks.json", f.Folder, entryID, eventID)
}

func (f *FPLFile) GetEntryGameWeekAllFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-entry-picks-full.json", f.Folder, entryID)
}

func (f *FPLFile) GetMyTeamFilename(entryID int) string {
	return fmt.Sprintf("%s/%d-my-team.json", f.Folder, entryID)
}

func (f *FPLFile) GetClassicLeagueFilename(leagueID int) string {
	return fmt.Sprintf("%s/%d-my-league-classic.json", f.Folder, leagueID)
}

func (f *FPLFile) GetTeamLargeShirtFilename(teamID int, teamName string) string {
	team := strings.ReplaceAll(teamName, " ", "-")
	return fmt.Sprintf("%s/%d-shirt-large-%s-220.png", f.Shirt, teamID, team)
}

func (f *FPLFile) GetTeamMediumShirtFilename(teamID int, teamName string) string {
	team := strings.ReplaceAll(teamName, " ", "-")
	return fmt.Sprintf("%s/%d-shirt-medium-%s-110.png", f.Shirt, teamID, team)
}

func (f *FPLFile) GetTeamSmallShirtFilename(teamID int, teamName string) string {
	team := strings.ReplaceAll(teamName, " ", "-")
	return fmt.Sprintf("%s/%d-shirt-small-%s-60.png", f.Shirt, teamID, team)
}

func (f *FPLFile) GetPlayerImageFilename(elementID int, playerName string, teamName string) string {
	player := strings.ReplaceAll(playerName, " ", "-")
	team := strings.ReplaceAll(teamName, " ", "-")
	return fmt.Sprintf("%s/%d-player-%s(%s).png", f.Player, elementID, player, team)
}

// SaveBodyToFile save response body to file
func (f *FPLFile) SaveBodyToFile(body io.ReadCloser, filename string) error {

	byteValue, readErr := ioutil.ReadAll(body)
	if readErr != nil {
		return readErr
	}

	fErr := ioutil.WriteFile(filename, byteValue, 0644)
	if fErr != nil {
		return fErr
	}
	return nil
}

// SaveByteArrayToFile save byte array to file
func (f *FPLFile) SaveByteArrayToFile(b []byte, filename string) error {

	fErr := ioutil.WriteFile(filename, b, 0644)
	if fErr != nil {
		return fErr
	}
	return nil
}

func (f *FPLFile) SaveURLtoFile(url string, filename string) error {
	// wont work for FPL - use ExecuteFPLGet
	// likely due to lack of user-client
	response, e := http.Get(url)
	if e != nil {
		return e
	}

	f.SaveBodyToFile(response.Body, filename)

	// byteValue, readErr := ioutil.ReadAll(response.Body)
	// if readErr != nil {
	// 	return readErr
	// }

	// fErr := ioutil.WriteFile(filename, byteValue, 0644)
	// if fErr != nil {
	// 	return fErr
	// }

	defer response.Body.Close()

	return nil
}
