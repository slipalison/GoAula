/*
	go run main.go
	go build main.go
*/

package main

import "fmt"

func main() {
	separador("- Variables")
	Variable()

	separador("- Conditions")
	conditions()

	separador("- Functions and returns")
	fmt.Println("resultado da soma é:", soma(2, 3))
	num, str := somaConStr(2, 3)
	fmt.Println(num)
	fmt.Println(str)

	separador("- Collections")
	Collections()

	separador("- Loops")
	loop()

	fmt.Println("Fim")
}

func separador(valor string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(valor, "************************************************************************")
	fmt.Println("")
}
func Variable() {
	//formas de declarar variaveis
	var name string = "Alison"
	var str2 = "outra string"
	str3 := "mais uma string"
	const strConstante = "essa variavel é constante"
	// em bloco
	num1, num2 := 2, 4
	const (
		age  = 27
		yaer = 2018
	)

	fmt.Println(name)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(strConstante)
	fmt.Println(age)
	fmt.Println(yaer)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num2 * num1)
}

func conditions() {
	b := 3
	if b > 1 {
		fmt.Printf("%v é o valor da variavel\n", b)
	} else {
		fmt.Println("não esta na condição")
	}

	//variavel de escopo
	if n := 30; n > 9 {
		fmt.Printf("%v é maior que 9\n e não pode ser acessada de fora do escopo\n", n)
	}

	// default é obrigatorio
	switch b {
	case 2:
		fmt.Printf("o valor da variavel %v", b)
	default:
		fmt.Println("valor não encontrado")
	}
}

//Collections
func Collections() {
	// Array de 10 posições
	// inicializa com 0
	t := [10]int{0, 5, 10}
	t[3] = 15

	fmt.Println(t)
	fmt.Println("o tamanho do array é: ", len(t))

	// MAPS / dictionary
	user := map[string]string{
		"nome": "alison",
		"nick": "slip",
	}
	fmt.Println(user)

	//slice ou list
	slice := []int{0, 5, 10}
	slice = append(slice, 3)

	printSlice("slice", slice)

	//slice travado em 5 posições e com 80 cabacidade
	sl := make([]int, 5, 80)
	sl[0] = 0
	sl[1] = 1
	sl[2] = 2
	sl[3] = 3
	sl[4] = 4
	printSlice("sl", sl)

	//aponte tudo de 1 slice para outro
	sl2 := sl[:]
	fmt.Println(sl2)

	// aponte da posição 2 em diante
	sl2 = sl[2:]
	fmt.Println(sl2)

	// aponte até a posição 2
	sl2 = sl[:2]
	fmt.Println(sl2)

	// aponte da posição 1 até a posição 3
	sl2 = sl[1:3]
	fmt.Println(sl2)

	//Como isso se trata de um apontamento, ao alterar o valor de 1 slice, acaba impactando o outro
	sl2[1] = 200
	fmt.Println(sl)
	fmt.Println(sl2)

	//como o slice foi criado com 5 posiçoes, ao adicionar mais 1 ele irar gerar panic/erro de execução
	// sl[5] = 5
	// fmt.Println(sl)
	// Porem, se usar o append, ele irar criar um outro slice com mais posição, e irá parar de apontar para o slice principal
	sl2 = append(sl2, 5, 6, 7, 8, 9)
	sl2[1] = 2
	fmt.Println(sl2)
	fmt.Println(sl)
	printSlice("sl2", sl2)

}

func loop() {
	num := 3
	// tudo é for!
	// while
	for num > 0 {
		fmt.Printf("While %v\n", num)
		num--
	}

	//for
	for n := 0; n < 3; n++ {
		fmt.Println("for...")
	}

	// Do While
	n := 1
	for ok := true; ok; ok = n > 0 {
		fmt.Println("Do While")
		n--
	}

	arr := [...]int{1, 2, 3}
	//foreach index, value, colocase _ para ignorar o valor
	for index, value := range arr {
		fmt.Println(index, value)
	}

	user := map[string]string{
		"name": "Alison",
		"nick": "Slipalison",
	}
	languages := []string{
		"Go",
		"JS",
		"C#",
		"Python",
		"Elixir",
	}

	for key, value := range user {
		fmt.Printf("O campo \"%s\" tem o valor igual a \"%s\"\n", key, value)
	}

	for _, value := range languages {
		fmt.Println(value)
	}
}

//funçao com retorno
func soma(num1, num2 int) int {
	return num1 + num2
}

//funçao com mais de 1 retorno
func somaConStr(num1, num2 int) (int, string) {
	return num1 + num2, "fez a soma e retornou uma string"
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
