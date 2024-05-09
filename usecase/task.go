package usecase

import (
	"intikom-test-be/model"
	"intikom-test-be/service"
)

type TaskUsecase interface {
	ReadAll() ([]*model.Task, error)
	Create(req model.TaskInput) (*model.Task, error)
	ReadById(id int) (*model.Task, error)
	Update(id int, req *model.TaskInput) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskService service.TaskService
}

func NewTaskUsecase(taskService service.TaskService) TaskUsecase {
	return &taskUsecase{taskService: taskService}
}

func (u *taskUsecase) ReadAll() ([]*model.Task, error) {
	return u.taskService.ReadAll()
}

func (u *taskUsecase) Create(req model.TaskInput) (*model.Task, error) {
	m := &model.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserID:      uint(req.UserID),
	}

	return u.taskService.Create(m)
}

func (u *taskUsecase) ReadById(id int) (*model.Task, error) {
	return u.taskService.ReadById(id)
}

func (u *taskUsecase) Update(id int, req *model.TaskInput) (*model.Task, error) {
	return u.taskService.Update(id, req)
}

func (u *taskUsecase) Delete(id int) error {
	return u.taskService.Delete(id)
}
