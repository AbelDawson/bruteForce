package main

import (
	"testing"
)

// check for every element of a slice
func TestFind(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	for i := range arr {
		if !find(arr, i) {
			t.Fatalf("Number %d not in range", i)
		}
	}
}

// check for an element not in slice
func TestFindOutside(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	num := 6
	if find(arr, num) {
		t.Fatalf("Number %d should not have been found", num)
	}
}

// check for the correct empty spot
func TestGetNextEmpty(t *testing.T) {
	slice := [][]int{{1, 1, 0}, {1, 1, 1}, {1, 1, 1}}
	loc := [2]int{0, 0}
	if !getNextEmpty(slice, &loc) && loc[1] != 2 {
		t.Fatalf("getNextEmpty returned incorrect value of %d", loc[1])
	}
}

// check for no empty spot
func TestGetNextEmptyNoEmpty(t *testing.T) {
	slice := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	loc := [2]int{0, 0}
	if getNextEmpty(slice, &loc) {
		t.Fatalf("Found empty sport at %d and %d. This should have been empty", loc[0], loc[1])
	}
}

// check to accept in the correct position
func TestValidBox(t *testing.T) {

	slice := [][]int{{0, 1, 3}, {4, 5, 6}, {7, 8, 9}}
	loc := [2]int{0, 0}

	if validBox(&slice, 2, loc) {
		t.Fatalf("2 not accepted where it is supposed to go")
	}
}

// check it doesn't accept incorrect locations
func TestInvalidBox(t *testing.T) {

	slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	loc := [2]int{0, 0}

	for i := 1; i <= 9; i++ {
		if !validBox(&slice, i, loc) {
			t.Fatalf("%d was accepted where wasn't supposed to go", i)
		}
	}

}

// check to accept in the correct position
func TestValidRow(t *testing.T) {
	slice := [][]int{{0, 2, 3, 4, 5, 6, 7, 8, 9}}
	loc := [2]int{0, 0}

	if validRow(&slice, 1, loc) {
		t.Fatalf("1 was accepted where wasn't supposed to go")
	}
}

// check it doesn't accept incorrect locations
func TestInvalidRow(t *testing.T) {
	slice := [][]int{{0, 2, 3, 4, 5, 6, 7, 8, 9}}
	loc := [2]int{0, 0}

	for i := 2; i < 9; i++ {
		if !validRow(&slice, i, loc) {
			t.Fatalf("%d was accepted where wasn't supposed to go", i)
		}
	}
}

// check to accept in the correct position
func TestValidColumn(t *testing.T) {
	slice := [][]int{{0}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	loc := [2]int{0, 0}

	if validColumn(&slice, 1, loc) {
		t.Fatalf("1 was accepted where wasn't supposed to go")
	}
}

// check it doesn't accept incorrect locations
func TestInvalidColumn(t *testing.T) {
	slice := [][]int{{0}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	loc := [2]int{0, 0}

	for i := 2; i < 9; i++ {
		if !validColumn(&slice, i, loc) {
			t.Fatalf("%d was accepted where wasn't supposed to go", i)
		}
	}
}

func equal(s [][]int, w [][]int) bool {
	if len(s) != len(w) {
		return false
	}
	for i, v := range s {
		for k, key := range v {
			if key != w[i][k] {
				return false
			}
		}
	}
	return true
}

func TestBruteForce(t *testing.T) {
	sudoku := [][]int{{0, 4, 0, 0, 0, 0, 0, 0, 7},
		{5, 0, 0, 0, 8, 0, 0, 6, 0},
		{0, 0, 0, 0, 5, 0, 4, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 8, 0},
		{0, 3, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 0, 6, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 2},
		{6, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 4}}

	want := [][]int{{2, 4, 1, 3, 6, 9, 8, 5, 7},
		{5, 7, 3, 1, 8, 4, 2, 6, 9},
		{8, 6, 9, 2, 5, 7, 4, 1, 3},
		{1, 2, 4, 5, 7, 3, 9, 8, 6},
		{9, 3, 6, 8, 1, 2, 7, 4, 5},
		{7, 5, 8, 9, 4, 6, 3, 2, 1},
		{4, 1, 5, 7, 9, 8, 6, 3, 2},
		{6, 9, 2, 4, 3, 1, 5, 7, 8},
		{3, 8, 7, 6, 2, 5, 1, 9, 4}}

	if bruteForce(&sudoku) {
		if !(equal(sudoku, want)) {
			t.Fatalf("Didn't solve correct")
		}
	} else {
		t.Fatalf("Failed to solve a solveable board")
	}

}
