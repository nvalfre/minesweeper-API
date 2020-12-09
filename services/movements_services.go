package services

import (
	"errors"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"minesweeper-API/utils"
)

const maxPorcentageOfBoardOpened = 0.3
const base = 10
const initCount = 0

type MovementServiceInterface interface {
	ClickCell() error
	FlagClickCell() error
}
type MovementService struct {
	game                *domain.Game
	grid                *[]domain.CellGrid
	clickPosition       *domain.CellPos
	clickOpenCellsCount int64
}

func (s *MovementService) ClickCell() error {
	pos, err := s.getClickPosition()
	if err != nil {
		return err
	}
	if pos.Flagged {
		return errors.New("cell already flagged")
	}
	s.game.Clicks++
	pos.Clicked = true
	if pos.Mine {
		s.game.GameStatus.Status = game_status.Finished
		s.game.GameStatus.Won = game_status.Lose
		s.game.GameStatus.Alive = false
		return nil
	}
	s.clickOpenCellsCount = initCount
	s.openAdjacent(pos.Column, pos.Row)
	if s.isGameFinished() {
		s.game.GameStatus.Status = game_status.Finished
		s.game.GameStatus.Won = game_status.Won
	}

	return nil
}

func (s *MovementService) FlagClickCell() error {
	pos, err := s.getClickPosition()
	if err != nil {
		return err
	}

	if s.clickPosition.Flag {
		if pos.Flagged {
			s.game.Flags--
			pos.Flagged = false
			return nil
		}
		s.game.Flags++
		pos.Flagged = true
		return nil
	}
	if pos.Flagged {
		return errors.New("cell already flagged")
	}
	return errors.New("internal error")
}

func (s *MovementService) getClickPosition() (*domain.Cell, error) {
	if !utils.IsValidPosition(s.game, s.clickPosition.Row, s.clickPosition.Col) {
		return nil, errors.New("invalid position click")
	}
	pos := s.getCurrentPos()
	if pos.Covered {
		return nil, errors.New("cell already clicked and covered")
	}
	if pos.Opened {
		return nil, errors.New("cell already openAdjacent and free covered")
	}
	return pos, nil
}

func (s *MovementService) getCurrentPos() *domain.Cell {
	return &s.game.Grid[s.clickPosition.Row][s.clickPosition.Col]
}

func (s *MovementService) getPositionOffset(row, column int64) *domain.Cell {
	return &s.game.Grid[row][column]
}

func (s *MovementService) openAdjacent(row, column int64) {
	if maximum := (s.game.Rows * s.game.Cols) / (maxPorcentageOfBoardOpened * base); s.clickOpenCellsCount == maximum {
		return
	}

	if !utils.IsValidPosition(s.game, row, column) {
		return
	}
	pos := s.getPositionOffset(row, column)
	if pos.Covered || pos.Opened || pos.Mine {
		return
	}
	if s.isCurrentPosition(row, column) {
		pos.Covered = true
		s.game.CoveredCells++
	} else {
		pos.Opened = true
		s.game.OpenCells++
		s.clickOpenCellsCount++
	}

	if mineCount := utils.MineCount(s.game, *s.grid, s.getCurrentPos()); mineCount > 0 {
		s.getCurrentPos().AdjacentBombs = mineCount
		return
	}
	s.openAdjacent(row, column-1)
	s.openAdjacent(row, column+1)

	s.openAdjacent(row-1, column)
	s.openAdjacent(row+1, column)

	s.openAdjacent(row-1, column-1)
	s.openAdjacent(row+1, column+1)
	s.openAdjacent(row-1, column+1)
	s.openAdjacent(row+1, column-1)
}

func (s *MovementService) isGameFinished() bool {
	totalCells := s.game.Rows * s.game.Cols
	totalPlayed := s.game.CoveredCells + s.game.OpenCells + s.game.Flags

	playedWellDiff := totalCells - totalPlayed
	cellsWithoutMines := totalCells - s.game.Mines

	r1 := playedWellDiff == s.game.Mines
	r2 := cellsWithoutMines == totalPlayed
	return r1 || r2
}

func (s *MovementService) isCurrentPosition(column, row int64) bool {
	return column == s.clickPosition.Col && row == s.clickPosition.Row
}
