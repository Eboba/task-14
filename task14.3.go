package main

import (
	"fmt"
)

func addition(x int) (a int) {
	a += x
	return
}

func multiplication(у int) (b int) {
	b = 2 * у
	addition(b)
	return
}

func main() {
	fmt.Println("Задание 3. Именованные возвращаемые значения")

	fmt.Println(multiplication(3))
}
