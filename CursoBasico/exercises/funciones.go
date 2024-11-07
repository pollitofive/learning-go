package exercises

import (
	"fmt"
	"math"
)

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func areaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoide(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}

func ExerciseFunctions() {
	fmt.Printf("Circulo %.2f \n", areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", areaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoide(10, 5, 3))
}

/*
func main() {
}
*/
// Se pueden devolver todos los valores que se quiera de una funcion. Si no se quiere recibir alguno de la respuesta, hay que poner _ en el lugar
