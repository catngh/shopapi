package database

import (
	//"github.com/BerIincat/shopapi/utils" import cycle
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func Connect() {
	db_name := getEnv("DB_NAME")
	db_user := getEnv("DB_USER")
	db_pass := getEnv("DB_PASS")
	db_port := getEnv("DB_PORT")
	db_url := db_user + ":" + db_pass + "@tcp(127.0.0.1:" + db_port + ")/" + db_name
	db, _ := sqlx.Connect("mysql", db_url)
	DB = db
}
func Close() {
	DB.Close()
}
func getEnv(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
