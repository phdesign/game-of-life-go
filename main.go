package main

import (
	"fmt"
	//"github.com/phdesign/game-of-life-go/display"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func main() {
	// Get a greeting message and print it.
	message := Hello("Gladys")
	fmt.Println(message)
	//w, _ := display.Init()
	//fmt.Println(w)
	//display.Close()
}
