package main

import (
	"html/template"
	"net/http"
)

func MostrarPlantilla(respuesta http.ResponseWriter, solicitud *http.Request) {
	p, _ := template.ParseFiles("template.html")
	p.Execute(respuesta, "Ya casi nos vamoooos!!!!!!")
}

func main() {
	http.HandleFunc("/plantilla", MostrarPlantilla)
	http.ListenAndServe(":8080", nil)
}
