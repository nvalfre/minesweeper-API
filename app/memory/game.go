package memory

import (
	"errors"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
)

type GameStoreInterface interface {
	Insert(game *domain.Game) error
	Update(game *domain.Game) error
	GetAll() (map[string]*domain.Game, error)
	GetByName(name string) (*domain.Game, error)
	insertGrid(game *domain.Game) error
	updateGrid(game *domain.Game) error
	GetByNameGrid(name string) (*[]domain.CellGrid, error)
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
	if g.Clicks == 0 {
		if err := s.insertGrid(&g); err != nil {
			return errors.New("can't insert grid on memory")
		} else {
			return nil
		}
	}
	if err := s.updateGrid(&g); err != nil {
		return errors.New("can't update grid on memory")
	}
	return nil
}

func (s *GameStore) GetByName(name string) (*domain.Game, error) {
	if game, ok := s.db.games[name]; ok {
		g := *game
		if g.GameStatus.Status != game_status.WaitingForStart {
			grid, err := s.GetByNameGrid(game.Name)
			if err != nil {
				return nil, errors.New("grid1 not found")
			}
			g.Grid = *grid
		}
		return &g, nil
	}
	return nil, errors.New("game not found")
}

func (s *GameStore) GetAll() (map[string]*domain.Game, error) {
	return s.db.games, nil
}

func (s *GameStore) insertGrid(game *domain.Game) error {
	var grid *[]domain.CellGrid
	if _, ok := s.db.grids[game.Name]; ok {
		return errors.New("grid already exist")
	}
	grid = &game.Grid
	s.db.grids[game.Name] = grid
	return nil
}

func (s *GameStore) updateGrid(game *domain.Game) error {
	g := *game
	if _, ok := s.db.grids[game.Name]; !ok {
		return errors.New("grid do not exist")
	}
	s.db.grids[game.Name] = &g.Grid
	return nil
}

func (s *GameStore) GetByNameGrid(name string) (*[]domain.CellGrid, error) {
	if grid, ok := s.db.grids[name]; ok {
		g := *grid
		return &g, nil
	}
	return nil, errors.New("grid not found")
}
