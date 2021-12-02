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
	if err := ctx.BindJSON(&data); err != nil || len(data.Key) == 0 {
		c.Response(model.BadRequest, nil)
		return
	}
	v, _ := service.Get(data.Key)
	data.Value = v
	c.Response(model.Success, data)
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
