// Package main provides ...
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.String(http.StatusOK, "%s", hostname)
	})
	r.Run(":8080")

}
