/*
D. Ферзи
Известно, что на доске 8×8 можно расставить 8 ферзей так, чтобы они не били друг друга.
Вам дана расстановка 8 ферзей на доске, определите, есть ли среди них пара бьющих друг друга.
*/

package main

import (
	"fmt"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func check(board [8][2]int, n int) string {
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if board[i][0] == board[j][0] || board[i][1] == board[j][1] ||
				(diff(board[i][0], board[j][0]) == diff(board[i][1], board[j][1])) {
				return "YES"
			}
		}
	}
	return "NO"
}

func main() {
	const n = 8
	const k = 2

	var coordinates [n][k]int
	for i := 1; i < n; i++ {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			return
		}
		coordinates[i][0], coordinates[i][1] = x, y
	}

	fmt.Println(check(coordinates, n))
}
