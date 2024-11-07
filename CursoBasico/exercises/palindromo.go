package exercises

import (
	"fmt"
	"strings"
)

func isPalindromo(word string) {
	var wordReverse string
	word = strings.ToUpper(word)

	for i := len(word) - 1; i >= 0; i-- {
		wordReverse += string(word[i])
	}

	if word == wordReverse {
		fmt.Println(word + " Es Palindromo")
	} else {
		fmt.Println(word + " No es Palindromo")
	}
}

func ExerciseIsPalindromo() {
	isPalindromo("Amas")
	isPalindromo("Ama")
}
