package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
)

const (
	defaultCols  = 8
	defaultRows  = 8
	defaultMines = 10
	maxCols      = 16
	maxRows      = 30
	minMines     = 10
	maxMines     = 99
)

type GameServiceInterface interface {
	Create(game *domain.Game) error
	StartGame(game *domain.Game) (*domain.Game, error)
	Start(name string) (*domain.Game, error)
	Click(name string, i, j int64, flag bool) (*domain.Game, error)
}

type GameService struct {
	Store           memory.GameStoreInterface
	MovementService MovementService
}

func (s *GameService) Create(game *domain.Game) error {
	if game.Name == "" {
		return errors.New("no Game name")
	}

	if game.Rows == 0 {
		game.Rows = defaultRows
	}
	if game.Cols == 0 {
		game.Cols = defaultCols
	}
	if game.Mines == 0 {
		game.Mines = defaultMines
	}

	if game.Rows > maxRows {
		game.Rows = maxRows
	}
	if game.Cols > maxCols {
		game.Cols = maxCols
	}

	if game.Mines > maxMines {
		game.Mines = maxMines
	}

	if game.Mines < minMines {
		game.Mines = minMines
	}

	game.GameStatus = game_status.GameStatus{
		Alive:  true,
		Status: game_status.WaitingForStart,
	}

	err := s.Store.Insert(game)
	return err
}

func (s *GameService) StartGame(game *domain.Game) (*domain.Game, error) {
	logrus.WithFields(logrus.Fields{
		"file":      "game_controller",
		"service":   "start",
		"method":    "StartGame",
		"game_name": game.Name,
		"game_uuid": game.UUID,
	})

	game, err := s.Start(game.Name)
	return game, err
}

func (s *GameService) Start(name string) (*domain.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	BuildBoard(game)

	game.GameStatus.Status = game_status.Started

	err = s.Store.Update(game)

	grid, _ := json.Marshal(game.Grid)
	fmt.Printf("GameGrid: %s", string(grid))

	return game, err
}

func (s *GameService) Click(name string, pos *domain.CellPos) (*domain.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	grid, err := s.Store.GetByNameGrid(name)
	game.Grid = *grid
	if err != nil {
		return nil, errors.New("cannot get grid on click")

	}
	s.MovementService = MovementService{
		game: game,
		grid: grid,
		pos:  pos,
	}
	if err := s.MovementService.ClickCell(); err != nil {
		return nil, err
	}

	if err := s.Store.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}
