package controller

import (
	"fmt"
	"log"

	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/validation"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("CreateUser controller called")
		var userRequest request.UserRequest

		if err := c.ShouldBindJSON(&userRequest); err != nil {
			log.Printf("Error trying to marshal object,error=%s\n", err.Error())
			errRest := validation.ValidateUserError(err)
		

				c.JSON(errRest.Code,errRest)
				return
		}

		fmt.Println(userRequest)

}
