/* Треугольник задан координатами своих вершин. Найти периметр и площадь
треугольника.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Println("Введите координаты вершин треугольника (x1 y1 x2 y2 x3 y3):")
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	// Вычисляем длины сторон треугольника
	a := distance(x1, y1, x2, y2)
	b := distance(x2, y2, x3, y3)
	c := distance(x3, y3, x1, y1)

	// Вычисляем периметр
	perimeter := a + b + c

	// Вычисляем полупериметр для формулы Герона
	p := perimeter / 2

	// Вычисляем площадь по формуле Герона
	area := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	// Выводим результаты
	fmt.Printf("Периметр треугольника: %.2f\n", perimeter)
	fmt.Printf("Площадь треугольника: %.2f\n", area)
}

// Функция для вычисления расстояния между двумя точками
func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
