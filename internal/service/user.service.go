package service

import (
	"ArticleVista/internal/model"
	"ArticleVista/internal/repository"
)

type UserService interface {
	CreateUser(user *model.Users) error
	UpdateUserById(id int, user model.Users) error
	GetUserById(id int) (*model.Users, error)
	GetUsers() ([]model.Users, error)
	DeleteUserById(id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) CreateUser(user *model.Users) error {

	return u.repo.CreateUser(user)
}

func (u *userService) UpdateUserById(id int, user model.Users) error {
	return u.repo.UpdateUserById(id, user)
}

func (u *userService) GetUserById(id int) (*model.Users, error) {
	return u.GetUserById(id)
}

func (u *userService) GetUsers() ([]model.Users, error) {
	return u.GetUsers()
}

func (u *userService) DeleteUserById(id int) error {
	return u.DeleteUserById(id)
}
