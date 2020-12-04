package utils

import "minesweeper-API/domain"

func MineCount(game *domain.Game, grid []domain.CellGrid, pos *domain.Cell) int {
	var count int
	count += isBomb(game, grid, pos.Column-1, pos.Row)
	count += isBomb(game, grid, pos.Column+1, pos.Row)

	count += isBomb(game, grid, pos.Column, pos.Row-1)
	count += isBomb(game, grid, pos.Column, pos.Row+1)

	count += isBomb(game, grid, pos.Column-1, pos.Row-1)
	count += isBomb(game, grid, pos.Column+1, pos.Row+1)
	count += isBomb(game, grid, pos.Column+1, pos.Row-1)
	count += isBomb(game, grid, pos.Column-1, pos.Row+1)
	return count
}

func isBomb(game *domain.Game, grid []domain.CellGrid, column, row int64) int {
	isValidPositionInGrid := IsValidPosition(game, column, row)
	if isValidPositionInGrid && grid[column][row].Mine {
		return 1
	}
	return 0
}

func IsValidPosition(game *domain.Game, x, y int64) bool {
	return x >= 0 && x < game.Cols && y >= 0 && y < game.Rows
}
