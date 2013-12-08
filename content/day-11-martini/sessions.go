package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/sessions"

)

func main() {
  m := martini.Classic()

  //START OMIT
  store := sessions.NewCookieStore([]byte("secret123"))
  m.Use(sessions.Sessions("my_session", store))

  m.Get("/set", func(session sessions.Session) string {
    session.Set("hello", "world")
    return "OK"
  })

  m.Get("/get", func(session sessions.Session) string {
    v := session.Get("hello")
    if v == nil {
      return ""
    }
    return v.(string)
  })
  //END OMIT

  m.Run()
}
