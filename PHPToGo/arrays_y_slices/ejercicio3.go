package arrays_y_slices

import "fmt"

func filtrarPares(slice []int) []int {
	var slice2 []int

	for i := 0; i < len(slice); i++ {
		if (slice[i] % 2) == 0 {
			slice2 = append(slice2, slice[i])
		}
	}

	return slice2
}

func Ejercicio3() {
	/*
		Ejercicio 3: Filtrar Números Pares en un Slice
		Implementa una función que reciba un slice de enteros y devuelva otro slice que contenga únicamente los números pares.
	*/

	slice := []int{10, 11, 12, 13, 14}
	result := filtrarPares(slice)
	fmt.Println(result)
}
