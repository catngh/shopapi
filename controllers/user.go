package controllers

import (
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c *gin.Context) {
	var newUser, queriedUser models.User
	err := c.BindJSON(&newUser)

	// Validating request frorm and user input
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request"}, c) {
		return
	}
	if !utils.ValidateEmailPwd(newUser.Email, newUser.Password, c) {
		return
	}
	//queriedUser, err = database.User.GetByEmail(newUser.Email)
	queriedUser, err = h.db.User.GetByEmail(newUser.Email)
	// Email not found
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "email not found"}, c) {
		return
	}

	// Comparing input password and db record
	err = bcrypt.CompareHashAndPassword([]byte(queriedUser.Password), []byte(newUser.Password))
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "password incorrect"}, c) {
		return
	}
	// Correct email and pass
	c.JSON(200, gin.H{"userId": queriedUser.UserID, "email": queriedUser.Email, "role": queriedUser.Role})

}
func (h *Handler) Register(c *gin.Context) {
	var newUser models.User
	err := c.BindJSON(&newUser)

	// Validating user input
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}
	if !utils.ValidateEmailPwd(newUser.Email, newUser.Password, c) {
		return
	}

	// Hash user password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 8)
	newUser.Password = string(hashed)

	// Populated newUser struct
	//newUser, err = h.db.User().Create(newUser)
	newUser, err = h.db.User.Create(newUser)
	// Duplicated entry
	if err != nil {
		if err.Error()[6:10] == "1062" {
			c.JSON(400, gin.H{"error": "email existed"})
		} else {
			c.JSON(500, gin.H{"error": "database error"})
		}
		return
	}

	// Return registered info
	c.JSON(201, gin.H{"userId": newUser.UserID, "email": newUser.Email, "role": newUser.Role})

}
