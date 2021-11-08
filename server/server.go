package main

import (
	"log"
	"net/http"

	"github.com/buyme/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:9000"}
	r.Use(cors.New(config))

	err := models.InitDatabase()
	if err != nil {
		log.Fatalln("could not connect database")
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	r.Run(":2222")
}
