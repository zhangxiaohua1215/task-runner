package model

import "time"

type Script struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"not null"` // 脚本文件名
	Hash        string `gorm:"unique;not null"` // 脚本文件哈希值
	Path        string `gorm:"not null"`        // 本地存储路径
	Description string // 脚本描述
	CreatedAt   time.Time
}
