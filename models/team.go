package models

import (
	"fmt"
	s "strings"
)

// FPLTeams represents all teams in the FPL
type FPLTeams struct {
	Teams map[int]FPLTeam `mapstructure:"teams"`
}

func (p *FPLTeams) New(teams []FPLTeam) {
	ts := map[int]FPLTeam{}
	for _, team := range teams {
		ts[team.Code] = team
	}
	p.Teams = ts
}

func (p *FPLTeams) GetTeamByName(name string) (FPLTeam, error) {
	var ret FPLTeam

	for _, team := range p.Teams {
		if s.ToLower(team.Name) == s.ToLower(name) {
			return team, nil
		}
	}
	return ret, fmt.Errorf("No team called %s found", name)
}

func (p *FPLTeams) GetTeamByCode(code int) (FPLTeam, error) {
	var ret FPLTeam

	team, found := p.Teams[code]
	if !found {
		return ret, fmt.Errorf("No team found with code %s", code)
	}

	return team, nil
}

// FPLTeam represents a Premier League team
type FPLTeam struct {
	Code                int    `mapstructure:"code"`
	Draw                int    `mapstructure:"draw"`
	Form                int    `mapstructure:"form"`
	ID                  int    `mapstructure:"id"`
	Lost                int    `mapstructure:"loss"`
	Name                string `mapstructure:"name"`
	Played              int    `mapstructure:"played"`
	Points              int    `mapstructure:"points"`
	Position            int    `mapstructure:"position"`
	ShortName           string `mapstructure:"short_name"`
	Strength            int    `mapstructure:"strength"`
	TeamDivision        int    `mapstructure:"team_division"`
	Unavailable         bool   `mapstructure:"unavailable"`
	Win                 int    `mapstructure:"win"`
	StrengthOverallHome int    `mapstructure:"strength_overall_home"`
	StrengthOverallAway int    `mapstructure:"strength_overall_away"`
	StrengthAttackHome  int    `mapstructure:"strength_attack_home"`
	StrengthAttackAway  int    `mapstructure:"strength_attack_away"`
	StrengthDefenceHome int    `mapstructure:"strength_defence_home"`
	StrengthDefenceAway int    `mapstructure:"strength_defence_away"`
}

// GetShirtSmall returns url to small verion of the team shirt image
func (p *FPLTeam) GetShirtSmall() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/image/shirts/shirt_%d-66.png", p.Code)
}

// GetShirtMedium returns url to medium verion of the team shirt image
func (p *FPLTeam) GetShirtMedium() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/image/shirts/shirt_%d-110.png", p.Code)
}

// GetShirtLarge returns url to large verion of the team shirt image
func (p *FPLTeam) GetShirtLarge() string {
	return fmt.Sprintf("https://fantasy.premierleague.com/dist/image/shirts/shirt_%d-220.png", p.Code)
}
