package models

import (
	"encoding/json"
)

func NewMyTeam(b []byte, players *Players, teams *Teams, positions *PlayerTypes) (*MyTeam, error) {

	myteam := MyTeam{}
	json.Unmarshal([]byte(b), &myteam)

	for i := 0; i < len(myteam.Picks); i++ {
		myteam.Picks[i].Player = players.PlayersByID[myteam.Picks[i].PlayerID]
		myteam.Picks[i].Team = teams.Teams[myteam.Picks[i].Player.TeamCode]
	}

	return &myteam, nil
}
