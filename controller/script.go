package controller

import (
	"fmt"
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

func isSupportedFile(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".sh", ".exe":
		return true
	default:
		return false
	}
}

// 上传脚本文件
func (s *Script) UploadScript(c *gin.Context) {
	file, _ := c.FormFile("file")
	description := c.PostForm("description")
	// 验证文件类型
	if !isSupportedFile(file.Filename) {
		response.Fail(c, "不支持的文件类型", nil)
		return
	}

	// 检查脚本是否存在，存在则返回已存在的脚本ID
	// 计算文件哈希
	hash, err := utils.GetMd5FromFile(file)
	if err != nil {
		response.Fail(c, "计算文件哈希失败", err.Error())
		return
	}

	// 存在
	script, err := appService.Script.FindByHash(hash)
	if err != nil && err != gorm.ErrRecordNotFound {
		response.Fail(c, "查询脚本失败", err.Error())
		return
	}
	if script.ID!= 0 {
		response.Success(c, "脚本已存在", gin.H{"script_id": script.ID})
		return
	}

	// 不存在，保存文件
	id := utils.GenID()
	path := utils.GenPath(strconv.FormatInt(id, 10) + "-" + file.Filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		response.Fail(c, "保存文件失败", err.Error())
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

	response.Success(c, "上传成功", gin.H{"script_id": script.ID})
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
	for i := range scripts {
		scripts[i].Url = genDownloadUrl(c.Request.Host, scripts[i].ID)
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
	script, err := appService.Script.First(scriptID)
	if err != nil {
		response.Fail(c, "脚本id不存在")
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+script.Name)
	c.Header("Content-Type", "application/octet-stream")
	c.File(script.Path)

}

func genDownloadUrl(hostName string, id int64) string {
	return fmt.Sprintf("http://%s/script/download/%x", hostName, id)
}
