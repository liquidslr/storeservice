package main

import (
	"fmt"

	"github.com/liquidslr/storeservice/db"

	"github.com/liquidslr/storeservice/routes"
)

func createDB() {
	routes.DBClient = &db.BoltDB{}
	routes.DBClient.Initialize()
	fmt.Println("Db instance created")
}

func main() {
	createDB()
	routes.Server("6767")
}
