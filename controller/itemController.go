package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rasyidev/dts-h8-assignment-ii/models/domain"
	"gorm.io/gorm"
)

func (db *gorm.DB) Create(ctx *gin.Context) {
	item := domain.Item{}
	ctx.ShouldBindJSON(&item)

	db.Create(item)

}
