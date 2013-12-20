func work() {
    exitChan := make(chan int)
    go task1(exitChan)
    go task2(exitChan)
    time.Sleep(5 * time.Second)
    close(exitChan)
}

func task1(exitChan chan int) {
    <-exitChan
    log.Printf("task1 exiting")
}

func task2(exitChan chan int) {
    <-exitChan
    log.Printf("task2 exiting")
}
