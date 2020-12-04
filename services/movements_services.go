package services

import (
	"errors"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
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
	s.game.Clicks++
	s.getCurrentPos().Clicked = true
	if s.getCurrentPos().Covered {
		return errors.New("cell already clicked and covered")
	}
	if s.getCurrentPos().Opened {
		return errors.New("cell already open and free covered")
	}
	if s.pos.Flag {
		if s.getCurrentPos().Flagged {
			s.game.Flags -= 1
			s.getCurrentPos().Flagged = false
			return nil
		}
		s.game.Flags += 1
		s.getCurrentPos().Flagged = true
		return nil
	}
	if s.getCurrentPos().Mine {
		s.game.GameStatus.Status = game_status.Finished
		s.game.GameStatus.Won = game_status.Lose
		return nil
	}
	s.getCurrentPos().Covered = true
	s.game.CoveredCells++
	s.open(invalidColumn, invalidColumn)
	if s.isGameFinished() {
		s.game.GameStatus.Status = game_status.Finished
		s.game.GameStatus.Won = game_status.Won
	}

	return nil
}

func (s *MovementService) getCurrentPos() *domain.Cell {
	return &s.game.Grid[s.pos.Col][s.pos.Row]
}

func (s *MovementService) open(column, row int) {
	if !s.isValidPosition(s.getCurrentPos().Column, s.getCurrentPos().Row) {
		return
	}
	if s.getCurrentPos().Opened {
		return
	}
	s.game.OpenCells++
	s.getCurrentPos().Opened = true

	if s.mineCount() > 0 {
		return
	}
	if column > invalidColumn && row > invalidColumn {
		s.open(column-1, row)
		s.open(column+1, row)

		s.open(column, row-1)
		s.open(column, row+1)

		s.open(column-1, row-1)
		s.open(column+1, row+1)
		s.open(column+1, row-1)
		s.open(column-1, row+1)
	}
}

func (s *MovementService) mineCount() int {
	s.isBomb(s.getCurrentPos().Column-1, s.getCurrentPos().Row)
	s.isBomb(s.getCurrentPos().Column+1, s.getCurrentPos().Row)

	s.isBomb(s.getCurrentPos().Column, s.getCurrentPos().Row-1)
	s.isBomb(s.getCurrentPos().Column, s.getCurrentPos().Row+1)

	s.isBomb(s.getCurrentPos().Column-1, s.getCurrentPos().Row-1)
	s.isBomb(s.getCurrentPos().Column+1, s.getCurrentPos().Row+1)
	s.isBomb(s.getCurrentPos().Column+1, s.getCurrentPos().Row-1)
	s.isBomb(s.getCurrentPos().Column-1, s.getCurrentPos().Row+1)
	return s.getCurrentPos().AdjacentBombs
}

func (s *MovementService) isBomb(column, row int64) {
	isValidPositionInGrid := s.isValidPosition(column, row)
	grid := *s.grid
	if isValidPositionInGrid && grid[column][row].Mine {
		s.getCurrentPos().AdjacentBombs++
	}
}

func (s *MovementService) isValidPosition(x, y int64) bool {
	return x >= 0 && x < s.game.Cols && y >= 0 && y < s.game.Rows
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
