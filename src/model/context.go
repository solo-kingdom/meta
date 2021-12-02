package model

import "github.com/gin-gonic/gin"

type Ctx struct {
	C *gin.Context
}

func (c *Ctx) Response(err Status, data interface{}) {
	c.C.JSON(err.HttpCode, Response{
		Code:    err.Code,
		Message: err.Message,
		Data:    data,
	})
}
