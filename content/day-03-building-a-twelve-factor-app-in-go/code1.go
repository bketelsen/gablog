package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	debug bool
)

func main() {
	raw_debug := os.Getenv("MYAPP_DEBUG")
	debug, err := strconv.ParseBool(raw_debug)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Debug is set to: %v\n", debug)
}
