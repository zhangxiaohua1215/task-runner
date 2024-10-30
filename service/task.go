package service

import (
	"task-runner/gobal"
	"task-runner/model"
	"time"
)

type Task struct {
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

// 更新任务状态为正在执行
func (t *Task) Start(taskID int64) {
	gobal.DB.Model(&model.Task{ID: taskID}).Updates(map[string]any{"status": TaskStatusRunning, "started_at": time.Now()})
}

// 更新任务状态为已完成
func (t *Task) Complete(taskID int64, status TaskStatus, stdout []byte, exitCode int) {
	gobal.DB.Model(&model.Task{ID: taskID}).Updates(map[string]any{
		"status":       status,
		"completed_at": time.Now(),
		"std_out":      stdout,
		"exit_code":    exitCode,
	})
}

func (t *Task) Create(task *model.Task) {
	gobal.DB.Create(&task)
}
