package main

import (
	"github.com/Lincolnbiancard/gin-api/database"
	"github.com/Lincolnbiancard/gin-api/routes"
)

func main() {
	database.ConnectToDB()
	routes.HandleRequests()
}