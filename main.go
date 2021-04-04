package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "check /metrics")
	})
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(":5000")
}
