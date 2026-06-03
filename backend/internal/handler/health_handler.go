package handler

import (
	"github.com/absolute-achilles/plato/pkg/response"
	"github.com/gin-gonic/gin"
)

func RegisterHealthCheck(api *gin.RouterGroup) {
	api.GET("/health", func(ctx *gin.Context) {
		response.OK(ctx, gin.H{
			"health": "ok",
		})
	})
}
