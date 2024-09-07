package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/usecases"
	"github.com/Arasy41/go-gin-quiz-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	UserUc usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) UserHandler {
	return &userHandler{
		UserUc: uc,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /cms/user [get]
func (h *userHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UserUc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// FindByUserID godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/user/{id} [get]
func (h *userHandler) GetUserByID(c *gin.Context) {
	var user *models.User
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserUc.GetUserByID(uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "record not found"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser godoc
// @Summary Create new user
// @Description Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param user body models.User true "Create user"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/user [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.NewValidator()
	if err := validator.ValidateStruct(validate, &input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUc.CreateUser(input.Username, input.Email, input.Password, input.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "User ID"
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /cms/user/{id} [put]
func (h *userHandler) UpdateUser(c *gin.Context) {
	var input models.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userId, err := h.UserUc.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	input.ID = userId.ID

	if input.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.NewValidator()
	if err := validator.ValidateStruct(validate, &input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUc.UpdateUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Delete user
func (h *userHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.UserUc.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = h.UserUc.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
