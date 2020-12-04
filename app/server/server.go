package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"log"
	"minesweeper-API/app/router"
	"time"
)

func Start() {
	logrus.StandardLogger()
	r := gin.Default()
	router.InitRoutes(r)

	n := negroni.Classic()
	n.UseHandler(r)
	r.HandleMethodNotAllowed = true
	r.RedirectTrailingSlash = false
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	log.Print("Server running on port :3000")
	r.Run()
}
