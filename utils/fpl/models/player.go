package models

import (
	"fmt"
	s "strings"
)

// Players represents all players in the FPL via bootstrap
type Players struct {
	Players     map[int]Player `json:"players"`
	PlayersByID map[int]Player `json:"players"`
}

// GetPlayerByFullName Get FPL element by name
func (p *Players) GetPlayerByFullName(fullname string) (Player, error) {
	var ret Player

	for _, player := range p.Players {
		if s.ToLower(player.GetFullName()) == s.ToLower(fullname) {
			return player, nil
		}
	}
	return ret, fmt.Errorf("No player called %s found", fullname)
}

// Player FPL player
type Player struct {
	ChanceOfPlayingNextRound int     `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound int     `json:"chance_of_playing_this_round"`
	Code                     int     `json:"code"`
	CostChangeEvent          int     `json:"cost_change_event"`
	CostChnageFall           int     `json:"cost_change_event_fall"`
	CostChangeStart          int     `json:"cost_change_start"`
	CostChangeStartFall      int     `json:"cost_change_start_fall"`
	DreamPlayerCount         int     `json:"dreamplayer_count"`
	PlayerTypeID             int     `json:"element_type"`
	EPNext                   float64 `json:"ep_next"` //float64
	EPThis                   float64 `json:"ep_this"` //float64
	EventPoints              int     `json:"event_points"`
	FirstName                string  `json:"first_name"`
	Form                     float64 `json:"form"` //float64
	ID                       int     `json:"id"`
	InDreamTeam              bool    `json:"in_dreamteam"`
	News                     string  `json:"news"`
	NewsAdded                string  `json:"news_added"` //time.Time
	NowCost                  int     `json:"now_cost"`
	Photo                    string  `json:"photo"`
	PointsPerGame            float64 `json:"points_per_game"` //float64
	SecondName               string  `json:"second_name"`
	SelectedByPercent        float64 `json:"selected_by_percent"` //float64
	Special                  bool    `json:"special"`
	SquadNumber              int     `json:"squad_number"`
	Status                   string  `json:"status"`
	TeamID                   int     `json:"team"`
	TeamCode                 int     `json:"team_code"`
	TotalPoints              int     `json:"total_points"`
	TransfersIn              int     `json:"transfers_in"`
	TransfersInEvent         int     `json:"transfers_in_event"`
	TransfersOut             int     `json:"transfers_out"`
	TransfersOutEvent        int     `json:"transfers_out_event"`
	ValueForm                float64 `json:"value_form"`   //float64
	ValueSeason              float64 `json:"value_season"` //float64
	WebName                  string  `json:"web_name"`
	Minutes                  int     `json:"minutes"`
	GoalsScored              int     `json:"goals_scored"`
	Assists                  int     `json:"assists"`
	CleanSheets              int     `json:"clean_sheets"`
	GoalsConceded            int     `json:"goals_conceded"`
	OwnGoals                 int     `json:"own_goals"`
	PenaltiesSaved           int     `json:"penalties_saved"`
	PenaltiesMissed          int     `json:"penalties_missed"`
	YellowCards              int     `json:"yellow_cards"`
	RedCards                 int     `json:"red_cards"`
	Saved                    int     `json:"saves"`
	Bonus                    int     `json:"bonus"`
	BPS                      int     `json:"bps"`
	Influence                float64 `json:"influence"`  //float64
	Creativity               float64 `json:"creativity"` //float64
	Threat                   float64 `json:"threat"`     //float64
	ICTIndex                 float64 `json:"ict_index"`  //float64
}

// GetFullName returns the player's fullname
func (p *Player) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.SecondName)
}

// GetPhotoURL returns the full URL to the players photo
func (p *Player) GetPhotoURL() string {
	return fmt.Sprintf("https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/110x140/p%d.png", p.Code)
}

// GetShirtSmall returns url to small verion of the team shirt image
func (p *Player) GetShirtSmall() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/img/shirts/standard/shirt_%d-66.png", p.TeamCode)
}

// GetShirtMedium returns url to medium verion of the team shirt image
func (p *Player) GetShirtMedium() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/img/shirts/standard/shirt_%d-110.png", p.TeamCode)
}

// GetShirtLarge returns url to large verion of the team shirt image
func (p *Player) GetShirtLarge() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/img/shirts/standard/shirt_%d-220.png", p.TeamCode)
}

func (p *Player) GetTeam(teams *Teams) *Team {
	t := teams.TeamsByCode[p.TeamCode]
	return &t
}

func (p *Player) GetPlayerType(types *PlayerTypes) *PlayerType {
	t := types.Positions[p.PlayerTypeID]
	return &t
}
