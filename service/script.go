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

func (s *Script) First(id int64)( *model.Script, error){
	var script model.Script
	err := gobal.DB.First(&script, id).Error
	return &script, err
}

func (s *Script) FindByHash(hash string) (*model.Script, error) {
	var script model.Script
	err := gobal.DB.Where("hash =?", hash).First(&script).Error
	return &script, err

}

// 分页查询
func (s *Script) List(pageNum, pageSize int, ext, name, sortField, sortOrder string) (scripts []model.ScriptWithUrl, cnt int64, err error) {
	db := gobal.DB.Model(&model.Script{})
	if ext!= "" {
		db = db.Where("ext =?", ext)
	}
	if name!= "" {
		db = db.Where("name like ?", "%"+ name + "%")
	}
	if sortField == "" {
		sortField = "id"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	err = db.Count(&cnt).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order(sortField + " " + sortOrder).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&scripts).Error
	return scripts, cnt, err
}

