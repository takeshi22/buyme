package main

import (
	"log"
	"net/http"

	"github.com/buyme/api/auth"
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

	err = models.InitRedis()
	if err != nil {
		log.Fatalln("could not connect redis")
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	api := r.Group("/api")
	{
		api.POST("/signUp", auth.SignUp)
		api.POST("/login", auth.Login)
	}

	r.Run(":2222")
}
