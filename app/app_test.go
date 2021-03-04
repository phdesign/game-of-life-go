package app

import "testing"

/*
0 1 2.3
4 5.6 7
8.9.0 1
*/
func TestCountNeighbours(t *testing.T) {
	w := 4
	h := 3
	board := make(map[int]bool)
	for i := 0; i < w*h; i++ {
		board[i] = false
	}
	board[2] = true
	board[5] = true
	board[9] = true
	board[8] = true

	t.Run("should count number of alive neighbours", func(t *testing.T) {
		want := 3
		got := countNeighbours(board, 6, w, h)

		if got != want {
			t.Errorf("Want %d, got %d", want, got)
		}
	})

	t.Run("should count number of alive neighbours in row only", func(t *testing.T) {
		want := 1
		got := countNeighbours(board, 7, w, h)

		if got != want {
			t.Errorf("Want %d, got %d", want, got)
		}
	})
}

func TestTick(t *testing.T) {
	w := 4
	h := 3
	board := make(map[int]bool)
	for i := 0; i < w*h; i++ {
		board[i] = false
	}
	board[2] = true
	board[5] = true
	board[9] = true
	board[8] = true

	t.Run("should evolve by one generation", func(t *testing.T) {
		want := true
		Tick(board, w, h)
		got := board[4]

		if got != want {
			t.Errorf("Want %v, got %v", want, got)
		}
	})
}
