package utils

import (
	"os"

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
