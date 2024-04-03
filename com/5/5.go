package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите целое число:")
	fmt.Scan(&num)

	square := num * num

	fmt.Println("Квадрат числа:", square)
}
