package models

import ()
import "github.com/jonnoking/vidukavindaloo/utils/fpl"

//https://fantasy.premierleague.com/api/entry/4719576/transfers/

// EntryTransfers tbc
type EntryTransfersMap struct {
	Transfers map[string]*Transfer `json:"transfers"`
}

// EntryTransfers tbc
type EntryTransfers struct {
	Transfers []Transfer `json:"transfers"`
}

// Transfer A team transfer
type Transfer struct {
	ElementIn      int    `json:"element_in"`
	ElementInCost  int    `json:"element_in_cost"`
	ElementOut     int    `json:"element_out"`
	ElementOutCost int    `json:"element_out_cost"`
	Entry          int    `json:"entry"`
	Event          int    `json:"event"`
	Time           string `json:"time"`
}

func (t *Transfer) getPlayerIn() Player {
	return fpl.Players.PlayersByID[t.ElementIn]
}
