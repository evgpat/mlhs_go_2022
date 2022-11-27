/*
C. Кегельбан

N кеглей выставили в один ряд, занумеровав их слева направо числами от 1 до N. Затем по этому ряду бросили K шаров,
при этом i-й шар сбил все кегли с номерами от l_i до r_i включительно. Определите, какие кегли остались стоять на месте.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n, k, x, y int
		skittles   string
	)
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	skittles = strings.Repeat("I", n)
	for i := 0; i < k; i++ {
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			return
		}
		x--
		y--
		skittles = skittles[:x] + strings.Repeat(".", y-x+1) + skittles[y+1:]
	}

	fmt.Println(skittles)
}
