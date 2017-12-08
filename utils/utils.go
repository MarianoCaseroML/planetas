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

func Fileexists(fileName string) (bool) {
	_, err := os.Stat(fileName)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return true
}