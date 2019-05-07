package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
)

var caja *sessions.CookieStore

func init() {
	caja = sessions.NewCookieStore([]byte("llave"))
}

func revisar(respuesta http.ResponseWriter, solicitud *http.Request) {
	var e error
	sesion, e := caja.Get(solicitud, "nombre")
	if e != nil {
		log.Fatal("error al ejecutar la consulta: ", e)
	}
	if (sesion.Values["nombre"] == nil) {
		fmt.Fprint(respuesta, "Aqui ño")
	} else {
		fmt.Fprint(respuesta, "Wel come gallegos")
	}
}


func activar(respuesta http.ResponseWriter, solicitud *http.Request) {
	var e error
	sesion, e := caja.Get(solicitud, "nombre")
	if e != nil {
		log.Fatal("error al ejecutar la consulta: ", e)
	}
	sesion.Values["nombre"] = "Hello its me"
	sesion.Save(solicitud, respuesta)
}

func main() {
	http.HandleFunc("/revisar", revisar)
	http.HandleFunc("/activar", activar)
	http.ListenAndServe(":8080", nil)
}
