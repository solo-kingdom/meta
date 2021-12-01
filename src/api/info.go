package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solo-kingdom/meta/pkg/settings"
	"time"
)

func Info(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"service":  "meta",
		"datetime": time.Now(),
		"config": map[string]interface{}{
			"server": settings.ServerConfig,
		},
	})
}
