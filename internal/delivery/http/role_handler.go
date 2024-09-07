package http

import (
	"net/http"
	"strconv"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/usecases"
	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	GetAllRoles(c *gin.Context)
	GetRoleByID(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}

type roleHandler struct {
	RoleUc usecases.RoleUsecase
}

func NewRoleHandler(uc usecases.RoleUsecase) RoleHandler {
	return &roleHandler{
		RoleUc: uc,
	}
}

func (h *roleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.RoleUc.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

func (h *roleHandler) GetRoleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.RoleUc.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

// Godoc
// @Summary Create role
// @Description Create role
// @Tags Role
// @Accept json
// @Produce json
// @Param name string true "Create role"
// @Success 201 {object} models.Role
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
func (h *roleHandler) CreateRole(c *gin.Context) {
	var input models.Role

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.RoleUc.CreateRole(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"role": role})
}

// Update role
func (h *roleHandler) UpdateRole(c *gin.Context) {
	var input models.Role
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	roleId, err := h.RoleUc.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	input.ID = roleId.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.RoleUc.UpdateRole(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

// Delete role
func (h *roleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.RoleUc.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	err = h.RoleUc.DeleteRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
