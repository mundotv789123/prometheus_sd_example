package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/prometheus/sd", func(c *gin.Context) {
		services := []gin.H{
			{
				"targets": []string{"app-java:8080"},
				"labels": gin.H{
					"__meta_prometheus_job": "example",
					"__metrics_path__":      "/actuator/prometheus",
				},
			},
		}
		c.JSON(http.StatusOK, services)
	})

	r.Run()
}
