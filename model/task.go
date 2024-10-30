package model

import "time"

type Task struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	ScriptID    int64     `json:"script_id,omitempty"` // 关联脚本文件ID
	Name        string    `json:"name,omitempty"`      // 任务名称
	Arguments   string    `json:"arguments,omitempty"` // 脚本参数
	Status      string    `json:"status,omitempty"`    // 任务状态 (pending, running, completed, failed)
	Input       string    `json:"input,omitempty"`
	Output      string    `json:"output,omitempty"`
	ExitCode    int       `json:"exit_code,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`   // 任务创建时间
	StartedAt   time.Time `json:"started_at,omitempty"`   // 任务开始时间
	CompletedAt time.Time `json:"completed_at,omitempty"` // 任务完成时间
}

// 任务详情
type TaskDetail struct {
	Task
	
	Script Script `json:"script,omitempty"` // 关联脚本文件
}
