package main

//START1 OMIT
type Event int

func NewEvent() Event                   { return 0 }
func (e Event) Merge(other Event) Event { return e + other }

func produce(out chan<- Event) {
	for {
		delay := time.Duration(rand.Intn(5)+1) * time.Second
		nMessages := rand.Intn(10) + 1
		time.Sleep(delay)
		for i := 0; i < nMessages; i++ {
			out <- Event(rand.Intn(10))
		}
	}
}

//END1 OMIT

//START2 OMIT
func coalesce500(in <-chan Event, out chan<- Event) {
	ticker := time.NewTicker(500 * time.Millisecond)
	event := NewEvent()
	for {
		select {
		case <-ticker.C:
			out <- event
			event = NewEvent()
		case e := <-in:
			event = event.Merge(e)
		}
	}
}

//END2 OMIT

//START3 OMIT
func coalesceSlow(in <-chan Event, out chan<- Event) {
	event := NewEvent()
	for e := range in {
		event = event.Merge(e)
	loop:
		for {
			select {
			case e := <-in:
				event = event.Merge(e)
			case out <- event:
				event = NewEvent()
				break loop
			}
		}
	}
}

//END3 OMIT

//START4 OMIT
func coalesce(in <-chan Event, out chan<- Event) {
	event := NewEvent()
	timer := time.NewTimer(0)

	var timerCh <-chan time.Time
	var outCh chan<- Event

	for {
		select {
		case e := <-in:
			event = event.Merge(e)
			if timerCh == nil {
				timer.Reset(500 * time.Millisecond)
				timerCh = timer.C
			}
		case <-timerCh:
			outCh = out
			timerCh = nil
		case outCh <- event:
			event = NewEvent()
			outCh = nil
		}
	}
}

//END4 OMIT
