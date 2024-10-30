package model

import "time"

// Script 表结构
type Script struct {
	ID          int64     `json:"id,omitempty" gorm:"primaryKey" `
	Name        string    `json:"name,omitempty"`        // 脚本文件名
	Ext         string    `json:"-"`         // 脚本扩展名
	Hash        string    `json:"hash,omitempty"`        // 脚本文件哈希值
	Path        string    `json:"-"`        // 本地存储路径
	Description string    `json:"description,omitempty"` // 脚本描述
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type ScriptWithUrl struct {
	Script
	Url string `json:"url"`

	// Tasks []Task `json:"tasks"`
}

