package main
import(
	"fmt" //library for Print
	"net/http"
	"time"
	"html/template"
)


func mostrarHtml(respuesta http.ResponseWriter,solicitud *http.Request){

	http.ServeFile(respuesta,solicitud,"cookiesH.html")

}
func putCookie(respuesta http.ResponseWriter,solicitud *http.Request){

	oreo1:= http.Cookie{

		Name: "language",
		Value: "frances",
		Expires: time.Now().Add(10*time.Second),

	}
	oreo2:= http.Cookie{
		Name:"other",
		Value: "galletita",

	}
	http.SetCookie(respuesta,&oreo1)
	http.SetCookie(respuesta,&oreo2)

}
func obtenerCookie(respuesta http.ResponseWriter,solicitud *http.Request){
	datos:=solicitud.Header["Cookie"]
	fmt.Fprint(respuesta, datos)
}

func localStorage(respuesta http.ResponseWriter,solicitud *http.Request){
 	if solicitud.Method == "POST" {
 		ajax_post_data := solicitud.FormValue("ajax_post_data")
 		fmt.Println("Receive ajax post data string ", ajax_post_data)
 	}
}

func main() {
	http.HandleFunc("/storage",localStorage)
	http.HandleFunc("/poner",putCookie)
	http.HandleFunc("/obtener",obtenerCookie)
	fmt.Printf("hello, world\n")
	http.HandleFunc("/vue",mostrarHtml)
	http.ListenAndServe(":8080", nil )
}
