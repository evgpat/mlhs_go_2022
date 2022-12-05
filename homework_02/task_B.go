/*
B. Теорема Лагранжа

Теорема Лагранжа утверждает, что любое натуральное число можно представить в виде суммы четырех точных квадратов.

По данному числу n найдите такое представление:
напечатайте от 1 до 4 натуральных чисел, квадраты которых дают в сумме данное число.
*/

package main

import (
	"fmt"
)

func fourSquareTheorem(n int) {
	var a, b, c, d int

	for a = 1; a*a <= n; a++ {
		if a*a == n {
			fmt.Println(a)
			return
		}
		for b = 1; b <= n-a*a; b++ {
			if a*a+b*b == n {
				fmt.Println(a, b)
				return
			}
			for c = 1; c <= n-a*a-b*b; c++ {
				if a*a+b*b+c*c == n {
					fmt.Println(a, b, c)
					return
				}
				for d = 1; d <= n-a*a-b*b-c*c; d++ {
					if a*a+b*b+c*c+d*d == n {
						fmt.Println(a, b, c, d)
						return
					}
				}
			}
		}
	}
}

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fourSquareTheorem(n)
}
