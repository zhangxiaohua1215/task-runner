package service

import (
	"task-runner/model"
	"time"

	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)


func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

// 更新任务状态为正在执行
func (t *TaskService) Start(taskID int64) {
	t.db.Model(&model.Task{ID: taskID}).Updates(map[string]any{"status": TaskStatusRunning, "started_at": time.Now()})
}

// 更新任务状态为已完成
func (t *TaskService) Complete(taskID int64, status TaskStatus, stdout []byte, exitCode int) {
	t.db.Model(&model.Task{ID: taskID}).Updates(map[string]any{
		"status": status, 
		"completed_at": time.Now(), 
		"std_out": stdout,
		"exit_code": exitCode,
	})
}

func (t *TaskService) Create(task *model.Task) {
	t.db.Create(&task)
}
