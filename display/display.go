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

func Draw(board [][]int8, seed int64, generation, w, h int) {
	for row := 0; row < len(board)-2; row += 2 {
		for col := 0; col < len(board[0])-2; col += 2 {
			x := col / 2
			y := row / 2
			quad := [4]int8{
				board[row][col],
				board[row][col+1],
				board[row+1][col],
				board[row+1][col+1],
			}
			drawPixels(quad, x, y)
		}
	}
	printText(fmt.Sprintf("seed: %d, w: %d, h: %d, gen: %d", seed, w, h, generation), 0)
	printText("Press ESC to quit.", 1)
	termbox.Flush()
}

const (
	pixelTop                    rune = 0x2580
	pixelBottom                 rune = 0x2584
	pixelLeft                   rune = 0x258C
	pixelRight                  rune = 0x2590
	pixelLowerLeft              rune = 0x2596
	pixelLowerRight             rune = 0x2597
	pixelUpperLeft              rune = 0x2598
	pixelLeftAndLowerRight      rune = 0x2599
	pixelUpperLeftAndLowerRight rune = 0x259A
	pixelLeftAndUpperRight      rune = 0x259B
	pixelRightAndUpperLeft      rune = 0x259C
	pixelUpperRight             rune = 0x259D
	pixelUpperRightAndLowerLeft rune = 0x259E
	pixelRightAndLowerLeft      rune = 0x259F
	pixelFull                   rune = 0x2588
)

var pixelMap = []rune{
	' ',
	pixelUpperLeft,
	pixelUpperRight,
	pixelTop,
	pixelLowerLeft,
	pixelLeft,
	pixelUpperRightAndLowerLeft,
	pixelLeftAndUpperRight,
	pixelLowerRight,
	pixelUpperLeftAndLowerRight,
	pixelRight,
	pixelRightAndUpperLeft,
	pixelBottom,
	pixelLeftAndLowerRight,
	pixelRightAndLowerLeft,
	pixelFull,
}

func drawPixels(quad [4]int8, x, y int) {
	var idx int8 = 0
	for i, a := range quad {
		idx += a << i
	}
	termbox.SetChar(x, y, pixelMap[idx])
}

func printText(text string, row int) {
	for i, char := range text {
		termbox.SetChar(i, row, rune(char))
	}
}

func Wait() {
	termbox.PollEvent()
}

func Close() {
	termbox.Close()
}
