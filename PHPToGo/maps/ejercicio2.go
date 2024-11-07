package maps

import "fmt"

func filtrar(lista map[string]int, minimo int) map[string]int {
	nuevo_map := make(map[string]int)

	for key, value := range lista {
		if value >= minimo {
			nuevo_map[key] = value
		}
	}

	return nuevo_map

}

func Ejercicio2() {
	/* Filtrar Estudiantes por Notas
	Imagina que tienes un map de estudiantes donde la clave es el nombre del estudiante (string) y el valor es su calificación (int).
	Escribe una función que reciba este map y un valor de calificación mínimo, y devuelva un nuevo map con solo los estudiantes que tienen
	calificaciones mayores o iguales al valor mínimo.*/

	lista := map[string]int{
		"Carlos": 8,
		"Pedro":  7,
		"Damian": 10,
		"Emma":   6,
		"Diego":  4,
		"Lucas":  10,
	}

	filtrados := filtrar(lista, 8)
	for clave, valor := range filtrados {
		fmt.Printf("%s: %d\n", clave, valor)
	}

}
