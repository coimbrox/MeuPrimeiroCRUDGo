package view

import (
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/model/response"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		Age:      userDomain.GetAge(),
	}
}