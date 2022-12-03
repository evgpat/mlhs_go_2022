/*
A. Выборы в США

Как известно, в США президент выбирается не прямым голосованием, а путем двухуровневого голосования.
Сначала проводятся выборы в каждом штате и определяется победитель выборов в данном штате.

Затем проводятся государственные выборы:
на этих выборах каждый штат имеет определенное число голосов — число выборщиков от этого штата.
На практике, все выборщики от штата голосуют в соответствии с результами голосования внутри штата,
то есть на заключительной стадии выборов в голосовании участвуют штаты, имеющие различное число голосов.

Вам известно за кого проголосовал каждый штат и сколько голосов было отдано данным штатом.
Подведите итоги выборов: для каждого из участника голосования определите число отданных за него голосов.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	counter := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		tokens := strings.Split(strings.Trim(strings.TrimSpace(scanner.Text()), "\n"), " ")

		value, _ := strconv.Atoi(tokens[1])
		counter[tokens[0]] += value
	}

	keys := make([]string, 0, len(counter))

	for key, _ := range counter {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, counter[key])
	}
}
