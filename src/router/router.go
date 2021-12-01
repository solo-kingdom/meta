package router

import (
	"github.com/gin-gonic/gin"
	"github.com/solo-kingdom/meta/src/api"
	v1 "github.com/solo-kingdom/meta/src/api/v1"
)

func GetRouter() *gin.Engine {
	r := gin.New()

	r.GET("info", api.Info)
	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("version", v1.Version)
	}
	return r
}
