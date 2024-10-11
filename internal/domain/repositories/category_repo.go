package repositories

import (
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	UpdateCategory(category *models.Category) (*models.Category, error)
	DeleteCategory(category *models.Category) error
	GetCategoryByID(id uint) (*models.Category, error)
	GetCategoryByName(name string) (*models.Category, error)
	GetAllCategories() ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	return category, r.db.Create(category).Error
}

func (r *categoryRepository) UpdateCategory(category *models.Category) (*models.Category, error) {
	return category, r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(category *models.Category) error {
	return r.db.Delete(category).Error
}

func (r *categoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	category := &models.Category{}
	return category, r.db.Where("id = ?", id).First(category).Error
}

func (r *categoryRepository) GetCategoryByName(name string) (*models.Category, error) {
	category := &models.Category{}
	return category, r.db.Where("name = ?", name).First(category).Error
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	categories := []models.Category{}
	return categories, r.db.Find(&categories).Error
}
