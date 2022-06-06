package utils

import (
	"database/sql"
	"net/mail"
	"os"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/gin-gonic/gin"
	"github.com/go-passwd/validator"
	"github.com/joho/godotenv"
)

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func GetEnv(key string) string {
	err := godotenv.Load(".env")
	CheckError(err)
	return os.Getenv(key)
}
func UserIdExist(uid string) bool {
	db := database.DB
	user := models.User{}
	// q := "SELECT * FROM usr WHERE userId=" + uid
	// row := db.QueryRow(q)
	res := db.Where("userId=?", uid).First(&user)
	if res.Error == sql.ErrNoRows || res.Error != nil {
		return false
	}
	return true
}
func ProductIdExist(pid string) bool {
	db := database.DB
	product := models.Product{}
	// q := "SELECT * FROM product WHERE productId=" + pid
	// row := db.QueryRow(q)
	res := db.Where("productId=?", pid).First(&product)
	if res.Error == sql.ErrNoRows || res.Error != nil {
		return false
	}
	return true
}

func PrintErrIfAny(err error, code int, mess gin.H, c *gin.Context) bool {
	if err != nil {
		c.JSON(code, mess)
		return true
	}
	return false
}
func EmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func PwdValid(pwd string) bool {
	passwordValidator := validator.New(validator.MinLength(6, nil))
	err := passwordValidator.Validate(pwd)
	return err == nil
}

func ValidateEmailPwd(email string, pwd string, c *gin.Context) bool {
	if !EmailValid(email) {
		c.JSON(400, gin.H{"error": "invalid email"})
		return false
	}
	if !PwdValid(pwd) {
		c.JSON(400, gin.H{"error": "password required at least 6 characters"})
		return false
	}
	return true
}
