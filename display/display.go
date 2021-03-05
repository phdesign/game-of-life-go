package display

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

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

func Loop(callback func()) {
	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	callback()
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:
			callback()
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func Draw(board [][]int8, seed int64, generation int) {
	for row, cells := range board {
		for col, state := range cells {
			colour := termbox.ColorDefault
			if state == 1 {
				colour = termbox.ColorGreen
			}
			x := col
			y := row
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, colour)
		}
	}
	text := fmt.Sprintf("seed: %d, gen: %d", seed, generation)
	for i, char := range text {
		termbox.SetChar(i, 0, rune(char))
	}
	termbox.Flush()
}

func Close() {
	termbox.Close()
}
