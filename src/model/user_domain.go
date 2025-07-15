package model

import (
	"crypto/md5"
	"encoding/hex"
)


type UserDomainInterface interface {
	GetAge() int
	GetEmail() string
	GetName() string
	GetPassword() string
	EncryptPassword()
}


func NewUserDomain(
	email, password, name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
type userDomain struct {
	email    string
	password string
	name     string
	age      int
}

func (ud *userDomain)  GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int {
	return ud.age
}




func(ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

