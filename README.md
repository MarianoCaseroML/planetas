# Examen ML
Mariano Casero

## Información
- Para almacenar la información del clima use boltDB, una base key/value

https://github.com/boltdb/bolt

- Use como httpServer la implementación Gorilla

https://github.com/gorilla/http

- Para ajustar coordenadas tome presición de 1 decimal, redondeando con Ceil y Floor.

- Agregue algunas clases de testeo para los principales metodos

- Se incluye un Dockerfile con el que se creo el container.

- Use Beanstalk para subir el container.

## Fórmulas usadas:
- Posicion XY dentro del circulo:

http://www.universoformulas.com/fisica/cinematica/posicion-movimiento-circular/
- Distancia entre puntos:

https://es.wikihow.com/encontrar-la-distancia-entre-dos-puntos
- Perímetro del triángulo:

https://es.wikihow.com/encontrar-el-per%C3%ADmetro-de-un-tri%C3%A1ngulo
- Puntos alineados:

https://www.vitutor.com/geo/vec/a_9.html
- Orientación del triángulo:

http://www.dma.fi.upm.es/personal/mabellanas/tfcs/kirkpatrick/Aplicacion/algoritmos.htm#puntoInteriorAlgoritmo


## instalación
```
cd $GO_HOME/src
git clone https://github.com/MarianoCaseroML/planetas.git
```

## compilación (version Docker)
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application . 
```

## APIs

#### GET /health
Healthcheck
```json
[
    {
        "Status": "OK"
    }
]
```

#### GET /api/1/regenerarDB
Regenera la base de datos con información de los 10 años
```json
[
    {
        "Status": "OK"
    }
]
```

#### GET /api/1/clima?dia=all
Informacion del clima de todos los dias
```json
[
    {
        "dia":0,
        "clima":"Sequia"
    },
	,{
		"dia":1,
		"clima":"Nollueve"
	}
]
```

#### GET /api/1/clima?dia=33
Informacion del clima de un dia
```json
[
    {
        "dia":33,
        "clima":"Sequia"
    }
]
```

#### GET /api/1/periodo?clima=[sequia|Lluvia|CondicionesOptimas|Nollueve]
Informacion de un periodo de clima	
```json
[
    {
		"periodo":"lluvia",
		"cantidad":605,
		"PicoMaximo":71
	}
]
```
