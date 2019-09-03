package fpl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jonnoking/vidukavindaloo/models"
	"github.com/jonnoking/vidukavindaloo/utils/cache"
)

var Players *models.Players
var Teams *models.Teams
var PlayerTypes *models.PlayerTypes

func Load() {
	// load globals
	b, _ := LoadBootsrapFromCache()
	Teams, _ = models.NewTeamsFromBootStrapByteArray(b)
	Players, _ = models.NewPlayersFromBootStrapByteArray(b)
	PlayerTypes, _ = models.NewPlayerTypesFromByteArray(b)

	//Teams := models.NewFPLTeams(t)

	// ts := []models.FPLTeam{}

	// var result map[string]interface{}
	// json.Unmarshal([]byte(b), &result)

	// teams := result["teams"].([]interface{})

	// for _, v := range teams {

	// 	//t, ok := v.(models.FPLTeam) // check interface against type
	// 	// if !ok {
	// 	// 	println("Not ok")
	// 	// 	log.Println(v)
	// 	// }
	// 	var team models.FPLTeam
	// 	mapstructure.Decode(v, &team)
	// 	ts = append(ts, team)
	// 	//log.Println(team)
	// }

	// Teams = models.NewFPLTeams(ts)

}

func main() {

	// cookies, err := RefreshCookies()
	// if err != nil {
	// 	log.Println(err)
	// }
	// CacheCookies(cookies)

	// cookies, _ := ReadCookieCache()
	// log.Println(cookies["pl_profile"].Value)
	// log.Println(cookies["pl_profile"].RawExpires)

	// isValid, _ := ValidateCookies(cookies)
	// log.Println(isValid)

	//GetMyTeam()

}

func GetMyTeam(teamID int, players *models.Players, teams *models.Teams, playerTypes *models.PlayerTypes) (*models.MyTeam, error) {

	// add teamID to file name
	f, err := ioutil.ReadFile("./fpl-myteam.json")
	if err != nil {
		return nil, err
	}

	myteam, _ := models.NewMyTeam(f, players, teams, playerTypes)

	return myteam, nil

}

//RefreshMyTeam retrive my team from FPL
func RefershMyTeam(teamID int) (*models.MyTeam, error) {

	var myteam models.MyTeam
	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/my-team/%d/", teamID)

	client := &http.Client{}

	r, _ := BuildFPLRequest(apiURL, "GET")

	resp, respErr := client.Do(r)
	if respErr != nil {
		return &myteam, respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return &myteam, fmt.Errorf("MyTeam : status code: %d - %s", resp.StatusCode, resp.Status)
	}

	byteValue, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	myteam = models.MyTeam{}
	json.Unmarshal([]byte(byteValue), &myteam)

	cache.SaveBodyToFile(resp.Body, "./fpl-myteam.json")

	return &myteam, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
