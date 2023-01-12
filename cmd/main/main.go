package main

import (
	"net/http"
	"fmt"
	"github.com/dee-d-dev/gosql/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":7000", r)
	fmt.Println("Server running on port 7000")
}
