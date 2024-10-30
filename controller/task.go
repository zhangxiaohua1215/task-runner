package controller

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"task-runner/gobal"
	"task-runner/gobal/response"
	"task-runner/job"
	"task-runner/model"
	"task-runner/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task struct{}

// 执行任务
func (t *Task) ExecuteTask(c *gin.Context) {
	id := c.PostForm("script_id")

	arg := c.PostFormArray("arg")
	name := c.PostForm("name")
	input := c.PostForm("input")
	fmt.Println(arg)
	scriptID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)

		c.JSON(400, gin.H{"error": "invalid script_id"})
		return
	}

	// 任务入库
	task := model.Task{
		ID:        utils.GenID(),
		ScriptID:  scriptID,
		Name:      name,
		Arguments: strings.Join(arg, " "),
		Status:    "pending",
		Input:     input,
	}
	gobal.DB.Create(&task)

	script := appService.Script.Find(scriptID)
	// 加入任务队列
	job.TaskQueue <- job.Task{
		ID:        task.ID,
		ScriptID:  script.ID,
		Arguments: arg,
		FilePath:  script.Path,
		Ext:       path.Ext(script.Name),
		Input:     input,
	}

	c.JSON(200, gin.H{"task_id": task.ID})
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
	var task model.Task
	err = gobal.DB.Joins("Script").First(&task, taskID).Error
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
