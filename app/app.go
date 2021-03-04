package app

import (
	"math/rand"
)

func countNeighbours(board map[int]bool, pos int, w int, h int) (count int) {
	count = 0
	neighbours := [8]int{pos - w - 1, pos - w, pos - w + 1, pos - 1, pos + 1, pos + w - 1, pos + w, pos + w + 1}
	for _, n := range neighbours {
		if n >= 0 && n < w*h {
			if board[n] {
				count++
			}
		}
	}
	return
}

func Seed(w int, h int, seed int64) map[int]bool {
	board := make(map[int]bool)

	for i := 0; i < w*h; i++ {
		board[i] = false
	}

	rand.Seed(seed)
	for i := 0; i < rand.Intn(w*h); i++ {
		board[rand.Intn(w*h)] = true
	}

	return board
}

func Tick(board map[int]bool, w int, h int) {
	for pos, alive := range board {
		neighbours := countNeighbours(board, pos, w, h)
		if alive {
			if neighbours < 2 || neighbours > 3 {
				board[pos] = false
			}
		} else {
			if neighbours == 3 {
				board[pos] = true
			}
		}
	}
}
