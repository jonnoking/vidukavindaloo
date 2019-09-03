package models

import (
	"encoding/json"
	"log"
)

func NewMyTeam(b []byte, players *Players, teams *Teams, positions *PlayerTypes) (*MyTeam, error) {

	myteam := MyTeam{}
	json.Unmarshal([]byte(b), &myteam)

	log.Println("Player Start", myteam.Picks[0].Player.FirstName)
	for _, p := range myteam.Picks {
		p.Player = players.PlayersByID[p.PlayerID]
		//log.Println(len(players.Players))
		//log.Println(players.PlayersByID[p.PlayerID])
		//log.Println(p.Player.FirstName)
		p.Team = teams.Teams[p.Player.TeamCode]
		log.Println(p.Team.Name)
		p.PlayerType = positions.Positions[p.Position]
	}
	log.Println("Player End", myteam.Picks[0].Player.FirstName)
	log.Println("Team End", myteam.Picks[0].Team.Name)

	return &myteam, nil
}
