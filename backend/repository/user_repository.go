package repository

import (
	"BakeMeets/api/models"
	"gorm.io/gorm"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepostory interface {
	SignUp(user *models.User) (models.UserResponse, error)
	Login(user *models.User) (string, error)
	CreateUser(user *models.User) error
	GetUserByEmail(user *models.User, email string) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepostory {
	return &UserRepository{db}
}

func (ur *UserRepository) SignUp(user *models.User) (models.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.UserResponse{}, err
	}
	newUser := models.User{Email: user.Email, Password: string(hash)}
	if err := ur.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}
	resUser := models.UserResponse{
		ID:		newUser.ID,
		Email:	newUser.Email,
	}
	return resUser, nil
}

func (ur *UserRepository) Login(user *models.User) (string, error) {
	storedUser := models.User{}
	if err := ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// passwordの検証
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// jwtトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	if err := ur.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(user *models.User, email string) error {
	if err := ur.DB.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}