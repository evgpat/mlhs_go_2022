/*
F. Самое частое слово

Дан текст. Выведите слово, которое в этом тексте встречается чаще всего. Если таких слов несколько, выведите то,
которое меньше в лексикографическом порядке.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var maxValue int
	var maxKey string
	counter := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		words := strings.Split(strings.Trim(strings.TrimSpace(scanner.Text()), "\n"), " ")
		for _, key := range words {
			counter[key] += 1
		}
	}

	for key, value := range counter {
		if value < maxValue {
			continue
		} else if value > maxValue {
			maxValue = value
			maxKey = key
		} else if maxKey > key {
			maxKey = key
		} else {
			continue
		}
	}

	fmt.Println(maxKey)
}
