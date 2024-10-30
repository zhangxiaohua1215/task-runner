package service

import (
	"task-runner/gobal"
	"task-runner/model"
)

type Script struct {
}

func (s *Script) Create(script *model.Script) {
	gobal.DB.Create(&script)
}

func (s *Script) Find(id int64) *model.Script {
	var script model.Script
	err := gobal.DB.First(&script, id).Error
	if err != nil {
		return nil
	}

	return &script
}

func (s *Script) FindByHash(hash string) *model.Script {
	var script model.Script
	err := gobal.DB.Where("hash =?", hash).First(&script).Error
	if err != nil {
		return nil
	}
	return &script
}
