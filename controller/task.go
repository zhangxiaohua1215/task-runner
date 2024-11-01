package controller

import (
	"path"
	"strconv"
	"strings"
	"task-runner/gobal"
	"task-runner/gobal/response"
	"task-runner/job"
	"task-runner/model"
	"task-runner/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task struct{}

// TODO 实现文件的输入，结果文件的输出
// 执行任务
func (t *Task) ExecuteTask(c *gin.Context) {
	args := c.PostFormArray("arg")
	name := c.PostForm("name")
	inputFile, err := c.FormFile("input")
	if err != nil {
		response.Fail(c, "文件上传失败", err)
		return
	}
	scriptID, err := strconv.ParseInt(c.PostForm("script_id"), 10, 64)
	if err != nil {
		response.Fail(c, "脚本id解析错误", err)
		return
	}
	// 验证脚本ID
	script, err := appService.Script.First(scriptID)
	if err != nil {
		response.Fail(c, "找不到指定的脚本文件", err.Error())
		return
	}
	// 保存输入文件到本地
	// taskID := utils.GenID()
	// desPath := utils.GenInputFilePath(taskID, inputFile.Filename)
	// if err := c.SaveUploadedFile(inputFile, desPath); err != nil {
	// 	response.Fail(c, "保存输入文件失败", err)
	// 	return
	// }

	// 任务入库
	task := model.Task{
		ID:            utils.GenID(),
		ScriptID:      scriptID,
		Name:          name,
		Arguments:     strings.Join(args, " "),
		Status:        "pending",
		InputFileName: inputFile.Filename,
		CreatedAt:     time.Now(),
	}
	gobal.DB.Omit("CompletedAt", "ExitCode", "ExecuteTime", "StartedAt").Create(&task)

	f, err := inputFile.Open()
	if err != nil {
		response.Fail(c, "打开输入文件失败", err)
		return
	}
	defer f.Close()
	
	// 加入任务队列
	job.TaskQueue <- job.Task{
		ID:            task.ID,
		ScriptID:      script.ID,
		ScriptPath:    script.Path,
		Arguments:     args,
		InputFileName: inputFile.Filename,
		Ext:           path.Ext(script.Name),
		Input:         f,
	}

	response.Success(c, "任务已加入队列", gin.H{"task_id": task.ID})
}

// 任务列表
func (t *Task) ListTask(c *gin.Context) {
	type Req struct {
		PageSize  int
		Page      int
		Status    string
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

	tasks, count, err := appService.Task.List(req.Page, req.PageSize, req.Status, req.SortField, req.SortOrder)

	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "", response.PageResult{
		List:     tasks,
		Total:    count,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}

// 任务详情
func (t *Task) DetailTask(c *gin.Context) {
	id := c.PostForm("task_id")
	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, "任务id解析错误", err)
		return
	}
	var task model.TaskDetail
	err = gobal.DB.Model(&model.TaskDetail{}).Joins("Script").First(&task, taskID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, "任务id不存在", err)
			return
		}
		response.Fail(c, "", err)
		return
	}
	response.Success(c, "", task)
}
