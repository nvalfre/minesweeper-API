package utils

import "minesweeper-API/domain"

func MineCount(game *domain.Game, grid []domain.CellGrid, pos *domain.Cell) int {
	var count int
	count += isBomb(game, grid, pos.Row, pos.Column-1)
	count += isBomb(game, grid, pos.Row, pos.Column+1)

	count += isBomb(game, grid, pos.Row-1, pos.Column)
	count += isBomb(game, grid, pos.Row+1, pos.Column)

	count += isBomb(game, grid, pos.Row-1, pos.Column-1)
	count += isBomb(game, grid, pos.Row+1, pos.Column+1)
	count += isBomb(game, grid, pos.Row-1, pos.Column+1)
	count += isBomb(game, grid, pos.Row+1, pos.Column-1)
	return count
}

func isBomb(game *domain.Game, grid []domain.CellGrid, row, column int64) int {
	isValidPositionInGrid := IsValidPosition(game, row, column)
	if isValidPositionInGrid && grid[row][column].Mine {
		return 1
	}
	return 0
}

func IsValidPosition(game *domain.Game, row, column int64) bool {
	return row >= 0 && row < game.Rows && column >= 0 && column < game.Cols
}
