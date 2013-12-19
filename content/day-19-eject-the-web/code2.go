package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"log"
	"time"
)

func print_tb(x, y int, msg string) {
	for _, c := range []rune(msg) {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
		x += runewidth.RuneWidth(c)
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	event := make(chan termbox.Event)
	go func() {
		for {
			// Post events to channel
			event <- termbox.PollEvent()
		}
	}()

	print_tb(1, 1, "Hit any key")
loop:
	for {
		// Poll key event or timeout
		select {
		case ev := <-event:
			print_tb(1, 2, fmt.Sprintf("Key typed: %v", ev.Ch))
			break loop
		case <-time.After(5 * time.Second):
			print_tb(1, 2, "Timeout")
			break loop
		}
	}
	close(event)
	time.Sleep(1 * time.Second)
	termbox.Close()
}
