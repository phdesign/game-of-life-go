package app

import (
	"math/rand"
)

type CellPos struct {
	col int
	row int
}

func countNeighbours(board Board, row int, col int) int {
	count := 0
	neighbours := [8]CellPos{
		{col - 1, row - 1}, {col, row - 1}, {col + 1, row - 1},
		{col - 1, row}, {col + 1, row},
		{col - 1, row + 1}, {col, row + 1}, {col + 1, row + 1},
	}

	for _, c := range neighbours {
		if c.row >= 0 && c.row < len(board) &&
			c.col >= 0 && c.col < len(board[0]) {
			count += int(board[c.row][c.col])
		}
	}
	return count
}

func cloneBoard(orig Board) Board {
	clone := make(Board, len(orig))
	for i := range orig {
		clone[i] = make(Row, len(orig[i]))
		copy(clone[i], orig[i])
	}
	return clone
}

func NewBoard(w int, h int) Board {
	board := make(Board, h)
	for i := range board {
		board[i] = make(Row, w)
	}
	return board
}

func Seed(board Board, seed int64) Board {
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

func Tick(board Board) Board {
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
