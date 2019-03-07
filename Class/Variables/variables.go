package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Image string
	Title string
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	css := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	js := http.FileServer(http.Dir("vendor"))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", js))
	im := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", im))
	tmpl := template.Must(template.ParseFiles("variables.html"))

	http.HandleFunc("/variables", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Image: "/img/balinese-cat.jpeg", Title: "Balinese"},
				{Image: "/img/bengal-cat.jpg", Title: "Bengal"},
				{Image: "/img/russian-blue.jpg", Title: "Russian"},
				{Image: "/img/persian-cat.jpg", Title: "Persian"},
				{Image: "/img/british-shorthair-cat.jpeg", Title: "British shorthair"},
				{Image: "/img/himalayan-cat.jpg", Title: "Himalayan"},
				{Image: "/img/scottish-fold-cat.jpg", Title: "Scottish fold"},
				{Image: "/img/siamese-cat.jpg", Title: "Siamese"},
				{Image: "/img/siberian-cat.jpg", Title: "Siberian"},
				{Image: "/img/snowshoe.jpg", Title: "Snowshoe"},
				{Image: "/img/sphynx-cat.png", Title: "Sphynx"},
				{Image: "/img/toyger-cat.jpeg", Title: "Toyger"},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}