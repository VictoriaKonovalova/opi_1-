package main

import "fmt"

func main() {
	var d int
	fmt.Println("Введите угол поворота часовой стрелки (0 < d < 360):")
	fmt.Scan(&d)

	// Находим количество часов
	hours := d / 30

	// Находим количество минут
	minutes := (d % 30) * 2

	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
