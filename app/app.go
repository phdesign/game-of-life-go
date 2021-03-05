package app

import (
	"math/rand"
)

type coord struct {
	col int
	row int
}

func countNeighbours(board [][]int8, row int, col int) int {
	count := 0
	neighbours := [8]coord{
		coord{col - 1, row - 1}, coord{col, row - 1}, coord{col + 1, row - 1},
		coord{col - 1, row}, coord{col + 1, row},
		coord{col - 1, row + 1}, coord{col, row + 1}, coord{col + 1, row + 1},
	}

	for _, c := range neighbours {
		if c.row >= 0 && c.row < len(board) &&
			c.col >= 0 && c.col < len(board[0]) {
			count += int(board[c.row][c.col])
		}
	}
	return count
}

func cloneBoard(orig [][]int8) [][]int8 {
	clone := make([][]int8, len(orig))
	for i := range orig {
		clone[i] = make([]int8, len(orig[i]))
		copy(clone[i], orig[i])
	}
	return clone
}

func NewBoard(w int, h int) [][]int8 {
	board := make([][]int8, h)
	for i := range board {
		board[i] = make([]int8, w)
	}
	return board
}

func Seed(board [][]int8, seed int64) [][]int8 {
	result := cloneBoard(board)
	w := len(board[0])
	rand.Seed(seed)
	for row := range board {
		for num := 0; num < rand.Intn(w); num++ {
			col := rand.Intn(w)
			result[row][col] = 1
		}
	}
	return result
}

func Tick(board [][]int8) [][]int8 {
	result := cloneBoard(board)
	for row, cells := range board {
		for col, alive := range cells {
			neighbours := countNeighbours(board, row, col)
			if alive == 1 {
				if neighbours < 2 || neighbours > 3 {
					result[row][col] = 0
				}
			} else {
				if neighbours == 3 {
					result[row][col] = 1
				}
			}
		}
	}
	return result
}
