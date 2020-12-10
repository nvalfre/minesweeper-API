package game

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"minesweeper-API/app/memory"
	"minesweeper-API/services"
	"minesweeper-API/services/builder"
	"minesweeper-API/services/validator"
	"net/http"
	"net/http/httptest"
	"testing"
)

var store = memory.NewGameStore(memory.New())

var builderService = builder.GameBuilderService{Store: store, Validator: validator.GameValidatorService{}}

func getMocketCtx() (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return w, c
}

func Test_gameController_ClickPosition(t *testing.T) {
	rec, c := getMocketCtx()

	s := &services.GameService{
		Store:           builderService.Store,
		MovementService: services.MovementService{},
	}
	type fields struct {
		GameService        services.GameService
		GameBuilderService builder.GameBuilderService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"test_click_position:ok", fields{
			GameService:        *s,
			GameBuilderService: builderService,
		}, args{c}},
	}

	builderService.BuildNewGame("my_test_game", 3, 2, 1)
	game, _ := builderService.Store.GetByName("my_test_game")
	s.StartGame(game)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &gameController{
				GameService:        tt.fields.GameService,
				GameBuilderService: tt.fields.GameBuilderService,
			}
			body := bytes.NewBufferString("{\n    \"row\": 2,\n    \"column\": 5,\n    \"flag\": false\n}")
			c.Request, _ = http.NewRequest("PUT", "http://127.0.0.1/game/movement?name=my_test_game&uuid=my_test_game_uuid", body)
			controller.ClickPosition(tt.args.c)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func Test_gameController_FlagClickPosition(t *testing.T) {
	rec, c := getMocketCtx()

	s := &services.GameService{
		Store:           builderService.Store,
		MovementService: services.MovementService{},
	}
	type fields struct {
		GameService        services.GameService
		GameBuilderService builder.GameBuilderService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"test_click_position:ok", fields{
			GameService:        *s,
			GameBuilderService: builderService,
		}, args{c}},
	}

	builderService.BuildNewGame("my_test_game", 3, 2, 1)
	game, _ := builderService.Store.GetByName("my_test_game")
	s.StartGame(game)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &gameController{
				GameService:        tt.fields.GameService,
				GameBuilderService: tt.fields.GameBuilderService,
			}
			body := bytes.NewBufferString("{\n    \"row\": 2,\n    \"column\": 5,\n    \"flag\": true\n}")
			c.Request, _ = http.NewRequest("PUT", "http://127.0.0.1/game/movement/flag?name=my_test_game&uuid=my_test_game_uuid", body)
			controller.ClickPosition(tt.args.c)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}

func Test_gameController_PauseGame(t *testing.T) {
	rec, c := getMocketCtx()

	type fields struct {
		GameService        services.GameService
		GameBuilderService builder.GameBuilderService
	}
	type args struct {
		c *gin.Context
	}

	s := &services.GameService{
		Store:           builderService.Store,
		MovementService: services.MovementService{},
	}
	test := struct {
		name   string
		fields fields
		args   args
	}{
		"test_pause_game:ok", fields{
			GameService:        *s,
			GameBuilderService: builderService,
		}, args{c: c},
	}
	controller := &gameController{
		GameService:        test.fields.GameService,
		GameBuilderService: test.fields.GameBuilderService,
	}
	builderService.BuildNewGame("my_test_game", 3, 2, 1)
	game, _ := builderService.Store.GetByName("my_test_game")
	s.StartGame(game)
	c.Request, _ = http.NewRequest("PUT", "http://127.0.0.1/game/pause?name=my_test_game", nil)

	t.Run(test.name, func(t *testing.T) {

		controller.PauseGame(test.args.c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
func Test_gameController_ResumeGame(t *testing.T) {
	rec, c := getMocketCtx()

	type fields struct {
		GameService        services.GameService
		GameBuilderService builder.GameBuilderService
	}
	type args struct {
		c *gin.Context
	}
	s := &services.GameService{
		Store:           builderService.Store,
		MovementService: services.MovementService{},
	}
	test := struct {
		name   string
		fields fields
		args   args
	}{
		"test_resume_game:ok", fields{
			GameService:        *s,
			GameBuilderService: builderService,
		}, args{c: c},
	}
	controller := &gameController{
		GameService:        test.fields.GameService,
		GameBuilderService: test.fields.GameBuilderService,
	}
	builderService.BuildNewGame("my_test_game", 3, 2, 1)
	game, _ := builderService.Store.GetByName("my_test_game")
	s.StartGame(game)
	c.Request, _ = http.NewRequest("PUT", "http://127.0.0.1/game/resume?name=my_test_game", nil)
	t.Run(test.name, func(t *testing.T) {
		controller.ResumeGame(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func Test_gameController_StartNewGame(t *testing.T) {
	rec, c := getMocketCtx()
	s := &services.GameService{
		Store:           builderService.Store,
		MovementService: services.MovementService{},
	}
	type fields struct {
		GameService        services.GameService
		GameBuilderService builder.GameBuilderService
	}
	type args struct {
		c *gin.Context
	}
	test := struct {
		name   string
		fields fields
		args   args
	}{
		"test_start_new_game:ok", fields{
			GameService:        *s,
			GameBuilderService: builderService,
		}, args{c: c},
	}
	controller := &gameController{
		GameService:        test.fields.GameService,
		GameBuilderService: test.fields.GameBuilderService,
	}
	c.Request, _ = http.NewRequest("POST", "http://127.0.0.1/game/new?name=my_test_game_new&rows=8&cols=8&mines=12", nil)

	t.Run(test.name, func(t *testing.T) {
		controller.StartNewGame(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

/*
func Test_gameController_GetHistory(t *testing.T) {
	type fields struct {
		GameService        services.GameService
		GameBuilderService builder.GameBuilderService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &gameController{
				GameService:        tt.fields.GameService,
				GameBuilderService: tt.fields.GameBuilderService,
			}
		})
	}
}*/
