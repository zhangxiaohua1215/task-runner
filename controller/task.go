package controller

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"task-runner/job"
	"task-runner/model"
	"task-runner/utils"

	"github.com/gin-gonic/gin"
)

type Task struct{}

// 执行任务
func (t *Task) ExecuteTask(c *gin.Context) {
	id := c.PostForm("script_id")

	arg := c.PostFormArray("arg")
	name := c.PostForm("name")
	stdin := c.PostForm("stdin")
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
		StdIn:     []byte(stdin),
	}
	appService.Task.Create(&task)

	script := appService.Script.Find(scriptID)
	// 加入任务队列
	job.TaskQueue <- job.Task{
		ID:        task.ID,
		ScriptID:  script.ID,
		Arguments: arg,
		FilePath:  script.Path,
		Ext:       path.Ext(script.Name),
		StdIn:     []byte(stdin),
	}

	c.JSON(200, gin.H{"task_id": task.ID})
}
