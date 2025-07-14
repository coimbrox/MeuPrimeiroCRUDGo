package model

import (
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/logger"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "CreateUser"))

	ud.EncryptPassword()
	return nil 
}