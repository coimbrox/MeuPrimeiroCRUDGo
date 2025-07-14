package controller

import (
	"fmt"

	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/logger"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/validation"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

		logger.Info(
		"CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

		fmt.Println(userRequest)

}
