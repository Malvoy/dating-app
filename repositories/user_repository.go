package repositories

import (
	"dating-app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUsersWithPagination(limit, offset int) ([]models.User, int64, error)
	GetUsersByFilters(filters map[string]interface{}) ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// Get user by ID
func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Get user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUsersWithPagination(limit, offset int) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	err := r.db.Model(&models.User{}).Count(&total).Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *userRepository) GetUsersByFilters(filters map[string]interface{}) ([]models.User, error) {
	var users []models.User
	query := r.db.Model(&models.User{})
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(user *models.User) error {
	return r.db.Delete(user).Error
}
