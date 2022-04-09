package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

// El uso de los punteros se ve reflejado aqui en la funcion
// donde se pasa la variable de persona como referencia y no como valor
// ya que si se pasa como valor solo se hace una copia de este y no lo modifica
// realmente en la variable de la funcion main.
func (obj *Persona) incrementarEdad(num int) {
	obj.edad += num
}

func main() {
	persona := Persona{nombre: "Santiago", edad: 22}
	fmt.Println(persona)
	persona.incrementarEdad(10)
	fmt.Println("Edad incrementada", persona)

	// Otro ejemplo
	a := 10
	fmt.Println("valor:", a)
	duplicar(&a)
	fmt.Println("Duplicado:", a)

}

func duplicar(a *int) {
	*a *= 2
}

// min 22:36 https://drive.google.com/file/d/1NEdqSJSzFlbrT1qzOaoMv74gFy2-NGq3/view
