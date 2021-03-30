package main

import "fmt"

func printBoard(sudoku [][]int) {
	l := len(sudoku)
	w := len(sudoku[0])
	for i := 1; i <= l; i++ {
		for k := 1; k <= w; k++ {
			fmt.Printf("%d ", sudoku[i-1][k-1])
			if k == l {
				fmt.Printf("\n")
			}
		}
	}
}

func find(in []int, key int) bool {
	for i := range in {
		if in[i] == key {
			return true
		}
	}
	return false
}

func getNextEmpty(s [][]int, loc *[2]int) bool {
	l := len(s)
	w := len(s[0])
	for i := 0; i < l; i++ {
		for k := 0; k < w; k++ {
			if (s)[i][k] == 0 {
				(*loc)[0] = i
				(*loc)[1] = k
				return true
			}
		}
	}
	return false
}

func validBox(s *[][]int, num int, loc [2]int) bool {
	col := (loc[0] / 3) * 3
	row := (loc[1] / 3) * 3
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if (*s)[i+col][k+row] == num {
				return true
			}
		}
	}
	return false
}

func validRow(s *[][]int, num int, loc [2]int) bool {
	i := loc[0]
	w := len((*s)[0])
	for k := 0; k < w; k++ {
		if (*s)[i][k] == num {
			return true
		}
	}
	return false
}

func validColumn(s *[][]int, num int, loc [2]int) bool {
	i := loc[1]
	w := len((*s))
	for k := 0; k < w; k++ {
		if (*s)[k][i] == num {
			return true
		}
	}
	return false
}

func valid(s *[][]int, num int, loc [2]int) bool {
	return !validBox(s, num, loc) && !validRow(s, num, loc) && !validColumn(s, num, loc)
}

func bruteForce(s *[][]int) bool {
	loc := [2]int{0, 0}
	if !getNextEmpty(*s, &loc) {
		return true
	}

	for i := 1; i < 10; i++ {
		if valid(s, i, loc) {
			(*s)[loc[0]][loc[1]] = i
			if bruteForce(s) {
				return true
			}
			(*s)[loc[0]][loc[1]] = 0
		}
	}

	return false
}

func main() {
	sudoku := [][]int{{0, 4, 0, 0, 0, 0, 0, 0, 7},
		{5, 0, 0, 0, 8, 0, 0, 6, 0},
		{0, 0, 0, 0, 5, 0, 4, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 8, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 0, 6, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 2},
		{6, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 4}}

	if bruteForce(&sudoku) {
		printBoard(sudoku)
	} else {
		fmt.Println("No solution exists")
	}

}
