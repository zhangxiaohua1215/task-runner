package handler

import (
	"net/http"
	"strconv"
	"task-runner/model"
	"task-runner/service"
	"task-runner/utils"

	"github.com/gin-gonic/gin"
)

type ScriptHandler struct{}

// 上传脚本文件
func (s *ScriptHandler) UploadScript(c *gin.Context) {
	file, _ := c.FormFile("file")
	description := c.PostForm("description")
	// 检查脚本是否存在，存在则返回已存在的脚本ID
	
	id := utils.GenID()

	path := utils.GenPath(strconv.FormatInt(id, 10) + file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "保存文件失败",
		})
		return
	}
	c.SaveUploadedFile(file, path)

	script := model.Script{ID: id, Name: file.Filename, Path: path, Description: description}
	service.ScriptServiceInstance.Create(&script)

	c.JSON(200, gin.H{"script_id": script.ID})
}
