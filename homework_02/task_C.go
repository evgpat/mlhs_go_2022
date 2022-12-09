/*
C. Банковские счета

Некоторый банк хочет внедрить систему управления счетами клиентов, поддерживающую следующие операции:
- Пополнение счета клиента
- Снятие денег со счета
- Запрос остатка средств на счете
- Перевод денег между счетами клиентов
- Начисление процентов всем клиентам

Вам необходимо реализовать такую систему. Клиенты банка идентифицируются именами
(уникальная строка, не содержащая пробелов). Первоначально у банка нет ни одного клиента.
Как только для клиента проводится операция пололнения, снятия или перевода денег, ему заводится счет с нулевым балансом.
Все дальнейшие операции проводятся только с этим счетом. Сумма на счету может быть как положительной,
так и отрицательной, при этом всегда является целым числом.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	accounts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		tokens := strings.Split(strings.Trim(strings.TrimSpace(scanner.Text()), "\n"), " ")
		command := tokens[0]
		client := tokens[1]

		switch command {
		case "DEPOSIT":
			amount, _ := strconv.Atoi(tokens[2])
			accounts[client] += amount
		case "WITHDRAW":
			amount, _ := strconv.Atoi(tokens[2])
			accounts[client] -= amount
		case "BALANCE":
			value, ok := accounts[client]
			if !ok {
				fmt.Println("ERROR")
			} else {
				fmt.Println(value)
			}
		case "TRANSFER":
			amount, _ := strconv.Atoi(tokens[3])
			client2 := tokens[2]
			accounts[client] -= amount
			accounts[client2] += amount
		case "INCOME":
			rate, _ := strconv.ParseFloat(tokens[1], 64)
			rate = 1 + rate/100
			for key, value := range accounts {
				if value > 0 {
					accounts[key] = int(float64(accounts[key]) * rate)
				}
			}
		}
	}
}
