package services

import (
	"dating-app/models"
	"dating-app/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepository.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.UserRepository.GetUserByID(uint(id))
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	user, err := s.UserRepository.GetUserByID(uint(id))
	if err != nil {
		return err
	}
	return s.UserRepository.DeleteUser(user)
}
