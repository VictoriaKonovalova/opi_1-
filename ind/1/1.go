/* Объем и площадь поверхности куба: Задайте переменную для длины
ребра куба. Рассчитайте и выведите объем и площадь его поверхности.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var sideLength float64
	fmt.Println("Введите длину ребра куба:")
	fmt.Scan(&sideLength)

	// Вычисляем объем куба
	volume := math.Pow(sideLength, 3)

	// Вычисляем площадь его поверхности
	surfaceArea := 6 * math.Pow(sideLength, 2)

	// Выводим результаты
	fmt.Printf("Объем куба: %.2f\n", volume)
	fmt.Printf("Площадь поверхности куба: %.2f\n", surfaceArea)
}
