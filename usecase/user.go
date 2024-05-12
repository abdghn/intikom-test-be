package usecase

import (
	"fmt"
	"intikom-test-be/helper"
	"intikom-test-be/model"
	"intikom-test-be/service"

	"gorm.io/gorm"
)

type UserUsecase interface {
	ReadAll() ([]*model.User, error)
	Create(req model.UserInput) (*model.User, error)
	ReadById(id int) (*model.User, error)
	ReadByEmail(email string) (*model.User, error)
	Update(id int, req *model.UserInput) (*model.User, error)
	Delete(id int) error
}

type userUsecase struct {
	userService service.UserService
}

func NewUserUsecase(userService service.UserService) UserUsecase {
	return &userUsecase{userService: userService}
}

func (u *userUsecase) ReadAll() ([]*model.User, error) {
	return u.userService.ReadAll()
}

func (u *userUsecase) Create(req model.UserInput) (*model.User, error) {

	user, err := u.userService.CheckByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user.ID > 0 {
		return nil, fmt.Errorf("email is exists, please use different email!")
	}

	helper.HashPassword(&req.Password)

	newUser := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	m, err := u.userService.Create(newUser)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (u *userUsecase) ReadByEmail(email string) (*model.User, error) {
	return u.userService.CheckByEmail(email)
}

func (u *userUsecase) ReadById(id int) (*model.User, error) {
	return u.userService.ReadById(id)
}

func (u *userUsecase) Update(id int, req *model.UserInput) (*model.User, error) {
	user, err := u.userService.CheckByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user.ID > 0 {
		return nil, fmt.Errorf("email is exists, please use different email!")
	}

	if req.Password != "" {
		helper.HashPassword(&req.Password)
	}
	return u.userService.Update(id, req)
}

func (u *userUsecase) Delete(id int) error {
	return u.userService.Delete(id)
}
