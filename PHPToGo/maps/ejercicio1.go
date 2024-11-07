package maps

import (
	"fmt"
	"strings"
)

func contarPalabras(palabras string) map[string]int {
	array := strings.Split(palabras, " ")
	contador := make(map[string]int)
	for _, palabra := range array {
		contador[palabra]++
	}

	return contador

}

func Ejercicio1() {
	/* Contar Ocurrencias de Palabras
	Escribe una función que reciba una cadena de texto y devuelva un map donde las claves sean las palabras y
	los valores el número de veces que cada palabra aparece en el texto.*/

	palabras := "Escribe una función que reciba una cadena de texto y devuelva un map donde las claves sean las palabras y los valores el número de veces que cada palabra aparece en el texto"

	totales := contarPalabras(palabras)
	for clave, valor := range totales {
		fmt.Printf("%s: %d\n", clave, valor)
	}
}
