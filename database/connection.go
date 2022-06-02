package database

import (
	"github.com/BerIincat/shopapi/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() {
	db_name := utils.GetEnv("DB_NAME")
	db_user := utils.GetEnv("DB_USER")
	db_pass := utils.GetEnv("DB_PASS")
	db_port := utils.GetEnv("DB_PORT")
	db_url := db_user + ":" + db_pass + "@tcp(127.0.0.1:" + db_port + ")/" + db_name
	db, err := sqlx.Connect("mysql", db_url)
	utils.CheckError(err)
	DB = db
}
func Close() {
	DB.Close()
}
