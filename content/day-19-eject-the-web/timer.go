package main

import (
	"time"
)

func main() {
	// START OMIT
	dirty := false
	timer := time.AfterFunc(0, func() {
		if dirty {
			redraw_full()
		} else {
			redraw_part()
		}
	})
	for {
		select {
		case ev := <-event:
			// handle terminal event
			switch ev.Type {
			case termbox.EventKey:
				// handle key event
				switch ev.Key {
				case termbox.KeyEnter:
					// update internal condition
					update_condition(ev)
					dirty = false
					// redraw immediately
					timer.Reset(1 * time.Microsecond)
				default:
					// update candidate
					update_candidate(ev)
					dirty = true
					// redrwa in later
					timer.Reset(200 * time.Microsecond)
				}
			}
			// handle another events
		}
	}
	// END OMIT
}
