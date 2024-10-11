package http

import (
	"net/http"
	"strconv"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/usecases"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryByName(c *gin.Context)
	GetAllCategories(c *gin.Context)
}

type categoryHandler struct {
	usecase usecases.CategoryUsecase
}

func NewCategoryHandler(uc usecases.CategoryUsecase) CategoryHandler {
	return &categoryHandler{
		usecase: uc,
	}
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} models.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/categories [get]
func (h *categoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.usecase.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary Create category
// @Description Create category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/category [post]
func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var category *models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := h.usecase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/category [put]
func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var input models.Category
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	categoryID, err := h.usecase.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	input.ID = categoryID.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.usecase.UpdateCategory(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

// Get Category By ID
// @Summary Get Category by id
// @Description Get Category by id
// @Tags categories
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/category/{category_id} [get]
func (h *categoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.usecase.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

// Get Category By Name
// @Summary Get Category by name
// @Description Get Category by name
// @Tags categories
// @Accept json
// @Produce json
// @Param category_name path string true "Category Name"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/category/{category_name} [get]
func (h *categoryHandler) GetCategoryByName(c *gin.Context) {
	name := c.Param("name")
	category, err := h.usecase.GetCategoryByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})

}

// Delete Category By ID
// @Summary Delete Category by id
// @Description Delete Category by id
// @Tags categories
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/category/{category_id} [delete]
func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	categoryID, err := h.usecase.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	err = h.usecase.DeleteCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
