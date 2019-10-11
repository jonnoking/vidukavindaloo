package api

import (
	"encoding/json"
	"fmt"
	"github.com/jonnoking/vidukavindaloo/utils/cache"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/config"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
	//	"io/ioutil"
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
func GetCompleteEntry(entryID int) (*models.Entry, error) {

	entry, _ := GetEntry(entryID)
	history, _ := GetEntryHistory(entryID)
	transfers, _ := GetEntryTransfers(entryID)
	_, picks, _ := GetAllEntryPicks(entryID)

	entry.History = history
	entry.Transfers = transfers
	entry.EventPicks = &picks

	byteValue, _ := json.Marshal(entry)

	cache.SaveByteArrayToFile(byteValue, config.GetEntryFullFilename(entryID))

	return entry, nil
}

//GetEntry retrive my team from FPL
func GetEntry(entryID int) (*models.Entry, error) {

	var entry models.Entry

	byteValue, readErr := ExecuteFPLGet(config.GetEntryAPI(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	entry = models.Entry{}
	json.Unmarshal([]byte(byteValue), &entry)

	cache.SaveByteArrayToFile(byteValue, config.GetEntryFilename(entryID))

	return &entry, nil
}

//GetEntryHistory retrive my team from FPL
func GetEntryHistory(entryID int) (*models.EntryHistory, error) {

	var entryHistory models.EntryHistory

	byteValue, readErr := ExecuteFPLGet(config.GetEntryHistoryAPI(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	entryHistory = models.EntryHistory{}
	json.Unmarshal([]byte(byteValue), &entryHistory)

	cache.SaveByteArrayToFile(byteValue, config.GetEntryHistoryFilename(entryID))

	return &entryHistory, nil
}

//GetEntryTransfers retrive my team from FPL
func GetEntryTransfers(entryID int) (*models.EntryTransfers, error) {

	var entryTransfers models.EntryTransfers
	var t []models.Transfer

	byteValue, readErr := ExecuteFPLGet(config.GetEntryTransfersAPI(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	entryTransfers = models.EntryTransfers{}

	t = []models.Transfer{}
	json.Unmarshal([]byte(byteValue), &t)

	entryTransfers.Transfers = t

	cache.SaveByteArrayToFile(byteValue, config.GetEntryTransfersFilename(entryID))

	log.Printf("Transfers 1 Length: %d\n", len(entryTransfers.Transfers))

	return &entryTransfers, nil
}

//GetEntryPicks retrive my team from FPL
func GetEntryPicks(entryID int, eventID int) (models.EntryPicks, error) {

	var entryPicks models.EntryPicks

	byteValue, readErr := ExecuteFPLGet(config.GetEntryGameweekAPI(entryID, eventID))
	if readErr != nil {
		return entryPicks, readErr
	}

	entryPicks = models.EntryPicks{}
	json.Unmarshal([]byte(byteValue), &entryPicks)

	cache.SaveByteArrayToFile(byteValue, config.GetEntryGameWeekFilename(entryID, eventID))

	return entryPicks, nil
}

//GetAllEntryPicks Get all 38 event picks
func GetAllEntryPicks(entryID int) ([]models.EntryPicks, models.EntryPicksMap, error) {

	maxEvent := config.MAX_EVENT_WEEK

	eps := []models.EntryPicks{}

	etm := models.EntryPicksMap{}
	etm.EntryEventPicks = map[string]models.EntryPicks{}

	// could move to goroutines - would then need to sort
	for i := 1; i <= maxEvent; i++ {
		ep, e := GetEntryPicks(entryID, i)
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

	cache.SaveByteArrayToFile(byteValue, config.GetEntryGameWeekAllFilename(entryID))

	return eps, etm, nil
}

// CreateTransferMap Return a map from an array with event id as index
func CreateTransferMap(transfers *models.EntryTransfers) (*models.EntryTransfersMap, error) {

	etm := models.EntryTransfersMap{}
	etm.Transfers = map[string]*models.Transfer{}

	log.Printf("Transfers Length: %d\n", len(transfers.Transfers))

	for i, t := range transfers.Transfers {
		etm.Transfers[fmt.Sprintf("event-%d", t.Event)] = &t
		log.Println(i)
	}

	return &etm, nil
}
