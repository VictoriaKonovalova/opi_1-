package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите неотрицательное натуральное число:")
	fmt.Scan(&num)

	// Убираем последнюю цифру (единицы)
	num = num / 10

	// Находим новую последнюю цифру (десятки)
	tens := num % 10

	fmt.Println("Число десятков:", tens)
}
