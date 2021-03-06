package app

import (
	"math/rand"
)

type CellPos struct {
	col int
	row int
}

type Board struct {
	Rows []Row
}

type Row = []uint8

func NewBoard(w int, h int) *Board {
	board := new(Board)
	board.Rows = make([]Row, h)
	for i := range board.Rows {
		board.Rows[i] = make(Row, w)
	}
	return board
}

func (b *Board) Seed(seed int64) {
	w := len(b.Rows[0])
	rand.Seed(seed)
	for row := range b.Rows {
		for num := 0; num < rand.Intn(w); num++ {
			col := rand.Intn(w)
			b.Rows[row][col] = 1
		}
	}
}

func (b *Board) Clone() *Board {
	clone := new(Board)
	clone.Rows = make([]Row, len(b.Rows))
	for i := range b.Rows {
		clone.Rows[i] = make(Row, len(b.Rows[i]))
		copy(clone.Rows[i], b.Rows[i])
	}
	return clone
}

func (b *Board) Tick() *Board {
	result := b.Clone()
	for row, cells := range b.Rows {
		for col, alive := range cells {
			neighbours := countNeighbours(b, row, col)
			if alive == 1 {
				if neighbours < 2 || neighbours > 3 {
					result.Rows[row][col] = 0
				}
			} else {
				if neighbours == 3 {
					result.Rows[row][col] = 1
				}
			}
		}
	}
	return result
}

func countNeighbours(board *Board, row int, col int) int {
	count := 0
	neighbours := [8]CellPos{
		{col - 1, row - 1}, {col, row - 1}, {col + 1, row - 1},
		{col - 1, row}, {col + 1, row},
		{col - 1, row + 1}, {col, row + 1}, {col + 1, row + 1},
	}

	for _, c := range neighbours {
		if c.row >= 0 && c.row < len(board.Rows) &&
			c.col >= 0 && c.col < len(board.Rows[0]) {
			count += int(board.Rows[c.row][c.col])
		}
	}
	return count
}
