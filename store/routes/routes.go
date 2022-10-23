package routes

import (
	"net/http"
	"studies/alura/store/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/products", controllers.Products)
	http.HandleFunc("/products/delete", controllers.Delete)
	http.HandleFunc("/products/edit", controllers.Edit)
	http.HandleFunc("/products/insert", controllers.Insert)
	http.HandleFunc("/products/register", controllers.Register)
	http.HandleFunc("/products/update", controllers.Update)
}
