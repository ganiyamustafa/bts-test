package services

import (
	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/jinzhu/copier"
)

type UserService struct {
	Handler *utils.Handler
}

// create user function
func (u *UserService) CreateUser(payload requests.RegisterRequest) (*models.User, *apperror.AppError) {
	// copy payload to user models
	var user models.User
	copier.Copy(&user, &payload)

	// hash password
	user.HashPassword()

	return &user, apperror.FromError(u.Handler.Postgre.Create(&user).Error)
}

// get user by email function
func (u *UserService) GetUserByEmail(email string) (*models.User, *apperror.AppError) {
	var user models.User
	return &user, apperror.FromError(u.Handler.Postgre.Where("email = ?", email).First(&user).Error)
}
