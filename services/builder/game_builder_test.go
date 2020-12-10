package builder

import (
	"github.com/stretchr/testify/assert"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"minesweeper-API/domain/mock"
	"testing"
)

var Game domain.Game

func TestGameBuilderService_BuildNewGame(t *testing.T) {
	Game = domain.Game{
		Rows:       3,
		Cols:       2,
		Mines:      2,
		Grid:       mock.Grid,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
	type fields struct {
		Store memory.GameStoreInterface
	}
	type args struct {
		name    string
		rows    int64
		columns int64
		mines   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test_build_new_game:ok", fields{Store: memory.NewGameStore(memory.New())}, args{
			name:    "my test game",
			rows:    2,
			columns: 2,
			mines:   1,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GameBuilderService{
				Store: tt.fields.Store,
			}
			got, err := s.BuildNewGame(tt.args.name, tt.args.rows, tt.args.columns, tt.args.mines)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildNewGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			assert.NotNil(t, got.Timer)
			assert.Nil(t, got.Grid)
			assert.NotEmpty(t, got.Name)
			assert.NotEmpty(t, got.UUID)
			assert.Zero(t, got.Clicks)
			assert.NotZero(t, got.Rows)
			assert.NotZero(t, got.Cols)
			assert.NotZero(t, got.Mines)
		})
	}
}
