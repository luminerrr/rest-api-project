package main

import (
	"rest-api-project/database"
	"rest-api-project/routers"
)


func main() {
	database.StartDB()
	PORT := ":8080"

	routers.StartServer().Run(PORT)

}