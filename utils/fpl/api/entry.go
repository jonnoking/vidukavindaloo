package api

import (
	"encoding/json"
	"fmt"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
	"log"
	"strings"
)

func GetEntryFromCache(teamID int, players *models.Players, teams *models.Teams, playerTypes *models.PlayerTypes) (*models.Entry, error) {

	// f, err := ioutil.ReadFile(fmt.Sprintf("./fpl-json/fpl-entry-%d.json", teamID))
	// if err != nil {
	// 	return nil, err
	// }

	// entry, _ := models.Entry.(f, players, teams, playerTypes)

	// return entry, nil

	return nil, fmt.Errorf("Not implemented")

}

//GetCompleteEntry Get complete entry details
func (api *API) GetCompleteEntry(entryID int) (*models.Entry, error) {

	entry, _ := api.GetEntry(entryID)
	history, _ := api.GetEntryHistory(entryID)
	transfers, _ := api.GetEntryTransfers(entryID)
	_, picks, _ := api.GetAllEntryPicks(entryID)

	entry.History = history
	entry.Transfers = transfers
	entry.EventPicks = &picks

	byteValue, _ := json.Marshal(entry)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetEntryFullFilename(entryID))

	return entry, nil
}

//GetEntry retrive my team from FPL
func (api *API) GetEntry(entryID int) (*models.Entry, error) {

	var entry models.Entry

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetEntryAPI(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	entry = models.Entry{}
	json.Unmarshal([]byte(byteValue), &entry)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetEntryFilename(entryID))

	return &entry, nil
}

//GetEntryHistory retrive my team from FPL
func (api *API) GetEntryHistory(entryID int) (*models.EntryHistory, error) {

	var entryHistory models.EntryHistory

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetEntryHistoryAPI(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	entryHistory = models.EntryHistory{}
	json.Unmarshal([]byte(byteValue), &entryHistory)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetEntryHistoryFilename(entryID))

	return &entryHistory, nil
}

//GetEntryTransfers retrive my team from FPL
func (api *API) GetEntryTransfers(entryID int) (*models.EntryTransfers, error) {

	var entryTransfers models.EntryTransfers
	var t []models.Transfer

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetEntryTransfersAPI(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	entryTransfers = models.EntryTransfers{}

	t = []models.Transfer{}
	json.Unmarshal([]byte(byteValue), &t)

	entryTransfers.Transfers = t

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetEntryTransfersFilename(entryID))

	log.Printf("Transfers 1 Length: %d\n", len(entryTransfers.Transfers))

	return &entryTransfers, nil
}

//GetEntryPicks retrive my team from FPL
func (api *API) GetEntryPicks(entryID int, eventID int) (models.EntryPicks, error) {

	var entryPicks models.EntryPicks

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetEntryGameweekAPI(entryID, eventID))
	if readErr != nil {
		return entryPicks, readErr
	}

	entryPicks = models.EntryPicks{}
	json.Unmarshal([]byte(byteValue), &entryPicks)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetEntryGameWeekFilename(entryID, eventID))

	return entryPicks, nil
}

//GetAllEntryPicks Get all 38 event picks
func (api *API) GetAllEntryPicks(entryID int) ([]models.EntryPicks, models.EntryPicksMap, error) {

	maxEvent := api.Config.MaxEventWeek

	eps := []models.EntryPicks{}

	etm := models.EntryPicksMap{}
	etm.EntryEventPicks = map[string]models.EntryPicks{}

	// could move to goroutines - would then need to sort
	for i := 1; i <= maxEvent; i++ {
		ep, e := api.GetEntryPicks(entryID, i)
		if e != nil {
			// break if picks returns 404 as the event week is not active
			if strings.Contains(e.Error(), "status code: 404") {
				break
			}
			return nil, etm, e
		}
		eps = append(eps, ep)
		etm.EntryEventPicks[fmt.Sprintf("event-%d", i)] = ep
	}

	byteValue, _ := json.Marshal(eps)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetEntryGameWeekAllFilename(entryID))

	return eps, etm, nil
}

// CreateTransferMap Return a map from an array with event id as index
func (api *API) CreateTransferMap(transfers *models.EntryTransfers) (*models.EntryTransfersMap, error) {

	etm := models.EntryTransfersMap{}
	etm.Transfers = map[string]*models.Transfer{}

	log.Printf("Transfers Length: %d\n", len(transfers.Transfers))

	for i, t := range transfers.Transfers {
		etm.Transfers[fmt.Sprintf("event-%d", t.Event)] = &t
		log.Println(i)
	}

	return &etm, nil
}
