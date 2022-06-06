package utils

import (
	"net/mail"
	"os"

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
