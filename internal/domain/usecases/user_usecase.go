package usecases

import (
	"errors"
	"time"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/repositories"
	"github.com/Arasy41/go-gin-quiz-api/pkg/utils"
	"gorm.io/gorm"
)

type UserUsecase interface {
	CreateUser(username, email, password string, roleId uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.UserList, error)
	GetUsersByRoleID(roleID uint) ([]models.User, error)
	ChangePassword(userID uint, oldPassword, newPassword string) error
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) CreateUser(username, email, password string, roleId uint) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		RoleID:    roleId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) UpdateUser(user *models.User) (*models.User, error) {
	if user.ID == 0 {
		return nil, errors.New("user id is required")
	}

	// Cek apakah password diubah, jika ya maka hash ulang
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	// Set waktu update ke saat ini
	user.UpdatedAt = time.Now()

	// Gunakan repository untuk melakukan update
	updatedUser, err := u.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *userUsecase) DeleteUser(user *models.User) error {
	return u.userRepo.DeleteUser(user)
}

func (u *userUsecase) GetUserByID(id uint) (*models.User, error) {
	user, err := u.userRepo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u *userUsecase) GetUserByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	return u.userRepo.FindUserByUsername(username)
}

func (u *userUsecase) GetUserByEmail(email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	return u.userRepo.FindUserByEmail(email)
}

func (u *userUsecase) GetAllUsers() ([]models.UserList, error) {
	user, err := u.userRepo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	users := []models.UserList{}
	for _, u := range user {
		users = append(users, models.UserList{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
			RoleName: u.Role.Name,
		})
	}
	return users, nil
}

func (u *userUsecase) GetUsersByRoleID(roleID uint) ([]models.User, error) {
	return u.userRepo.FindUserByRoleID(roleID)
}

func (u *userUsecase) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := u.userRepo.FindUserByID(userID)
	if err != nil {
		return err
	}

	ok := utils.CheckPasswordHash(oldPassword, user.Password)
	if !ok {
		return errors.New("old password does not match")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	_, err = u.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
