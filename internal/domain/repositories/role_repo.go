package repositories

import (
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) (*models.Role, error)
	Update(role *models.Role) (*models.Role, error)
	Delete(role *models.Role) error
	FindRoleByID(id uint) (*models.Role, error)
	FindRoleByName(name string) (*models.Role, error)
	FindAllRoles() ([]models.Role, error)
}

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{DB: db}
}

func (r *roleRepository) Create(role *models.Role) (*models.Role, error) {
	err := r.DB.Create(role).Error
	
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) FindAllRoles() ([]models.Role, error) {
	var roles []models.Role

	err := r.DB.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) Update(role *models.Role) (*models.Role, error) {
	err := r.DB.Save(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) Delete(role *models.Role) error {
	err := r.DB.Delete(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) FindRoleByID(id uint) (*models.Role, error) {
	role := &models.Role{}
	err := r.DB.First(role, id).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) FindRoleByName(name string) (*models.Role, error) {
	role := &models.Role{}
	err := r.DB.Where("name = ?", name).First(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}
