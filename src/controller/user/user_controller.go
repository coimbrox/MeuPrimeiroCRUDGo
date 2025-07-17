package controller

import (
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/model/service"
	"github.com/gin-gonic/gin"
)


func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}


type userControllerInterface struct {
	service service.UserDomainService
}