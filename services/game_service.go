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

type GameServiceInterface interface {
	History() (map[string]*domain.Game, error)
	StartGame(game *domain.Game) (*domain.Game, error)
	Start(name string) (*domain.Game, error)
	Click(name string, pos *domain.CellPos) (*domain.Game, error)
	FlagClick(name string, pos *domain.CellPos) (*domain.Game, error)
	Pause(name string) (*domain.Game, error)
	Resume(name string) (*domain.Game, error)
}

type GameService struct {
	Store           memory.GameStoreInterface
	MovementService MovementService
}

func (s *GameService) History() (map[string]*domain.Game, error) {
	logrus.WithFields(logrus.Fields{
		"file":    "game_controller",
		"service": "history",
		"method":  "GetHistory",
	})
	gameHistory, err := s.Store.GetAll()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "game_controller",
			"service": "history",
			"method":  "GetHistory",
			"err":     err,
			"message": "cannot get game history",
		})
		return nil, err
	}
	return gameHistory, err
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
		game:          game,
		grid:          grid,
		clickPosition: pos,
	}
	if err := s.MovementService.ClickCell(); err != nil {
		return nil, err
	}

	if err := s.Store.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *GameService) FlagClick(name string, pos *domain.CellPos) (*domain.Game, error) {
	if !pos.Flag {
		return nil, errors.New("request should be flagged")
	}
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
		game:          game,
		grid:          grid,
		clickPosition: pos,
	}
	if err := s.MovementService.FlagClickCell(); err != nil {
		return nil, err
	}

	if err := s.Store.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *GameService) Pause(name string) (*domain.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	game.GameStatus.Status = game_status.Paused
	if err := s.Store.Update(game); err != nil {
		return nil, err
	}
	return game, nil
}

func (s *GameService) Resume(name string) (*domain.Game, error) {
	game, err := s.Store.GetByName(name)
	if err != nil {
		return nil, err
	}

	game.GameStatus.Status = game_status.Started
	if err := s.Store.Update(game); err != nil {
		return nil, err
	}
	return game, nil
}
