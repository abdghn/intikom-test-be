package service

import (
	"fmt"
	"intikom-test-be/model"

	"gorm.io/gorm"
)

type TaskService interface {
	ReadAll() ([]*model.Task, error)
	Create(req *model.Task) (*model.Task, error)
	Update(id int, req *model.TaskInput) (*model.Task, error)
	ReadById(id int) (*model.Task, error)
	Delete(id int) error
}

type taskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) TaskService { return &taskService{db: db} }

func (s *taskService) ReadAll() ([]*model.Task, error) {
	var tasks []*model.Task

	err := s.db.Order("updated_at DESC").Preload("User").Find(&tasks).Error
	if err != nil {
		return nil, fmt.Errorf("failed view all data: %v", err.Error())
	}

	return tasks, nil
}

func (s *taskService) Create(req *model.Task) (*model.Task, error) {

	tx := s.db.Begin()
	defer tx.Rollback()
	err := tx.Save(&req).Error
	if err != nil {
		return nil, fmt.Errorf("failed insert data")
	}

	tx.Commit()

	return req, nil
}

func (s *taskService) Update(id int, req *model.TaskInput) (*model.Task, error) {

	var upTask = model.Task{}
	err := s.db.Preload("User").Table("tasks").Where("id = ?", id).First(&upTask).Updates(&req).Error
	if err != nil {
		fmt.Printf("[task.service.Update] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &upTask, nil
}

func (s *taskService) ReadById(id int) (*model.Task, error) {

	var task = model.Task{}
	err := s.db.Preload("User").Table("tasks").Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, fmt.Errorf("id is not exists")
	}
	return &task, nil
}

func (s *taskService) Delete(id int) error {

	var task = model.Task{}
	err := s.db.Table("tasks").Where("id = ?", id).First(&task).Delete(&task).Error
	if err != nil {
		return fmt.Errorf("id is not exists")
	}
	return nil
}
