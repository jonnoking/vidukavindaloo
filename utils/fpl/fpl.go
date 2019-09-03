package fpl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jonnoking/vidukavindaloo/models"
	"github.com/jonnoking/vidukavindaloo/utils/cache"
)

var Players map[int]models.FPLPlayer
var Teams *models.FPLTeams

func Load() {
	// load globals
	b, _ := LoadBootsrapFromCache()
	Teams, _ = models.NewFPLTeamsFromBootStrapByteArray(b)

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

//GetMyTeam retrive my team from FPL
func GetMyTeam(teamID int) error {
	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/my-team/%d/", teamID)

	client := &http.Client{}

	r, _ := BuildFPLRequest(apiURL, "GET")

	resp, respErr := client.Do(r)
	if respErr != nil {
		return respErr
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Panic("\nStatus Code: ", resp.StatusCode)
		log.Println(resp.Status)
		return nil
	}

	cache.SaveBodyToFile(resp.Body, "./fpl-myteam.json")

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
