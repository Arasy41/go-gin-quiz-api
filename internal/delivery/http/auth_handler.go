package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/usecases"
	"github.com/Arasy41/go-gin-quiz-api/pkg/constant"
	"github.com/Arasy41/go-gin-quiz-api/pkg/jwt"
	"github.com/Arasy41/go-gin-quiz-api/pkg/utils"
	"github.com/Arasy41/go-gin-quiz-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	ChangePassword(c *gin.Context)
	GetCurrentUser(c *gin.Context)
}

type authHandler struct {
	userUsecase usecases.UserUsecase
}

func NewAuthHandler(uc usecases.UserUsecase) AuthHandler {
	return &authHandler{
		userUsecase: uc,
	}
}

// LoginUser godoc
// @Summary Login as user.
// @Description Logging in to get jwt token to access admin or user API by roles.
// @Tags Auth
// @Param Body body models.LoginRequest true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var user *models.User
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.NewValidator()
	if err := validator.ValidateStruct(validate, &req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.GetUserByUsername(req.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user == nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		log.Println("Invalid username or password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := jwt.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// RegisterUser godoc
// @Summary Register as user
// @Description Register a new user to the system with username, email, password, and role name.
// @Tags Auth
// @Accept json
// @Produce json
// @Param username body string true "Username for the new user"
// @Param email body string true "Email address for the new user"
// @Param password body string true "Password for the new user (8-32 characters)"
// @Param role_name body string true "Role name for the new user (e.g., 'user', 'admin')"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		RoleName string `json:"role_name" binding:"required"`
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

	_, err := h.userUsecase.GetUserByUsername(input.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	_, err = h.userUsecase.GetUserByEmail(input.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Check length password
	if len(input.Password) < constant.MinPasswordLength || len(input.Password) > constant.MaxPasswordLength {
		log.Println(len(input.Password), "Password must be at least between 8 and 32 characters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least between 8 and 32 characters"})
		return
	}

	roleId, err := defineRoles(input.RoleName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.CreateUser(input.Username, input.Email, input.Password, roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// ChangePassword godoc
// @Summary ChangePassword as user
// @Description This API is for change passsword user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param old_password body string true "Old password from user data"
// @Param new_password body string true "New Password for user""
// @Success 201 {string} Password changed susccesfully
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/auth/change-password [post]
func (h *authHandler) ChangePassword(c *gin.Context) {
	type ChangePasswordInput struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	var input ChangePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id") // Assuming JWT middleware sets userID in context
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := h.userUsecase.ChangePassword(userID.(uint), input.OldPassword, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// GetCurrentUser godoc
// @Summary Get user if logged in
// @Description Get user information by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/detail-user [get]
func (h *authHandler) GetCurrentUser(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert userID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Get user data by user ID
	user, err := h.userUsecase.GetUserByID(userIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send user data as response
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func defineRoles(role string) (uint, error) {
	var err error

	switch role {
	case constant.RoleStudent:
		return constant.RoleStudentID, err
	case constant.RoleTeacher:
		return constant.RoleTeacherID, err
	default:
		return 0, errors.New("Invalid role name")
	}
}
