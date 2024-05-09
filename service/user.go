package service

import (
	"fmt"
	"intikom-test-be/model"

	"gorm.io/gorm"
)

type UserService interface {
	ReadAll() ([]*model.User, error)
	Create(req *model.User) (*model.User, error)
	Update(id int, req *model.UserInput) (*model.User, error)
	ReadById(id int) (*model.User, error)
	Delete(id int) error
	CheckByEmail(email string) (*model.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService { return &userService{db: db} }

func (s *userService) ReadAll() ([]*model.User, error) {
	var users []*model.User

	err := s.db.Order("created_at DESC").Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed view all data: %v", err.Error())
	}

	return users, nil
}

func (s *userService) Create(req *model.User) (*model.User, error) {
	tx := s.db.Begin()
	defer tx.Rollback()
	err := tx.Save(&req).Error
	if err != nil {
		return nil, fmt.Errorf("failed insert data")
	}

	tx.Commit()

	return req, nil
}

func (s *userService) CheckByEmail(email string) (*model.User, error) {
	var user = model.User{}
	err := s.db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("[user.service.CheckByEmail] error execute query %v \n", err)
		return nil, err
	}
	return &user, nil
}

func (s *userService) Update(id int, req *model.UserInput) (*model.User, error) {
	var upUser = model.User{}
	err := s.db.Table("users").Where("id = ?", id).First(&upUser).Updates(&req).Error
	if err != nil {
		fmt.Printf("[user.service.Update] error execute query %v \n", err)
		return nil, fmt.Errorf("failed update data")
	}
	return &upUser, nil
}

func (s *userService) ReadById(id int) (*model.User, error) {

	var user = model.User{}
	err := s.db.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("id is not exists")
	}
	return &user, nil
}

func (s *userService) Delete(id int) error {
	var user = model.User{}
	err := s.db.Table("users").Where("id = ?", id).First(&user).Delete(&user).Error
	if err != nil {
		//helper.CommonLogger().Error(err)
		//fmt.Printf("[user.service.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exists")
	}
	return nil
}
