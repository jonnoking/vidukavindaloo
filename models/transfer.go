package models

import ()

//https://fantasy.premierleague.com/api/entry/4719576/transfers/

// EntryTransfers tbc
type EntryTransfers struct {
	Transfers []Transfer
}

// Transfer A team transfer
type Transfer struct {
	ElementIn      int `json:"element_in"`
	PlayerIn       Player
	ElementInCost  int `json:"element_in_cost"`
	ElementOut     int `json:"element_out"`
	PlayerOut      Player
	ElementOutCost int    `json:"element_out_cost"`
	Entry          int    `json:"entry"`
	Event          int    `json:"event"`
	Time           string `json:"time"`
}
