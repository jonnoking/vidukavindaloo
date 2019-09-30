package models

import (
//fmt
)

// /api/my-team

// MyTeam Object returned via /api/my-team/{team-id}
type MyTeam struct {
	Picks     []Picks   `json:"picks"`
	Chips     []Chips   `json:"chips"`
	Transfers Transfers `json:"transfers"`
}

// Transfers Status of my team transfers
type Transfers struct {
	Cost   int    `json:"cost"`
	Status string `json:"status"`
	Limit  int    `json:"limit"`
	Made   int    `json:"made"`
	Bank   int    `json:"bank"`
	Value  int    `json:"value"`
}

// Chips Status of chips of my team
type Chips struct {
	StatusForEntry string `json:"status_for_entry"`
	PlayedByEntry  []int  `json:"played_by_entry"`
	Name           string `json:"name"`
	Number         int    `json:"number"`
	StartEvent     int    `json:"start_event"`
	StopEvent      int    `json:"stop_event"`
	ChipType       string `json:"chip_type"`
}

// Picks Players selected in my team
type Picks struct {
	ElementID     int `json:"element"`
	Player        Player
	Team          Team
	PlayerType    PlayerType
	Position      int  `json:"position"`
	SellingPrice  int  `json:"selling_price"`
	Multiplier    int  `json:"multiplier"`
	PurhcasePrice int  `json:"purchase_price"`
	IsCaptain     bool `json:"is_captain"`
	IsViceCaptain bool `json:"is_vice_captain"`
}
