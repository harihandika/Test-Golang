package main

import (
	"fmt"
	"golang-test/database"
	"golang-test/pkg/mysql"
	"golang-test/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/").Subrouter())

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	fmt.Println("server running localhost:3030")
	http.ListenAndServe("localhost:3030", handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
