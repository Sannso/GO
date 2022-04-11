package main

import "fmt"

type Animal interface {
	Mover()
	Comer(comida string) string
}

type Perro struct {
	nombre string
	patas  int
}

type Serpent struct {
	nombre  string
	especie string
}

func (obj Serpent) Mover() {
	fmt.Println("Reptando")
}

func (obj Perro) Mover() {
	fmt.Println("Corriendo")
}

func (obj Perro) Comer(comida string) string {
	fmt.Println("Comiendo", comida)
	return "Comío"
}

// Funcion clave
func funcionAnimal(animal Animal) {
	//animal.Comer("wsad")
	animal.Mover()
}

func main() {
	perro := Perro{"Lucas", 4}
	funcionAnimal(perro) // Si no esta el metodo implementado no compila
}
