package main

import "fmt"

func main() {
	separador("- Estruturas - OO classes")
	Structs()
	fmt.Println("Fim")
}

func Structs() {
	var user User
	user.Init()

	fmt.Println(user)
	for user.Age < 80 {
		user.Birthday()
		fmt.Println("Fez ", user.Age, " anos")
	}

	user.Die()
	fmt.Println(user)

	su := SuperUser{
		User{
			"Espermatoboys",
			32,
			true,
		},
		"Censurado",
	}

	fmt.Println(su)

	Show(user)
	Show(su)
}

func Show(u IUser) {
	fmt.Println(u.Show())
}

/** "CLASS" *********************/
type User struct {
	Name Name
	Age  Age
	Live Live
}
type Name string
type Age int
type Live bool

/*Herança*/
type SuperUser struct {
	User       // herança
	SuperPoder string
}

/*Interfaces*/
type IUser interface {
	Show() string
}

func (u User) Show() string {
	return fmt.Sprintf("Ola, meu nome é %s, e eu tenho %v anos", u.Name, u.Age)
}
func (u SuperUser) Show() string {
	return fmt.Sprintf("Ola, meu nome é %s, e eu tenho %v anos. Meu super poder é %s.", u.Name, u.Age, u.SuperPoder)
}

func (u *User) Init() {
	u.Age = 27
	u.Name = "Alison"
	u.Live = true
}

func (u *User) Die() {
	u.Live.Morrer()
}

func (u *User) Birthday() {
	u.Age.Envelhecer()
}

func (a *Age) Envelhecer() {
	*a++
}

func (l *Live) Morrer() {
	fmt.Println("Esse ai foi se encontrar com Deus pessoalmente")
	*l = false
}

func separador(valor string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(valor, "************************************************************************")
	fmt.Println("")
}
