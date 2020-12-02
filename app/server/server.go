package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"log"
	"minesweeper-API/app/router"
)

func Start() {
	logrus.StandardLogger()
	r := gin.Default()
	router.InitRoutes(r)

	n := negroni.Classic()
	n.UseHandler(r)

	log.Print("Server running on port :3000")
	r.Run()
}
