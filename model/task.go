package model

import "time"

type Task struct {
	ID          int64  `gorm:"primaryKey"`
	ScriptID    int64  `gorm:"not null"` // 关联脚本文件ID
	Name        string // 任务名称
	Arguments   string // 脚本参数
	Status      string `gorm:"default:'pending'"` // 任务状态 (pending, running, completed, failed)
	StdIn       []byte
	StdOut      []byte
	ExitCode    int
	CreatedAt   time.Time // 任务创建时间
	StartedAt   time.Time // 任务开始时间
	CompletedAt time.Time // 任务完成时间

	// 关联脚本文件
	Script Script `gorm:"foreignKey:ScriptID"`
}
