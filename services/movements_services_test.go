package services

import (
	"minesweeper-API/domain"
	"minesweeper-API/domain/mock"
	"reflect"
	"testing"
)

func TestMovementService_getCurrentPos(t *testing.T) {
	type fields struct {
		game                *domain.Game
		grid                *[]domain.CellGrid
		clickPosition       *domain.CellPos
		clickOpenCellsCount int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.Cell
	}{
		{"get_current_pos:Ok", fields{
			game: &mock.Game,
			grid: &mock.Grid,
			clickPosition: &domain.CellPos{
				Row:  1,
				Col:  1,
				Flag: false,
			},
			clickOpenCellsCount: 0,
		}, &mock.MockCell12},
		{"get_current_pos:Ok", fields{
			game: &mock.Game,
			grid: &mock.Grid,
			clickPosition: &domain.CellPos{
				Row:  1,
				Col:  1,
				Flag: false,
			},
			clickOpenCellsCount: 0,
		}, &mock.MockCell12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MovementService{
				game:                tt.fields.game,
				grid:                tt.fields.grid,
				clickPosition:       tt.fields.clickPosition,
				clickOpenCellsCount: tt.fields.clickOpenCellsCount,
			}
			if got := s.getCurrentPos(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCurrentPos() = %v, want22 %v", got, tt.want)
			}
		})
	}
}

func TestMovementService_getPositionOffset(t *testing.T) {
	type fields struct {
		game                *domain.Game
		grid                *[]domain.CellGrid
		clickPosition       *domain.CellPos
		clickOpenCellsCount int64
	}
	type args struct {
		column int64
		row    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.Cell
	}{
		{"get_position_offset: Ok", fields{
			game:                &mock.Game,
			grid:                &mock.Grid,
			clickPosition:       &domain.CellPos{Row: 1, Col: 1, Flag: false},
			clickOpenCellsCount: 0,
		}, args{
			column: 0,
			row:    0,
		}, &mock.MockCell0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MovementService{
				game:                tt.fields.game,
				grid:                tt.fields.grid,
				clickPosition:       tt.fields.clickPosition,
				clickOpenCellsCount: tt.fields.clickOpenCellsCount,
			}
			if got := s.getPositionOffset(tt.args.column, tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPositionOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovementService_openAdjacent(t *testing.T) {
	type fields struct {
		game                *domain.Game
		grid                *[]domain.CellGrid
		clickPosition       *domain.CellPos
		clickOpenCellsCount int64
	}
	type args struct {
		column int64
		row    int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus string
	}{
		{"openAdjacent: Bomb", fields{
			game:          &mock.GameBomb,
			grid:          &mock.GridWithBomb,
			clickPosition: &domain.CellPos{Row: 1, Col: 1, Flag: false},
		}, args{
			column: 0,
			row:    0}, "Game in progress",
		},
		{"openAdjacent: Ok", fields{
			game:          &mock.Game,
			grid:          &mock.Grid,
			clickPosition: &domain.CellPos{Row: 1, Col: 1, Flag: false},
		}, args{
			column: 0,
			row:    0}, "Game in progress",
		},
		{"openAdjacent: InvalidPosition", fields{
			game:          &mock.Game,
			grid:          &mock.Grid,
			clickPosition: &domain.CellPos{Row: 1, Col: 1, Flag: false},
		}, args{
			column: 0,
			row:    0}, "Game in progress",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MovementService{
				game:                tt.fields.game,
				grid:                tt.fields.grid,
				clickPosition:       tt.fields.clickPosition,
				clickOpenCellsCount: tt.fields.clickOpenCellsCount,
			}
			s.openAdjacent(tt.fields.clickPosition.Row, tt.fields.clickPosition.Col)
			if s.game.GameStatus.Status != tt.wantStatus {
				t.Errorf("openAdjacent() = %v, want %v", s.game.GameStatus.Status, tt.wantStatus)
			}

		})
	}
}

func TestMovementService_isCurrentPosition(t *testing.T) {
	type fields struct {
		game                *domain.Game
		grid                *[]domain.CellGrid
		clickPosition       *domain.CellPos
		clickOpenCellsCount int64
	}
	type args struct {
		column int64
		row    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"test_is_current_position", fields{
			game: &mock.Game,
			grid: &mock.Grid,
			clickPosition: &domain.CellPos{
				Row: 1,
				Col: 1,
			},
		}, args{
			column: 1,
			row:    1,
		}, true},
		{"test_is_not_current_position", fields{
			game: &mock.Game,
			grid: &mock.Grid,
			clickPosition: &domain.CellPos{
				Row: 1,
				Col: 1,
			},
		}, args{
			column: 2,
			row:    1,
		}, false},
		{"test_is_not_2_current_position", fields{
			game: &mock.Game,
			grid: &mock.Grid,
			clickPosition: &domain.CellPos{
				Row: 1,
				Col: 1,
			},
		}, args{
			column: 2,
			row:    2,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MovementService{
				game:                tt.fields.game,
				grid:                tt.fields.grid,
				clickPosition:       tt.fields.clickPosition,
				clickOpenCellsCount: tt.fields.clickOpenCellsCount,
			}
			if got := s.isCurrentPosition(tt.args.column, tt.args.row); got != tt.want {
				t.Errorf("isCurrentPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
