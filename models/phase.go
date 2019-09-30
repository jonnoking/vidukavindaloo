package models

// Phases all the phases
type Phases struct {
	PhasesArray []Phase
	Phases      map[int]Phase
}

// Phase the phase
type Phase struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StartEvent int    `json:"start_event"`
	StopEvent  int    `json:"stop_event"`
}
