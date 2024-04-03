package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Введите значения параметров a и b:")
	fmt.Scan(&a, &b)

	// Площадь поверхности
	surfaceArea := 2*math.Pi*b*b + 2*math.Pi*a*b

	// Объем тела
	volume := (4.0 / 3.0) * math.Pi * a * b * b

	fmt.Printf("Площадь поверхности: %.2f\n", surfaceArea)
	fmt.Printf("Объем тела: %.2f\n", volume)
}
