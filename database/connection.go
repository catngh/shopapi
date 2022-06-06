package database

import (
	//"github.com/BerIincat/shopapi/utils" import cycle
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DC *DbControllers

func Connect() {
	db_name := getEnv("DB_NAME")
	db_user := getEnv("DB_USER")
	db_pass := getEnv("DB_PASS")
	db_port := getEnv("DB_PORT")
	db_url := db_user + ":" + db_pass + "@tcp(127.0.0.1:" + db_port + ")/" + db_name
	db, _ := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	DB = db
}
func getEnv(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
func (d *DbControllers) Init(ins *DbControllers) {
	d.UserContr = userControl{}
	d.ProductContr = productControl{}
	d.CartContr = cartControl{}
	d.OrderContr = orderControl{}
	DC = ins
}

type DbControllers struct {
	UserContr    userControl
	ProductContr productControl
	CartContr    cartControl
	OrderContr   orderControl
}
