package repository

import (
	"github.com/caoquy2000/meeting-app/infrastructure"
	"github.com/caoquy2000/meeting-app/models"
	"github.com/caoquy2000/meeting-app/utils"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(user models.UserRegister) error {
	var dbUser models.User
	dbUser.Email = user.Email
	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	dbUser.Password = user.Password
	dbUser.IsActive = true
	return u.db.DB.Create(&dbUser).Error
}

func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {
	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := utils.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}
