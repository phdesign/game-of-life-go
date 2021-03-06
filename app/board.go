package app

type Board = []Row
type Row = []uint8

type Generation struct {
	width      int
	height     int
	generation int
	board      Board
}
