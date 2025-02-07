package main

import (
	// "log"

	"github.com/Desmond51/go-CRUD/initializers"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func init() {
  initializers.LoadEnvVariables() 
  initializers.ConnectToDb()
}

func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}