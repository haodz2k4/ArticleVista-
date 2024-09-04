package repository

import (
	"ArticleVista/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUserByQuery() ([]model.Users, error)
	GetUserById(id int) (*model.Users, error)
	UpdateUserById(id int, user *model.Users) error
	DeleteUserById(id int) error
	CreateUser(user *model.Users) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) GetAllUserByQuery() ([]model.Users, error) {
	var users []model.Users
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) GetUserById(id int) (*model.Users, error) {
	var user *model.Users
	if err := u.db.Take(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) UpdateUserById(id int, user *model.Users) error {

	if err := u.db.Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) DeleteUserById(id int) error {
	var user *model.Users
	if err := u.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CreateUser(user *model.Users) error {

	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
