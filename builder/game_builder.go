package builder

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"time"
)

const (
	defaultCols         = 8
	defaultRows         = 8
	defaultMines        = 10
	maxCols             = 16
	maxRows             = 30
	minMines            = 10
	maxMines            = 99
	defaultDurationTime = 10 * time.Minute
	defaultStartClick   = 0
)

type GameBuilderService struct {
	Store memory.GameStoreInterface
}

type GameServiceInterface interface {
	BuildNewGame(name string, rows, columns, mines int64) (*domain.Game, error)
	create(game *domain.Game) error
}

func (s *GameBuilderService) create(game *domain.Game) error {
	if game.Name == "" {
		return errors.New("no Game name")
	}

	if game.Rows < defaultRows {
		game.Rows = defaultRows
	}
	if game.Cols < defaultCols {
		game.Cols = defaultCols
	}
	if game.Mines < defaultMines {
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
func (s *GameBuilderService) BuildNewGame(name string, rows, columns, mines int64) (*domain.Game, error) {
	game := &domain.Game{
		Timer:  time.NewTimer(defaultDurationTime),
		Name:   name,
		UUID:   fmt.Sprintf("%v", uuid.New()),
		Rows:   rows,
		Cols:   columns,
		Mines:  mines,
		Grid:   nil,
		Clicks: defaultStartClick,
	}
	logrus.WithFields(logrus.Fields{
		"file":    "game_controller",
		"service": "create",
		"method":  "StartNewGame",
		"game":    game,
	})
	err := s.create(game)
	if err != nil {
		return nil, err
	}
	return game, err
}
