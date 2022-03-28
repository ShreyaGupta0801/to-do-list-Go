package main

import (
	"fmt"
	"golang-react-to-do/server/database"
	"golang-react-to-do/server/router"
	"log"
	"net/http"
)

func main() {

	database.CreateDbInstance()
	r := router.Router()
	fmt.Println("Server started on port 3306..")
	log.Fatal(http.ListenAndServe(":3306", r))

}
