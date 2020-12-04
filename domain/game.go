package domain

import (
	"minesweeper-API/domain/game_status"
	"time"
)

type Cell struct {
	Column int64 `json:"column"`
	Row    int64 `json:"row"`

	StartingAdjacentBombs int `json:"value"`
	AdjacentBombs         int `json:"adjacent_bombs"`

	Mine bool `json:"mine"`

	Clicked bool `json:"clicked"`
	Flagged bool `json:"flagged"`
	Opened  bool `json:"opened"`
	Covered bool `json:"covered"`
}

type CellGrid []Cell

type Game struct {
	Timer        *time.Timer            `json:"-"`
	Name         string                 `json:"name"`
	UUID         string                 `json:"uuid"`
	Rows         int64                  `json:"rows"`
	Cols         int64                  `json:"cols"`
	Mines        int64                  `json:"mines"`
	Grid         []CellGrid             `json:"grid,omitempty"`
	Clicks       int64                  `json:"clicks"`
	Flags        int64                  `json:"flags"`
	OpenCells    int64                  `json:"open_cells"`
	CoveredCells int64                  `json:"covered_cells"`
	GameStatus   game_status.GameStatus `json:"-"`
}

type ClickResult struct {
	Cell       Cell
	Game       Game
	GameStatus game_status.GameStatus
}

type CellPos struct {
	Row  int64 `json:"row"`
	Col  int64 `json:"column"`
	Flag bool  `json:"flag"`
}
