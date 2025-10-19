package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	listTemplate, errTemp := template.ParseGlob("./templates/*.html")
	if errTemp != nil {
		fmt.Println(errTemp.Error())
		os.Exit(1)
	}

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	type Product struct {
		ID    int
		Name  string
		Price float64
		Image string
	}

	products := []Product{
		{ID: 1, Name: "PULL A CAPUCHE BEIGE", Price: 89.99, Image: "/static/img/products/sweatcap1.webp"},
		{ID: 2, Name: "PULL A CAPUCHE NOIR", Price: 79.99, Image: "/static/img/products/sweatcap2.webp"},
		{ID: 3, Name: "PULL A CAPUCHE VERT", Price: 79.99, Image: "/static/img/products/sweatcap3.webp"},
		{ID: 4, Name: "PULL A CAPUCHE BLEU", Price: 99.99, Image: "/static/img/products/sweatcap4.webp"},
		{ID: 5, Name: "SWEAT NOIR", Price: 69.99, Image: "/static/img/products/sweat1.webp"},
		{ID: 6, Name: "JEAN DROIT", Price: 39.99, Image: "/static/img/products/jean1.webp"},
		{ID: 7, Name: "PANTALON CARGO", Price: 49.99, Image: "/static/img/products/pants1.webp"},
	}

	http.HandleFunc("/temp/Homepage", func(w http.ResponseWriter, r *http.Request) {
		listTemplate.ExecuteTemplate(w, "Homepage", products)
	})

	http.HandleFunc("/temp/Product", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Product ID missing", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var selected *Product
		for _, p := range products {
			if p.ID == id {
				selected = &p
				break
			}
		}

		if selected == nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		listTemplate.ExecuteTemplate(w, "Product", selected)
	})

	http.HandleFunc("/temp/Add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			listTemplate.ExecuteTemplate(w, "AddProduct", nil)
			return
		}

		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			priceStr := r.FormValue("price")

			if name == "" || priceStr == "" {
				http.Error(w, "Nom et prix sont requis", http.StatusBadRequest)
				return
			}

			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil {
				http.Error(w, "Prix invalide", http.StatusBadRequest)
				return
			}

			name = strings.TrimSpace(strings.ToUpper(name))

			newID := len(products) + 1
			newProduct := Product{
				ID:    newID,
				Name:  name,
				Price: price,
				Image: "/static/img/products/sweatcap5.webp",
			}

			products = append(products, newProduct)

			http.Redirect(w, r, "/temp/Homepage", http.StatusSeeOther)
			return
		}
	})

	http.ListenAndServe("localhost:8000", nil)
}
