package api

import (
	"encoding/json"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
	//	"io/ioutil"
	"log"
)

func (api *API) GetClassicLeague(leagueID int) (*models.ClassicLeague, error) {
	var league models.ClassicLeague

	byteValue, readErr := api.ExecuteFPLGet(api.Config.API.GetClassicLeagueAPI(leagueID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	league = models.ClassicLeague{}
	json.Unmarshal([]byte(byteValue), &league)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetClassicLeagueFilename(leagueID))

	return &league, nil
}
