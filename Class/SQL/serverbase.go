package main
import(
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Timeline struct {
    Id int
    Titulo string
}

var db *sql.DB

func init() {
	var e error
	db, e = sql.Open("mysql", "root@/golang")
	if e != nil {
		log.Fatal("Error al conectar con la base de datos: ", e)
	} 
}

func basedatos(respuesta http.ResponseWriter, solicitud *http.Request) {
	var e error
	filas, e := db.Query("SELECT * FROM bestmovies")
	if e != nil {
		log.Fatal("Error al ejecutar la consulta: ", e)
	}
	defer filas.Close()
	for filas.Next() {
		timeline := Timeline{}
		e = filas.Scan(&timeline.Id, &timeline.Titulo)
		if e != nil {
			panic(e)
		}
		fmt.Println(timeline)
	}
	e = filas.Err()
	if e != nil {
		panic(e)
	}
	
}

func main() {
	http.HandleFunc("/basedatos", basedatos)
	http.ListenAndServe(":8080", nil)
}