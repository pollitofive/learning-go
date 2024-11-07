package maps

import "fmt"

func contar(frase string) map[rune]int {
	contador := make(map[rune]int)
	for _, char := range frase {
		contador[char]++
	}

	return contador

}

func Ejercicio5() {
	/*
		Contar Frecuencia de Caracteres
		Escribe una función que reciba una cadena de texto y devuelva un map en el que las claves sean los caracteres individuales de la cadena y
		los valores representen cuántas veces aparece cada carácter.
	*/

	frase := "Escribe una función que reciba una cadena de texto y devuelva un map en el que las claves sean los caracteres individuales"

	contador := contar(frase)
	for clave, valor := range contador {
		fmt.Println(string(clave), ":", valor)
	}

}
