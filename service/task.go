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
		"output":      stdout,

		"exit_code":    exitCode,
	})
}

// 任务列表
func (t *Task) List(pageNum, pageSize int, status, sortField, sortOrder string) (tasks []model.Task, cnt int64, err error) {
	db := gobal.DB.Model(&tasks)
	if status!= "" {
		db = db.Where("status = ?", status)
	}
	err = db.Count(&cnt).Error
	if err!= nil {
		return nil, 0, err
	}
	if sortField == "" {
		sortField = "id"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	err = db.Order(sortField + " " + sortOrder).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&tasks).Error
	return tasks, cnt, err
}
