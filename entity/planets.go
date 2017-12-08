package entity

import "math"

//Estructura de planeta
type Planeta struct {
	Velocidad int
	Distancia int
	Sentido   int
	Coordenada struct {
		Dia int
		PosicionX float64
		PosicionY float64
	}
}

//funcion para setear la posicion xy por dia
func (p Planeta) SetPosicion(dia int)  {
	p.Coordenada.Dia = dia
	p.Coordenada.PosicionX = 0
	p.Coordenada.PosicionY = 0
}

//Calculo la posicion xy para cada planeta segun el dia
//Use esta formula
//http://www.universoformulas.com/fisica/cinematica/posicion-movimiento-circular/
func (p *Planeta) CalcularPosicionXY (dia int){
	//primero calculo el angulo
	var angulo  float64 = float64(p.Velocidad) * float64(dia) * float64(p.Sentido)

	posX := float64(p.Distancia) * math.Cos(angulo)
	posY := float64(p.Distancia) * math.Sin(angulo)

	p.Coordenada.Dia = dia
	p.Coordenada.PosicionX = posX
	p.Coordenada.PosicionY = posY

}