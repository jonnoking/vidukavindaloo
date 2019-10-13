package fpl

import (
	"fmt"
	api "github.com/jonnoking/vidukavindaloo/utils/fpl/api"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/config"
	"github.com/jonnoking/vidukavindaloo/utils/fpl/models"
)

var players *models.Players
var teams *models.Teams
var playerTypes *models.PlayerTypes
var events *models.Events
var phases *models.Phases

type FPL struct {
	Config    *config.FPLConfig
	Bootstrap *Bootstrap
	API       *api.API
}

type Bootstrap struct {
	Players     *models.Players
	Teams       *models.Teams
	PlayerTypes *models.PlayerTypes
	Events      *models.Events
	Phases      *models.Phases
}

func New(config *config.FPLConfig) *FPL {

	fplapi := api.New(config)

	return &FPL{
		Config: config,
		API:    fplapi,
	}
}

func (f *FPL) LoadBoostrapLive() {

	// var players *models.Players
	// var teams *models.Teams
	// var playerTypes *models.PlayerTypes
	// var events *models.Events
	// var phases *models.Phases

	bs := f.API.RefreshBootstrap()
	events, _ = models.NewEventsFromBootStrapMap(bs)
	phases, _ = models.NewPhasesFromBootStrapMap(bs)
	playerTypes, _ = models.NewPlayerTypesFromBootStrapMap(bs)
	players, _ = models.NewPlayersFromBootStrapMap(bs)
	teams, _ = models.NewTeamsFromBootStrapMap(bs)

	f.Bootstrap = &Bootstrap{
		Events:      events,
		Phases:      phases,
		PlayerTypes: playerTypes,
		Players:     players,
		Teams:       teams,
	}
}

func (f *FPL) LoadBootstrapCache() {

	b, _ := f.API.LoadBootsrapFromCache()
	events, _ = models.NewEventsFromBootStrapByteArray(b)
	phases, _ = models.NewPhasesFromByteArray(b)
	playerTypes, _ = models.NewPlayerTypesFromByteArray(b)
	players, _ = models.NewPlayersFromBootStrapByteArray(b)
	teams, _ = models.NewTeamsFromBootStrapByteArray(b)

	f.Bootstrap = &Bootstrap{
		Events:      events,
		Phases:      phases,
		PlayerTypes: playerTypes,
		Players:     players,
		Teams:       teams,
	}
}

func init() {
	//LoadFromLive()
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

func (f *FPL) CacheAllTeamShirts() error {
	for _, team := range f.Bootstrap.Teams.Teams {

		large, _ := f.GetFPLImage(team.GetShirtLarge())
		medium, _ := f.GetFPLImage(team.GetShirtMedium())
		small, _ := f.GetFPLImage(team.GetShirtSmall())

		f.Config.Files.SaveByteArrayToFile(large, f.Config.Files.GetTeamLargeShirtFilename(team.ID, team.Name))
		f.Config.Files.SaveByteArrayToFile(medium, f.Config.Files.GetTeamMediumShirtFilename(team.ID, team.Name))
		f.Config.Files.SaveByteArrayToFile(small, f.Config.Files.GetTeamSmallShirtFilename(team.ID, team.Name))
	}

	return nil
}

func (f *FPL) CacheAllPlayerShirts() error {
	for _, player := range f.Bootstrap.Players.Players {

		photo, _ := f.GetFPLImage(player.GetPhotoURL())
		if len(photo) == 0 {
			fmt.Println(player.GetFullName() + " has no photo")
		}
		f.Config.Files.SaveByteArrayToFile(photo, f.Config.Files.GetPlayerImageFilename(player.ID, player.GetFullName(), f.Bootstrap.Teams.TeamsByCode[player.TeamCode].Name))
	}

	return nil
}

func (f *FPL) GetFPLImage(url string) ([]byte, error) {
	return f.API.ExecuteFPLGet(url)
}

// func (f *FPL) GetTeamFromBootstrap(code int) *models.Team {
// 	t := f.Bootstrap.Teams.TeamsByCode[code]
// 	return &t
// }
