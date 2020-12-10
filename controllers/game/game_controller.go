package game

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"minesweeper-API/app/memory"
	"minesweeper-API/domain"
	"minesweeper-API/services"
	"minesweeper-API/services/builder"
	"net/http"
	"strconv"
)

var db = memory.New()
var gameStore = memory.NewGameStore(db)
var Controller gameControllerInterface = &gameController{
	services.GameService{Store: gameStore},
	builder.GameBuilderService{
		Store: gameStore,
	}}

type gameControllerInterface interface {
	StartNewGame(c *gin.Context)
	PauseGame(c *gin.Context)
	ResumeGame(c *gin.Context)
	ClickPosition(c *gin.Context)
	FlagClickPosition(c *gin.Context)
	GetHistory(c *gin.Context)
}
type gameController struct {
	GameService        services.GameService
	GameBuilderService builder.GameBuilderService
}

func (controller *gameController) StartNewGame(c *gin.Context) {
	name := c.Query("name")
	rows, _ := strconv.ParseInt(c.Query("rows"), 10, 64)
	columns, _ := strconv.ParseInt(c.Query("columns"), 10, 64)
	mines, _ := strconv.ParseInt(c.Query("mines"), 10, 64)

	game, err := controller.GameBuilderService.BuildNewGame(name, rows, columns, mines)
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
func (controller *gameController) PauseGame(c *gin.Context) {
	name := c.Query("name")

	game, err := controller.GameService.Pause(name)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "game_controller",
			"service": "pause",
			"method":  "PauseGame",
			"error":   err,
		})
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: game.GameStatus,
	})
	return
}

func (controller *gameController) ResumeGame(c *gin.Context) {
	name := c.Query("name")

	game, err := controller.GameService.Resume(name)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "game_controller",
			"service": "resume",
			"method":  "ResumeGame",
			"error":   err,
		})
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: game.GameStatus,
	})
	return
}

func (controller *gameController) ClickPosition(c *gin.Context) {
	var cellPos domain.CellPos
	gameName := c.Query("name")
	gameUUID := c.Query("uuid")

	logrus.WithFields(logrus.Fields{
		"file":      "game_controller",
		"service":   "click",
		"method":    "ClickPosition",
		"game_name": gameName,
		"game_uuid": gameUUID,
	})

	if err := json.NewDecoder(c.Request.Body).Decode(&cellPos); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Game Name: %v, UUID: %v, Error: %v", gameName, gameUUID, err),
		})
		return
	}

	game, err := controller.GameService.Click(gameName, &cellPos)
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
			Message: fmt.Sprintf("Game Name: %v, UUID: %v, Error: %v", gameName, gameUUID, err),
		})
		return
	}
	cell := game.Grid[cellPos.Row][cellPos.Col]

	clickResult := domain.ClickResult{
		GameStatus: game.GameStatus,
		Cell:       cell,
		Game:       *game,
	}
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: &clickResult,
	})
}

func (controller *gameController) FlagClickPosition(c *gin.Context) {
	var cellPos domain.CellPos
	gameName := c.Query("name")
	gameUUID := c.Query("uuid")

	logrus.WithFields(logrus.Fields{
		"file":      "game_controller",
		"service":   "flag_click",
		"method":    "FlagClickPosition",
		"game_name": gameName,
		"game_uuid": gameUUID,
	})

	if err := json.NewDecoder(c.Request.Body).Decode(&cellPos); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("Game Name: %v, UUID: %v, Error: %v", gameName, gameUUID, err),
		})
		return
	}

	game, err := controller.GameService.FlagClick(gameName, &cellPos)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file":    "game_controller",
			"service": "click",
			"method":  "FlagClickPosition",
			"err":     err,
			"message": "cannot flag cell",
		})

		c.JSON(http.StatusInternalServerError, Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Game Name: %v, UUID: %v, Error: %v", gameName, gameUUID, err),
		})
		return
	}
	cell := game.Grid[cellPos.Row][cellPos.Col]

	flagClickResult := domain.ClickResult{
		GameStatus: game.GameStatus,
		Cell:       cell,
		Game:       *game,
	}
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: &flagClickResult,
	})
}

func (controller *gameController) GetHistory(c *gin.Context) {
	gameHistory, err := controller.GameService.History()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprint("Can not get game history"),
		})

	}
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: &gameHistory,
	})

}
