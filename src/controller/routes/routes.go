package routes

import (
	controller "github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/user/:userId", userController.GetUserById)
	r.GET("/userEmail/:userEmail", userController.GetUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
}