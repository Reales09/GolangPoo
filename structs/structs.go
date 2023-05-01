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

type Empleado struct {
	Persona
	sueldo float64
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

	empleado := Empleado{
		sueldo: 500,
	}

	empleado.nombre = "Reales"
	empleado.edad = 32
	empleado.imprimir()
	fmt.Println(empleado)

}
