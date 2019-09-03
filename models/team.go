package models

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	s "strings"
)

// FPLTeams represents all teams in the FPL
type FPLTeams struct {
	Teams map[int]FPLTeam `json:"teams"`
}

func NewFPLTeams(teams []FPLTeam) (*FPLTeams, error) {
	ts := map[int]FPLTeam{}
	for _, team := range teams {
		ts[team.Code] = team
	}

	t := new(FPLTeams)
	t.Teams = ts

	return t, nil
}

func NewFPLTeamsFromBootStrap(filename string) (*FPLTeams, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	r, e := NewFPLTeamsFromBootStrapByteArray(f)

	return r, e
}

func NewFPLTeamsFromBootStrapByteArray(bootstrap []byte) (*FPLTeams, error) {

	var result map[string]interface{}
	json.Unmarshal([]byte(bootstrap), &result)

	r, e := NewFPLTeamsFromBootStrapMap(result)

	return r, e
}

func NewFPLTeamsFromBootStrapMap(bootstrap map[string]interface{}) (*FPLTeams, error) {

	ts := []FPLTeam{}
	teams := bootstrap["teams"].([]interface{})

	for _, v := range teams {
		var team FPLTeam
		mapstructure.Decode(v, &team)
		ts = append(ts, team)
	}

	r, e := NewFPLTeams(ts)

	return r, e
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

// FPLTeam represents a Premier League team
type FPLTeam struct {
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

// GetByCode returns a team
func (p *FPLTeam) GetByCode(code int) FPLTeam {
	t := FPLTeam{}
	return t
}

// func getTeamsFromCache() map[string]FPLTeam {

// }
