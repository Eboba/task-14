package main

import (
	"fmt"
	"math/rand"
)

func random(x1, y1, x2, y2, x3, y3 int) (int, int, int, int, int, int) {
	return rand.Intn(x1), rand.Intn(y1), rand.Intn(x2), rand.Intn(y2), rand.Intn(x3), rand.Intn(y3)
}

func conversion(x, y int) (int, int) {
	return 2*x + 10, -3*y - 5
}

func main() {
	fmt.Println("Задание 2. Функция, возвращающая несколько значений")

	x1, y1, x2, y2, x3, y3 := random(100, 100, 100, 100, 100, 100)

	fmt.Println(conversion(x1, y1))
	fmt.Println(conversion(x2, y2))
	fmt.Println(conversion(x3, y3))
}
