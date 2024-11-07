package arrays_y_slices

import "fmt"

func invertir(slice []int) []int {

	var slice2 []int

	for i := len(slice) - 1; i >= 0; i-- {
		slice2 = append(slice2, slice[i])
	}

	return slice2
}

func Ejercicio2() {
	/*
		Ejercicio 2: Invertir un Slice
		Escribe una función que tome un slice de enteros y lo invierta, es decir, el último elemento se convierte en el primero, el penúltimo en el segundo, etc.
	*/
	slice := []int{10, 20, 30, 40, 50}
	result := invertir(slice)
	fmt.Println(result)
}
