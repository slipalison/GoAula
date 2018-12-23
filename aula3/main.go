package main

import "fmt"

func main() {
	separador("- Tipos - OO")
	tipos()
	fmt.Println("Fim")
}

// criando uma lista que recebe qualquer coisa
type List []interface{}

func (l List) Show() {
	fmt.Println(l)
}

// inicializa de forma imutavel
func (l List) Init() List {
	l = []interface{}{
		"Alison",
		27,
		true,
		1.73,
	}
	return l
}

// inicializa como ponteiro
func (l *List) InitCursor() {
	*l = []interface{}{
		"Alison2",
		28,
		true,
		1.73,
	}
}

func tipos() {

	var l List
	l = l.Init()
	l.Show()
	l.InitCursor()
	l.Show()

}

func separador(valor string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(valor, "************************************************************************")
	fmt.Println("")
}
