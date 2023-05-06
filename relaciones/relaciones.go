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

// Relacion uno a muchos

type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	// Relacion uno a uno

	/*
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
	*/

	// Relacion uno a muchos
	video1 := Video{titulo: "01-Introduccion"}
	video2 := Video{titulo: "02-Instalacion"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}

}
