package main

import "fmt"

// Struc persona

type Persona struct {
	nombre string
	edad   int
}

func main() {

	persona1 := Persona{"Juan", 20}
	fmt.Println(persona1)
	persona1.nombre = "Reales"

	fmt.Println(persona1)

	persona2 := Persona{nombre: "Adriana", edad: 20}

	fmt.Println(persona2)
}
