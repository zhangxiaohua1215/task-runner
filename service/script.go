package service

import (
	"task-runner/model"

	"gorm.io/gorm"
)

type ScriptService struct {
	db *gorm.DB
}

func NewScriptService(db *gorm.DB) *ScriptService {
	return &ScriptService{db: db}
}

func (s *ScriptService) Create(script *model.Script) {
	s.db.Create(&script)
}

func (s *ScriptService) Find(id int64) *model.Script {
	var script model.Script
	s.db.First(&script, id)
	return &script
}