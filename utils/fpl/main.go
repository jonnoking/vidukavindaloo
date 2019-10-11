package fpl

import (
	api "github.com/jonnoking/vidukavindaloo/utils/fpl/api"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
)

var Players *models.Players
var Teams *models.Teams
var PlayerTypes *models.PlayerTypes
var Events *models.Events
var Phases *models.Phases

func init() {
	LoadFromLive()
}

func LoadFromLive() {
	bs := api.RefreshBootstrap()
	Teams, _ = models.NewTeamsFromBootStrapMap(bs)
	Players, _ = models.NewPlayersFromBootStrapMap(bs)
	PlayerTypes, _ = models.NewPlayerTypesFromBootStrapMap(bs)
	Events, _ = models.NewEventsFromBootStrapMap(bs)
	Phases, _ = models.NewPhasesFromBootStrapMap(bs)
}

func LoadFromCache() {
	// load globals
	b, _ := api.LoadBootsrapFromCache()
	Teams, _ = models.NewTeamsFromBootStrapByteArray(b)
	Players, _ = models.NewPlayersFromBootStrapByteArray(b)
	PlayerTypes, _ = models.NewPlayerTypesFromByteArray(b)
	Events, _ = models.NewEventsFromBootStrapByteArray(b)
	Phases, _ = models.NewPhasesFromByteArray(b)
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
