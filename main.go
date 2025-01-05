// Executable file will be created from this package main
// main.go
// Basically, when main.go runs then the application is started, this process includes the database connection and mux router instance connection
// as described in this file.
package main

import (
	"log"

	"myvet-v2-api/context"
	"myvet-v2-api/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	c := context.Base{}
	log.Println("MYVET VERSION 2.0 API")
	log.Println("Server is starting")
	// Set nicer log output.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error

	// Read the service configuration.
	err = c.ReadConf()
	check(err)

	// Initialize the database connection.
	err = c.DBInit()
	check(err)

	c.Router = mux.NewRouter()
	err = handlers.InitRoutes(&c)
	check(err)
	c.Run()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
