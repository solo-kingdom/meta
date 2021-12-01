package v1

import "github.com/gin-gonic/gin"

func Version(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"version": "api/v1",
	})
}
