package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a) // считываем значение первого числа
	fmt.Scan(&b) // считываем значение второго числа

	a = a * a // находим квадрат первого числа
	b = b * b // находим квадрат второго числа
	c = a + b // находим сумму квадратов двух чисел

	fmt.Println("Результат:", c)
}
