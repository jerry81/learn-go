package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":   c.Param("id"),
			"name": "Ada Lovelace",
		})
	})

	r.Run(":8080")
}
