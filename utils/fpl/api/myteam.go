package api

import (
	"encoding/json"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
	"io/ioutil"
	"log"
)

func (api *API) GetMyTeamFromCache(entryID int, players *models.Players, teams *models.Teams, playerTypes *models.PlayerTypes) (*models.MyTeam, error) {

	// add teamID to file name
	f, err := ioutil.ReadFile(api.Config.API.GetMyTeamAPI(entryID))
	if err != nil {
		return nil, err
	}

	myteam, _ := models.NewMyTeam(f, players, teams, playerTypes)

	return myteam, nil

}

//GetMyTeam retrive my team from FPL
func (api *API) GetMyTeam(entryID int) (*models.MyTeam, error) {

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

	byteValue, readErr := api.ExecuteFPLGet(api.Config.Files.GetMyTeamFilename(entryID))
	if readErr != nil {
		log.Fatal(readErr)
	}

	myteam = models.MyTeam{}
	json.Unmarshal([]byte(byteValue), &myteam)

	api.Config.Files.SaveByteArrayToFile(byteValue, api.Config.Files.GetMyTeamFilename(entryID))

	return &myteam, nil
}
