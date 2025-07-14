package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/rest_err"
)


func NewUserDomain(email, password, name string, age int) *UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int
}


func(ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	GetUserByEmail(string) (*userDomain, *rest_err.RestErr)
	GetUserById(string) (*userDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}