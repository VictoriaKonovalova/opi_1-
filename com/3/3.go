package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Введите число: ")
	fmt.Scan(&input)

	result := input * 2
	result += 100

	fmt.Print(result)
}
