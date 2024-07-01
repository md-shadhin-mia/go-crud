package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/md-shadhin-mia/go-crud/models"
)

type UserController struct {
	DB *gorm.DB
}

// NewUserController creates a new instance of UserController with dependencies injected.
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (u *UserController) GetAll(c *gin.Context) {
	var users []models.User
	if err := u.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Get a user by id
// @Description Get a user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (u *UserController) GetById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Create a user
// @Description Create a user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Router /users [post]
func (u *UserController) Create(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// @Summary Update a user by id
// @Description Update a user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (u *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Delete a user by id
// @Description Delete a user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /users/{id} [delete]
func (u *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	if err := u.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
