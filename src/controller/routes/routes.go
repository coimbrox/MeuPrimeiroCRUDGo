package routes

import (
	controller "github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.GetUserById)
	r.GET("/userEmail/:userEmail", controller.GetUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}