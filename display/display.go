package display

import "github.com/nsf/termbox-go"

func Init() (int, int) {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	return w, h
}

func Close() {
	termbox.PollEvent()
	termbox.Close()
}
