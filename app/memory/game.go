package memory

import (
	"errors"
	"minesweeper-API/domain"
)

type GameStoreInterface interface {
	Insert(game *domain.Game) error
	Update(game *domain.Game) error
	GetByName(name string) (*domain.Game, error)
}

type GameStore struct {
	db *DB
}

func NewGameStore(db *DB) *GameStore {
	return &GameStore{db: db}
}

func (s *GameStore) Insert(game *domain.Game) error {
	if _, ok := s.db.games[game.Name]; ok {
		return errors.New("game already exist")
	}
	s.db.games[game.Name] = game
	return nil
}

func (s *GameStore) Update(game *domain.Game) error {
	g := *game
	if _, ok := s.db.games[game.Name]; !ok {
		return errors.New("game do not exist")
	}
	s.db.games[game.Name] = &g
	return nil
}

func (s *GameStore) GetByName(name string) (*domain.Game, error) {
	if game, ok := s.db.games[name]; ok {
		g := *game
		return &g, nil
	}
	return nil, errors.New("game not found")
}
