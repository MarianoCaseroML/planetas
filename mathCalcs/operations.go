package mathCalcs

import (
	"planetas/entity"
	"math"
)

//Funcion para redondar valores a x decimales
//La necesite para poder calcular la alineacion de los planetas, sin redondear nunca me quedaban sobre la misma recta
func Round(val float64, roundOn float64, places int) float64 {

	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)

	var round float64
	if val > 0 {
		if div >= roundOn {
			round = math.Ceil(digit)
		} else {
			round = math.Floor(digit)
		}
	} else {
		if div >= roundOn {
			round = math.Floor(digit)
		} else {
			round = math.Ceil(digit)
		}
	}

	return round / pow
}

//Recibo los tres planetas y en base a las coordenadas xy voy calculando si estan en la misma recta.
//Use esta formula
//https://www.vitutor.com/geo/vec/a_9.html
func EstanAlineados (a entity.Planeta, b entity.Planeta, c entity.Planeta) bool {

	var ba = b.Coordenada.PosicionY - a.Coordenada.PosicionY
	var cb = c.Coordenada.PosicionY - b.Coordenada.PosicionY
	if(ba==0 || cb==0) {
		return (ba==0 && cb==0)  //dia 0 estan todos alineados
	}
	var p1 = (b.Coordenada.PosicionX - a.Coordenada.PosicionX) / ba
	var p2 = (c.Coordenada.PosicionX - b.Coordenada.PosicionX) / cb

	//redondeo a un decimal
	p1 = Round(p1, .5, 1)
	p2 = Round(p2, .5, 1)
	return p1 == p2
}

//uso la misma funcion de alineacion, sacando un planeta y poniendo al sol
func ElSolTambienEstaAlineado (a entity.Planeta, b entity.Planeta, c entity.Planeta, s entity.Planeta) bool {
	v1 := EstanAlineados(s, c, a)
	v2 := EstanAlineados(s, b, c)
	return v1 && v2
}

//Calculo si el Sol estan en medio del triangulo
//primero calculo la orientacion entre los distitos planetas y despues con el sol.
//Si la orientacion de todos es igual, es porque el Sol esta adentro.
func ElSolEstaEnMedioDelTriagulo (a entity.Planeta, b entity.Planeta, c entity.Planeta, p entity.Planeta) bool {
	abc := orientacionTriangulo(a, b, c);
	abp := orientacionTriangulo(a, b, p);
	bcp := orientacionTriangulo(b, c, p);
	cap := orientacionTriangulo(c, a, p);
	if(abc==1 && abp==1 && bcp==1 && cap==1) {
		return true;
	} else if (abc==-1 && abp==-1 && bcp==-1 && cap==-1) {
		return true;
	}

	return false;
}

//calculo la orientacion del triangulo
//Use esta formula
//http://www.dma.fi.upm.es/personal/mabellanas/tfcs/kirkpatrick/Aplicacion/algoritmos.htm#puntoInteriorAlgoritmo
//es una auxiliar para determinar si el Sol esta en medio del triangulo
func orientacionTriangulo (a entity.Planeta, b entity.Planeta, c entity.Planeta) int {
	x := (a.Coordenada.PosicionX - c.Coordenada.PosicionX)*( b.Coordenada.PosicionY- c.Coordenada.PosicionY) - ( a.Coordenada.PosicionY- c.Coordenada.PosicionY)*(b.Coordenada.PosicionX - c.Coordenada.PosicionX);
	if (x >= 0.0) {
		return  1
	} else {
		return  0
	}
}

//Calculo el perimetro total del triangulo
//Use esta formula
//https://es.wikihow.com/encontrar-el-per%C3%ADmetro-de-un-tri%C3%A1ngulo
func PerimetroTriangulo (a entity.Planeta, b entity.Planeta, c entity.Planeta) float64{
	// P = L1 + L2 + L3
	ab := distanciaEntrePuntos(a,b);
	bc := distanciaEntrePuntos(b,c);
	ca := distanciaEntrePuntos(c,a);
	return ab + bc + ca;
}

//funcion auxiliar para calcular el pemimetro
//Use esta formula
//https://es.wikihow.com/encontrar-la-distancia-entre-dos-puntos
func distanciaEntrePuntos(a entity.Planeta, b entity.Planeta) float64 {
	return math.Sqrt( math.Pow( (a.Coordenada.PosicionX - b.Coordenada.PosicionX), 2 ) + math.Pow( (a.Coordenada.PosicionY - b.Coordenada.PosicionY), 2 ) );
}

