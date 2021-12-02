package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solo-kingdom/meta/pkg/settings"
	"github.com/solo-kingdom/meta/src/model"
	"log"
	"path"
)

func Upload(ctx *gin.Context) {
	c := model.Ctx{C: ctx}
	file, err := ctx.FormFile("file")
	if err != nil || file == nil {
		log.Printf("parse upload request failed. [error=%v]", err.Error())
		c.Response(model.BadRequest, nil)
		return
	}

	dst := path.Join(settings.AppConfig.UploadPath, file.Filename)
	log.Printf("save file. [file=%v, size=%v, dst=%s]",
		file.Filename, file.Size, dst)

	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		log.Printf("save file failed. [error=%v]", err.Error())
		c.Response(model.Error, nil)
		return
	}
	c.Response(model.Success, nil)
}
