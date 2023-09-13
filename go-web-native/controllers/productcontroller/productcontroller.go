package productcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()

	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Detail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		idKategori, err := strconv.Atoi(r.FormValue("id_kategori"))

		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))

		if err != nil {
			panic(err)
		}

		product.Nama = r.FormValue("nama")
		product.Kategori.Id = uint(idKategori)
		product.Stock = int64(stock)
		product.Deskripsi = r.FormValue("deskripsi")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product":    product,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		idString := r.FormValue("id")

		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		idKategori, err := strconv.Atoi(r.FormValue("id_kategori"))

		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))

		if err != nil {
			panic(err)
		}

		product.Nama = r.FormValue("nama")
		product.Kategori.Id = uint(idKategori)
		product.Stock = int64(stock)
		product.Deskripsi = r.FormValue("deskripsi")
		product.UpdatedAt = time.Now()

		if ok := productmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if err := productmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)

}
