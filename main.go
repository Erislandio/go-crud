package main

import (
	"fmt"
	"net/http"

	"github.com/alura-curso-crud-api/api/routes"
	_ "github.com/lib/pq"
)

func main() {

	routes.LoadRoutes()
	fmt.Println("Server online")
	http.ListenAndServe(":8080", nil)

}
