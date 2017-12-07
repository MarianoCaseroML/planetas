package mathCalcs

import (
	"planetas/entity"
	"math"
)

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

//uso la misma funcion, sacando un planeta y poniendo al sol
func ElSolTambienEstaAlineado (a entity.Planeta, b entity.Planeta, c entity.Planeta, s entity.Planeta) bool {
	v1 := EstanAlineados(s, c, a)
	v2 := EstanAlineados(s, b, c)
	return v1 && v2
}
func orientacionTriangulo (a entity.Planeta, b entity.Planeta, c entity.Planeta) int {
	x := (a.Coordenada.PosicionX - c.Coordenada.PosicionX)*( b.Coordenada.PosicionY- c.Coordenada.PosicionY) - ( a.Coordenada.PosicionY- c.Coordenada.PosicionY)*(b.Coordenada.PosicionX - c.Coordenada.PosicionX);
	if (x >= 0.0) {
		return  1
	} else {
		return  0
	}
}

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

func PerimetroTriangulo (a entity.Planeta, b entity.Planeta, c entity.Planeta) float64{
	// P = L1 + L2 + L3
	ab := distanciaEntrePuntos(a,b);
	bc := distanciaEntrePuntos(b,c);
	ca := distanciaEntrePuntos(c,a);
	return ab + bc + ca;
}

func distanciaEntrePuntos(a entity.Planeta, b entity.Planeta) float64 {
	return math.Sqrt( math.Pow( (a.Coordenada.PosicionX - b.Coordenada.PosicionX), 2 ) + math.Pow( (a.Coordenada.PosicionY - b.Coordenada.PosicionY), 2 ) );
}