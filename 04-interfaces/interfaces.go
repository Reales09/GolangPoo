package main

import "fmt"

type Area interface {
	area() float64
	perimetro() float64
}

type cuadrado struct {
	lado float64
}

type circulo struct {
	radio float64
}

type Animal interface {
	mover() string
}

func (a *cuadrado) area() float64 {
	return a.lado * a.lado
}

func (a *cuadrado) perimetro() float64 {
	return 2*a.lado + 2*a.lado
}

func (a *circulo) area() float64 {
	return a.radio * a.radio * 3.14
}

func (c *circulo) perimetro() float64 {
	return 2 * 3.14 * c.radio
}

type Perro struct{}
type Pez struct{}
type Pajaro struct{}

func (*Perro) mover() string {
	return "Soy un perro y camino"
}

func (*Pez) mover() string {
	return "Soy un pez y nado"
}

func (*Pajaro) mover() string {
	return "Soy un pajaro y vuelo"
}

func moverAnimal(animal Animal) {

	fmt.Println(animal.mover())
}

func printArea(area Area) {
	fmt.Printf("El area del cuadrado es: %v \n", area.area())
	fmt.Printf("El perimetro de un cuadrado es: %v \n", area.perimetro())

}

func printAreaCirculo(area Area) {
	fmt.Printf("El perimetro del cirsulo es: %v \n", area.area())
	fmt.Printf("El perimetro del circulo es: %v \n", area.perimetro())

}

func main() {
	perro := Perro{}
	moverAnimal(&perro)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)

	areaCuadradro := cuadrado{lado: 10}
	printArea(&areaCuadradro)

	areaCirculo := circulo{radio: 10}
	printAreaCirculo(&areaCirculo)

}
