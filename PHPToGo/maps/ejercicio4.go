package maps

import "fmt"

func unir(lista1 map[string]int, lista2 map[string]int) map[string]int {

	nuevo_map := make(map[string]int)

	for key, value := range lista1 {
		if nuevo_map[key] != 0 {
			nuevo_map[key] += value
		} else {
			nuevo_map[key] = value
		}
	}

	for key, value := range lista2 {
		if nuevo_map[key] != 0 {
			nuevo_map[key] += value
		} else {
			nuevo_map[key] = value
		}
	}

	return nuevo_map
}

func Ejercicio4() {
	/*
		Unir dos Maps
			Escribe una función que reciba dos maps de tipo map[string]int y los una en uno solo. Si una clave está en ambos maps, suma sus valores.
			La función debe devolver el map resultante.
	*/

	lista1 := map[string]int{
		"Carlos": 8,
		"Pedro":  7,
		"Emma":   6,
		"Diego":  4,
		"Lucas":  10,
		"Damian": 10,
	}

	lista2 := map[string]int{
		"Damian": 10,
		"Jose":   5,
		"Alexis": 8,
		"Nahuel": 3,
		"Emma":   3,
	}

	unidos := unir(lista1, lista2)
	for clave, valor := range unidos {
		fmt.Printf("%s: %d\n", clave, valor)
	}

}
