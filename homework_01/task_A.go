// A. Минимум 4 чисел
//
// Напишите функцию min4(a, b, c, d), вычисляющую минимум четырех чисел, которая не содержит инструкции if,
// а использует стандартную функцию min. Считайте четыре целых числа и выведите их минимум.

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	_, i := fmt.Scan(&a, &b, &c, &d)
	if i != nil {
		return
	}
	fmt.Println(int(math.Min(math.Min(a, b), math.Min(c, d))))
}
