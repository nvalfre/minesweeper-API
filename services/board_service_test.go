package services

import (
	"github.com/stretchr/testify/assert"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"minesweeper-API/domain/mock"
	"testing"
)

func TestBuildBoard(t *testing.T) {
	type args struct {
		game *domain.Game
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_build_board:ok", args{game: &domain.Game{
			Rows:       3,
			Cols:       2,
			Mines:      2,
			Grid:       nil,
			GameStatus: game_status.GameStatus{Status: game_status.Started},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BuildBoard(tt.args.game)
			assert.NotNil(t, tt.args.game.Grid)
			assert.Equal(t, tt.args.game.Rows, int64(len(tt.args.game.Grid)))
			assert.Equal(t, tt.args.game.Cols, int64(len(tt.args.game.Grid[0])))
		})
	}
}

func Test_generateRandomPositions(t *testing.T) {
	type args struct {
		game *domain.Game
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_generate_random_positions:ok", args{game: &domain.Game{
			Rows:       3,
			Cols:       2,
			Mines:      2,
			Grid:       nil,
			GameStatus: game_status.GameStatus{Status: game_status.Started},
		}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateRandomPositions(tt.args.game)
			assert.NotNil(t, tt.args.game.Grid[0])
			assert.NotNil(t, tt.args.game.Grid[0][0])
			assert.NotNil(t, tt.args.game.Grid[0][1])
			assert.NotNil(t, tt.args.game.Grid[1])
			assert.NotNil(t, tt.args.game.Grid[1][0])
			assert.NotNil(t, tt.args.game.Grid[1][1])
			assert.NotNil(t, tt.args.game.Grid[2])
			assert.NotNil(t, tt.args.game.Grid[2][0])
			assert.NotNil(t, tt.args.game.Grid[2][1])
		})
	}
}

func Test_setCellValues(t *testing.T) {
	type args struct {
		game *domain.Game
	}
	tests := []struct {
		name string
		args args
	}{
		{"test_set_cell_value:ok", args{game: &domain.Game{
			Rows:       3,
			Cols:       2,
			Mines:      2,
			Grid:       mock.Grid,
			GameStatus: game_status.GameStatus{Status: game_status.Started},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setCellValues(tt.args.game)
			assert.Equal(t, tt.args.game.Grid[0][0].Row, int64(0))
			assert.Equal(t, tt.args.game.Grid[0][1].Row, int64(0))
			assert.Equal(t, tt.args.game.Grid[1][0].Row, int64(1))
			assert.Equal(t, tt.args.game.Grid[1][1].Row, int64(1))
			assert.Equal(t, tt.args.game.Grid[2][0].Row, int64(2))
			assert.Equal(t, tt.args.game.Grid[2][1].Row, int64(2))
			assert.Equal(t, tt.args.game.Grid[0][0].Column, int64(0))
			assert.Equal(t, tt.args.game.Grid[0][1].Column, int64(1))
			assert.Equal(t, tt.args.game.Grid[1][0].Column, int64(0))
			assert.Equal(t, tt.args.game.Grid[1][1].Column, int64(1))
			assert.Equal(t, tt.args.game.Grid[2][0].Column, int64(0))
			assert.Equal(t, tt.args.game.Grid[2][1].Column, int64(1))
		})
	}
}
