package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/md-shadhin-mia/go-crud/models"
	"gorm.io/gorm"
)

type DemoController struct {
	DB *gorm.DB
}

// NewDemoController creates a new instance of DemoController with dependencies injected.
func NewDemoController(db *gorm.DB) *DemoController {
	return &DemoController{DB: db}
}

// @Summary Get all demos
// @Description Get all demos
// @Tags Demo
// @Accept json
// @Produce json
// @Success 200 {array} models.Demo
// @Router /users [get]
func (u *DemoController) GetAll(c *gin.Context) {
	var demos []models.Demo
	if err := u.DB.Find(&demos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch demos"})
		return
	}
	c.JSON(http.StatusOK, demos)
}

// @Summary Get a demo by id
// @Description Get a demo by id
// @Tags Demo
// @Accept json
// @Produce json
// @Param id path string true "Demo ID"
// @Success 200 {object} models.Demo
// @Router /users/{id} [get]
func (u *DemoController) GetById(c *gin.Context) {
	id := c.Param("id")
	var demo models.Demo
	if err := u.DB.First(&demo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Demo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch demo"})
		}
		return
	}
	c.JSON(http.StatusOK, demo)
}

// @Summary Create a demo
// @Description Create a demo
// @Tags Demo
// @Accept json
// @Produce json
// @Param demo body models.Demo true "Demo data"
// @Success 201 {object} models.Demo
// @Router /users [post]
func (u *DemoController) Create(c *gin.Context) {
	var input models.Demo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create demo"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// @Summary Update a demo by id
// @Description Update a demo by id
// @Tags Demo
// @Accept json
// @Produce json
// @Param id path string true "Demo ID"
// @Param demo body models.Demo true "Demo data"
// @Success 200 {object} models.Demo
// @Router /users/{id} [put]
func (u *DemoController) Update(c *gin.Context) {
	id := c.Param("id")
	var demo models.Demo
	if err := u.DB.First(&demo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Demo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch demo"})
		}
		return
	}

	if err := c.ShouldBindJSON(&demo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.DB.Save(&demo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update demo"})
		return
	}

	c.JSON(http.StatusOK, demo)
}

// @Summary Delete a demo by id
// @Description Delete a demo by id
// @Tags Demo
// @Accept json
// @Produce json
// @Param id path string true "Demo ID"
// @Success 200
// @Router /users/{id} [delete]
func (u *DemoController) Delete(c *gin.Context) {
	id := c.Param("id")
	var demo models.Demo
	if err := u.DB.First(&demo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Demo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch demo"})
		}
		return
	}

	if err := u.DB.Delete(&demo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete demo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Demo deleted"})
}
