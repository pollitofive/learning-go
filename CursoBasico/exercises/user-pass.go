package exercises

import "fmt"

func validar(user, pass string) bool {
	return user == "damian" && pass == "1234"
}

func ExerciseValidatePassword() {
	user := "damian"
	pass := "1234"

	if validar(user, pass) {
		fmt.Println("El usuario y password es correcto")
	} else {
		fmt.Println("El usuario y password es incorrecto")
	}
}
