package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/wishes", func(r render.Render) {
		r.HTML(200, "list", nil)
	})

	m.Run()
}
