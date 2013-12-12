package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		fmt.Println("Received Interrupt")
		os.Exit(1)
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
