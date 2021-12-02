package model

import "github.com/gin-gonic/gin"

type Ctx struct {
	C *gin.Context
}

func (c *Ctx) Response(status Status, data interface{}) {
	c.C.JSON(status.HttpCode, Response{
		Code:    status.Code,
		Success: status.Success,
		Message: status.Message,
		Data:    data,
	})
}
