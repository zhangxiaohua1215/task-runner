package service

import "task-runner/gobal"

var (
	TaskServiceInstance = NewTaskService(gobal.DB)
	ScriptServiceInstance = NewScriptService(gobal.DB)
)

func init() {

}