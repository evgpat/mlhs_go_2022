/*
E. Встречалось ли число раньше

Во входной строке записана последовательность чисел через пробел.
Для каждого числа выведите слово YES (в отдельной строке), если это число ранее встречалось в последовательности или NO,
если не встречалось.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	numbers := strings.Split(strings.Trim(strings.TrimSpace(scanner.Text()), "\n"), " ")

	counter := make(map[string]int)

	for _, key := range numbers {
		_, isAvailable := counter[key]
		if isAvailable {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		counter[key] = 0
	}
}
