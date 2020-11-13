package routes

import (
	"net/http"

	"github.com/alura-curso-crud-api/api/controller"
)

func LoadRoutes() {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)

}
