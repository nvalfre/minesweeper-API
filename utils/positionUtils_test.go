package utils

import (
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"minesweeper-API/domain/mock"
	"testing"
)

var (
	Grid          = []domain.CellGrid{{mock.MockCell0, mock.MockCell01}, {mock.MockCell11, mock.MockCell12}, {mock.MockCell21, mock.MockCell22}}
	GridWithBomb  = []domain.CellGrid{{mock.MockCell0, mock.MockCell01}, {mock.MockCell11, mock.MockCell12}, {mock.MockCell21bomb, mock.MockCell22}}
	GridWithBomb3 = []domain.CellGrid{{mock.MockCell0, mock.MockCell01Bomb}, {mock.MockCell11, mock.MockCell12Bomb}, {mock.MockCell21bomb, mock.MockCell22}}

	Game = domain.Game{
		Rows:       3,
		Cols:       2,
		Mines:      2,
		Grid:       Grid,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
	GameBomb = domain.Game{
		Rows:       3,
		Cols:       2,
		Mines:      2,
		Grid:       GridWithBomb,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
	GameBomb3 = domain.Game{
		Rows:       3,
		Cols:       2,
		Mines:      2,
		Grid:       GridWithBomb3,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
)

func TestIsValidPosition(t *testing.T) {
	type args struct {
		game   *domain.Game
		row    int64
		column int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test_valid_position", args{
			game:   &mock.Game,
			row:    0,
			column: 0,
		}, true},
		{"test_NOT_valid_position", args{
			game:   &mock.Game,
			row:    3,
			column: 0,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidPosition(tt.args.game, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("IsValidPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMineCount(t *testing.T) {
	type args struct {
		game *domain.Game
		grid []domain.CellGrid
		pos  *domain.Cell
	}
	posEsc1 := &domain.Cell{
		Row:    1,
		Column: 1,
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_mine_count:ok", args{
			game: &GameBomb,
			grid: GridWithBomb,
			pos:  posEsc1,
		}, 1},
		{"test_mine_count:ok", args{
			game: &GameBomb3,
			grid: GridWithBomb3,
			pos:  posEsc1,
		}, 2},
		{"test_mine_count:0", args{
			game: &Game,
			grid: Grid,
			pos:  posEsc1,
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MineCount(tt.args.game, tt.args.grid, tt.args.pos); got != tt.want {
				t.Errorf("MineCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBomb(t *testing.T) {
	type args struct {
		game   *domain.Game
		grid   []domain.CellGrid
		row    int64
		column int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_is_bomb", args{
			game:   &GameBomb3,
			grid:   GridWithBomb3,
			row:    1,
			column: 1,
		}, 1},
		{"test_is_not_bomb", args{
			game:   &GameBomb3,
			grid:   GridWithBomb3,
			row:    1,
			column: 2,
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBomb(tt.args.game, tt.args.grid, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("isBomb() = %v, want %v", got, tt.want)
			}
		})
	}
}
