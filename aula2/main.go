package main

import "fmt"

func main() {
	separador("- Funções variaticas")
	fmt.Println(soma(1, 2, 3, 4, 5, 6))

	separador("- Funções recurciva")
	recurcivo(10)

	fmt.Println("Fim")
}

// função variaticas
func soma(numbers ...int) int {
	theSum := 0
	for _, value := range numbers {
		theSum += value
	}
	return theSum
}

func recurcivo(i int) {
	fmt.Println(i)
	i--
	if i > 0 {
		recurcivo(i)
	}
}

func separador(valor string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(valor, "************************************************************************")
	fmt.Println("")
}
