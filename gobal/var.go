package gobal

import (
	"task-runner/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	if err := DB.AutoMigrate(&model.Script{}, &model.Task{}); err != nil {
		panic(err)
	}
}
