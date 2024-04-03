/* Треугольник задан координатами своих вершин. Найти периметр и площадь треугольника. */

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
	a := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	b := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
	c := math.Sqrt(math.Pow(x1-x3, 2) + math.Pow(y1-y3, 2))

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
