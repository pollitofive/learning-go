package exercises

import (
	"fmt"
	"strconv"
)

func decirSiEsPar(value int) bool {
	return (value % 2) == 0
}

func ExerciseParImpar() {
	number := 12

	if decirSiEsPar(number) {
		fmt.Println(strconv.Itoa(number) + " Es par")
	} else {
		fmt.Println(strconv.Itoa(number) + " Es impar")
	}

}
