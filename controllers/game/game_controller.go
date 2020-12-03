package game

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
	"minesweeper-API/services"
	"net/http"
	"strconv"
	"time"
)

const defaultDurationTime = 10 * time.Minute
const defaultStartClick = 0

var db = memory.New()
var gameStore = memory.NewGameStore(db)
var Controller gameControllerInterface = &gameController{services.GameService{Store: gameStore}}

type gameControllerInterface interface {
	StartNewGame(c *gin.Context)
	ClickPosition(c *gin.Context)
}
type gameController struct {
	GameService services.GameService
}

func (controller *gameController) StartNewGame(c *gin.Context) {
	name := c.Query("name")
	rows, _ := strconv.ParseInt(c.Query("rows"), 10, 64)
	columns, _ := strconv.ParseInt(c.Query("columns"), 10, 64)
	mines, _ := strconv.ParseInt(c.Query("mines"), 10, 64)

	game, err := controller.buildNewGame(name, rows, columns, mines)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "game_controller",
			"service": "create",
			"method":  "StartNewGame",
			"error":   err,
		})
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	game, err = controller.GameService.StartGame(game)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "game_controller",
			"service": "start",
			"method":  "StartNewGame",
			"err":     err,
			"message": "cannot start game",
		})
		c.JSON(http.StatusInternalServerError, Response{
			Status:  http.StatusInternalServerError,
			Message: game,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: game,
	})
	return
}

func (controller *gameController) buildNewGame(name string, rows, columns, mines int64) (*domain.Game, error) {
	game := &domain.Game{
		Timer:  time.NewTimer(defaultDurationTime),
		Name:   name,
		UUID:   fmt.Sprintf("%v", uuid.New()),
		Rows:   rows,
		Cols:   columns,
		Mines:  mines,
		Flags:  rows * columns,
		Grid:   nil,
		Clicks: defaultStartClick,
	}
	logrus.WithFields(logrus.Fields{
		"file":    "game_controller",
		"service": "create",
		"method":  "StartNewGame",
		"game":    game,
	})
	err := controller.GameService.Create(game)
	if err != nil {
		return nil, err
	}
	return game, err
}

func (controller *gameController) ClickPosition(c *gin.Context) {
	var CellPos struct {
		Row  int64
		Col  int64
		Flag bool
	}
	gameName := c.Query("name")
	gameUUID := c.Query("uuid")

	logrus.WithFields(logrus.Fields{
		"file":      "game_controller",
		"service":   "click",
		"method":    "ClickPosition",
		"game_name": gameName,
		"game_uuid": gameUUID,
	})

	if err := json.NewDecoder(c.Request.Body).Decode(&CellPos); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Game Name: %v, UUID: %v, Error: %v", &gameName, &gameUUID, err),
		})
		return
	}

	game, err := controller.GameService.Click(gameName, CellPos.Row, CellPos.Col, CellPos.Flag)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "game_controller",
			"service": "click",
			"method":  "ClickPosition",
			"err":     err,
			"message": "cannot click cell",
		})

		c.JSON(http.StatusInternalServerError, Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Game Name: %v, UUID: %v, Error: %v", &gameName, &gameUUID, err),
		})
		return
	}
	cell := game.Grid[CellPos.Row][CellPos.Col]

	if game.Status != "over" && game.Status != "won" {
		game.Grid = nil
	}
	clickResult := domain.ClickResult{
		Cell: cell, Game: *game,
	}
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: &clickResult,
	})
}
