package main

import "fmt"

//Lista de tareas
type ListaTareas struct {
	task []*Tarea
}

func (tl *ListaTareas) agregarTarea(t *Tarea) {
	tl.task = append(tl.task, t)
}

func (tl *ListaTareas) removeTarea(index int) {
	tl.task = append(tl.task[:index], tl.task[index+1:]...)
}

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

	lista := ListaTareas{}
	lista.agregarTarea(&t1)
	lista.agregarTarea(&t2)

	fmt.Println(lista)
	// t1.imprimir()
	// t2.imprimir()

	t3 := Tarea{
		nombre:      "Curso de CSS",
		descripcion: "Terminar el curso de CSS",
		completado:  true,
	}

	lista.agregarTarea(&t3)

	fmt.Println(lista)

	lista.removeTarea(1)

	for i, tarea := range lista.task {
		fmt.Println(i, tarea)
		tarea.imprimir()
	}
}
