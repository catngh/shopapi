package controllers

import (
	"net/mail"
	"strconv"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/gin-gonic/gin"
	"github.com/go-passwd/validator"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	db := database.DB
	var newUser models.User
	err := c.BindJSON(&newUser)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
	} else if !emailValid(newUser.Email) {
		c.JSON(400, gin.H{"error": "invalid email"})
	} else if !pwdValid(newUser.Password) {
		c.JSON(400, gin.H{"error": "password required at least 6 characters"})
	} else {
		q := "SELECT * FROM usr WHERE email='" + newUser.Email + "'"
		var pwd, id, email, role string
		row := db.QueryRow(q)
		err := row.Scan(&id, &email, &pwd, &role)
		// Email not found
		if err != nil {
			c.JSON(400, gin.H{"error": "user not found"})
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(newUser.Password))
		if err != nil {
			// Email found but wrong pass
			c.JSON(400, gin.H{"error": "password incorrect"})
		} else {
			// Correct email and pass
			c.JSON(200, gin.H{"userId": id, "email": email, "role": role})
		}

	}
}
func Register(c *gin.Context) {
	db := database.DB
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid register form"})
	} else if !emailValid(newUser.Email) {
		c.JSON(400, gin.H{"error": "invalid email"})
	} else if !pwdValid(newUser.Password) {
		c.JSON(400, gin.H{"error": "password required at least 6 characters"})
	} else {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 8)
		newUser.Password = string(hashed)
		q := "INSERT INTO usr (email,password,role) VALUES ('" + newUser.Email + "','" + newUser.Password + "','" + newUser.Role + "')"
		res, err := db.Exec(q)
		// Duplicate entry
		if err != nil {
			if err.Error()[6:10] == "1062" {
				c.JSON(400, gin.H{"error": "email existed"})
			} else {
				c.JSON(500, gin.H{"error": "database error"})
			}
		}
		usrid, err := res.LastInsertId()
		c.JSON(201, gin.H{"userId": string(strconv.FormatInt(usrid, 10)), "email": newUser.Email, "role": newUser.Role})
	}

}

func emailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func pwdValid(pwd string) bool {
	passwordValidator := validator.New(validator.MinLength(6, nil))
	err := passwordValidator.Validate(pwd)
	return err == nil
}
