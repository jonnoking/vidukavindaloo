package models

//Events Array of Events
type Events struct {
	Events map[int]Event
}

// Event The event week
type Event struct {
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

// EventChipPlays The chips played in an event
type EventChipPlays struct {
	ChipPlayed string `json:"chip_played"`
	NumPlayed  int    `json:"num_played"`
}
