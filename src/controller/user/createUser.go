package controller

import (
	"net/http"

	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/logger"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/validation"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/model/request"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/model"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
		var userRequest request.UserRequest

		if err := c.ShouldBindJSON(&userRequest); err != nil {
			logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
			errRest := validation.ValidateUserError(err)
		

				c.JSON(errRest.Code,errRest)
				return
		}


		domain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)


		service := service.NewUserDomainService()

		if err := service.CreateUser(domain); err != nil {
			c.JSON(err.Code, err)
			return
		}



		logger.Info(
		"CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

		c.JSON(http.StatusOK,"")
}
