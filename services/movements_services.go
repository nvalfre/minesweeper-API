package services

import (
	"errors"
	"minesweeper-API/domain"
)

const finished = "finished: won"
const finishedLose = "finished: lose"

func ClickCell(game *domain.Game, i, j int64, flag bool) error {
	if game.Grid[i][j].Clicked {
		return errors.New("cell already clicked")
	}
	if flag {
		if game.Grid[i][j].Flagged {
			game.Flags -= 1
			game.Grid[i][j].Flagged = false
			return nil
		}
		game.Flags += 1
		game.Grid[i][j].Flagged = true
		return nil
	}
	game.Grid[i][j].Clicked = true
	if game.Grid[i][j].Mine {
		game.Status = finishedLose
		return nil
	}
	game.Clicks += 1
	if isGameFinished(game) {
		game.Status = finished
	}

	return nil
}

func isGameFinished(game *domain.Game) bool {
	return game.Clicks == ((game.Rows * game.Cols) - game.Mines)
}
