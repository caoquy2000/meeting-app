package service

import (
	"github.com/caoquy2000/meeting-app/api/repository"
	"github.com/caoquy2000/meeting-app/models"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(user models.UserRegister) error {
	return u.repo.CreateUser(user)
}

func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}
