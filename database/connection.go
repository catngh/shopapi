package database

import (
	"database/sql"

	"github.com/BerIincat/shopapi/utils"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	db_name := utils.GetEnv("DB_NAME")
	db_user := utils.GetEnv("DB_USER")
	db_pass := utils.GetEnv("DB_PASS")
	db_port := utils.GetEnv("DB_PORT")
	db_url := db_user + ":" + db_pass + "@tcp(127.0.0.1:" + db_port + ")/" + db_name
	db, err := sql.Open("mysql", db_url)
	utils.CheckError(err)
	DB = db

}
func Close() {
	DB.Close()
}
