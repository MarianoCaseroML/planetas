package server

import (
	"log"
	"planetas/external/github.com/boltdb/bolt"
	"encoding/json"
	"strconv"
	"planetas/entity"
	"planetas/db"
	"planetas/mathCalcs"
	"strings"
)


const cantidadDiasModelo = 3650

var Ferengi entity.Planeta
var Betasoide entity.Planeta
var Vulcano entity.Planeta
var Sol entity.Planeta


//Inicializo los planetas segun los datos el ejercicio
func inicializarPlanetas () {

	Ferengi.Velocidad = 1
	Ferengi.Distancia= 500
	Ferengi.Sentido = -1

	Betasoide.Velocidad = 3
	Betasoide.Distancia = 2000
	Betasoide.Sentido = -1

	Vulcano.Velocidad = 5
	Vulcano.Distancia = 1000
	Vulcano.Sentido = 1

	Sol.Velocidad= 0
	Sol.Distancia = 0
	Sol.Sentido = 0
}

//Para ejecutar al inicio, si la base no existe la creo cuando se inicia
func CrearBaseSiNoExiste() {
	if (!db.CheckExistsBucket()) {
		log.Println("Regenerando....")
		RegenerarBase()
	}
}

//Genero la base datos con toda la info de todos los dias
func RegenerarBase() {
	//Abro la base
	dataBase := db.InitBolt()
	defer dataBase.Close()
	log.Println("Regenerando....")
	//Creo los planetas con sus datos
	inicializarPlanetas()

	//por cada dia llamo a la funcion para generar los datos, hago los calculos y grabo en la base
	for dia := 0; dia <= cantidadDiasModelo; dia++ {
		log.Println("Generando dia: ", dia)
		generarDia(dataBase, dia)
	}
}

//genero la informacion para un dia especifico
func generarDia (dataBase *bolt.DB ,dia int) {

	//calculo la posicion de cada planeta para cada dia
	Betasoide.CalcularPosicionXY(dia)
	Vulcano.CalcularPosicionXY(dia)
	Ferengi.CalcularPosicionXY(dia)

	//Todos los calculos matematicos
	if (mathCalcs.EstanAlineados(Betasoide, Vulcano, Ferengi)) {
		if (mathCalcs.ElSolTambienEstaAlineado(Betasoide, Vulcano, Ferengi, Sol)) {
			//Sequia
			grabarDatos(dataBase, dia, entity.Sequia, 0)
		} else {
			//CondicionesOptimas
			grabarDatos(dataBase, dia, entity.Optimo, 0)
		}
	} else { //hay triangulo
		//calculo el perimetro del triangulo para saber cuando es el mas grande
		perimetroTriangulo := mathCalcs.PerimetroTriangulo(Vulcano, Ferengi, Betasoide)
		if (mathCalcs.ElSolEstaEnMedioDelTriagulo(Betasoide, Vulcano, Ferengi, Sol)) {
			//Luvia
			grabarDatos(dataBase, dia, entity.Lluvia, perimetroTriangulo)
		} else {
			//NoLlueve
			grabarDatos(dataBase, dia, entity.NoLluvia, perimetroTriangulo)
		}
	}
}

//Grabo datos en la base
func grabarDatos (dataBase *bolt.DB, dia int, tipoClima string, perimetro float64) error {

	//Paso el dia a int para usarlo como key
	var s string = strconv.Itoa(dia)

	//genero la entidad para grabar el la base
	var values = entity.Clima{dia, tipoClima, perimetro}

	encoded, err := json.Marshal(values)
	if err != nil {
		return err
	}
	//guardo estructura
	err = db.Put(dataBase, []byte(s), encoded)
	return  err
}

//para consultar el clima para un dia especifico
//agregue una opcion para devolver un array con la info de todos los dias
func ClimaPorDia (dia string) [] byte {

	//Abro la base
	dataBase := db.InitBolt()
	defer dataBase.Close()
	//Para sacar el perimetro
	var rr entity.ClimaResult

	if (dia == "all") {
		array := []entity.ClimaResult{}
		for i := 0; i <= cantidadDiasModelo; i++ {
			var s string = strconv.Itoa(i)
			result2, _ := db.GetClimaPorDia(dataBase, s)
			json.Unmarshal(result2, &rr)
			array = append(array, rr)
		}
		p,_ := json.Marshal(array)
		return p
	} else {
		valor, _ := db.GetClimaPorDia(dataBase, dia)
		json.Unmarshal(valor, &rr)
		response, _ := json.Marshal(rr)
		return response
	}
}

//Para consultar por un tipo de clima especifico
func ClimaPorPeriodo (climaBuscado string) [] byte {
	//Abro la base
	dataBase := db.InitBolt()
	defer dataBase.Close()

	//Controlo el pedido de clima
	if (strings.ToUpper(climaBuscado) == strings.ToUpper(entity.Lluvia) ||
		strings.ToUpper(climaBuscado) == strings.ToUpper(entity.NoLluvia) ||
		strings.ToUpper(climaBuscado) == strings.ToUpper(entity.Optimo) ||
		strings.ToUpper(climaBuscado) == strings.ToUpper(entity.Sequia)) {

		cantidadPeriodos, diaMaximo  := db.GetCantidadPeriodos(dataBase, climaBuscado)
		//armo una nueva entidad
		resp := entity.Periodo{climaBuscado, cantidadPeriodos, diaMaximo}
		//paso a json
		response, _ := json.Marshal(resp)
		return response
	} else {

		type resultError struct {
			Status string `json:"status"`
			Valoresposibles string `json:"valoresPosibles"`
		}
		resp := resultError{"error", entity.Lluvia + "-" + entity.NoLluvia + "-" + entity.Optimo + "-" + entity.Sequia}
		response, _ := json.Marshal(resp)
		return response
	}

}