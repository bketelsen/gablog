package main

import "..."

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Merry Christmas!"
	})

	m.Run()

}
