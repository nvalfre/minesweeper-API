package services

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
)

const (
	defaultRows  = 10
	defaultCols  = 10
	defaultMines = 12
	maxRows      = 30
	maxCols      = 30

	gameCreatedStatus  = "Created"
	gameStartedStatus  = "Started"
	gameFinishedStatus = "Finished"
)

type GameServiceInterface interface {
	Create(game *domain.Game) error
	StartGame(game *domain.Game) (*domain.Game, error)
	Start(name string) (*domain.Game, error)
	Click(name string, i, j int64, flag bool) (*domain.Game, error)
}

type GameService struct {
	Store memory.GameStoreInterface
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
	if game.Mines > (game.Cols * game.Rows) {
		game.Mines = (game.Cols * game.Rows)
	}
	game.Status = gameCreatedStatus

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

	game.Status = gameStartedStatus

	err = s.Store.Update(game)
	fmt.Printf("GameGrid: %v", game.Grid)
	game.Grid = nil

	return game, err
}

func (s *GameService) Click(name string, i, j int64, flag bool) (*domain.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	if err := ClickCell(game, i, j, flag); err != nil {
		return nil, err
	}

	if err := s.Store.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}
