package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Context struct {
	Gin *gin.Engine
	DB  *gorm.DB
}
