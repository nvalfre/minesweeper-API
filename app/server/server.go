package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"log"
	"minesweeper-API/app/memory"
	"minesweeper-API/app/router"
)

var GameStore *memory.GameStore

func Start() {
	db := memory.New()
	GameStore = memory.NewGameStore(db)
	logrus.StandardLogger()
	r := gin.Default()
	router.InitRoutes(r)

	n := negroni.Classic()
	n.UseHandler(r)

	log.Print("Server running on port :3000")
	r.Run()
}
