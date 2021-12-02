package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solo-kingdom/meta/pkg/service"
	"github.com/solo-kingdom/meta/src/model"
)

type KVData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Get(ctx *gin.Context) {
	c := model.Ctx{C: ctx}
	data := KVData{}

	data.Key = ctx.Param("key")
	if len(data.Key) == 0 {
		if err := ctx.BindJSON(&data); err != nil || len(data.Key) == 0 {
			c.Response(model.BadRequest, nil)
			return
		}
	}

	if len(data.Key) == 0 {
		c.Response(model.BadRequest, nil)
	}

	exists, v, err := service.Get(data.Key)
	if !exists {
		c.Response(model.NotFound, nil)
	} else if err != nil {
		c.Response(model.Error, nil)
	} else {
		data.Value = v
		c.Response(model.Success, data)
	}
}

func Set(ctx *gin.Context) {
	c := model.Ctx{C: ctx}
	data := KVData{}
	if err := ctx.BindJSON(&data); err != nil || len(data.Key) == 0 {
		c.Response(model.BadRequest, nil)
		return
	}
	_ = service.Set(data.Key, data.Value)
	c.Response(model.Success, nil)
}
