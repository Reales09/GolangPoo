package main

import "fmt"

// Struc persona

type Persona struct {
	nombre string
	edad   int
}

//Metodos

func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s, Edad: %d ", p.nombre, p.edad)

}

func (p *Persona) editarNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

func main() {

	persona1 := Persona{"Reales", 20}
	// fmt.Println(persona1)

	persona1.imprimir()
	// fmt.Println(persona1)
	persona1.editarNombre("Pedro")

	persona1.imprimir()

	persona2 := Persona{nombre: "Adriana", edad: 20}

	// fmt.Println(persona2)

	persona2.imprimir()

	persona2.editarNombre("Maria")
	persona2.imprimir()
}
