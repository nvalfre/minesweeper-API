package memory

import "minesweeper-API/domain"

type DB struct {
	games map[string]*domain.Game
}

func New() *DB {
	return &DB{
		games: make(map[string]*domain.Game),
	}
}
