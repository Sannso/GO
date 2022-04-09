package main

import (
	"fmt"
)

type Estudiante struct {
	nombre     string
	nota_mate  float32
	nota_caste float32
	nota_bio   float32
}

func main() {

	estudiantes := make([]Estudiante, 0, 10)
	estudiantes = append(estudiantes,
		Estudiante{nombre: "Juan Perez", nota_mate: 4, nota_caste: 4.5, nota_bio: 3},
		Estudiante{nombre: "Sebastian Velez", nota_mate: 2, nota_caste: 2.5, nota_bio: 5})

	end := false

	for {
		fmt.Println("\n\n--------- Estudiantes ------------")
		fmt.Println(" 1. Registrar estudiante\n",
			"2. Ver estudiantes\n",
			"3. Mostrar notas finales de cada estudiante\n",
			"4. Mostrar promedio grupal\n",
			"5. Mostrar estudiante con nota más alta\n",
			"6. Salir\n",
			"Ingrese su opcion: ")
		var opcion int
		fmt.Scanf("%d\n", &opcion)

		switch opcion {
		case 1:
			estudiantes = registrar(estudiantes)
		case 2:
			verEstudiantes(estudiantes)
		case 3:
			notasFinales(estudiantes)
		case 4:
			promedioGrupal(estudiantes)
		case 5:
			notaMasAlta(estudiantes)
		case 6:
			end = true
		}

		if end {
			break
		}
	}
}

func registrar(ar []Estudiante) (result []Estudiante) {
	var nombre string
	var mate, caste, bio float32
	fmt.Println("Ingrese el nombre:")
	fmt.Scanf("%s\n", &nombre)
	fmt.Println("Ingrese nota de matematicas:")
	fmt.Scanf("%f\n", &mate)
	fmt.Println("Ingrese nota de castellano:")
	fmt.Scanf("%f\n", &caste)
	fmt.Println("Ingrese nota de biologia:")
	fmt.Scanf("%f\n", &bio)

	result = append(ar, Estudiante{nombre: nombre,
		nota_mate: mate, nota_caste: caste, nota_bio: bio})

	return
}

func verEstudiantes(ar []Estudiante) {
	fmt.Println("\n---- Estudiantes ingresados:")
	for i := range ar {
		fmt.Println("\nNombre:", ar[i].nombre)
	}
}

func notasFinales(ar []Estudiante) {
	fmt.Println("\n---- Notas finales:")
	for i := range ar {
		fmt.Println("\nNombre:", ar[i].nombre, "Notas:\n", "Matematicas:", ar[i].nota_mate,
			"\nCastellano:", ar[i].nota_caste, "\nBiologia:", ar[i].nota_bio)
	}
}

func promedioGrupal(ar []Estudiante) {
	fmt.Println("\n---- Promedio grupal:")
	var pro_mate, pro_cas, pro_bio float32
	for i := range ar {
		pro_mate = pro_mate + ar[i].nota_mate
		pro_cas = pro_cas + ar[i].nota_caste
		pro_bio = pro_bio + ar[i].nota_bio
	}

	fmt.Println("\nPromedio Matematicas:", (pro_mate / float32(len(ar))))
	fmt.Println("Promedio Castellano:", (pro_cas / float32(len(ar))))
	fmt.Println("Promedio Biologia:", (pro_bio / float32(len(ar))))
}

func notaMasAlta(ar []Estudiante) {
	fmt.Println("\n---- Promedio grupal:")
	var n_mate, n_cas, n_bio float32
	var na_mate, na_cas, na_bio string
	for i := range ar {
		if i == 0 {
			n_mate = ar[i].nota_mate
			n_cas = ar[i].nota_caste
			n_bio = ar[i].nota_bio

			na_mate = ar[i].nombre
			na_cas = ar[i].nombre
			na_bio = ar[i].nombre
		} else {
			if ar[i].nota_mate > n_mate {
				n_mate = ar[i].nota_mate
				na_mate = ar[i].nombre
			}
			if ar[i].nota_caste > n_cas {
				n_cas = ar[i].nota_caste
				na_cas = ar[i].nombre
			}
			if ar[i].nota_bio > n_bio {
				n_bio = ar[i].nota_bio
				na_bio = ar[i].nombre
			}
		}

	}

	fmt.Println("\nNota mas alta en Matematicas:", na_mate, "Con:", n_mate)
	fmt.Println("Nota mas alta en Castellano:", na_cas, "Con:", n_cas)
	fmt.Println("Nota mas alta en Biologia:", na_bio, "Con:", n_bio)
}
