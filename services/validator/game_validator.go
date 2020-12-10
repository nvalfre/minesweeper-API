package validator

import (
	"minesweeper-API/domain"
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

type GameValidatorService struct{}

type GameValidatorInterface interface {
	Validate(game *domain.Game)
}

func (v *GameValidatorService) Validate(game *domain.Game) {
	validateRows(game)
	validateColumns(game)
	validateMines(game)
}

func validateMines(game *domain.Game) {
	if game.Mines < defaultMines {
		game.Mines = defaultMines
	}
	if game.Mines > maxMines {
		game.Mines = maxMines
	}
	if game.Mines < minMines {
		game.Mines = minMines
	}
}

func validateColumns(game *domain.Game) {
	if game.Cols < defaultCols {
		game.Cols = defaultCols
	}
	if game.Cols > maxCols {
		game.Cols = maxCols
	}
}

func validateRows(game *domain.Game) {
	if game.Rows < defaultRows {
		game.Rows = defaultRows
	}
	if game.Rows > maxRows {
		game.Rows = maxRows
	}
}
