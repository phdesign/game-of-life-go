package main

import (
	"flag"
	"github.com/phdesign/game-of-life-go/app"
	"github.com/phdesign/game-of-life-go/display"
	"strconv"
	"time"
)

const delay int = 150

func main() {
	flag.Parse()
	seedString := flag.Arg(0)
	var seed int64
	if seedString != "" {
		var err error
		seed, err = strconv.ParseInt(seedString, 10, 64)
		if err != nil {
			panic(err)
		}
	} else {
		seed = time.Now().UnixNano()
	}

	generation := 0
	w, h := display.Init()
	width := w * 2
	height := h * 2

	board := app.NewBoard(width, height)
	board.Seed(seed)
	display.Draw(board, seed, generation, width, height)

	display.Loop(func() {
		generation++
		time.Sleep(time.Duration(delay) * time.Millisecond)
		board = board.Tick()
		display.Draw(board, seed, generation, width, height)
	})

	display.Close()
}
