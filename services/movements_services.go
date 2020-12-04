package services

import (
	"errors"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"minesweeper-API/utils"
)

const invalidColumn = -1

type MovementServiceInterface interface {
	ClickCell() error
}
type MovementService struct {
	game *domain.Game
	grid *[]domain.CellGrid
	pos  *domain.CellPos
}

func (s *MovementService) ClickCell() error {
	pos := s.getCurrentPos()
	s.game.Clicks++
	pos.Clicked = true
	if pos.Covered {
		return errors.New("cell already clicked and covered")
	}
	if pos.Opened {
		return errors.New("cell already open and free covered")
	}
	if s.pos.Flag {
		if pos.Flagged {
			s.game.Flags -= 1
			pos.Flagged = false
			return nil
		}
		s.game.Flags += 1
		pos.Flagged = true
		return nil
	}
	if pos.Mine {
		s.game.GameStatus.Status = game_status.Finished
		s.game.GameStatus.Won = game_status.Lose
		return nil
	}
	s.open(pos.Column, pos.Row)
	if s.isGameFinished() {
		s.game.GameStatus.Status = game_status.Finished
		s.game.GameStatus.Won = game_status.Won
	}

	return nil
}

func (s *MovementService) getCurrentPos() *domain.Cell {
	return &s.game.Grid[s.pos.Col][s.pos.Row]
}

func (s *MovementService) getPos(column, row int64) *domain.Cell {
	return &s.game.Grid[column][row]
}

func (s *MovementService) open(column, row int64) {
	if !utils.IsValidPosition(s.game, column, row) {
		return
	}
	pos := s.getPos(column, row)
	if pos.Covered || pos.Opened || pos.Mine {
		return
	}

	if s.isCurrentPosition(column, row) {
		pos.Covered = true
		s.game.CoveredCells++
	} else {
		pos.Opened = true
		s.game.OpenCells++
	}

	if mineCount := utils.MineCount(s.game, *s.grid, s.getCurrentPos()); mineCount > 0 {
		s.getCurrentPos().AdjacentBombs = mineCount
		return
	}
	s.open(column-1, row)
	s.open(column+1, row)

	s.open(column, row-1)
	s.open(column, row+1)

	s.open(column-1, row-1)
	s.open(column+1, row+1)
	s.open(column+1, row-1)
	s.open(column-1, row+1)
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
	return column == s.pos.Col && row == s.pos.Row
}
