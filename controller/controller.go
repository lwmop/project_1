package controller

import (
	"github.com/gin-gonic/gin"
	"project_1/server"
)

func Router(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.POST("/text", server.CountFrequency)
	}
}
