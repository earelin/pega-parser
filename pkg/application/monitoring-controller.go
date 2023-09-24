package application

import "github.com/gin-gonic/gin"

func MonitoringConfig(e *gin.Engine) {
	e.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
