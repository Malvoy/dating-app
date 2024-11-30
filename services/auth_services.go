package services

import (
	"dating-app/models"
	"dating-app/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{UserRepository: repo}
}

func (s *AuthService) Signup(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepository.CreateUser(&user)
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}
