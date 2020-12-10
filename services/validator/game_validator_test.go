package validator

import (
	"minesweeper-API/domain"
	"reflect"
	"testing"
)

var v = &GameValidatorService{}

func TestGameValidatorService_Validate(t *testing.T) {
	type args struct {
		game *domain.Game
	}
	tests := []struct {
		name        string
		args        args
		wantRows    int64
		wantColumns int64
		wantMines   int64
	}{
		{"test_validate_ok_min", args{game: &domain.Game{
			Rows:  3,
			Cols:  3,
			Mines: 2,
		}}, 8, 8, 10},
		{"test_validate_ok_def", args{game: &domain.Game{
			Rows:  10,
			Cols:  10,
			Mines: 20,
		}}, 10, 10, 20},
		{"test_validate_ok_max", args{game: &domain.Game{
			Rows:  99,
			Cols:  99,
			Mines: 999,
		}}, 30, 16, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v.Validate(tt.args.game)
			if !reflect.DeepEqual(tt.args.game.Rows, tt.wantRows) {
				t.Errorf("Validate() got = %v, want %v", tt.args.game.Rows, tt.wantRows)
			}
			if !reflect.DeepEqual(tt.args.game.Cols, tt.wantColumns) {
				t.Errorf("Validate() got = %v, want %v", tt.args.game.Cols, tt.wantColumns)
			}
			if !reflect.DeepEqual(tt.args.game.Mines, tt.wantMines) {
				t.Errorf("Validate() got = %v, want %v", tt.args.game.Mines, tt.wantMines)
			}
		})
	}
}
