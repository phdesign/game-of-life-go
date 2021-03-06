package app

import (
	"reflect"
	"testing"
)

func TestSeed(t *testing.T) {
	w := 3
	h := 3

	t.Run("should initalise a board with some live cells", func(t *testing.T) {
		want := Board{[]Row{
			{0, 0, 0},
			{0, 1, 0},
			{1, 0, 0},
		}}
		board := NewBoard(w, h)
		board.Seed(12)

		assertBoardEqual(t, &want, board)
	})
}

func TestNewBoard(t *testing.T) {
	w := 4
	h := 3

	t.Run("should count number of alive neighbours", func(t *testing.T) {
		board := NewBoard(w, h)

		got := len(board.Rows[0])
		want := 4
		if got != want {
			t.Errorf("Expected board width to be %d, got %d", want, got)
		}
		got = len(board.Rows)
		want = 3
		if got != want {
			t.Errorf("Expected board height to be %d, got %d", want, got)
		}
	})
}

func TestCountNeighbours(t *testing.T) {
	board := Board{[]Row{
		{1, 0, 0},
		{1, 0, 0},
		{1, 0, 0},
	}}

	t.Run("should count number of alive neighbours", func(t *testing.T) {
		assertIntEqual(t, 3, countNeighbours(&board, 1, 1))
		assertIntEqual(t, 0, countNeighbours(&board, 1, 2))
		assertIntEqual(t, 2, countNeighbours(&board, 0, 1))
	})
}

func TestTick(t *testing.T) {
	board := Board{[]Row{
		{1, 0, 0},
		{1, 0, 0},
		{1, 0, 0},
	}}

	t.Run("should evolve by one generation", func(t *testing.T) {
		want := Board{[]Row{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0},
		}}
		got := board.Tick()

		assertBoardEqual(t, &want, got)
	})
}

func assertIntEqual(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("Want %d, got %d", want, got)
	}
}

func assertBoardEqual(t *testing.T, want, got *Board) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v, got %v", want, got)
	}
}
