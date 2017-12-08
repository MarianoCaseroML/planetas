package utils

import (
	"os"
)

//Para obtener el directorio actual donde se esta ejecutando el binario
//para no andar cambiando entre modo debug y el container
func CurrentWF() string {
	wf, _ := os.Getwd()
	return wf
}
