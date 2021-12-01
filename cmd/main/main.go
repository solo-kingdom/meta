package main

import (
	"github.com/gin-gonic/gin"
	"github.com/solo-kingdom/meta/pkg/settings"
	"strconv"
	"time"
)

func init() {
	settings.SetUp()
}

func main() {
	r := gin.Default()
	r.GET("info", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"service":  "meta",
			"datetime": time.Now(),
			"config": map[string]interface{}{
				"server": settings.ServerConfig,
			},
		})
	})
	err := r.Run(":" + strconv.Itoa(settings.ServerConfig.Port))
	if err != nil {
		println(err.Error())
	}
}
