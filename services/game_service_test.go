package services

import (
	"github.com/go-playground/assert/v2"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
	"minesweeper-API/domain/mock"
	"minesweeper-API/services/builder"
	"minesweeper-API/services/validator"
	"reflect"
	"testing"
)

var service = MovementService{}
var store = memory.NewGameStore(memory.New())
var builderService = builder.GameBuilderService{Store: store, Validator: validator.GameValidatorService{}}

func TestGameService_Click(t *testing.T) {
	type fields struct {
		Store           memory.GameStoreInterface
		MovementService MovementService
	}
	type args struct {
		name string
		pos  *domain.CellPos
	}
	test := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		"test_click:ok", fields{
			Store:           store,
			MovementService: service,
		}, args{
			name: "my_test_game",
			pos: &domain.CellPos{
				Row:  1,
				Col:  1,
				Flag: false,
			},
		}, false,
	}
	builderService.BuildNewGame("my_test_game", 3, 2, 1)
	s := &GameService{
		Store:           test.fields.Store,
		MovementService: test.fields.MovementService,
	}
	game, _ := test.fields.Store.GetByName("my_test_game")
	s.StartGame(game)
	t.Run(test.name, func(t *testing.T) {
		got, err := s.Click(test.args.name, test.args.pos)
		if (err != nil) != test.wantErr {
			t.Errorf("Click() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		assert.Equal(t, got.Clicks, int64(1))
	})
}

func TestGameService_FlagClick(t *testing.T) {
	type fields struct {
		Store           memory.GameStoreInterface
		MovementService MovementService
	}
	type args struct {
		name string
		pos  *domain.CellPos
	}
	test := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		"test_flag_click:ok", fields{
			Store:           store,
			MovementService: service,
		}, args{
			name: "my_test_game_flag",
			pos: &domain.CellPos{
				Row:  1,
				Col:  1,
				Flag: true,
			},
		}, false,
	}
	_, _ = builderService.BuildNewGame("my_test_game_flag", 3, 2, 1)
	s := &GameService{
		Store:           test.fields.Store,
		MovementService: test.fields.MovementService,
	}
	game, _ := test.fields.Store.GetByName("my_test_game_flag")
	s.StartGame(game)
	t.Run(test.name, func(t *testing.T) {
		got, err := s.FlagClick(test.args.name, test.args.pos)
		if (err != nil) != test.wantErr {
			t.Errorf("FlagClick() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		assert.Equal(t, got.Clicks, int64(0))
		assert.Equal(t, got.Flags, int64(1))
		assert.Equal(t, got.Grid[1][1].Flagged, true)
	})
}

func TestGameService_History(t *testing.T) {
	type fields struct {
		Store           memory.GameStoreInterface
		MovementService MovementService
	}
	test := struct {
		name    string
		fields  fields
		want    map[string]*domain.Game
		wantErr bool
	}{
		"test_history", fields{
			Store:           memory.NewGameStore(memory.New()),
			MovementService: MovementService{},
		}, map[string]*domain.Game{mock.Game.Name: &mock.Game, mock.GameBomb.Name: &mock.GameBomb}, false,
	}

	t.Run(test.name, func(t *testing.T) {
		s := &GameService{
			Store:           test.fields.Store,
			MovementService: test.fields.MovementService,
		}
		_ = s.Store.Insert(&mock.Game)
		_ = s.Store.Insert(&mock.GameBomb)
		got, err := s.History()
		if (err != nil) != test.wantErr {
			t.Errorf("History() error = %v, wantErr %v", err, test.wantErr)
			return
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("History() got = %v, want %v", got, test.want)
		}
	})
}

func TestGameService_Pause(t *testing.T) {

	type fields struct {
		Store           memory.GameStoreInterface
		MovementService MovementService
	}
	type args struct {
		name string
	}
	game := &domain.Game{
		Name:       "my_test_game",
		GameStatus: game_status.GameStatus{Status: game_status.Paused},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Game
		wantErr bool
	}{
		{"test_pause:ok", fields{
			Store:           memory.NewGameStore(memory.New()),
			MovementService: MovementService{},
		}, args{name: "my_test_game"}, game, false},
	}
	s := &GameService{
		Store: builderService.Store,
	}
	builderService.BuildNewGame("my_test_game", 3, 2, 1)
	game, _ = builderService.Store.GetByName("my_test_game")
	s.StartGame(game)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Pause(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GameStatus.Status, tt.want.GameStatus.Status) {
				t.Errorf("Pause() got = %v, game %v", got, tt.want)
			}
		})
	}
}

func TestGameService_Resume(t *testing.T) {
	game := &domain.Game{
		Name:       "my_test_game_resume",
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
	type fields struct {
		Store           memory.GameStoreInterface
		MovementService MovementService
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Game
		wantErr bool
	}{
		{"test_resume:ok", fields{
			Store:           memory.NewGameStore(memory.New()),
			MovementService: MovementService{},
		}, args{name: "my_test_game_resume"}, game, false},
	}
	s := &GameService{
		Store: builderService.Store,
	}
	builderService.BuildNewGame("my_test_game_resume", 3, 2, 1)
	game, _ = builderService.Store.GetByName("my_test_game_resume")
	s.StartGame(game)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Resume(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got.GameStatus.Status, tt.want.GameStatus.Status)
		})
	}
}
