package domain

import (
	"time"
)

type Cell struct {
	Mine    bool `json:"mine"`
	Clicked bool `json:"clicked"`
	Flagged bool `json:"flagged"`
	Value   int  `json:"value"`
}

type CellGrid []Cell

type Game struct {
	Timer  *time.Timer `json:"-"`
	Name   string      `json:"name"`
	UUID   string      `json:"uuid"`
	Rows   int64       `json:"rows"`
	Cols   int64       `json:"cols"`
	Mines  int64       `json:"mines"`
	Status string      `json:"status"`
	Grid   []CellGrid  `json:"grid,omitempty"`
	Clicks int64       `json:"-"`
	Flags  int64       `json:"-"`
}

type ClickResult struct {
	Cell Cell
	Game Game
}
