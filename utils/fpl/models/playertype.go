package models

// PlayerTypes Types of players
type PlayerTypes struct {
	Positions   map[int]PlayerType
	PlayerTypes []PlayerType
}

// New Create a new PlayerType
func (p *PlayerTypes) New() {

	p.Positions = map[int]PlayerType{}

	for _, player := range p.PlayerTypes {
		p.Positions[player.ID] = player
	}
}

//PlayerType The type of player, defender, midfielder, forward, goalkeeper
type PlayerType struct {
	ID                 int    `json:"id"`
	PluralName         string `json:"plural_name"`
	PluralNameShort    string `json:"plural_name_short"`
	SingularName       string `json:"singular_name"`
	SinguarlNameShort  string `json:"singular_name_short"`
	SquadSelect        int    `json:"squad_select"`
	SquadMinPlay       int    `json:"squad_min_plan"`
	SquadMaxPlay       int    `json:"squad_max_plan"`
	UIShirtSpecific    bool   `json:"ui_shirt_specific"`
	SubPositionsLocked []int  `json:"sub_positions_locked"`
}
