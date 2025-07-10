package controller

import (
	"fmt"

	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/model/request"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

		var userRequest request.UserRequest

		if err := c.ShouldBindJSON(&userRequest); err != nil {
			restErr := rest_err.NewBadRequestError(
				fmt.Sprintf("There are some incorrect fields, error =%s",err.Error()))

				c.JSON(restErr.Code,restErr)
				return
		}

		fmt.Println(userRequest)

}
