package main 

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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
		Name  string
		Price float64
		Image string
	}

	products := []Product{
		{Name: "Product 1", Price: 10.99, Image: "/static/img/products/sweatcap1.webp"},
		{Name: "Product 2", Price: 15.99, Image: "/static/img/products/sweatcap2.webp"},
		{Name: "Product 3", Price: 20.99, Image: "/static/img/products/sweatcap3.webp"},
		{Name: "Product 4", Price: 25.99, Image: "/static/img/products/sweatcap4.webp"},
		{Name: "Product 3", Price: 20.99, Image: "/static/img/products/sweat1.webp"},
		{Name: "Product 3", Price: 20.99, Image: "/static/img/products/jean1.webp"},
		{Name: "Product 3", Price: 20.99, Image: "/static/img/products/pants1.webp"},
	}

	http.HandleFunc("/temp/Homepage", func(w http.ResponseWriter, r *http.Request) {
		listTemplate.ExecuteTemplate(w, "Homepage", products)
	})

	http.ListenAndServe("localhost:8000", nil)
}