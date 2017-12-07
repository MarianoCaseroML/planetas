package controllers

import (
	"net/http"
	//"planetas/entity"
	//"encoding/json"
	"planetas/external/github.com/gorilla/mux"
	"log"
	"planetas/server"
	"path/filepath"
	"os"
)

const serverPort string = "8080"

type Service struct {
	router *mux.Router
	port   int
}

func (service *Service) Start() {
	//new router
	router := mux.NewRouter()
	//registro APIs en router
	registerRoutes(router)

	//Cuando inicio verifico si la base existe, si no existe la creo
	server.CreateDBifNotExists()

	log.Println("Listening on port " + serverPort)
	//inicio server
	log.Fatalln(http.ListenAndServe(":" + serverPort, router))

}

//Por si hay que regenerar la base
const regenerarDB = "/api/1/regenerarDB"
const clima ="/api/1/clima"
const periodo = "/api/1/periodo"

//registro APIs
func registerRoutes(router *mux.Router) {

	//get para healthCheck para saber si estoy vivo
	router.Path("/health").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte(`{ "status": "ok" }`))
	})

	//regenerarDB
	router.Path(regenerarDB).Methods("GET").HandlerFunc(RecreateDB)
	//Consulta de clima por dia
	router.Path(clima).Methods("GET").HandlerFunc(ConsultaClimaPorDia)
	//Consulta por periodo de clima
	router.Path(periodo).Methods("GET").HandlerFunc(consultaPorPeriodo)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir + "/static/"))))

}

func RecreateDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	server.RegenerarBase()
	w.Write([]byte(`{ "status": "ok" }`))
}


func ConsultaClimaPorDia (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	q := r.URL.Query()
	var dia = q.Get("dia")
	response := server.ClimaPorDia(dia)
	w.Write(response)
}

func consultaPorPeriodo (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	q := r.URL.Query()
	var climaBuscado = q.Get("clima")
	response := server.ClimaPorPeriodo(climaBuscado)
	w.Write(response)
}