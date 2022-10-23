package controllers

import (
	"net/http"
	"strconv"
	"studies/alura/store/models"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

const RedirectStatus = 301

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index", nil)
}

func Products(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()

	temp.ExecuteTemplate(w, "products", allProducts)
}

func Register(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "register", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, _ := strconv.ParseFloat(price, 64)
		quantityConverted, _ := strconv.Atoi(quantity)

		models.CreateProduct(name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/products", RedirectStatus)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteProduct(id)

	http.Redirect(w, r, "/products", RedirectStatus)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	product := models.GetProduct(id)

	temp.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, _ := strconv.ParseFloat(price, 64)
		quantityConverted, _ := strconv.Atoi(quantity)

		models.UpdateProduct(id, name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/products", RedirectStatus)
}
