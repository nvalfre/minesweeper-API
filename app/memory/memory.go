package memory

import "minesweeper-API/domain"

type DB struct {
	games map[string]*domain.Game
	grids map[string]*[]domain.CellGrid
}

func New() *DB {
	return &DB{
		games: make(map[string]*domain.Game),
		grids: make(map[string]*[]domain.CellGrid),
	}
}
