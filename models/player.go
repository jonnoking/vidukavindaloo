package models

import (
	"fmt"
	s "strings"
)

// FPLPlayers represents all players in the FPL
type FPLPlayers struct {
	Players map[int]FPLPlayer `mapstructure:"players"`
}

func (p *FPLPlayers) GetPlayerByFullName(fullname string) (FPLPlayer, error) {
	var ret FPLPlayer

	for _, player := range p.Players {
		if s.ToLower(player.GetFullName()) == s.ToLower(fullname) {
			return player, nil
		}
	}
	return ret, fmt.Errorf("No player called %s found", fullname)
}

type FPLPlayer struct {
	ChanceOfPlayingNextRound int     `mapstructure:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound int     `mapstructure:"chance_of_playing_this_round"`
	Code                     int     `mapstructure:"code"`
	CostChangeEvent          int     `mapstructure:"cost_change_event"`
	CostChnageFall           int     `mapstructure:"cost_change_event_fall"`
	CostChangeStart          int     `mapstructure:"cost_change_start"`
	CostChangeStartFall      int     `mapstructure:"cost_change_start_fall"`
	DreamPlayerCount         int     `mapstructure:"dreamplayer_count"`
	ElementType              int     `mapstructure:"element_type"`
	EPNext                   float64 `mapstructure:"ep_next"` //float64
	EPThis                   float64 `mapstructure:"ep_this"` //float64
	EventPoints              int     `mapstructure:"event_points"`
	FirstName                string  `mapstructure:"first_name"`
	Form                     float64 `mapstructure:"form"` //float64
	ID                       int     `mapstructure:"id"`
	InDreamTeam              bool    `mapstructure:"in_dreamteam"`
	News                     string  `mapstructure:"news"`
	NewsAdded                string  `mapstructure:"news_added"` //time.Time
	NowCost                  int     `mapstructure:"now_cost"`
	Photo                    string  `mapstructure:"photo"`
	PointsPerGame            float64 `mapstructure:"points_per_game"` //float64
	SecondName               string  `mapstructure:"second_name"`
	SelectedByPercent        float64 `mapstructure:"selected_by_percent"` //float64
	Special                  bool    `mapstructure:"special"`
	SquadNumber              int     `mapstructure:"squad_number"`
	Status                   string  `mapstructure:"status"`
	Team                     int     `mapstructure:"team"`
	TeamCode                 int     `mapstructure:"team_code"`
	TotalPoints              int     `mapstructure:"total_points"`
	TransfersIn              int     `mapstructure:"transfers_in"`
	TransfersInEvent         int     `mapstructure:"transfers_in_event"`
	TransfersOut             int     `mapstructure:"transfers_out"`
	TransfersOutEvent        int     `mapstructure:"transfers_out_event"`
	ValueForm                float64 `mapstructure:"value_form"`   //float64
	ValueSeason              float64 `mapstructure:"value_season"` //float64
	WebName                  string  `mapstructure:"web_name"`
	Minutes                  int     `mapstructure:"minutes"`
	GoalsScored              int     `mapstructure:"goals_scored"`
	Assists                  int     `mapstructure:"assists"`
	CleanSheets              int     `mapstructure:"clean_sheets"`
	GoalsConceded            int     `mapstructure:"goals_conceded"`
	OwnGoals                 int     `mapstructure:"own_goals"`
	PenaltiesSaved           int     `mapstructure:"penalties_saved"`
	PenaltiesMissed          int     `mapstructure:"penalties_missed"`
	YellowCards              int     `mapstructure:"yellow_cards"`
	RedCards                 int     `mapstructure:"red_cards"`
	Saved                    int     `mapstructure:"saves"`
	Bonus                    int     `mapstructure:"bonus"`
	BPS                      int     `mapstructure:"bps"`
	Influence                float64 `mapstructure:"influence"`  //float64
	Creativity               float64 `mapstructure:"creativity"` //float64
	Threat                   float64 `mapstructure:"threat"`     //float64
	ICTIndex                 float64 `mapstructure:"ict_index"`  //float64
}

// GetFullName returns the player's fullname
func (p *FPLPlayer) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.SecondName)
}

// GetPhotoURL returns the full URL to the players photo
func (p *FPLPlayer) GetPhotoURL() string {
	return fmt.Sprintf("https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/110x140/p%d.png", p.Code)
}
