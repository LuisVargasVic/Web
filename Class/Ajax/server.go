package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ajax_text.html")
}

func mostrarTexto(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "nojson.txt")
}

func mostrarArreglo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Method)
}

func atenderForma(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Fprint(w, r.Form)
	} else {
		http.ServeFile(w, r, "ajax_text.html")
	}
}

func escogerDelArray(w http.ResponseWriter, r *http.Request) {
	arreglo := [4]string{"Is", "An", "Array", "?"}
	x, err := strconv.Atoi(r.FormValue("position"))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	if x == 0 {
		fmt.Fprintf(w, arreglo[0])
	} else if x == 1 {
		fmt.Fprintf(w, arreglo[1])
	} else if x == 2 {
		fmt.Fprintf(w, arreglo[2])
	} else if x == 3 {
		fmt.Fprintf(w, arreglo[3])
	} else {
		fmt.Fprintf(w, "Out of bounds")
	}
}

func main() {

	http.HandleFunc("/ajax", mostrarHTML)
	http.HandleFunc("/file", mostrarTexto)
	http.HandleFunc("/array", mostrarArreglo)
	http.HandleFunc("/login", atenderForma)
	http.HandleFunc("/postArray", escogerDelArray)
	http.ListenAndServe(":8080", nil)

}
