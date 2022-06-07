package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/routes"
	"github.com/BerIincat/shopapi/utils"

	_ "github.com/lib/pq"
)

func main() {
	database.Connect()
	//ins := database.DbControllers{}
	//ins.Init(&ins)
	dbCon := database.New()

	fmt.Print(utils.GetEnv("APP_PORT"))
	log.Fatal(http.ListenAndServe(utils.GetEnv("APP_PORT"), routes.Setup(dbCon)))
}
