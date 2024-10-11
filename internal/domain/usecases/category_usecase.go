package usecases

import (
	"time"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/repositories"
)

type CategoryUsecase interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	UpdateCategory(category *models.Category) (*models.Category, error)
	DeleteCategory(category *models.Category) error
	GetCategoryByID(id uint) (*models.Category, error)
	GetCategoryByName(name string) (*models.Category, error)
	GetAllCategories() ([]models.CategoryList, error)
}

type categoryUsecase struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryUsecase(repo repositories.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo: repo}
}

func (u *categoryUsecase) CreateCategory(category *models.Category) (*models.Category, error) {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	return u.categoryRepo.CreateCategory(category)
}

func (u *categoryUsecase) UpdateCategory(category *models.Category) (*models.Category, error) {
	category.UpdatedAt = time.Now()
	return u.categoryRepo.UpdateCategory(category)
}

func (u *categoryUsecase) DeleteCategory(category *models.Category) error {
	return u.categoryRepo.DeleteCategory(category)
}

func (u *categoryUsecase) GetCategoryByID(id uint) (*models.Category, error) {
	return u.categoryRepo.GetCategoryByID(id)
}

func (u *categoryUsecase) GetCategoryByName(name string) (*models.Category, error) {
	return u.categoryRepo.GetCategoryByName(name)
}

func (u *categoryUsecase) GetAllCategories() ([]models.CategoryList, error) {
	category, err := u.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	categories := []models.CategoryList{}
	for _, c := range category {
		categories = append(categories, models.CategoryList{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return categories, nil
}
