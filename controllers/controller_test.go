package controllers

import (
	"testing"
	"os"
	"net/http"
	"net/http/httptest"
	"fmt"
	"strings"
)

func TestMain(m *testing.M) {
	i := m.Run()
	os.Exit(i)
}

func TestServer_Start(t *testing.T) {
	s := Service{}
	go s.Start()
}

func TestRegenerarDB(t *testing.T) {
	//request a regenerarDB
	req, err := http.NewRequest("GET", regenerarDB, nil)
	if err != nil {
		t.Fatal(err)
	}

	//me guardo el resultado
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RecreateDB)

	handler.ServeHTTP(rr, req)

	//chequeo  codigo
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Resultado incorrecto, se obtuvo: %v, se esperaba: %v.",
			status, http.StatusOK)
	}

	//chequeo body
	var resultado = `{ "status": "ok" }`
	fmt.Println(rr.Body.String())
	if rr.Body.String() != resultado {
		t.Errorf("Resultado incorrecto, se obtuvo: %v, se esperaba: %v.",
			rr.Body.String(), resultado)
	}
}

func TestConsultaClimaPorDia(t *testing.T) {
	//request a clima
	req, err := http.NewRequest("GET", clima + "?dia=33", nil)
	if err != nil {
		t.Fatal(err)
	}

	//me guardo el resultado
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ConsultaClimaPorDia)

	handler.ServeHTTP(rr, req)

	//chequeo  codigo
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Resultado incorrecto, se obtuvo: %v, se esperaba: %v.",
			status, http.StatusOK)
	}

	//chequeo body
	var resultado string = `{"dia":33,"clima":"Lluvia"}`
	if strings.TrimSuffix(rr.Body.String(), "\n") != resultado {
		t.Errorf("Resultado incorrecto, se obtuvo: %v, se esperaba: %v.",
			rr.Body.String(), resultado)
	}
}

func TestConsultaPorPeriodo (t *testing.T) {
	//request a clima
	req, err := http.NewRequest("GET", periodo + "?clima=Lluvia", nil)
	if err != nil {
		t.Fatal(err)
	}

	//me guardo el resultado
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(consultaPorPeriodo)

	handler.ServeHTTP(rr, req)

	//chequeo  codigo
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Resultado incorrecto, se obtuvo: %v, se esperaba: %v.",
			status, http.StatusOK)
	}

	//chequeo body
	var resultado string = `{"periodo":"Lluvia","cantidad":605,"picomaximo":71}`
	if strings.TrimSuffix(rr.Body.String(), "\n") != resultado {
		t.Errorf("Resultado incorrecto, se obtuvo: %v, se esperaba: %v.",
			rr.Body.String(), resultado)
	}
}