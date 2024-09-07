package usecases

import (
	"errors"
	"time"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/repositories"
	"gorm.io/gorm"
)

type RoleUsecase interface {
	CreateRole(role *models.Role) (*models.Role, error)
	UpdateRole(role *models.Role) (*models.Role, error)
	DeleteRole(role *models.Role) error
	GetRoleByID(id uint) (*models.Role, error)
	GetRoleByName(rolename string) (*models.Role, error)
	GetAllRoles() ([]models.RoleList, error)
}

type roleUsecase struct {
	roleRepo repositories.RoleRepository
}

func NewRoleUsecase(repo repositories.RoleRepository) RoleUsecase {
	return &roleUsecase{roleRepo: repo}
}

func (u *roleUsecase) CreateRole(req *models.Role) (*models.Role, error) {
	role := &models.Role{
		ID:        req.ID,
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	if req.Name == "" {
		return nil, errors.New("role name is required")
	}
	return u.roleRepo.Create(role)
}

func (u *roleUsecase) UpdateRole(role *models.Role) (*models.Role, error) {
	if role.ID == 0 {
		return nil, errors.New("role id is required")
	}

	return u.roleRepo.Update(role)
}

func (u *roleUsecase) DeleteRole(role *models.Role) error {
	if role.ID == 0 {
		return errors.New("role id is required")
	}
	return u.roleRepo.Delete(role)
}

func (u *roleUsecase) GetRoleByID(id uint) (*models.Role, error) {
	return u.roleRepo.FindRoleByID(id)
}

func (u *roleUsecase) GetRoleByName(rolename string) (*models.Role, error) {
	return u.roleRepo.FindRoleByName(rolename)
}

func (u *roleUsecase) GetAllRoles() ([]models.RoleList, error) {
	user, err := u.roleRepo.FindAllRoles()
	if err != nil {
		return nil, err
	}

	roles := []models.RoleList{}
	for _, r := range user {
		roles = append(roles, models.RoleList{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	return roles, nil
}
