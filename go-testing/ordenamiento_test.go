package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	testSlice := []int{5, 3, 2, 4, 1}
	//fmt.Println("Sin ordenar", testSlice)

	Ordenar(&testSlice)
	//fmt.Println("Ordenado", testSlice)

	for i := 0; i < len(testSlice)-1; i++ {
		if testSlice[i] > testSlice[i+1] {
			t.Errorf("Los valores siguen desordenados")
		}
	}
}

//==============================================================//
func TestOrdenarConTestify(t *testing.T) {
	testSlice := []int{5, 3, 2, 4, 1}
	Ordenar(&testSlice)
	resultadoEsperado := []int{1, 2, 3, 4, 5}

	assert.Equal(t, testSlice, resultadoEsperado, "El slice debe estar ordenado")

}

//==============================================================//
