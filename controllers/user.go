package controllers

import (
	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	db := database.DB
	var newUser, queriedUser models.User
	err := c.BindJSON(&newUser)

	// Validating request frorm and user input
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request"}, c) {
		return
	}
	if !utils.ValidateEmailPwd(newUser.Email, newUser.Password, c) {
		return
	}
	result := db.Where("email=?", newUser.Email).First(&queriedUser)
	//err = db.Get(&queriedUser, "SELECT * FROM usr WHERE email=?", newUser.Email)

	// Email not found
	if utils.PrintErrIfAny(result.Error, 400, gin.H{"error": "email not found"}, c) {
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
func Register(c *gin.Context) {
	db := database.DB
	var newUser models.User
	err := c.BindJSON(&newUser)

	// Validating user input
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}
	if !utils.ValidateEmailPwd(newUser.Email, newUser.Password, c) {
		return
	}

	// Preparing query
	hashed, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 8)
	newUser.Password = string(hashed)

	// Populated newUser struct
	result := db.Select("Email", "Password", "Role").Create(&newUser)
	//res, err := db.Exec("INSERT INTO usr (email,password,role) VALUES ('?','?','?')", newUser.Email, newUser.Password, newUser.Role)
	// Duplicated entry
	if result.Error != nil {
		if result.Error.Error()[6:10] == "1062" {
			c.JSON(400, gin.H{"error": "email existed"})
		} else {
			c.JSON(500, gin.H{"error": "database error"})
		}
		return
	}

	// Return registered info
	c.JSON(201, gin.H{"userId": newUser.UserID, "email": newUser.Email, "role": newUser.Role})

}
