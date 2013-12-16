package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filepath.ToSlash(
		filepath.Join(append([]string{pwd}, os.Args[1:]...)...)))
}
