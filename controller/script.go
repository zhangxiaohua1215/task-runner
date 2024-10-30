package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"task-runner/gobal"
	"task-runner/gobal/response"
	"task-runner/model"
	"task-runner/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		Ext:         filepath.Ext(file.Filename),
		Path:        path,
		Hash:        hash,
		Description: description,
	}
	appService.Script.Create(script)

	c.JSON(200, gin.H{"script_id": script.ID})
}

// 脚本列表
func (s *Script) ListScript(c *gin.Context) {
	type Req struct {
		PageSize  int
		Page      int
		Ext       string
		Name      string
		SortField string
		SortOrder string
	}
	var req Req
	err := c.BindJSON(&req)
	if err != nil {
		response.Fail(c, "参数绑定错误", err.Error())
		return
	}

	if req.PageSize == 0 {
		req.PageSize = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}

	scripts, cnt, err := appService.Script.List(req.Page, req.PageSize, req.Ext, req.Name, req.SortField, req.SortOrder)
	if err != nil {
		response.Fail(c, "", err.Error())
		return
	}
	response.Success(c, "", response.PageResult{
		List:     scripts,
		Total:    cnt,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}

// 脚本详情
func (s *Script) DetailScript(c *gin.Context) {
	id := c.PostForm("script_id")
	scriptID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, "脚本id解析错误", err.Error())
		return
	}
	var script model.ScriptWithUrl
	err = gobal.DB.Model(&model.Script{}).First(&script, scriptID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, "脚本id不存在", err.Error())
			return
		}
		response.Fail(c, "", err.Error())
		return
	}
	// 生成下载链接
	script.Url = genDownloadUrl(c.Request.Host, script.ID)
	response.Success(c, "", script)
}

// 下载脚本文件
func (s *Script) DownloadScript(c *gin.Context) {
	id := c.Param("id")
	scriptID, err := strconv.ParseInt(id, 16, 64)
	if err != nil {
		response.Fail(c, "脚本id解析错误", err.Error())
		return
	}
	script := appService.Script.First(scriptID)
	if script == nil {
		response.Fail(c, "脚本id不存在", )
		return
	}

	c.File(script.Path)
}

func genDownloadUrl(hostName string, id int64) string {
	return fmt.Sprintf("%s/download/%x", hostName, id)
}
