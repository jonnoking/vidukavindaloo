package config

const (
	redirectURI = "https://fantasy.premierleague.com/"
	fplAPP      = "plfpl-web"
	ENTRY_ID    = "{entry_id}"
	ELEMENT_ID  = "{element_id}"
	LEAGUE_ID   = "{league_id}"
	EVENT_ID    = "{event_id}"
)

type FPLConfig struct {
	MaxEventWeek int
	Files        FPLFile
	API          FPLAPI
	Login        FPLLogin
}

type FPLLogin struct {
	User        string
	Password    string
	RedirectURI string
	App         string
}

func New(username string, password string, maxEventWeek int, folder string, shirtFolder string, playerFolder string) *FPLConfig {

	var maxEvent = 38
	if maxEventWeek > 0 || maxEventWeek < 39 {
		maxEvent = maxEventWeek
	}

	var fld = "./fpl-json"
	if folder != "" {
		fld = folder
	}

	var shirtsFld = "./fpl-shirts"
	if shirtFolder != "" {
		shirtsFld = shirtFolder
	}

	var playerFld = "./fpl-players"
	if playerFolder != "" {
		playerFld = playerFolder
	}

	return &FPLConfig{
		MaxEventWeek: maxEvent,
		Files: FPLFile{
			Folder: fld,
			Shirt:  shirtsFld,
			Player: playerFld,
		},
		Login: FPLLogin{
			User:        username,
			Password:    password,
			RedirectURI: redirectURI,
			App:         fplAPP,
		},
		API: FPLAPI{
			Bootstrap:      "https://fantasy.premierleague.com/api/bootstrap-static/",
			Fixtures:       "https://fantasy.premierleague.com/api/fixtures/",
			Element:        "https://fantasy.premierleague.com/api/element-summary/{element_id}/",
			ClassicLeague:  "https://fantasy.premierleague.com/api/leagues-classic/{league_id}/standings/",
			H2HLeague:      "https://fantasy.premierleague.com/api/leagues-h2h/{league_id}/standings/",
			Entry:          "https://fantasy.premierleague.com/api/entry/{entry_id}/",
			EntryHistory:   "https://fantasy.premierleague.com/api/entry/{entry_id}/history/",
			EntryGameweek:  "https://fantasy.premierleague.com/api/entry/{entry_id}/event/{event_id}/picks/",
			EntryTransfers: "https://fantasy.premierleague.com/api/entry/{entry_id}/transfers/",
			MyTeam:         "https://fantasy.premierleague.com/api/my-team/{entry_id}/",
			GameWeek:       "https://fantasy.premierleague.com/api/event/{entry_id}/live/",
		},
	}
}
