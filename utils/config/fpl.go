package config

import (
	"strings"
)

const (
	BOOTSTRAP                = "https://fantasy.premierleague.com/api/bootstrap-static/"
	FIXTURES                 = "https://fantasy.premierleague.com/api/fixtures/"
	PLAYERS                  = "https://fantasy.premierleague.com/api/element-summary/{element_id}/"
	LEAGUE_STANDINGS_CLASSIC = "https://fantasy.premierleague.com/api/leagues-classic/1132753/standings/"
	LEAGUE_STANDINGS_H2H     = "https://fantasy.premierleague.com/api/leagues-h2h/{league-id}/standings/"
	USER_DETAILS             = "https://fantasy.premierleague.com/api/entry/1759299/"
	USER_HISTORY             = "https://fantasy.premierleague.com/api/entry/1759299/history/"
	USER_GAMEWEEK            = "https://fantasy.premierleague.com/api/entry/{team_id}/event/{gameweek_id}/picks/"
	MY_TEAM                  = "https://fantasy.premierleague.com/api/my-team/1759299/"
	MY_TRANSFERS             = "https://fantasy.premierleague.com/api/entry/1759299/transfers-latest/"
	GAME_WEEK                = "https://fantasy.premierleague.com/api/event/{gameweek_id}/live/"
	LEAGUE_ENGLAND           = "261"
	LEAGUE_OVERALL           = "314"
	LEAGUE_VIDUKA_VINDALOO   = 1132753
)

func GetPlayersAPI(code int) string {
	api = PLAYERS
	return strings.Replace(api, "{element_id", code, -1)
}
