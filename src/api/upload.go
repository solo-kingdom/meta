package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solo-kingdom/meta/pkg/e"
	"github.com/solo-kingdom/meta/pkg/settings"
	"github.com/solo-kingdom/meta/pkg/utils"
	"github.com/solo-kingdom/meta/src/model"
	"io"
	"log"
	"os"
	"path"
)

func Upload(ctx *gin.Context) {
	c := model.Ctx{C: ctx}
	file, err := ctx.FormFile("file")
	if err != nil || file == nil {
		log.Printf("parse upload request failed. [e=%v]", e.ErrMsg(err))
		c.Response(model.BadRequest, nil)
		return
	}

	dir := path.Join(settings.AppConfig.UploadPath, ctx.Param("space"))
	err = utils.EnsureDir(dir)
	if err != nil {
		log.Printf("create dir failed. [e=%v]", err.Error())
		c.Response(model.Error, nil)
		return
	}

	dst := path.Join(dir, file.Filename)
	log.Printf("save file. [file=%v, size=%v, dst=%s]",
		file.Filename, file.Size, dst)

	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		log.Printf("save file failed. [e=%v]", err.Error())
		c.Response(model.Error, nil)
		return
	}
	c.Response(model.Success, nil)
}

func Download(ctx *gin.Context) {
	log.Printf("download file. [space=%s, file=%s]",
		ctx.Param("space"), ctx.Param("file"))

	fp := path.Join(settings.AppConfig.UploadPath, ctx.Param("space"), ctx.Param("file"))
	file, err := os.Open(fp)
	if err != nil || file == nil {
		log.Printf("open file failed. [e=%v]", e.ErrMsg(err))
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	ctx.Writer.Header().Add("Content-type", "application/octet-stream")
	_, err = io.Copy(ctx.Writer, file)
	ctx.Done()
}
