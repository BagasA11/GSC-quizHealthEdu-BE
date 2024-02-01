package service

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/dto"
	"BagasA11/GSC-quizHealthEdu-BE/api/models"
	"BagasA11/GSC-quizHealthEdu-BE/configs"
	"BagasA11/GSC-quizHealthEdu-BE/helpers"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{
		db: configs.GetDB(),
	}
}

func (as *AuthService) UserLogin(req *dto.UserLogin) (string, error) {
	var user models.User
	err := as.db.Where("email = ? AND admin = ?", req.Email, false).First(&user).Error
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	if user.Block {
		return "", errors.New("this user was blocked")
	}

	acessToken, err := helpers.GenerateAccessToken(user.ID, user.Username)
	return acessToken, err
}

func (as *AuthService) AdmiLogin(req *dto.AdminLogin) (string, error) {
	var user models.User
	err := as.db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	if user.Block {
		return "", errors.New("this user was blocked")
	}

	acessToken, err := helpers.GenerateAccessToken(user.ID, user.Username)
	return acessToken, err
}