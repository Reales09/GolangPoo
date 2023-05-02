package main

import "fmt"

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *Tarea) imprimir() {
	fmt.Printf("Nombre: %s, Descripcion: %s, Completado: %t", t.nombre, t.descripcion, t.completado)
}

func (t *Tarea) marcarCompleta() {
	t.completado = true
}

func main() {

	t1 := Tarea{
		nombre:      "Curso de Go",
		descripcion: "Terminar el curso de Go",
		completado:  false,
	}

	t2 := Tarea{
		nombre:      "Curso de HTML",
		descripcion: "Terminar el curso de HTML",
		completado:  true,
	}

	t1.imprimir()
	t2.imprimir()

}
