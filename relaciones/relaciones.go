package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

func main() {

	alex := User{
		nombre: "Alexys",
		email:  "djkdlk@gmail.com",
		activo: true,
	}

	reales := User{
		nombre: "reales",
		email:  "reales@gmail.com",
		activo: true,
	}

	realesStudent := Student{
		user:   reales,
		codigo: "0001",
	}

	fmt.Println(alex)
	fmt.Println(reales)
	fmt.Println(realesStudent.user.nombre)

}
