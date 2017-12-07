package entity

type Clima struct {
	Dia int `json:"dia"`
	TipoClima string `json:"clima"`
	PerimetroTriangulo float64 `json:"perimetrotriangulo,omitempty"`
}

//solo para devolver los datos sin perimetro
type ClimaResult struct {
	Dia int `json:"dia"`
	Clima string `json:"clima"`
}

const (
	Lluvia string = "Lluvia"
	Optimo string = "CondicionesOptimas"
	Sequia string = "Sequia"
	NoLluvia string = "Nollueve"
)

type Periodo struct {
	Periodo string `json:"periodo"`
	Cantidad int `json:"cantidad"`
	PicoMaximo int `json:"picomaximo,omitempty"`
}
