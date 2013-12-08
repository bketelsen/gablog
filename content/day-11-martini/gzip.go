package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/gzip"
)

func main() {
  m := martini.Classic()
  //START OMIT
  m.Use(gzip.All())
  //END OMIT
  m.Run()
}
