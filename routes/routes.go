package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes (router *gin.Engine) {
	router.POST("/users/signup")
	router.POST("/users/login")
	router.POST("/admin/product")
	router.GET("/users/product")
	router.GET("/users/search")
}