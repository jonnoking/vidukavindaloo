package models

import (
	"fmt"
	"log"
	s "strings"
)

// Teams represents all teams in the FPL
type Teams struct {
	TeamsByCode map[int]Team `json:"teams"`
	Teams       []Team
}

// New Create a new team
func (p *Teams) New(teams []Team) {
	ts := map[int]Team{}
	for _, team := range teams {
		ts[team.Code] = team
	}
	p.TeamsByCode = ts
}

// GetTeamByName Returns a team details via a team name
func (p *Teams) GetTeamByName(name string) *Team {
	var ret Team

	for _, team := range p.Teams {
		if s.ToLower(team.Name) == s.ToLower(name) {
			return &team
		}
	}
	log.Panicf("No team called %s found", name)
	return &ret
}

// GetTeamByCode Returns a team details via a team code
func (p *Teams) GetTeamByCode(code int) (Team, error) {
	var ret Team

	team, found := p.TeamsByCode[code]
	if !found {
		return ret, fmt.Errorf("No team found with code %d", code)
	}

	return team, nil
}

// Team represents a Premier League team
type Team struct {
	Code                int    `json:"code"`
	Draw                int    `json:"draw"`
	Form                int    `json:"form"`
	ID                  int    `json:"id"`
	Lost                int    `json:"loss"`
	Name                string `json:"name"`
	Played              int    `json:"played"`
	Points              int    `json:"points"`
	Position            int    `json:"position"`
	ShortName           string `json:"short_name"`
	Strength            int    `json:"strength"`
	TeamDivision        int    `json:"team_division"`
	Unavailable         bool   `json:"unavailable"`
	Win                 int    `json:"win"`
	StrengthOverallHome int    `json:"strength_overall_home"`
	StrengthOverallAway int    `json:"strength_overall_away"`
	StrengthAttackHome  int    `json:"strength_attack_home"`
	StrengthAttackAway  int    `json:"strength_attack_away"`
	StrengthDefenceHome int    `json:"strength_defence_home"`
	StrengthDefenceAway int    `json:"strength_defence_away"`
}

// GetShirtSmall returns url to small verion of the team shirt image
func (p *Team) GetShirtSmall() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/img/shirts/standard/shirt_%d-66.png", p.Code)
}

// GetShirtMedium returns url to medium verion of the team shirt image
func (p *Team) GetShirtMedium() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/img/shirts/standard/shirt_%d-110.png", p.Code)
}

// GetShirtLarge returns url to large verion of the team shirt image
func (p *Team) GetShirtLarge() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/img/shirts/standard/shirt_%d-220.png", p.Code)
}

func (t *Team) GetTeam(players *Players) map[int]Player {
	var squad = make(map[int]Player)

	for _, player := range players.Players {
		if player.TeamCode == t.Code {
			squad[player.Code] = player
		}
	}
	return squad
}

func (t *Team) GetSquad(players *Players) (goalkeepers, defenders, midfielders, forwards map[int]Player) {

	goalkeepers = make(map[int]Player)
	defenders = make(map[int]Player)
	midfielders = make(map[int]Player)
	forwards = make(map[int]Player)

	for _, player := range players.Players {
		if player.TeamCode == t.Code {
			// Goalkeepers
			if player.PlayerTypeID == 1 {
				goalkeepers[player.Code] = player
			}
			// Defenders
			if player.PlayerTypeID == 2 {
				defenders[player.Code] = player
			}
			// Midfielders
			if player.PlayerTypeID == 3 {
				midfielders[player.Code] = player
			}
			// Forwards
			if player.PlayerTypeID == 4 {
				forwards[player.Code] = player
			}
		}
	}
	return
}

//https://fantasy.premierleague.com/dist/img/shirts/standard/shirt_8-220.png
