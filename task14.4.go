// надеюсь правильно понял задание

package main

import (
	"fmt"
)

var a, b, c = 10, 20, 30

func first(x int) int {
	return x + a
}

func second(у int) int {
	return у + b
}

func third(z int) int {
	return z + c
}

func main() {
	fmt.Println("Задание 4. Область видимости переменных")

	fmt.Println(third(second(first(3))))
}
