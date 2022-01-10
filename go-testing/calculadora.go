package calculadora

func Sumar(a int, b int) int {
	return a + b
}

func Ordenar(s *[]int) []int {

	for i := 0; i < len(*s); i++ {
		for j := 0; j < len(*s)-1; j++ {
			if (*s)[j] > (*s)[j+1] {
				(*s)[j], (*s)[j+1] = (*s)[j+1], (*s)[j]
			}
		}
	}

	return *s
}
