package service

import (
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/logger"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/model"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "CreateUser"))

	userDomain.EncryptPassword()
	return nil 
}