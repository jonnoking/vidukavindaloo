package models

import (
	"fmt"
	"time"
)

type FPLPlayer struct {
	ChanceOfPlayingNextRound int       `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound int       `json:"chance_of_playing_this_round"`
	Code                     int       `json:"code"`
	CostChangeEvent          int       `json:"cost_change_event"`
	CostChnageFall           int       `json:"cost_change_event_fall"`
	CostChangeStart          int       `json:"cost_change_start"`
	CostChangeStartFall      int       `json:"cost_change_start_fall"`
	DreamTeamCount           int       `json:"dreamteam_count"`
	ElementType              int       `json:"element_type"`
	EPNext                   float32   `json:"ep_next"`
	EPThis                   float32   `json:"ep_this"`
	EventPoints              int       `json:"event_points"`
	FirstName                string    `json:"first_name"`
	Form                     float32   `json:"form"`
	ID                       int       `json:"id"`
	InDreamTeam              bool      `json:"in_dreamteam"`
	News                     string    `json:"news"`
	NewsAdded                time.Time `json:"news_added"`
	NowCost                  int       `json:"now_cost"`
	Photo                    string    `json:"photo"`
	PointsPerGame            float32   `json:"points_per_game"`
	SecondName               string    `json:"second_name"`
	SelectedByPercent        float32   `json:"selected_by_percent"`
	Special                  bool      `json:"special"`
	SquadNumber              int       `json:"squad_number"`
	Status                   string    `json:"status"`
	Team                     int       `json:"team"`
	TeamCode                 int       `json:"team_code"`
	TotalPoints              int       `json:"total_points"`
	TransfersIn              int       `json:"transfers_in"`
	TransfersInEvent         int       `json:"transfers_in_event"`
	TransfersOut             int       `json:"transfers_out"`
	TransfersOutEvent        int       `json:"transfers_out_event"`
	ValueForm                float32   `json:"value_form"`
	ValueSeason              float32   `json:"value_season"`
	WebName                  string    `json:"web_name"`
	Minutes                  int       `json:"minutes"`
	GoalsScored              int       `json:"goals_scored"`
	Assists                  int       `json:"assists"`
	CleanSheets              int       `json:"clean_sheets"`
	GoalsConceded            int       `json:"goals_conceded"`
	OwnGoals                 int       `json:"own_goals"`
	PenaltiesSaved           int       `json:"penalties_saved"`
	PenaltiesMissed          int       `json:"penalties_missed"`
	YellowCards              int       `json:"yellow_cards"`
	RedCards                 int       `json:"red_cards"`
	Saved                    int       `json:"saves"`
	Bonus                    int       `json:"bonus"`
	BPS                      int       `json:"bps"`
	Influence                float32   `json:"influence"`
	Creativity               float32   `json:"creativity"`
	Threat                   float32   `json:"threat"`
	ICTIndex                 float32   `json:"ict_index"`
}

// GetFullName returns the player's fullname
func (p *FPLPlayer) GetFullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.SecondName)
}

// GetPhotoURL returns the full URL to the players photo
func (p *FPLPlayer) GetPhotoURL() string {
	return fmt.Sprintf("https://platform-static-files.s3.amazonaws.com/premierleague/photos/players/110x140/p{%d}.png", p.Code)
}

// GetPlayersTeam returns the players team object
func (p *FPLPlayer) GetPlayersTeam() FPLTeam {
	t := FPLTeam{}

	return t.GetByCode(p.TeamCode)
}
