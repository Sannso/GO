package main

import (
	a "fmt"
)

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

	//---------- Funciones ---------------------

	//------------- Ciclos ---------------------
	// for j := 0; j < 10; j++ {}
	// for ; j < 10 ; {} -> While
	// for ; ; {} -> while (true)

}
