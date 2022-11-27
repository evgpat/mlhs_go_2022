// B. Минимальный делитель числа
//
// Дано натуральное число n>1. Выведите его наименьший делитель, отличный от 1.
// Алгоритм должен иметь сложность O(sqrt(n)).
//
// Указание. Если у числа n нет делителя не превосходящего sqrt(n), то число n — простое и ответом будет само число n.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	result := n
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			result = i
			break
		}
	}

	fmt.Println(result)
}
