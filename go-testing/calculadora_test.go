package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//==============================================================//
func TestSumar(t *testing.T) {
	num1 := 2
	num2 := 3

	resultadoEsperado := 5
	resultado := Sumar(num1, num2)
	if resultado != resultadoEsperado {
		t.Errorf("Sumar(2,3) debe dar 5")
	}

}

//==============================================================//
func TestSumarConTestify(t *testing.T) {
	num1 := 2
	num2 := 3

	resultadoEsperado := 5
	resultado := Sumar(num1, num2)

	assert.Equal(t, resultado, resultadoEsperado, "Sumar(2,3) debe dar 5")
}

//==============================================================//
