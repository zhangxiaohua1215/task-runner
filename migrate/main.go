package main

import (
	"fmt"
	"task-runner/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	if err := db.AutoMigrate(&model.Script{}, &model.Task{}); err != nil {
		panic(err) 
	}
	fmt.Println("success")
}