package main

import (
	"flag"
	"github.com/phdesign/game-of-life-go/app"
	"github.com/phdesign/game-of-life-go/display"
	"strconv"
	"time"
)

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

	board := app.NewBoard(w*2, h*2)
	board = app.Seed(board, seed)
	display.Draw(board, seed, generation, w*2, h*2)

	display.Loop(func() {
		generation++
		time.Sleep(200 * time.Millisecond)
		board = app.Tick(board)
		display.Draw(board, seed, generation, w*2, h*2)
	})

	display.Close()
}
