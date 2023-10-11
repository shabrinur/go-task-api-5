package main

import (
	"idstar-idp/rest-api/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()

	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run()
}
