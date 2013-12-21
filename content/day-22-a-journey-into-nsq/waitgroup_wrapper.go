type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

// can be used as follows:
wg := WaitGroupWrapper{}
wg.Wrap(func() { n.idPump() })
// ...
wg.Wait()
