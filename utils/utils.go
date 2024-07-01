package utils

import "github.com/gin-gonic/gin"

type Controller interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func Resources(r gin.RouterGroup, name string, c Controller) {
	r.GET("/"+name, c.GetAll)
	r.GET("/"+name+"/:id", c.GetById)
	r.POST("/"+name, c.Create)
	r.PUT("/"+name+"/:id", c.Update)
	r.DELETE("/"+name+"/:id", c.Delete)
}
