package controllers

import (
	"fmt"
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController struct to bind methods to.
type UserController struct {
	userRepo *repositories.UserRepository
}

// NewUserController creates a new UserController with the given UserRepository.
func NewUserController(r *repositories.UserRepository) *UserController {
	return &UserController{userRepo: r}
}

// GetUserByID handles the request to get a user by their ID.
func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.userRepo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByID handles the request to get a user by their ID.
func (uc *UserController) GetUserList(c *gin.Context) {

	userlist, err := uc.userRepo.ListUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userlist)
}

// CreateUser handles creating a new user.
func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("user: %+v", newUser)
	user, err := uc.userRepo.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles updating an existing user.
func (uc *UserController) UpdateUser(c *gin.Context) {
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.userRepo.UpdateUser(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles the deletion of a user.
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := uc.userRepo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// RegisterUserRoutes registers the user-related routes to the router.
func RegisterUserRoutes(router *gin.Engine, uc *UserController) {
	router.GET("/user/:id", uc.GetUserByID)
	router.POST("/user", uc.CreateUser)
	router.PUT("/user", uc.UpdateUser)
	router.DELETE("/user/:id", uc.DeleteUser)
	router.GET("/user/list", uc.GetUserList)
}
