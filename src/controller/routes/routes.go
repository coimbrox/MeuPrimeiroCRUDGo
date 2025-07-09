package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId")
	r.GET("/userEmail/:userEmail")
	r.POST("/user")
	r.PUT("/user")
	r.DELETE("/user/:userId")
}