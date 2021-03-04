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

	w, h := display.Init()

	board := app.Seed(w, h, seed)
	display.Draw(board, w)

	time.Sleep(500 * time.Millisecond)

	app.Tick(board, w, h)
	display.Draw(board, w)

	display.Close()
}
