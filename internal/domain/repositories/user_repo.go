package repositories

import (
	"errors"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
	FindUserByID(id uint) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindAllUsers() ([]models.User, error)
	FindUserByRoleID(id uint) ([]models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := ur.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	err := ur.DB.Model(user).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) DeleteUser(user *models.User) error {
	err := ur.DB.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User

	// Aktifkan debug untuk melihat query yang dijalankan
	err := ur.DB.Debug().Preload("Role").Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) FindUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := ur.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := ur.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindAllUsers() ([]models.User, error) {
	users := []models.User{}
	err := ur.DB.Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) FindUserByRoleID(id uint) ([]models.User, error) {
	users := []models.User{}
	err := ur.DB.Where("role_id = ?", id).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
