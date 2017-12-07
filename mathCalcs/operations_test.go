package mathCalcs

import (
	"testing"
	"planetas/entity"
)


var Ferengi entity.Planeta
var Betasoide entity.Planeta
var Vulcano entity.Planeta
var Sol entity.Planeta

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

func TestRound(t *testing.T) {

	var resultadoEsperado float64 = 12.1

	resultado := Round(12.14640546, 0.5, 1)

	if (resultado!=resultadoEsperado) {
		t.Errorf("Resultado incorrecto, se obtuvo: %d, se esperaba: %d.",
			resultado, resultadoEsperado)
	}
}

func TestEstanAlineados(t *testing.T) {
	inicializarPlanetas()

	//Sequia
	Betasoide.CalcularPosicionXY(11)
	Vulcano.CalcularPosicionXY(11)
	Ferengi.CalcularPosicionXY(11)

	if !EstanAlineados(Betasoide, Vulcano, Ferengi) {
		t.Errorf("Resultado incorrecto, para el dia 11 deben estar alineados")
	}

	//CondicionesOptimas
	Betasoide.CalcularPosicionXY(82)
	Vulcano.CalcularPosicionXY(82)
	Ferengi.CalcularPosicionXY(82)

	if !EstanAlineados(Betasoide, Vulcano, Ferengi) {
		t.Errorf("Resultado incorrecto, para el dia 82 deben estar alineados")
	}

	//Lluvia
	Betasoide.CalcularPosicionXY(75)
	Vulcano.CalcularPosicionXY(75)
	Ferengi.CalcularPosicionXY(75)

	if EstanAlineados(Betasoide, Vulcano, Ferengi) {
		t.Errorf("Resultado incorrecto, para el dia 75 no deben estar alineados")
	}

	//NoLlueve
	Betasoide.CalcularPosicionXY(76)
	Vulcano.CalcularPosicionXY(76)
	Ferengi.CalcularPosicionXY(76)

	if EstanAlineados(Betasoide, Vulcano, Ferengi) {
		t.Errorf("Resultado incorrecto, para el dia 76 no deben estar alineados")
	}
}

func TestElSolTambienEstaAlineado(t *testing.T) {
	inicializarPlanetas()

	//Sequia
	Betasoide.CalcularPosicionXY(344)
	Vulcano.CalcularPosicionXY(344)
	Ferengi.CalcularPosicionXY(344)

	if !ElSolTambienEstaAlineado(Betasoide, Vulcano, Ferengi, Sol) {
		t.Errorf("Resultado incorrecto, para el dia 344 deben estar alineados con el sol")
	}

	//NoLlueve
	Betasoide.CalcularPosicionXY(350)
	Vulcano.CalcularPosicionXY(350)
	Ferengi.CalcularPosicionXY(350)

	if ElSolTambienEstaAlineado(Betasoide, Vulcano, Ferengi, Sol) {
		t.Errorf("Resultado incorrecto, para el dia 350 no deben estar alineados con el sol")
	}
}

func TestElSolEstaEnMedioDelTriagulo(t *testing.T) {
	inicializarPlanetas()

	//Lluvia
	Betasoide.CalcularPosicionXY(137)
	Vulcano.CalcularPosicionXY(137)
	Ferengi.CalcularPosicionXY(137)

	if !ElSolEstaEnMedioDelTriagulo(Betasoide, Vulcano, Ferengi, Sol) {
		t.Errorf("Resultado incorrecto, para el dia 137 el sol esta en medio del triangulo")
	}

	//Lluvia
	Betasoide.CalcularPosicionXY(200)
	Vulcano.CalcularPosicionXY(200)
	Ferengi.CalcularPosicionXY(200)

	if ElSolEstaEnMedioDelTriagulo(Betasoide, Vulcano, Ferengi, Sol) {
		t.Errorf("Resultado incorrecto, para el dia 200 el sol no esta en medio del triangulo")
	}

}

func TestPerimetroTriangulo(t *testing.T) {
	inicializarPlanetas()

	//Lluvia
	Betasoide.CalcularPosicionXY(30)
	Vulcano.CalcularPosicionXY(30)
	Ferengi.CalcularPosicionXY(30)

	var resultadoEsperado float64 =(5763.198885224481)
	resultado := PerimetroTriangulo(Betasoide, Vulcano, Ferengi)
	if (resultado != resultadoEsperado) {
		t.Errorf("Resultado incorrecto, para el dia 30 se esperaba un perimetro de %f y se obtuvo %f", resultadoEsperado, resultado)
	}
}
