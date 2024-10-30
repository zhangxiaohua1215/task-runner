package controller

import (
	"net/http"
	"strconv"
	"task-runner/model"
	"task-runner/utils"

	"github.com/gin-gonic/gin"
)

type Script struct{}

// 上传脚本文件
func (s *Script) UploadScript(c *gin.Context) {
	file, _ := c.FormFile("file")
	description := c.PostForm("description")
	// 检查脚本是否存在，存在则返回已存在的脚本ID
	// 计算文件哈希
	hash, err := utils.GetMd5FromFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "计算文件哈希失败",
		})
		return
	}

	// 存在
	script := appService.Script.FindByHash(hash)
	if script != nil {
		c.JSON(200, gin.H{
			"script_id": script.ID,
			"msg":       "脚本已存在",
		})
		return
	}

	// 不存在，保存文件
	id := utils.GenID()
	path := utils.GenPath(strconv.FormatInt(id, 10) + "-" + file.Filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "保存文件失败",
		})
		return
	}

	script = &model.Script{
		ID:          id,
		Name:        file.Filename,
		Path:        path,
		Hash:        hash,
		Description: description,
	}
	appService.Script.Create(script)

	c.JSON(200, gin.H{"script_id": script.ID})
}
