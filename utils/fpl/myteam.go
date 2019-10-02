package fpl

import (
	"encoding/json"
	"fmt"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/config"
	"io/ioutil"
	"log"
	// "net/http"
	// "time"

	"github.com/jonnoking/vidukavindaloo/utils/cache"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
)

func GetMyTeamFromCache(entryID int, players *models.Players, teams *models.Teams, playerTypes *models.PlayerTypes) (*models.MyTeam, error) {

	// add teamID to file name
	f, err := ioutil.ReadFile(config.GetMyTeamAPI(entryID))
	if err != nil {
		return nil, err
	}

	myteam, _ := models.NewMyTeam(f, players, teams, playerTypes)

	return myteam, nil

}

//GetMyTeam retrive my team from FPL
func GetMyTeam(entryID int) (*models.MyTeam, error) {

	var myteam models.MyTeam
	// apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/my-team/%d/", teamID)

	// client := &http.Client{}

	// r, _ := BuildFPLRequest(apiURL, "GET")

	// resp, respErr := client.Do(r)
	// if respErr != nil {
	// 	return &myteam, respErr
	// }

	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	return &myteam, fmt.Errorf("MyTeam : status code: %d - %s", resp.StatusCode, resp.Status)
	// }

	// byteValue, readErr := ioutil.ReadAll(resp.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }

	byteValue, readErr := ExecuteFPLGet(config.GetMyTeamFilename(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	myteam = models.MyTeam{}
	json.Unmarshal([]byte(byteValue), &myteam)

	cache.SaveByteArrayToFile(byteValue, config.GetMyTeamFilename(entryID))

	return &myteam, nil
}
