package main

import (
	"net/http"
)

func showIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/index.html")
}

func showAbout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/about.html")
}

func showCatalog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/catalog.html")
}

func showService(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../pages/service.html")
}

func main() {
	css := http.FileServer(http.Dir("../css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	js := http.FileServer(http.Dir("../vendor"))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", js))
	im := http.FileServer(http.Dir("../img"))
	http.Handle("/img/", http.StripPrefix("/img/", im))
	pg := http.FileServer(http.Dir("../pages"))
	http.Handle("/pages/", http.StripPrefix("/pages/", pg))
	http.HandleFunc("/index", showIndex)
	http.HandleFunc("/about", showAbout)
	http.HandleFunc("/catalog", showCatalog)
	http.HandleFunc("/service", showService)
	http.ListenAndServe(":8080", nil)
	

}
