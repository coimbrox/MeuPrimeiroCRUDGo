package service

import (
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/model"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/rest_err"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}


type userDomainService struct{
	
}



type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string,model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}