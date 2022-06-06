package database

import (
	"github.com/BerIincat/shopapi/models"
)

//var User userControl

type userControl struct {
}

func User() *userControl {
	return &userControl{}
}

func (u userControl) GetByEmail(email string) (models.User, error) {
	var queriedUser models.User
	result := DB.Where("email=?", email).First(&queriedUser)
	return queriedUser, result.Error
}
func (u userControl) GetById(uid string) (models.User, error) {
	queriedUser := models.User{}
	result := DB.Where("userId=?", uid).First(&queriedUser)
	return queriedUser, result.Error
}

func (u userControl) Create(usr models.User) (models.User, error) {
	newUsr := usr
	result := DB.Select("Email", "Password", "Role").Create(&newUsr)
	return newUsr, result.Error
}

func (u userControl) IdExist(uid string) bool {
	user := models.User{}
	res := DB.Where("userId=?", uid).First(&user)
	if res.Error != nil {
		return false
	}
	return true
}
