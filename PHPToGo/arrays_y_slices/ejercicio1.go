package arrays_y_slices

import "fmt"

func sumarArray(lista [10]int) int {

	var sum int = 0
	for i := 0; i < len(lista); i++ {
		sum += lista[i]
	}

	return sum
}

func Ejercicio1() {
	/* Ejercicio 1: Sumar Elementos de un Array
	Crea una función en Go que reciba un array de enteros y devuelva la suma de todos sus elementos.
	*/
	lista := [10]int{5, 6, 7, 1, 5, 9, 3, 6, 5, 4}
	sum := sumarArray(lista)
	fmt.Println("El resultado es ", sum)
}
