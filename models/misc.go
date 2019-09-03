package models

import (
//fmt
)

// boostrap

type PlayerTypes struct {
	Positions   map[int]PlayerType
	PlayerTypes []PlayerType
}

func (p *PlayerTypes) New() {

	p.Positions = map[int]PlayerType{}

	for _, player := range p.PlayerTypes {
		p.Positions[player.ID] = player
	}
}

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

type Phases struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StartEvent int    `json:"start_event"`
	StopEvent  int    `json:"stop_event"`
}

type Events struct {
	ID                     int              `json:"id"`
	Name                   string           `json:"name"`
	DeadlineTime           string           `json:"deadline_time"` //time.Time
	AverageEntryScore      int              `json:"average_entry_score"`
	Finished               bool             `json:"finished"`
	DataChecked            bool             `json:"data_checked"`
	HighestScoringEntry    int              `json:"highest_scoring_entry"`
	DeadlineTimeEpoch      int              `json:"deadline_time_epoch"`
	DeadlineTimeGameOffset int              `json:"deadline_time_epoch_offset"`
	HighestScore           int              `json:"highest_score"`
	IsPrevious             bool             `json:"in_previous"`
	IsCurrent              bool             `json:"is_current"`
	IsNext                 bool             `json:"is_next"`
	ChipPlays              []EventChipPlays `json:"chip_plays"`
	MostedSelected         int              `json:"most_selected"`
	MostTransferredIn      int              `json:"most_transferred_in"`
	TopPlayer              int              `json:"top_element"`
	TransfersMade          int              `json:"transfers_made"`
	MostCaptained          int              `json:"most_captained"`
	MostViceCaptained      int              `json:"most_vice_captained"`
}

type EventChipPlays struct {
	ChipPlayed string `json:"chip_played"`
	NumPlayed  int    `json:"num_played"`
}

var PlayerStats = map[string]string{
	"minutes":          "Minutes players",
	"goals_scored":     "Goals scored",
	"assists":          "Assists",
	"clean_sheets":     "Clean Sheets",
	"goals_conceded":   "Goals condeded",
	"own_goals":        "Own goals",
	"penalties_saved":  "Penalties saved",
	"penalties_missed": "Penalties missed",
	"yellow_cards":     "Yellow cards",
	"red_cards":        "Red cards",
	"saves":            "Saves",
	"bonus":            "Bonus",
	"bps":              "Bonus Points System",
	"influence":        "Influence",
	"creativity":       "Creativity",
	"threat":           "Threat",
	"ict_index":        "ICT Index",
}
