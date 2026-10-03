package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/sd", func(c *gin.Context) {
		services := []gin.H{
			{
				"targets": []string{"10.0.10.2:9100", "10.0.10.3:9100", "10.0.10.4:9100", "10.0.10.5:9100"},
				"labels": gin.H{
					"__meta_prometheus_job": "node",
				},
			},
			{
				"targets": []string{"10.0.40.2:9100", "10.0.40.3:9100"},
				"labels": gin.H{
					"__meta_prometheus_job": "alertmanager",
				},
			},
			{
				"targets": []string{"10.0.40.2:9093", "10.0.40.3:9093"},
				"labels": gin.H{
					"__meta_prometheus_job": "alertmanager",
				},
			},
		}
		c.JSON(http.StatusOK, services)
	})

	r.Run()
}
