package main

import (
	"fmt"
)

func examination(e int) bool {
	if e%2 != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println("Задание 1. Функция, возвращающая результат")

	var a int

	fmt.Println("Введите число:")
	fmt.Scan(&a)

	fmt.Println(examination(a))
}
