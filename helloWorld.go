package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/comprarmamut", responder)
	http.ListenAndServe(":8080", nil)
}

func responder(respuesta http.ResponseWriter, solicitud *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(respuesta, currentTime.Format("Mon Jan _2 15:04:05 2006"))
}
