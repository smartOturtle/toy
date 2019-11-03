package main

import (
	"fmt"
	"strconv"
	"strings"
)

var sign = false

var sudoku = [9][9]int{}

// Sudoku is used to get a solution of given sudoku
func Sudoku(s string) string {
	if len(s) < 81 {
		return "Solution Not Found"
	}
	ss := strings.Split(s, "")
	for i := 0; i < 81; i++ {
		sudoku[i/9][i%9], _ = strconv.Atoi(ss[i])
	}
	sign = false
	dfs(0)
	return output(sudoku)
}

func output(sudoku [9][9]int) string {
	if !sign {
		return "Solution Not Found"
	}
	var sb strings.Builder
	for i := 0; i < 81; i++ {
		fmt.Fprintf(&sb, "%2d ", sudoku[i/9][i%9])
		if i != 0 && (i+1)%9 == 0 {
			fmt.Fprintln(&sb)
		}
	}
	return sb.String()
}

func check(n int, key int) bool {
	for i := 0; i < 9; i++ {
		r := n / 9
		if sudoku[r][i] == key {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		c := n % 9
		if sudoku[i][c] == key {
			return false
		}
	}

	x := n / 9 / 3 * 3
	y := n % 9 / 3 * 3

	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if sudoku[i][j] == key {
				return false
			}
		}
	}

	return true
}

func dfs(n int) {
	if n > 80 {
		sign = true
		return
	}

	if sudoku[n/9][n%9] != 0 {
		dfs(n + 1)
	} else {
		for v := 1; v <= 9; v++ {
			if check(n, v) {
				sudoku[n/9][n%9] = v
				dfs(n + 1)
				if sign {
					return
				}
				sudoku[n/9][n%9] = 0
			}
		}
	}
}
