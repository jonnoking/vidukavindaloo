package api

import (
	"encoding/json"
	"github.com/jonnoking/vidukavindaloo/utils/cache"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/config"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
	//	"io/ioutil"
	"log"
)

func GetClassicLeague(leagueID int) (*models.ClassicLeague, error) {
	var league models.ClassicLeague

	byteValue, readErr := ExecuteFPLGet(config.GetClassicLeagueAPI(leagueID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	league = models.ClassicLeague{}
	json.Unmarshal([]byte(byteValue), &league)

	cache.SaveByteArrayToFile(byteValue, config.GetClassicLeagueFilename(leagueID))

	return &league, nil
}
