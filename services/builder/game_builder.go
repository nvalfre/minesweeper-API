package builder

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"minesweeper-API/services/validator"
	"time"
)

const (
	defaultDurationTime = 10 * time.Minute
	defaultStartClick   = 0
)

type GameBuilderService struct {
	Store     memory.GameStoreInterface
	Validator validator.GameValidatorService
}

type GameBuilderServiceInterface interface {
	BuildNewGame(name string, rows, columns, mines int64) (*domain.Game, error)
	create(game *domain.Game) error
}

func (s *GameBuilderService) BuildNewGame(name string, rows, columns, mines int64) (*domain.Game, error) {
	game := &domain.Game{
		Timer:  time.NewTimer(defaultDurationTime),
		Time:   time.Now().Add(defaultDurationTime),
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

func (s *GameBuilderService) create(game *domain.Game) error {
	if game.Name == "" {
		return errors.New("no Game name")
	}

	s.Validator.Validate(game)

	game.GameStatus = game_status.GameStatus{
		Alive:  true,
		Status: game_status.WaitingForStart,
	}

	err := s.Store.Insert(game)
	return err
}
