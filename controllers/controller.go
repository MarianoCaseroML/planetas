package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"planetas/server"
	"planetas/utils"
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
	registrarRutas(router)

	//Cuando inicio verifico si la base existe, si no existe la creo
	server.CrearBaseSiNoExiste()

	log.Println("Listening on port " + serverPort)
	//inicio server
	log.Fatalln(http.ListenAndServe(":" + serverPort, router))

}

//Por si hay que regenerar la base
const regenerarDB = "/api/1/regenerarDB"
const clima ="/api/1/clima"
const periodo = "/api/1/periodo"

//registro APIs
func registrarRutas(router *mux.Router) {

	//get para healthCheck para saber si estoy vivo
	router.Path("/health").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte(`{ "status": "ok" }`))
	})

	//regenerarDB
	router.Path(regenerarDB).Methods("GET").HandlerFunc(RecrearDB)
	//Consulta de clima por dia
	router.Path(clima).Methods("GET").HandlerFunc(ConsultaClimaPorDia)
	//Consulta por periodo de clima
	router.Path(periodo).Methods("GET").HandlerFunc(consultaPorPeriodo)

	//Converti el readme a html para servirlo como pagina statica y mostrar "algo" cuando se accede al servicio
	wf := utils.CurrentWF()
	//agrego fix para que funcione el html al hacer debug
	var htmlFile string = wf + "/static/"
	if (!utils.Fileexists(htmlFile + "index.html")) {
		htmlFile = wf + "/planetas/static/"
	}
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(htmlFile))))

}

//para crear/regenerar la base
func RecrearDB(w http.ResponseWriter, r *http.Request) {
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