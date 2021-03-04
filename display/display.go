package display

import "github.com/nsf/termbox-go"

func Init() (int, int) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	return w, h
}

func Draw(board map[int]bool, width int) {
	for pos, state := range board {
		colour := termbox.ColorDefault
		if state {
			colour = termbox.ColorGreen
		}
		x := pos % width
		y := pos / width
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, colour)
	}
	termbox.Flush()
}

func Close() {
	termbox.PollEvent()
	termbox.Close()
}
