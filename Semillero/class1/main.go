package main

import (
	a "fmt"
)

type Persona struct {
	nombre string
	edad   string
}

//Heredar
type Empleado struct {
	Persona
	cargo string
}

func main() {
	a.Println("Hello world!")

	//---------- Variables ------------------------
	// var num int
	// num = 0
	// // Otra forma
	// num2 := 3

	// a.Println("Numeros: ", num, num2)

	//----------- Arreglo ---------------------
	// var arreglo [5]int
	// a.Println(arreglo) // Imprime arreglo de tamaño 5 con Ceros
	// arreglo[0] = 10
	// a.Println(arreglo)
	// len(arreglo) 			//Longitud arreglos

	//----------- Slice ------------------- Arreglos que actualiza tamaño
	// var slice []int
	// slice = append(slice, 15, 12, 13, 70)
	// slice = append(slice, 80)
	// a.Println(slice, " Tamanio: ", len(slice), " Capacidad: ", cap(slice))
	// var slice = make(int[], num elementos, cap)
	// Copiar slice
	// slice2 := int[]{1,2,3}
	// copy(slice2, slice) // Compara las posiciones y las va sobre escribiendo en s2

	//---------- Maps -------------------
	// mapa := make(map[string]int)
	// mapa["uno"] = 1

	//------------- Ciclos ---------------------
	// for j := 0; j < 10; j++ {}
	// for ; j < 10 ; {} -> While
	// for ; ; {} -> while (true)

	// Ejemplo
	// mapa["dos"] = 2
	// for k, v := range mapa{
	// 	a.Printf("%s -> %d \n", k, v)
	// }

	//---------- Funciones ---------------------
	//saludar("Santiago")
	//saludar2("Santiago", "Sosa", 20)

	// sum := suma(2, 3)
	// a.Println("La suma da:", sum)

	//----------- Estructuras -------------
	// todos los datos se inicializan
	persona := Persona{nombre: "Juan", edad: "23"}
	a.Println(persona)

	// HEREDAR
	// empleado := Empleado{}
	// empleado.cargo = "Programador"
	// empleado.nombre = "Juan"
	// a.Println(empleado)

	// --------- IO -----------
	a.Println("Ingrese su nombre: ")
	var name string
	a.Scanf("%s", &name) // %d numeros
	a.Println("Hola", name)

}

// func saludar(name string) {
// 	a.Println("Holaa", name)
// }

// func saludar2(name, last string, age int) {
// 	a.Println("Hola", name, last, "tienes", age, "anios?")
// }

// La forma de retornar es igual que siempre
// pero aqui puede cambiar
// Como en python puede retornar varios valores pero
// hay que asignarlo en la variable que los recibe
// func suma(x, y int) (suma int) {
// 	suma = x + y
// 	return
// }

//  CREACION DE "METODO" PARA UNA ESTRUCTURA
// func (obj Persona) saludar(){
// 	fmt.Println("Hola")
// }
