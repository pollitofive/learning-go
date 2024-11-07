package maps

import "fmt"

func invertir(lista map[string]int) map[int]string {
	invertido := make(map[int]string)

	for key, value := range lista {
		invertido[value] = key
	}

	return invertido

}

func Ejercicio3() {
	/* Invertir un Map
	Escribe una función que reciba un map donde las claves sean de tipo string y los valores de tipo int, e invierta sus claves y valores.
	Si dos claves tienen el mismo valor, solo debes conservar una de ellas en el map invertido*/

	lista := map[string]int{
		"Carlos": 8,
		"Pedro":  7,
		"Emma":   6,
		"Diego":  4,
		"Lucas":  10,
		"Damian": 10,
	}

	invertidos := invertir(lista)
	for clave, valor := range invertidos {
		fmt.Printf("%d: %s\n", clave, valor)
	}

}
