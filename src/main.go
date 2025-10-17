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

	http.HandleFunc("/temp/Homepage", func(w http.ResponseWriter, r *http.Request) {
		listTemplate.ExecuteTemplate(w, "Homepage", nil)
	})

	http.ListenAndServe("localhost:8000", nil)
}