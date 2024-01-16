package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	FILE_NAME = "leaderboard.json"
)

func main() {
	loadScores()
	router := gin.Default()

	router.GET("/scores", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "works"})
	})

	router.POST("/scores", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error reading request body"})
		}
		fmt.Print("Request Body:", string(body))
		c.JSON(200, gin.H{"message": "Request processed successfully"})
	})

	router.Run(":8080")
}

func loadScores() {

}
